package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {

	fmt.Println(name, "开始购物")
	time.Sleep(2 * time.Second)
	fmt.Println(name, "购物结束")
	wait.Done()

}

func main() {

	var wait sync.WaitGroup
	startTime := time.Now()

	wait.Add(3)

	go shopping("fire", &wait)
	go shopping("tang", &wait)
	go shopping("myl", &wait)

	time.Sleep(2 * time.Second)

	fmt.Println(time.Since(startTime))

	wait.Wait()

}
