package blog

import (
	"github.com/yuin/goldmark"
	"os"
	"path/filepath"
	"strings"
)

func Convert(inputFolder, outputFolder string) error {
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		return err
	}

	return filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // Skip directories
		}
		if strings.HasSuffix(info.Name(), ".md") {
			markdown, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var html strings.Builder
			if err := goldmark.Convert(markdown, &html); err != nil {
				return err
			}

			outputPath := filepath.Join(outputFolder, strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))+".html")
			return os.WriteFile(outputPath, []byte(html.String()), info.Mode())
		}
		return nil
	})
}
