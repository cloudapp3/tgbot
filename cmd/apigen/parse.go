package main

import (
	"fmt"
	"html"
	"regexp"
	"sort"
	"strings"
)

type fieldSpec struct {
	Name        string
	TypeExpr    string
	Required    bool
	Description string
}

type typeSpec struct {
	Name        string
	Anchor      string
	Description string
	BodyText    string
	Fields      []fieldSpec
}

type methodSpec struct {
	Name        string
	Anchor      string
	Description string
	BodyText    string
	Params      []fieldSpec
	ReturnsExpr string
}

type schema struct {
	Version     string
	VersionDate string
	Types       map[string]typeSpec
	Methods     map[string]methodSpec
	UnionRoots  map[string][]string
}

type docSection struct {
	Group  string
	Title  string
	Anchor string
	Body   string
}

var (
	headingPattern   = regexp.MustCompile(`(?s)<h([34])><a class="anchor" name="([^"]+)" href="#[^"]+"><i class="anchor-icon"></i></a>(.*?)</h[34]>`)
	tablePattern     = regexp.MustCompile(`(?s)<table class="table">(.*?)</table>`)
	thPattern        = regexp.MustCompile(`(?s)<th>(.*?)</th>`)
	trPattern        = regexp.MustCompile(`(?s)<tr>(.*?)</tr>`)
	cellPattern      = regexp.MustCompile(`(?s)<t[dh]>(.*?)</t[dh]>`)
	tagPattern       = regexp.MustCompile(`(?s)<[^>]+>`)
	typeTitlePattern = regexp.MustCompile(`^[A-Z][A-Za-z0-9]+$`)
	methodTitlePat   = regexp.MustCompile(`^[a-z][A-Za-z0-9]+$`)
	versionPattern   = regexp.MustCompile(`(?s)<h4><a class="anchor" name="[^"]+" href="#[^"]+"><i class="anchor-icon"></i></a>([^<]+)</h4>\s*<p><strong>(Bot API [^<]+)</strong></p>`)
)

func parseSchema(htmlData []byte) (*schema, error) {
	htmlText := string(htmlData)
	version, versionDate := extractLatestVersion(htmlText)
	sections := parseSections(htmlText)
	if len(sections) == 0 {
		return nil, fmt.Errorf("parse docs headings: no sections found")
	}

	types := map[string]typeSpec{}
	methods := map[string]methodSpec{}

	for _, section := range sections {
		headers, rows := extractTable(section.Body)
		bodyText := cleanText(section.Body)
		description := extractDescription(section.Body)

		if typeTitlePattern.MatchString(section.Title) && isTypeSection(headers, bodyText, section.Title) {
			fields := make([]fieldSpec, 0, len(rows))
			if len(headers) >= 3 && headers[0] == "Field" {
				for _, row := range rows {
					if len(row) < 3 {
						continue
					}
					fields = append(fields, fieldSpec{
						Name:        row[0],
						TypeExpr:    row[1],
						Required:    !strings.HasPrefix(row[2], "Optional."),
						Description: row[2],
					})
				}
			}
			types[section.Title] = typeSpec{
				Name:        section.Title,
				Anchor:      section.Anchor,
				Description: description,
				BodyText:    bodyText,
				Fields:      fields,
			}
			continue
		}

		if methodTitlePat.MatchString(section.Title) && isMethodSection(headers, bodyText) {
			params := make([]fieldSpec, 0, len(rows))
			if len(headers) >= 4 && headers[0] == "Parameter" {
				for _, row := range rows {
					if len(row) < 4 {
						continue
					}
					params = append(params, fieldSpec{
						Name:        row[0],
						TypeExpr:    row[1],
						Required:    strings.EqualFold(row[2], "Yes"),
						Description: row[3],
					})
				}
			}

			returnsExpr, err := extractReturnsExpr(section.Body, bodyText)
			if err != nil {
				return nil, fmt.Errorf("parse return type for %s: %w", section.Title, err)
			}
			methods[section.Title] = methodSpec{
				Name:        section.Title,
				Anchor:      section.Anchor,
				Description: description,
				BodyText:    bodyText,
				Params:      params,
				ReturnsExpr: returnsExpr,
			}
		}
	}

	if len(types) == 0 || len(methods) == 0 {
		return nil, fmt.Errorf("parse schema: got %d types and %d methods", len(types), len(methods))
	}

	unionRoots := detectUnionRoots(types)
	return &schema{
		Version:     version,
		VersionDate: versionDate,
		Types:       types,
		Methods:     methods,
		UnionRoots:  unionRoots,
	}, nil
}

