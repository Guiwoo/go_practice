package chapter26

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func StartPrinter() {
	if len(os.Args) < 2 {
		panic("The argument is not enough to run program")
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}

	for _, path := range files {
		findInfos = append(findInfos, findWordInAllFiles(word, path)...)
	}
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("=================================")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("==================================")
		fmt.Println()
	}
}

func findWordInAllFiles(word, path string) []FindInfo {
	FindInfos := []FindInfo{}

	fileList, err := getFileList(path)
	if err != nil {
		fmt.Printf("The file path is wrong , err : ", err)
		return FindInfos
	}
	ch := make(chan FindInfo)
	cnt := len(fileList)
	recvCnt := 0

	for _, filename := range fileList {
		go FindWordInFile(word, filename, ch)
	}

	for findInfo := range ch {
		FindInfos = append(FindInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return FindInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("can not find the file by filename")
		ch <- findInfo
	}
	defer file.Close()

	lineNo := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}
