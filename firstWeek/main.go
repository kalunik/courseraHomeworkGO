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
	var err error

	err = os.Chdir(path)
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {

		if file.IsDir() {
			fmt.Print(file.Name(), "\n")
			err = dirTree(out, file.Name(), printFiles)
			if err != nil {
				return err
			}
			err = os.Chdir("..")
			if err != nil {
				return err
			}
		} else {
			continue
		}
	}
	return err
}
