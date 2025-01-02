package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
}

func main() {
	user := User{Name: "fireshine", Age: 20, Password: "123456"}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData)) // 转为json
}
