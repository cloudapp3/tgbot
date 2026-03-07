package main

import (
	"fmt"
	"go/format"
	"regexp"
	"sort"
	"strings"
)

var discriminatorPattern = regexp.MustCompile(`(?i)\b(?:always|must be)\s+["“']?([a-z0-9_]+)["”']?`)

func (g *generator) collectUnionDecoderSpecs() []unionDecoderSpec {
	rootSet := map[string]bool{}
	for _, owner := range g.unionFieldOwners {
		for _, field := range owner.Fields {
			rootSet[field.RootName] = true
		}
	}
	for _, method := range g.schema.Methods {
		if rootName, _ := g.unionFieldInfo(method.ReturnsExpr); rootName != "" {
			rootSet[rootName] = true
		}
	}

	rootNames := make([]string, 0, len(rootSet))
	for rootName := range rootSet {
		rootNames = append(rootNames, rootName)
	}
	sort.Strings(rootNames)

	result := make([]unionDecoderSpec, 0, len(rootNames))
	for _, rootName := range rootNames {
		if rootName == "InputMessageContent" {
			result = append(result, unionDecoderSpec{RootName: rootName, SpecialKind: "input_message_content"})
			continue
		}
		if rootName == "MaybeInaccessibleMessage" {
			result = append(result, unionDecoderSpec{RootName: rootName, SpecialKind: "maybe_inaccessible_message"})
			continue
		}
		result = append(result, g.inferUnionDecoderSpec(rootName))
	}
	return result
}

func (g *generator) inferUnionDecoderSpec(rootName string) unionDecoderSpec {
	subtypes := g.schema.UnionRoots[rootName]
	common := map[string]int{}
	for index, subtype := range subtypes {
		fieldNames := map[string]bool{}
		for _, field := range g.schema.Types[subtype].Fields {
			fieldNames[field.Name] = true
		}
		for fieldName := range fieldNames {
			if index == 0 {
				common[fieldName] = 1
				continue
			}
			if common[fieldName] == index {
				common[fieldName]++
			}
		}
	}

	candidates := make([]string, 0, len(common))
	for _, preferred := range g.preferredDiscNames {
		if common[preferred] == len(subtypes) {
			candidates = append(candidates, preferred)
		}
	}
	for fieldName, count := range common {
		if count == len(subtypes) && !containsString(candidates, fieldName) {
			candidates = append(candidates, fieldName)
		}
	}
	sort.Strings(candidates)

	for _, fieldName := range candidates {
		cases := make([]unionCase, 0, len(subtypes))
		seenTags := map[string]bool{}
		ok := true
		for _, subtype := range subtypes {
			field, found := findField(g.schema.Types[subtype].Fields, fieldName)
			if !found {
				ok = false
				break
			}
			tagMatch := discriminatorPattern.FindStringSubmatch(field.Description)
			if len(tagMatch) != 2 {
				ok = false
				break
			}
			tag := strings.TrimSpace(tagMatch[1])
			if tag == "" || seenTags[tag] {
				ok = false
				break
			}
			seenTags[tag] = true
			cases = append(cases, unionCase{Subtype: subtype, Tag: tag})
		}
		if ok {
			return unionDecoderSpec{
				RootName:           rootName,
				DiscriminatorField: fieldName,
				Cases:              cases,
			}
		}
	}

	panic(fmt.Sprintf("unable to infer discriminator for union %s", rootName))
}

func (g *generator) collectUnionStructs() []unionStructSpec {
	typeNames := make([]string, 0, len(g.schema.Types))
	for typeName := range g.schema.Types {
		if g.isUnionRoot(typeName) || g.supportTypeNames[typeName] {
			continue
		}
		typeNames = append(typeNames, typeName)
	}
	sort.Strings(typeNames)

	result := make([]unionStructSpec, 0)
	for _, typeName := range typeNames {
		spec := g.schema.Types[typeName]
		unionFields := make([]unionFieldSpec, 0)
		for _, field := range spec.Fields {
			rootName, isSlice := g.unionFieldInfo(field.TypeExpr)
			if rootName == "" {
				continue
			}
			unionFields = append(unionFields, unionFieldSpec{
				JSONName: field.Name,
				GoName:   toPascal(field.Name),
				RootName: rootName,
				IsSlice:  isSlice,
			})
		}
		if len(unionFields) == 0 {
			continue
		}
		result = append(result, unionStructSpec{TypeName: typeName, Fields: unionFields})
	}
	return result
}

