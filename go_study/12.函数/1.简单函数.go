package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

func param1(id string) {
	fmt.Println("id:", id)
}

func param2(id, userName string) {
	fmt.Println("id:", id)
	fmt.Println("userName:", userName)
}

func param3(numberList ...int) {
	var sum int
	for _, i2 := range numberList {
		sum += i2
	}

	fmt.Println("sum:", sum)
}

func r1() {
	return
}

func r2() bool {
	return false
}

func r3() (string, bool) {
	var ok bool
	return "hello", ok
}

func r4() (val bool) {
	val = false
	return
}

func main() {
	sayHello()
	param1("tangshine")
	param2("tangshine", "future")
	param3(1, 2, 3, 4, 5)

	// 匿名函数
	var getName = func() string {
		return "tangfire"
	}

	var setName = func(data string) {
		fmt.Println(data)
		return
	}

	fmt.Println(getName())

	setName("myl")
}
