package main

import "fmt"

type SingInterface interface {
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
	c.Sing()
}

func main() {
	cat := Cat{Name: "miao"}
	sing(cat)

	dog := Dog{Name: "wang"}
	sing(dog)

}