func extractLatestVersion(htmlText string) (string, string) {
	match := versionPattern.FindStringSubmatch(htmlText)
	if len(match) != 3 {
		return "unknown", "unknown"
	}
	return cleanText(match[2]), cleanText(match[1])
}

func parseSections(htmlText string) []docSection {
	matches := headingPattern.FindAllStringSubmatchIndex(htmlText, -1)
	sections := make([]docSection, 0, len(matches))
	currentGroup := ""
	for index, match := range matches {
		level := htmlText[match[2]:match[3]]
		anchor := htmlText[match[4]:match[5]]
		title := cleanText(htmlText[match[6]:match[7]])
		bodyStart := match[1]
		bodyEnd := len(htmlText)
		if index+1 < len(matches) {
			bodyEnd = matches[index+1][0]
		}
		body := htmlText[bodyStart:bodyEnd]
		if level == "3" {
			currentGroup = title
			continue
		}
		sections = append(sections, docSection{
			Group:  currentGroup,
			Title:  title,
			Anchor: anchor,
			Body:   body,
		})
	}
	return sections
}

func extractTable(body string) ([]string, [][]string) {
	match := tablePattern.FindStringSubmatch(body)
	if len(match) < 2 {
		return nil, nil
	}
	tableHTML := match[1]
	headersHTML := thPattern.FindAllStringSubmatch(tableHTML, -1)
	headers := make([]string, 0, len(headersHTML))
	for _, header := range headersHTML {
		headers = append(headers, cleanText(header[1]))
	}

	rowMatches := trPattern.FindAllStringSubmatch(tableHTML, -1)
	rows := make([][]string, 0, len(rowMatches))
	for _, rowMatch := range rowMatches {
		cellsHTML := cellPattern.FindAllStringSubmatch(rowMatch[1], -1)
		if len(cellsHTML) == 0 {
			continue
		}
		row := make([]string, 0, len(cellsHTML))
		for _, cell := range cellsHTML {
			row = append(row, cleanText(cell[1]))
		}
		if len(row) == len(headers) && !equalStrings(row, headers) {
			rows = append(rows, row)
		}
	}
	return headers, rows
}

func cleanText(value string) string {
	value = strings.ReplaceAll(value, "<br>", " ")
	value = strings.ReplaceAll(value, "<br/>", " ")
	value = strings.ReplaceAll(value, "<br />", " ")
	value = tagPattern.ReplaceAllString(value, " ")
	value = html.UnescapeString(value)
	return strings.Join(strings.Fields(value), " ")
}

func extractDescription(body string) string {
	paragraphStart := strings.Index(body, "<p>")
	paragraphEnd := strings.Index(body, "</p>")
	if paragraphStart >= 0 && paragraphEnd > paragraphStart {
		return cleanText(body[paragraphStart+3 : paragraphEnd])
	}
	return cleanText(body)
}

func isTypeSection(headers []string, bodyText string, title string) bool {
	return true
}

func isMethodSection(headers []string, bodyText string) bool {
	if len(headers) >= 1 && headers[0] == "Parameter" {
		return true
	}
	return strings.Contains(bodyText, "Requires no parameters.") || strings.Contains(bodyText, "Returns ") || strings.Contains(bodyText, "returns ")
}

