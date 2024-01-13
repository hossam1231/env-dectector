package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyEnvFiles(src, destRoot string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".env" {
			// Determine relative path
			relPath, err := filepath.Rel(src, path)
			if err != nil {
				return err
			}

			// Create destination folder
			destFolder := filepath.Join(destRoot, relPath)
			if err := os.MkdirAll(destFolder, os.ModePerm); err != nil {
				return err
			}

			// Create destination file
			destPath := filepath.Join(destFolder, info.Name())
			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			// Open source file
			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			// Copy content
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}

			fmt.Printf("Copied %s to %s\n", path, destPath)
		}

		return nil
	})
}

func main() {
	srcPath := "../"
	destRoot := "../Environment_variables"

	err := copyEnvFiles(srcPath, destRoot)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
