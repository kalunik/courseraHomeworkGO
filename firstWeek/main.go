package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

//If there is no need to print files, it'll wipe out non directories entries. Also it sort fir entries
func contentFilter(nonFiltered []os.DirEntry, printFiles bool) []os.DirEntry {
	sort.Slice(nonFiltered, func(i, j int) bool {
		return nonFiltered[i].Name() < nonFiltered[j].Name()
	})
	if printFiles {
		return nonFiltered
	} else {
		var dirOnly []os.DirEntry
		for _, obj := range nonFiltered {
			if obj.IsDir() {
				dirOnly = append(dirOnly, obj)
			}
		}
		return dirOnly
	}
}

func sizePrinter(out io.Writer, file os.DirEntry) error {
	fi, _ := file.Info()

	if !file.IsDir() {
		if fi.Size() == 0 {
			_, err := fmt.Fprint(out, " (empty)")
			if err != nil {
				return fmt.Errorf("can't print file size")
			}
		} else {
			_, err := fmt.Fprint(out, " (", fi.Size(), "b)")
			if err != nil {
				return fmt.Errorf("can't print file size")
			}
		}
	}
	return nil
}

func searchingNode(out io.Writer, path string, printFiles bool, prevIndent string) error {
	allFiles, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("can't read dir")
	}

	files := contentFilter(allFiles, printFiles)

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
		_, err := fmt.Fprint(out, prevIndent, prefix, file.Name())
		if err != nil {
			return fmt.Errorf("can't print names")
		}

		err = sizePrinter(out, file)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(out, "\n")
		if err != nil {
			return fmt.Errorf("can't print")
		}

		if file.IsDir() {

			err := searchingNode(out, path+string(os.PathSeparator)+file.Name(), printFiles, prevIndent+indent)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	err := searchingNode(out, path, printFiles, "")
	return err
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run signer.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
