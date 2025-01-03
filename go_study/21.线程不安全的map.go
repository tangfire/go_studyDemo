package main

import "fmt"

func main() {
	var maps = map[int]string{}
	go func() {
		for {
			maps[1] = "tangfire"
		}
	}()

	go func() {
		for {
			fmt.Println(maps[1])
		}
	}()

	select {}
}