func (g *generator) unionFieldInfo(typeExpr string) (string, bool) {
	normalized := strings.TrimSpace(typeExpr)
	if strings.HasPrefix(normalized, "Array of ") {
		rootName, _ := g.unionFieldInfo(strings.TrimPrefix(normalized, "Array of "))
		if rootName == "" {
			return "", false
		}
		return rootName, true
	}
	if g.isUnionRoot(normalized) {
		return normalized, false
	}
	if tokens := splitAlternatives(normalized); len(tokens) > 1 {
		if rootName := g.commonUnionRoot(tokens); rootName != "" {
			return rootName, false
		}
	}
	return "", false
}

func (g *generator) generateUnionFile() ([]byte, error) {
	var builder strings.Builder
	builder.WriteString("// Telegram Bot API union decoding helpers aligned with the official docs.\n")
	builder.WriteString("// Source: ")
	builder.WriteString(officialBotAPIURL)
	builder.WriteString(" (")
	builder.WriteString(g.schema.Version)
	builder.WriteString(", ")
	builder.WriteString(g.schema.VersionDate)
	builder.WriteString(")\n\n")
	builder.WriteString("package ")
	builder.WriteString(g.packageName)
	builder.WriteString("\n\n")
	builder.WriteString("import (\n")
	builder.WriteString("\t\"bytes\"\n")
	builder.WriteString("\t\"encoding/json\"\n")
	builder.WriteString("\t\"fmt\"\n")
	builder.WriteString("\t\"strings\"\n")
	builder.WriteString(")\n\n")

	g.writeUnionResultSwitch(&builder)
	writeUnionHelpers(&builder)
	for _, spec := range g.unionDecoderSpecs {
		g.writeUnionDecoder(&builder, spec)
		g.writeUnionSliceDecoder(&builder, spec)
	}
	for _, unionStruct := range g.unionFieldOwners {
		g.writeUnionUnmarshalMethod(&builder, unionStruct)
	}

	return format.Source([]byte(builder.String()))
}

func (g *generator) writeUnionResultSwitch(builder *strings.Builder) {
	builder.WriteString("func decodeIntoKnownUnionResult(target any, raw json.RawMessage) (bool, error) {\n")
	builder.WriteString("\tswitch result := target.(type) {\n")
	for _, spec := range g.unionDecoderSpecs {
		fmt.Fprintf(builder, "\tcase *%s:\n", spec.RootName)
		fmt.Fprintf(builder, "\t\tvalue, err := decode%s(raw)\n", spec.RootName)
		builder.WriteString("\t\tif err != nil {\n")
		builder.WriteString("\t\t\treturn true, err\n")
		builder.WriteString("\t\t}\n")
		builder.WriteString("\t\t*result = value\n")
		builder.WriteString("\t\treturn true, nil\n")
		fmt.Fprintf(builder, "\tcase *[]%s:\n", spec.RootName)
		fmt.Fprintf(builder, "\t\tvalue, err := decode%sSlice(raw)\n", spec.RootName)
		builder.WriteString("\t\tif err != nil {\n")
		builder.WriteString("\t\t\treturn true, err\n")
		builder.WriteString("\t\t}\n")
		builder.WriteString("\t\t*result = value\n")
		builder.WriteString("\t\treturn true, nil\n")
	}
	builder.WriteString("\tdefault:\n")
	builder.WriteString("\t\treturn false, nil\n")
	builder.WriteString("\t}\n")
	builder.WriteString("}\n\n")
}

