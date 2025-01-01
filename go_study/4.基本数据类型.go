package main

import "fmt"

func main() {
	var age = 12
	var u8 uint8 = 255
	fmt.Println(age, u8)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", u8)

	var a byte = 'a' // ascil里面的字符
	fmt.Printf("%c %d\n", a, a)

	// 其实byte就是uint8
	var a1 uint8 = 97
	fmt.Printf("%c %d\n", a1, a1)

	var z rune = '中' // unicode码
	fmt.Printf("%c %d\n", z, z)

	// 其实rune就是int32
	var z1 int32 = 20013
	fmt.Printf("%c %d\n", z1, z1)

	fmt.Println("枫枫\t知道")              // 制表符
	fmt.Println("枫枫\n知道")              // 回车
	fmt.Println("\"枫枫\"知道")            // 双引号
	fmt.Println("枫枫\r知道")              // 回到行首
	fmt.Println("C:\\pprof\\main.exe") // 反斜杠

	fmt.Println(`hello \n 
world
`)

	var age_ int
	var b bool
	var s string
	//fmt.Println(age_, b, s)
	fmt.Printf("%#v\n", age_)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", s)
}
