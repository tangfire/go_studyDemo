package main

import "fmt"

// 组合
type Student struct {
	Class
	Name string
}

type Class struct {
	Name string
}

func (s Student) study() {
	fmt.Printf("%s 正在学习\n", s.Name)
}

func (s Student) Info() {
	fmt.Printf("名字: %s 班级: %s\n", s.Name, s.Class.Name)
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func main() {

	c1 := Class{
		Name: "三年一班",
	}

	s1 := Student{Name: "fire", Class: c1}
	s1.Info()
	s1.study()

	s2 := Student{Name: "tom", Class: c1}
	s2.study()
	s2.Info()
	s2.SetName("jack")
	s2.Info()

}
