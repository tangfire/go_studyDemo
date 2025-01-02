package main

import "fmt"

type MyCode int
type MyAliasCode = int

func (m MyCode) Code() {
	fmt.Println(m)
}

const myCode MyCode = 1
const myAliasCode MyAliasCode = 1

func main() {
	fmt.Printf("%v,%T\n", myCode, myCode)
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode)
	var age = 1
	if myCode == MyCode(age) {
		fmt.Println("hello world")
	}

	if myAliasCode == age {
		fmt.Println("hello world")
	}

}
