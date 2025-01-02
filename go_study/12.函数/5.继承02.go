package main

import "fmt"

type UserInfo struct {
	Name string `json:"name"`
}

func (u *UserInfo) SetName(name string) {
	u.Name = name

}

func main() {
	user := UserInfo{
		Name: "fireshine",
	}

	user.SetName("张三")
	fmt.Println(user.Name)

}
