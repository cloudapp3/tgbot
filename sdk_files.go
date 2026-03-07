package tgbot

import (
	"io"
	"path/filepath"
	"strings"
)

// InputFile represents Telegram file inputs: file_id, URL, local file path, or in-memory reader.
type InputFile struct {
	FileID   string
	URL      string
	FilePath string
	Reader   io.Reader
	FileName string
}

// FileFromID creates an InputFile that references an existing Telegram file_id.
func FileFromID(fileID string) InputFile {
	return InputFile{FileID: strings.TrimSpace(fileID)}
}

// FileFromURL creates an InputFile that references a public URL.
func FileFromURL(url string) InputFile {
	return InputFile{URL: strings.TrimSpace(url)}
}

// FileFromPath creates an InputFile that uploads a local file path.
func FileFromPath(path string) InputFile {
	trimmed := strings.TrimSpace(path)
	return InputFile{FilePath: trimmed, FileName: filepath.Base(trimmed)}
}

// FileFromReader creates an InputFile that uploads from a reader.
func FileFromReader(name string, reader io.Reader) InputFile {
	return InputFile{FileName: strings.TrimSpace(name), Reader: reader}
}

func (file InputFile) normalizedFileName(defaultName string) string {
	name := strings.TrimSpace(file.FileName)
	if name != "" {
		return name
	}
	if strings.TrimSpace(file.FilePath) != "" {
		return filepath.Base(strings.TrimSpace(file.FilePath))
	}
	if strings.TrimSpace(defaultName) != "" {
		return defaultName
	}
	return "upload.bin"
}

func (file InputFile) sourceCount() int {
	count := 0
	if strings.TrimSpace(file.FileID) != "" {
		count++
	}
	if strings.TrimSpace(file.URL) != "" {
		count++
	}
	if strings.TrimSpace(file.FilePath) != "" {
		count++
	}
	if file.Reader != nil {
		count++
	}
	return count
}

func (file InputFile) isEmpty() bool {
	return file.sourceCount() == 0
}
