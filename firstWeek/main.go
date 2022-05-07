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
	var lvl = 0
	err := searchingNode(out, path, printFiles, lvl)
	return err
}

func searchingNode(out *os.File, path string, printFiles bool, lvl int) error {
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
	//fmt.Println(len(files), "!!")

	for i, file := range files {
		for j := 0; j < lvl; j++ {
			if i != len(files)-1 {
				fmt.Print("│")
			}
			fmt.Print("\t")
		}
		if i == len(files)-1 {
			fmt.Print("└───")
		} else {
			fmt.Print("├───")
		}
		fmt.Print(file.Name(), " ", i, "\n")
		//fmt.Print("│")
		if file.IsDir() {
			//mb | and \t here, it could be better
			err := searchingNode(out, path+string(os.PathSeparator)+file.Name(), printFiles, lvl+1)
			if err != nil {
				return err
			}
		}
		//if files
		//fmt.Print("└───")
	}
	return err
}
