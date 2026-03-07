package main

import (
	"fmt"
	"go/format"
	"sort"
	"strings"
)

type generator struct {
	schema             *schema
	packageName        string
	unionParents       map[string][]string
	unionFieldOwners   []unionStructSpec
	unionDecoderSpecs  []unionDecoderSpec
	knownTypeNames     map[string]bool
	supportTypeNames   map[string]bool
	preferredDiscNames []string
}

type unionStructSpec struct {
	TypeName string
	Fields   []unionFieldSpec
}

type unionFieldSpec struct {
	JSONName string
	GoName   string
	RootName string
	IsSlice  bool
}

type unionDecoderSpec struct {
	RootName           string
	DiscriminatorField string
	Cases              []unionCase
	SpecialKind        string
}

type unionCase struct {
	Subtype string
	Tag     string
}

func newGenerator(schema *schema, packageName string) *generator {
	knownTypeNames := map[string]bool{}
	for name := range schema.Types {
		knownTypeNames[name] = true
	}
	supportTypeNames := map[string]bool{
		"InputFile":          true,
		"ResponseParameters": true,
	}
	unionParents := map[string][]string{}
	for root, subtypes := range schema.UnionRoots {
		for _, subtype := range subtypes {
			unionParents[subtype] = append(unionParents[subtype], root)
		}
	}

	g := &generator{
		schema:             schema,
		packageName:        packageName,
		unionParents:       unionParents,
		knownTypeNames:     knownTypeNames,
		supportTypeNames:   supportTypeNames,
		preferredDiscNames: []string{"type", "status", "source"},
	}
	g.unionFieldOwners = g.collectUnionStructs()
	g.unionDecoderSpecs = g.collectUnionDecoderSpecs()
	return g
}

func (g *generator) generateTypesFile() ([]byte, error) {
	var builder strings.Builder
	builder.WriteString("// Telegram Bot API type definitions aligned with the official docs.\n")
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

	typeNames := make([]string, 0, len(g.schema.Types))
	for name := range g.schema.Types {
		if g.supportTypeNames[name] {
			continue
		}
		typeNames = append(typeNames, name)
	}
	sort.Strings(typeNames)

	for _, typeName := range typeNames {
		spec := g.schema.Types[typeName]
		if _, isUnionRoot := g.schema.UnionRoots[typeName]; isUnionRoot {
			fmt.Fprintf(&builder, "// %s is a union type in Telegram Bot API.\n", typeName)
			fmt.Fprintf(&builder, "type %s interface {\n", typeName)
			fmt.Fprintf(&builder, "\tis%s()\n", typeName)
			builder.WriteString("}\n\n")
			continue
		}

		fmt.Fprintf(&builder, "// %s maps to Telegram Bot API type %q.\n", typeName, typeName)
		fmt.Fprintf(&builder, "type %s struct {\n", typeName)
		for _, field := range spec.Fields {
			goType, err := g.mapFieldType(field.TypeExpr)
			if err != nil {
				return nil, fmt.Errorf("map type field %s.%s: %w", typeName, field.Name, err)
			}
			tag := fmt.Sprintf("`json:\"%s\"`", field.Name)
			if !field.Required {
				tag = fmt.Sprintf("`json:\"%s,omitempty\"`", field.Name)
			}
			fmt.Fprintf(&builder, "\t%s %s %s\n", toPascal(field.Name), goType, tag)
		}
		builder.WriteString("}\n\n")

		for _, parent := range sortedStrings(g.unionParents[typeName]) {
			fmt.Fprintf(&builder, "func (*%s) is%s() {}\n\n", typeName, parent)
		}
	}

	return format.Source([]byte(builder.String()))
}

func (g *generator) generateMethodsFile() ([]byte, error) {
	var builder strings.Builder
	builder.WriteString("// Telegram Bot API method wrappers aligned with the official docs.\n")
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
	builder.WriteString("import \"context\"\n\n")

	methodNames := make([]string, 0, len(g.schema.Methods))
	for name := range g.schema.Methods {
		methodNames = append(methodNames, name)
	}
	sort.Strings(methodNames)

	for _, methodName := range methodNames {
		method := g.schema.Methods[methodName]
		goName := upperFirst(methodName)
		paramsName := goName + "Params"
		fmt.Fprintf(&builder, "// %s contains params for Telegram method %q.\n", paramsName, methodName)
		fmt.Fprintf(&builder, "type %s struct {\n", paramsName)
		for _, field := range method.Params {
			goType, err := g.mapValueType(field.TypeExpr)
			if err != nil {
				return nil, fmt.Errorf("map method param %s.%s: %w", methodName, field.Name, err)
			}
			tag := fmt.Sprintf("`json:\"%s\"`", field.Name)
			if !field.Required {
				tag = fmt.Sprintf("`json:\"%s,omitempty\"`", field.Name)
			}
			fmt.Fprintf(&builder, "\t%s %s %s\n", toPascal(field.Name), goType, tag)
		}
		builder.WriteString("}\n\n")

		resultType, err := g.mapValueType(method.ReturnsExpr)
		if err != nil {
			return nil, fmt.Errorf("map method return %s: %w", methodName, err)
		}
		fmt.Fprintf(&builder, "// %s calls Telegram method %q.\n", goName, methodName)
		fmt.Fprintf(&builder, "// Doc: %s#%s\n", officialBotAPIURL, method.Anchor)
		fmt.Fprintf(&builder, "func (bot *Bot) %s(ctx context.Context, params *%s) (%s, error) {\n", goName, paramsName, resultType)
		fmt.Fprintf(&builder, "\tvar result %s\n", resultType)
		fmt.Fprintf(&builder, "\tif err := bot.call(ctx, %q, params, &result); err != nil {\n", methodName)
		fmt.Fprintf(&builder, "\t\treturn %s, err\n", g.zeroValue(resultType))
		builder.WriteString("\t}\n")
		builder.WriteString("\treturn result, nil\n")
		builder.WriteString("}\n\n")
	}

	return format.Source([]byte(builder.String()))
}

