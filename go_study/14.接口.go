package main

import "fmt"

type SingInterface interface {
	GetName() string
	Sing()
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println(c.Name, "在唱歌")
}

type Dog struct {
	Name string
}

func (d Dog) Sing() {
	fmt.Println(d.Name, "在唱歌")
}

func sing(c SingInterface) {
	//v, ok := c.(Cat)
	//fmt.Println(v, ok)
	switch v := c.(type) {
	case Cat:

		fmt.Println(v)
	case Dog:
		fmt.Println(v)
	default:
		fmt.Println("其他")
	}
	c.Sing()
	fmt.Println(c.GetName())

}

func (c Cat) GetName() string {
	return c.Name
}

func (d Dog) GetName() string {
	return d.Name
}

// interface{}的别名为any
func printAll(val interface{}) {
	fmt.Println(val)
}

func main() {
	cat := Cat{Name: "miao"}
	sing(cat)

	dog := Dog{Name: "wang"}
	sing(dog)

	printAll(cat)
	printAll(dog)

}
