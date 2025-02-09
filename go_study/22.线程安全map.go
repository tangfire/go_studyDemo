package main

import (
	"fmt"
	"sync"
)

func main() {

	var maps = sync.Map{}
	go func() {
		for {
			maps.Store(1, "tangfire")

		}
	}()

	go func() {
		for {
			val, ok := maps.Load(1)
			fmt.Println(val, ok)
		}
	}()

	select {}
}