func extractReturnsExpr(body string, bodyText string) (string, error) {
	returnText := extractDescription(body)
	if strings.Contains(returnText, "otherwise True is returned") || strings.Contains(returnText, "otherwise True is returned.") {
		return "any", nil
	}
	patterns := []string{
		`(?i)returns basic information .* form of an? ([A-Z][A-Za-z0-9]+) object`,
		`(?i)on success, the sent ([A-Z][A-Za-z0-9]+) is returned`,
		`(?i)on success, the [a-z]+ ([A-Z][A-Za-z0-9]+) is returned`,
		`(?i)on success, an? ([A-Z][A-Za-z0-9]+) object is returned`,
		`(?i)on success, an array of ([A-Z][A-Za-z0-9]+) objects? .* is returned`,
		`(?i)on success, an array of ([A-Z][A-Za-z0-9]+) .* is returned`,
		`(?i)returns the ([A-Z][A-Za-z0-9]+) of the sent message on success`,
		`(?i)returns the uploaded ([A-Z][A-Za-z0-9]+) on success`,
		`(?i)returns the created .* as ([A-Z][A-Za-z0-9]+) on success`,
		`(?i)returns .* as an? ([A-Z][A-Za-z0-9]+) object`,
		`(?i)returns the .* as an? ([A-Z][A-Za-z0-9]+) object`,
		`(?i)returns the .* as ([A-Z][A-Za-z0-9]+) object`,
		`(?i)returns the .* as ([A-Z][A-Za-z0-9]+) on success`,
		`(?i)on success, returns an? ((?:Array of )?[A-Z][A-Za-z0-9]+(?:, [A-Z][A-Za-z0-9]+)*(?: and [A-Z][A-Za-z0-9]+)?) object`,
		`(?i)returns an? ((?:Array of )?[A-Z][A-Za-z0-9]+(?:, [A-Z][A-Za-z0-9]+)*(?: and [A-Z][A-Za-z0-9]+)?) object`,
		`(?i)on success, returns ((?:Array of )?[A-Z][A-Za-z0-9]+(?:, [A-Z][A-Za-z0-9]+)*(?: and [A-Z][A-Za-z0-9]+)?)`,
		`(?i)returns ((?:Array of )?[A-Z][A-Za-z0-9]+(?:, [A-Z][A-Za-z0-9]+)*(?: and [A-Z][A-Za-z0-9]+)?) on success`,
		`(?i)returns ((?:Array of )?[A-Z][A-Za-z0-9]+(?:, [A-Z][A-Za-z0-9]+)*(?: and [A-Z][A-Za-z0-9]+)?)`,
	}
	for _, pattern := range patterns {
		compiled := regexp.MustCompile(pattern)
		match := compiled.FindStringSubmatch(returnText)
		if len(match) == 2 {
			value := strings.TrimSpace(match[1])
			if strings.HasPrefix(strings.ToLower(value), "array of ") {
				value = "Array of " + strings.TrimSpace(value[len("Array of "):])
			}
			if !looksLikeReturnType(value) {
				continue
			}
			return value, nil
		}
	}
	if strings.Contains(returnText, "Returns True on success") || strings.Contains(returnText, "returns True on success") || strings.Contains(returnText, "On success, True is returned") {
		return "True", nil
	}
	return "", fmt.Errorf("unsupported return description: %q", returnText)
}

func looksLikeReturnType(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return false
	}
	if strings.HasPrefix(trimmed, "Array of ") {
		inner := strings.TrimSpace(strings.TrimPrefix(trimmed, "Array of "))
		return looksLikeReturnType(inner)
	}
	switch trimmed {
	case "True", "String", "Integer", "Boolean", "Float":
		return true
	}
	first := trimmed[0]
	return first >= 'A' && first <= 'Z'
}

func detectUnionRoots(types map[string]typeSpec) map[string][]string {
	names := make([]string, 0, len(types))
	for name := range types {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		if len(names[i]) == len(names[j]) {
			return names[i] < names[j]
		}
		return len(names[i]) > len(names[j])
	})

	patternParts := make([]string, 0, len(names))
	for _, name := range names {
		patternParts = append(patternParts, regexp.QuoteMeta(name))
	}
	namePattern := regexp.MustCompile(`\b(?:` + strings.Join(patternParts, `|`) + `)\b`)
	markers := []string{
		"It can be one of",
		"Currently, it can be one of",
		"currently support the following",
		"currently support results of the following",
		"Currently, one of",
		"It should be one of",
		"should be one of",
		"Currently, the following",
	}

	unionRoots := map[string][]string{}
	for name, spec := range types {
		markerIndex := -1
		for _, marker := range markers {
			index := strings.Index(spec.BodyText, marker)
			if index >= 0 {
				markerIndex = index
				break
			}
		}
		if markerIndex < 0 {
			continue
		}

		tail := spec.BodyText[markerIndex:]
		stopTokens := []string{"Note:", "Field Type Description", "Parameter Type Required Description"}
		for _, token := range stopTokens {
			if index := strings.Index(tail, token); index >= 0 {
				tail = tail[:index]
			}
		}

		matches := namePattern.FindAllString(tail, -1)
		if len(matches) < 2 {
			continue
		}
		seen := map[string]bool{}
		subtypes := make([]string, 0, len(matches))
		for _, match := range matches {
			if match == name || seen[match] {
				continue
			}
			seen[match] = true
			subtypes = append(subtypes, match)
		}
		if len(subtypes) >= 2 {
			unionRoots[name] = subtypes
		}
	}
	return unionRoots
}

func equalStrings(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
