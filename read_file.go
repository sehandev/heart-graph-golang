package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkOpenFile(err error) {
	// file이 올바르게 열렸는지 확인
	if err != nil {
		log.Fatalf("open file error: %v", err)
		panic(err)
	}
}

func checkScanFile(err error) {
	// EOF까지 올바르게 scan했는지 확인
	if err != nil {
		log.Fatalf("scan file error: %v", err)
		panic(err)
	}
}

func readFile(fileName string) {
	// beginTime := time.Now()

	fp, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	checkOpenFile(err)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() { // 한 줄 스캔, EOF -> 반복 종료
		line := scanner.Text() // 스캔한 문장을 string으로 변환

		if strings.Contains(line, "사랑") {
			// 문장에 단어 "사랑"이 포함되어 있으면
			fmt.Println(line)
		}
	}

	checkScanFile(scanner.Err())

	// fmt.Println("scan file, time :", time.Since(beginTime).Seconds())
}
