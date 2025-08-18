package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

func Scan(path string) {
	allowedExts := map[string]bool{
		".go":  true,
		".cs":  true,
		".py":  true,
		".js":  true,
		".ts":  true,
		".jsx": true,
		".tsx": true,
	}
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		cobra.CheckErr(fmt.Errorf("path does not exist: %s", path))
	}
	if err != nil {
		cobra.CheckErr(err)
	}

	if !fileInfo.IsDir() {
		ext := filepath.Ext(path)
		if allowedExts[ext] {
			fmt.Println("Processing file:", path)
		}
		return
	}

	if fileInfo.IsDir() {
		jobs := make(chan string, 10)
		var wg sync.WaitGroup
		numWorkers := runtime.NumCPU() * 3

		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go workerPool(&wg, jobs)
		}

		err := filepath.WalkDir(path, func(filePath string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() && (entry.Name() == "node_modules" || strings.HasPrefix(entry.Name(), ".")) {
				return filepath.SkipDir
			}
			ext := filepath.Ext(filePath)
			if allowedExts[ext] {
				jobs <- filePath
			}
			return nil
		})
		cobra.CheckErr(err)

		close(jobs)
		wg.Wait()
	}
}
