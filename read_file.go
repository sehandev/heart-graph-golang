package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkOpenFile(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() {
	fileName := "data/data.txt"
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	checkOpenFile(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		}
		fmt.Println(string(line))
	}
}
