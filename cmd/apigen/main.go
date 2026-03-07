package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const officialBotAPIURL = "https://core.telegram.org/bots/api"

func main() {
	var (
		htmlPath    = flag.String("html", "", "path to a downloaded Telegram Bot API HTML page")
		url         = flag.String("url", officialBotAPIURL, "Telegram Bot API documentation URL")
		outDir      = flag.String("out", ".", "output directory for generated SDK files")
		packageName = flag.String("package", "tgbot", "package name for generated files")
	)
	flag.Parse()

	htmlData, err := loadHTML(*htmlPath, *url)
	if err != nil {
		fail(err)
	}

	schema, err := parseSchema(htmlData)
	if err != nil {
		fail(err)
	}

	generator := newGenerator(schema, *packageName)
	typesFile, err := generator.generateTypesFile()
	if err != nil {
		fail(err)
	}
	methodsFile, err := generator.generateMethodsFile()
	if err != nil {
		fail(err)
	}
	unionFile, err := generator.generateUnionFile()
	if err != nil {
		fail(err)
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fail(fmt.Errorf("create output dir: %w", err))
	}

	typesPath := filepath.Join(*outDir, "sdk_types.go")
	methodsPath := filepath.Join(*outDir, "sdk_methods.go")
	unionPath := filepath.Join(*outDir, "sdk_unions.go")
	if err := os.WriteFile(typesPath, typesFile, 0o644); err != nil {
		fail(fmt.Errorf("write %s: %w", typesPath, err))
	}
	if err := os.WriteFile(methodsPath, methodsFile, 0o644); err != nil {
		fail(fmt.Errorf("write %s: %w", methodsPath, err))
	}
	if err := os.WriteFile(unionPath, unionFile, 0o644); err != nil {
		fail(fmt.Errorf("write %s: %w", unionPath, err))
	}

	fmt.Printf("generated %s, %s and %s\n", typesPath, methodsPath, unionPath)
}

func loadHTML(htmlPath string, url string) ([]byte, error) {
	if htmlPath != "" {
		return os.ReadFile(htmlPath)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url) //nolint:gosec // generator utility fetches a public docs page.
	if err != nil {
		return nil, fmt.Errorf("download docs: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("download docs: unexpected status %s", resp.Status)
	}
	return ioReadAll(resp.Body)
}

func fail(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "apigen: %v\n", err)
	os.Exit(1)
}
