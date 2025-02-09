package main

import "fmt"

type Code int

func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "成功"
	case ServiceErrCode:
		return "服务错误"
	case NetworkErrCode:
		return "网络错误"
	}
	return ""
}

func (c Code) Ok() (code Code, msg string) {
	return c, c.GetMsg()
}

const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 1001
	NetworkErrCode Code = 1002
)

func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode.Ok()
	}

	if name == "2" {
		return NetworkErrCode.Ok()
	}

	return NetworkErrCode.Ok()
}

func main() {
	fmt.Println(SuccessCode, ServiceErrCode, NetworkErrCode)
}
