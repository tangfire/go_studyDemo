package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"fireshine", "myl", "future"}

	fmt.Println(nameList)
	fmt.Println(nameList[2])
	fmt.Println(nameList[len(nameList)-1])
}
