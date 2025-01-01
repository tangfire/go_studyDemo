package main

import (
	"fmt"
	"go_studyDemo/go_study/version"
)

var age = 12

//age1 := 23 全局变量需要用var声明

func hello() {
	fmt.Println(age)
}

var (
	s1 string = "hello"
	s2 string = "world"
)

func main() {
	// 先声明
	var name string
	// 赋值
	name = "fireshine"
	fmt.Println(name)

	// 声明并赋值
	var name1 string = "fireshine"
	fmt.Println(name1)

	// 省略类型
	var name2 = "fireshine"
	fmt.Println(name2)

	// 声明并赋值 短声明符号
	name3 := "fireshine"
	fmt.Println(name3)

	hello()

	// 定义多个变量
	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

	fmt.Println(s1, s2)

	// 定义常量
	const Version = "2.0.1"

	fmt.Println(Version)
	fmt.Println(version.Version)

}
