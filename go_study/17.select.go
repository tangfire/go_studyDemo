package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int) // 声明并初始化一个长度为0的信道
var nameChan1 = make(chan string)
var doneChan = make(chan struct{})

func send(name string, money int, wait *sync.WaitGroup) {

	fmt.Println(name, "开始购物")
	time.Sleep(2 * time.Second)
	fmt.Println(name, "购物结束")

	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()

}

func main() {

	var wait sync.WaitGroup
	startTime := time.Now()

	wait.Add(3)

	go send("fire", 2, &wait)
	go send("tang", 3, &wait)
	go send("myl", 5, &wait)

	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)
		wait.Wait()

	}()

	//time.Sleep(2 * time.Second)

	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}

	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:

				return
			}
		}
	}

	//wait.Wait()

	event()
	fmt.Println(time.Since(startTime))
	fmt.Println(moneyList)
	fmt.Println(nameList)

}
