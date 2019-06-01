package main

import (
	"fmt"
	"time"
)

func test_ticker()  {
	ticker := time.NewTicker(3 * time.Second)//每次间隔3秒执行
	i := 0
	for {
		<- ticker.C
		i++;
		fmt.Println(i)
	}
}

func main()  {
	test_ticker()
}