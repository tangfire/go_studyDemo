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

}
