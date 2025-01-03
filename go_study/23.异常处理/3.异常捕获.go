package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	var list []int = []int{1, 2}
	fmt.Println(list[2])

}

func main() {
	read()
	// 正常逻辑
	fmt.Println("normal logic")
}
