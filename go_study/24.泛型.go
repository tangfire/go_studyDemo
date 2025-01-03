package main

import (
	"encoding/json"
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func plus[T Number](n1, n2 T) T {
	return n1 + n2
}

func myPrint[T int, K string | int](n1 T, n2 K) {
	fmt.Println(n1, n2)
}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main() {
	plus(1, 2)
	var u1, u2 uint
	plus(u1, u2)

	type User struct {
		Name string `json:"name"`
	}

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	//user := Response{
	//	Code: 0,
	//	Msg:  "success",
	//	Data: User{
	//		Name: "john",
	//	},
	//}
	//
	//byteData, _ := json.Marshal(user)
	//fmt.Println(string(byteData))
	//
	//userInfo := Response{
	//	Code: 0,
	//	Msg:  "success",
	//	Data: UserInfo{
	//		Name: "john",
	//		Age:  21,
	//	},
	//}
	//
	//byteData1, _ := json.Marshal(userInfo)
	//fmt.Println(string(byteData1))

	//{"code":0,"msg":"success","data":{"name":"john"}}
	//{"code":0,"msg":"success","data":{"name":"john","age":21}}

	fmt.Println("-------------------------")
	var userResponse Response[User]

	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name":"john"}}`), &userResponse)

	fmt.Println(userResponse)
	fmt.Println(userResponse.Data.Name)

	var userInfoResponse Response[UserInfo]

	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name":"john","age":21}}`), &userInfoResponse)

	fmt.Println(userInfoResponse)
	fmt.Println(userInfoResponse.Data.Age)

	type MySlice[T int | string] []T

	var mySlice = MySlice[int]{1, 2, 3}

	fmt.Println(mySlice)

	// map的key只能是基本数据类型
	type MyMap[T int | string, K any] map[T]K

	var myMap = MyMap[string, int]{"name": 12}

	fmt.Println(myMap)

}
