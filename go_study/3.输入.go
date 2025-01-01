package main

import "fmt"

func main() {
	fmt.Print("please input string: ")

	var name string

	fmt.Scan(&name)

	fmt.Println(name)

	fmt.Printf("%T\n", name)

	fmt.Print("please input your age:")
	var age int
	n, err := fmt.Scan(&age)

	fmt.Println(n, err, age)

}
