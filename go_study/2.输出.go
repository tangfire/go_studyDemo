package main

import "fmt"

func main() {
	fmt.Println("hello", "world")
	fmt.Println("hello", "fireshine")

	fmt.Printf("%s 哇，你好美!\n", "fireshine")

	fmt.Printf("%d\n", 3)

	fmt.Printf("%.2f\n", 4.4321421)

	fmt.Printf("%T %T\n", "hello", 2.5)

	fmt.Printf("%v\n", "f")

	fmt.Printf("%v\b", 3)

	fmt.Printf("%v\n", "")

	fmt.Printf("%#v\n", "")

	var f = fmt.Sprintf("%.2f", 4.2131241)

	fmt.Println(f)

}
