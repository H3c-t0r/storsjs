package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var failed int

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			failed++
			fmt.Printf("failed to read %v: %v\n", path, err)
			return nil
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		var header [1024]byte
		n, err := file.Read(header[:])
		if err != nil && err != io.EOF {
			fmt.Printf("failed to read %v: %v\n", path, err)
			return nil
		}

		if bytes.Contains(header[:n], []byte(`AUTOGENERATED`)) ||
			bytes.Contains(header[:n], []byte(`Code generated`)) {
			return nil
		}

		if !bytes.Contains(header[:n], []byte(`Copyright `)) {
			failed++
			fmt.Printf("missing copyright: %v\n", path)
		}

		return nil
	})

	if failed > 0 {
		os.Exit(1)
	}
}
