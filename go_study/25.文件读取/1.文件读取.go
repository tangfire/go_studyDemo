package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//byteData, err := os.ReadFile("D:\\GitRespority\\go_studyDemo\\go_study\\25.文件读取\\hello.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(string(byteData))

	file, err := os.Open("go_study/25.文件读取/hello.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//for {
	//	var byteData = make([]byte, 512)
	//	n, err := file.Read(byteData)
	//	if err == io.EOF {
	//		break
	//	}
	//
	//	fmt.Println(string(byteData), n)
	//
	//}

	// 按行读
	//buf := bufio.NewReader(file)
	//for {
	//	line, _, err := buf.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(line), err)
	//}

	buf := bufio.NewScanner(file)
	buf.Split(bufio.ScanWords)
	var index int
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())
	}

}