func writeUnionHelpers(builder *strings.Builder) {
	builder.WriteString(`func splitUnionFields(data []byte, fields ...string) (map[string]json.RawMessage, []byte, error) {
	if len(data) == 0 {
		return nil, []byte("{}"), nil
	}

	allFields := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &allFields); err != nil {
		return nil, nil, err
	}

	union := make(map[string]json.RawMessage, len(fields))
	for _, field := range fields {
		if raw, ok := allFields[field]; ok {
			union[field] = raw
			delete(allFields, field)
		}
	}

	base, err := json.Marshal(allFields)
	if err != nil {
		return nil, nil, err
	}
	if len(base) == 0 {
		base = []byte("{}")
	}
	return union, base, nil
}

func decodeStringTag(raw json.RawMessage, field string) (string, error) {
	probe := map[string]json.RawMessage{}
	if err := json.Unmarshal(raw, &probe); err != nil {
		return "", err
	}

	value, ok := probe[field]
	if !ok {
		return "", fmt.Errorf("missing discriminator %q", field)
	}

	var result string
	if err := json.Unmarshal(value, &result); err != nil {
		return "", err
	}
	return strings.TrimSpace(result), nil
}

func isJSONNull(raw json.RawMessage) bool {
	trimmed := bytes.TrimSpace(raw)
	return len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null"))
}

func decodeRawSlice(raw json.RawMessage) ([]json.RawMessage, error) {
	if isJSONNull(raw) {
		return nil, nil
	}
	var values []json.RawMessage
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil, err
	}
	return values, nil
}

`)
}

func (g *generator) writeUnionDecoder(builder *strings.Builder, spec unionDecoderSpec) {
	fmt.Fprintf(builder, "func decode%s(raw json.RawMessage) (%s, error) {\n", spec.RootName, spec.RootName)
	builder.WriteString("\tif isJSONNull(raw) {\n")
	builder.WriteString("\t\treturn nil, nil\n")
	builder.WriteString("\t}\n")
	switch spec.SpecialKind {
	case "maybe_inaccessible_message":
		builder.WriteString("\tprobe := struct {\n")
		builder.WriteString("\t\tDate int64 `json:\"date\"`\n")
		builder.WriteString("\t}{}\n")
		builder.WriteString("\tif err := json.Unmarshal(raw, &probe); err != nil {\n")
		builder.WriteString("\t\treturn nil, err\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tif probe.Date == 0 {\n")
		builder.WriteString("\t\tvar value InaccessibleMessage\n")
		builder.WriteString("\t\treturn &value, json.Unmarshal(raw, &value)\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tvar value Message\n")
		builder.WriteString("\treturn &value, json.Unmarshal(raw, &value)\n")
	case "input_message_content":
		builder.WriteString("\tprobe := map[string]json.RawMessage{}\n")
		builder.WriteString("\tif err := json.Unmarshal(raw, &probe); err != nil {\n")
		builder.WriteString("\t\treturn nil, err\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tif _, ok := probe[\"message_text\"]; ok {\n")
		builder.WriteString("\t\tvar value InputTextMessageContent\n")
		builder.WriteString("\t\treturn &value, json.Unmarshal(raw, &value)\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tif _, ok := probe[\"phone_number\"]; ok {\n")
		builder.WriteString("\t\tvar value InputContactMessageContent\n")
		builder.WriteString("\t\treturn &value, json.Unmarshal(raw, &value)\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tif _, ok := probe[\"payload\"]; ok {\n")
		builder.WriteString("\t\tvar value InputInvoiceMessageContent\n")
		builder.WriteString("\t\treturn &value, json.Unmarshal(raw, &value)\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tif _, ok := probe[\"prices\"]; ok {\n")
		builder.WriteString("\t\tvar value InputInvoiceMessageContent\n")
		builder.WriteString("\t\treturn &value, json.Unmarshal(raw, &value)\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tif _, hasLatitude := probe[\"latitude\"]; hasLatitude {\n")
		builder.WriteString("\t\tif _, hasLongitude := probe[\"longitude\"]; hasLongitude {\n")
		builder.WriteString("\t\t\tif _, hasAddress := probe[\"address\"]; hasAddress {\n")
		builder.WriteString("\t\t\t\tvar value InputVenueMessageContent\n")
		builder.WriteString("\t\t\t\treturn &value, json.Unmarshal(raw, &value)\n")
		builder.WriteString("\t\t\t}\n")
		builder.WriteString("\t\t\tvar value InputLocationMessageContent\n")
		builder.WriteString("\t\t\treturn &value, json.Unmarshal(raw, &value)\n")
		builder.WriteString("\t\t}\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\treturn nil, fmt.Errorf(\"unknown input message content shape\")\n")
	default:
		fmt.Fprintf(builder, "\ttypeTag, err := decodeStringTag(raw, %q)\n", spec.DiscriminatorField)
		builder.WriteString("\tif err != nil {\n")
		builder.WriteString("\t\treturn nil, err\n")
		builder.WriteString("\t}\n")
		builder.WriteString("\tswitch typeTag {\n")
		for _, item := range spec.Cases {
			fmt.Fprintf(builder, "\tcase %q:\n", item.Tag)
			fmt.Fprintf(builder, "\t\tvar value %s\n", item.Subtype)
			builder.WriteString("\t\treturn &value, json.Unmarshal(raw, &value)\n")
		}
		builder.WriteString("\tdefault:\n")
		fmt.Fprintf(builder, "\t\treturn nil, fmt.Errorf(\"unknown %s discriminator %%q\", typeTag)\n", spec.RootName)
		builder.WriteString("\t}\n")
	}
	builder.WriteString("}\n\n")
}

