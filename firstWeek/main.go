package main

import (
	"fmt"
	"os"
	//"io"
	//"path/filepath"
	//"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	err := searchingNode(out, path, printFiles, "")
	return err
}

func searchingNode(out *os.File, path string, printFiles bool, prevIndent string) error {
	allFiles, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var files []os.DirEntry
	if printFiles {
		files = allFiles
	} else {
		for _, file := range allFiles {
			if file.IsDir() {
				files = append(files, file)
			}
		}
	}

	var (
		lastInd = len(files) - 1
		prefix  = "├───"
		indent  = "│\t"
	)

	for i, file := range files {
		if i == lastInd {
			prefix = "└───"
			indent = "\t"
		}
		fmt.Print(prevIndent, prefix, file.Name(), "\n")
		if file.IsDir() {
			err := searchingNode(out, path+string(os.PathSeparator)+file.Name(), printFiles, prevIndent+indent)
			if err != nil {
				return err
			}
		}

	}
	return err
}
