package main

import "fmt"

func main() {

	var index int
	fmt.Scan(&index)

	var funcMap = map[int]func(){
		1: login,
		2: register,
		3: userCenter,
	}

	fun, ok := funcMap[index]
	if ok {
		fun()
	}

}

func login() {
	fmt.Println("login")
}

func register() {
	fmt.Println("register")
}

func userCenter() {
	fmt.Println("userCenter")
}
