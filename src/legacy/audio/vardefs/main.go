package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename> <function_name>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s audio_impl.go Sub_4523D0\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filename := os.Args[1]
	functionName := os.Args[2]

	// Make filename absolute if it's relative
	if !filepath.IsAbs(filename) {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
			os.Exit(1)
		}
		filename = filepath.Join(wd, filename)
	}

	varDecls, err := EnumerateVarDecls(filename, functionName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(varDecls) == 0 {
		fmt.Fprintf(os.Stderr, "No variable declarations found in function %s\n", functionName)
		os.Exit(1)
	}

	// Print variable declarations
	for _, varDecl := range varDecls {
		fmt.Printf("%s %s\n", varDecl.Name, varDecl.Type)
	}
}