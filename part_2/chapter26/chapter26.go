package chapter26

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Start() {
	a := os.Args
	if len(a) < 2 {
		panic("Should write a file path before running go")
	}
	word := a[1]
	files := a[2:]
	fmt.Println("Looking for word is : ", word)
	printAllFiles(files)
}

func getFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func printAllFiles(files []string) {
	for _, path := range files {
		filelist, err := getFileList(path)
		if err != nil {
			fmt.Println("Wrong File Path err : ", err)
			return
		}
		fmt.Println("Looking for file list")
		for _, file := range filelist {
			printFile(file)
		}
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't no find file by filename")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
