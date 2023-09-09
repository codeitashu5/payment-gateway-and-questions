package goroutines

import (
	"fmt"
)

func hello() {
	fmt.Println("hello go routine function")
}

func writeToTheChanel(ch chan int, ch1 chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("writen the %d\n", i)
	}
	ch1 <- 0
}

func readFromChan(ch chan int) {
	if len(ch) == 1 {
		fmt.Printf("Read %d\n", <-ch)
		return
	}
	// need to avoid this temp variable
	temp := <-ch
	readFromChan(ch)
	fmt.Printf("Read %d\n", temp)
}

func BasicsOfGoroutines() {
	// creating a buffer chanel
	ch := make(chan int, 100)
	ch1 := make(chan int)
	go writeToTheChanel(ch, ch1)
	<-ch1
	// one thing is to read using recursion
	readFromChan(ch)
}
