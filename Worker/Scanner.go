package scanner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func Scan(path string) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		cobra.CheckErr(fmt.Errorf("path does not exist: %s", path))
	}
	if err != nil {
		cobra.CheckErr(err)
	}

	if fileInfo.IsDir() {
		fmt.Println("Scanning directory:", path)

		err := filepath.WalkDir(path, func(filePath string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() {
				return nil
			}
			analyzeFile(filePath)
			return nil
		})
		cobra.CheckErr(err)

	} else {
		analyzeFile(path)
	}
}

func analyzeFile(path string) {
	ext := filepath.Ext(path)
	switch ext {
	case ".go":
		fmt.Println("run Go analysis on file:", path)
	case ".py":
		fmt.Println("run Python analysis on file:", path)
	case ".js", ".ts":
		fmt.Println("run JavaScript/TypeScript analysis on file:", path)
	default:
		fmt.Printf("Skipping unsupported file type: %s\n", path)
	}
}
