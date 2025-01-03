package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {

	fmt.Println(name, "开始购物")
	time.Sleep(2 * time.Second)
	fmt.Println(name, "购物结束")

	moneyChan <- money

	wait.Done()

}

func main() {

	var wait sync.WaitGroup
	startTime := time.Now()

	wait.Add(3)

	go pay("fire", 2, &wait)
	go pay("tang", 3, &wait)
	go pay("myl", 5, &wait)

	go func() {
		defer close(moneyChan)
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
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}

	//wait.Wait()

	fmt.Println(time.Since(startTime))
	fmt.Println(moneyList)

}