func (g *generator) mapFieldType(typeExpr string) (string, error) {
	goType, err := g.mapValueType(typeExpr)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(goType, "[]") || isPrimitiveGoType(goType) || goType == "any" || g.isUnionRoot(goType) || g.supportTypeNames[goType] {
		return goType, nil
	}
	return "*" + goType, nil
}

func (g *generator) mapValueType(typeExpr string) (string, error) {
	normalized := strings.TrimSpace(typeExpr)
	if normalized == "" {
		return "", fmt.Errorf("empty type expression")
	}
	if strings.HasPrefix(normalized, "Array of ") {
		inner := strings.TrimPrefix(normalized, "Array of ")
		mapped, err := g.mapValueType(inner)
		if err != nil {
			return "", err
		}
		return "[]" + mapped, nil
	}
	if primitive := mapPrimitiveType(normalized); primitive != "" {
		return primitive, nil
	}
	if normalized == "any" {
		return "any", nil
	}
	if g.supportTypeNames[normalized] {
		return normalized, nil
	}
	if g.isUnionRoot(normalized) || g.knownTypeNames[normalized] {
		return normalized, nil
	}
	if tokens := splitAlternatives(normalized); len(tokens) > 1 {
		if rootName := g.commonUnionRoot(tokens); rootName != "" {
			return rootName, nil
		}
		return "any", nil
	}
	return "", fmt.Errorf("unsupported type expression %q", typeExpr)
}

func (g *generator) commonUnionRoot(tokens []string) string {
	rootName := ""
	for _, token := range tokens {
		if primitive := mapPrimitiveType(token); primitive != "" {
			return ""
		}
		parents := g.unionParents[token]
		if len(parents) == 0 {
			return ""
		}
		if rootName == "" {
			rootName = parents[0]
			continue
		}
		if rootName != parents[0] {
			return ""
		}
	}
	return rootName
}

func (g *generator) zeroValue(goType string) string {
	switch goType {
	case "bool":
		return "false"
	case "string":
		return `""`
	case "int64", "float64":
		return "0"
	case "any":
		return "nil"
	}
	if strings.HasPrefix(goType, "[]") {
		return "nil"
	}
	if g.isUnionRoot(goType) {
		return "nil"
	}
	return goType + "{}"
}

func (g *generator) isUnionRoot(name string) bool {
	_, ok := g.schema.UnionRoots[name]
	return ok
}

func mapPrimitiveType(typeExpr string) string {
	switch strings.TrimSpace(typeExpr) {
	case "String":
		return "string"
	case "Integer", "Int":
		return "int64"
	case "Float":
		return "float64"
	case "Boolean", "True":
		return "bool"
	case "Integer or String":
		return "any"
	case "InputFile or String":
		return "any"
	}
	return ""
}

func splitAlternatives(typeExpr string) []string {
	replacer := strings.NewReplacer(
		", and ", ", ",
		" and ", ", ",
		" or ", ", ",
	)
	normalized := replacer.Replace(typeExpr)
	parts := strings.Split(normalized, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func upperFirst(value string) string {
	if value == "" {
		return ""
	}
	return strings.ToUpper(value[:1]) + value[1:]
}

func toPascal(value string) string {
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})
	initialisms := map[string]string{
		"api":  "API",
		"bio":  "Bio",
		"faq":  "FAQ",
		"id":   "ID",
		"ip":   "IP",
		"url":  "URL",
		"uuid": "UUID",
	}
	var builder strings.Builder
	for _, part := range parts {
		lower := strings.ToLower(part)
		if replacement, ok := initialisms[lower]; ok {
			builder.WriteString(replacement)
			continue
		}
		if part == "" {
			continue
		}
		builder.WriteString(strings.ToUpper(part[:1]))
		if len(part) > 1 {
			builder.WriteString(part[1:])
		}
	}
	return builder.String()
}

func sortedStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	clone := append([]string(nil), values...)
	sort.Strings(clone)
	return clone
}

func isPrimitiveGoType(goType string) bool {
	switch goType {
	case "bool", "string", "int64", "float64":
		return true
	default:
		return false
	}
}