func (g *generator) writeUnionSliceDecoder(builder *strings.Builder, spec unionDecoderSpec) {
	fmt.Fprintf(builder, "func decode%sSlice(raw json.RawMessage) ([]%s, error) {\n", spec.RootName, spec.RootName)
	builder.WriteString("\trawItems, err := decodeRawSlice(raw)\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn nil, err\n")
	builder.WriteString("\t}\n")
	fmt.Fprintf(builder, "\tresult := make([]%s, 0, len(rawItems))\n", spec.RootName)
	builder.WriteString("\tfor _, item := range rawItems {\n")
	fmt.Fprintf(builder, "\t\tvalue, err := decode%s(item)\n", spec.RootName)
	builder.WriteString("\t\tif err != nil {\n")
	builder.WriteString("\t\t\treturn nil, err\n")
	builder.WriteString("\t\t}\n")
	builder.WriteString("\t\tresult = append(result, value)\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn result, nil\n")
	builder.WriteString("}\n\n")
}

func (g *generator) writeUnionUnmarshalMethod(builder *strings.Builder, unionStruct unionStructSpec) {
	fmt.Fprintf(builder, "func (value *%s) UnmarshalJSON(data []byte) error {\n", unionStruct.TypeName)
	builder.WriteString("\tunion, base, err := splitUnionFields(data")
	for _, field := range unionStruct.Fields {
		fmt.Fprintf(builder, ", %q", field.JSONName)
	}
	builder.WriteString(")\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn err\n")
	builder.WriteString("\t}\n")
	fmt.Fprintf(builder, "\ttype alias %s\n", unionStruct.TypeName)
	builder.WriteString("\tvar temp alias\n")
	builder.WriteString("\tif err := json.Unmarshal(base, &temp); err != nil {\n")
	builder.WriteString("\t\treturn err\n")
	builder.WriteString("\t}\n")
	fmt.Fprintf(builder, "\t*value = %s(temp)\n", unionStruct.TypeName)
	for _, field := range unionStruct.Fields {
		fmt.Fprintf(builder, "\tif raw, ok := union[%q]; ok {\n", field.JSONName)
		decodeName := "decode" + field.RootName
		if field.IsSlice {
			decodeName += "Slice"
		}
		fmt.Fprintf(builder, "\t\tdecoded, err := %s(raw)\n", decodeName)
		builder.WriteString("\t\tif err != nil {\n")
		builder.WriteString("\t\t\treturn err\n")
		builder.WriteString("\t\t}\n")
		fmt.Fprintf(builder, "\t\tvalue.%s = decoded\n", field.GoName)
		builder.WriteString("\t}\n")
	}
	builder.WriteString("\treturn nil\n")
	builder.WriteString("}\n\n")
}

func findField(fields []fieldSpec, fieldName string) (fieldSpec, bool) {
	for _, field := range fields {
		if field.Name == fieldName {
			return field, true
		}
	}
	return fieldSpec{}, false
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
