package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		println("No files have been specified")
		return
	}

	for _, file := range files {
		// Glob patterns
		if strings.Contains(file, "*") {
			matches, err := filepath.Glob(file)

			if err != nil {
				panic(err)
			}

			for _, match := range matches {
				encode(match, match+".webm")
			}

			continue
		}

		// Single file
		encode(file, file+".webm")
	}
}
