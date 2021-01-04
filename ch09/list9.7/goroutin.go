package main

import (
	"fmt"
	"time"
)

func printNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	// チャネルにboolを入れて中断を解除
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	// チャネルにboolを入れて中断を解除
	w <- true
}

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers2(w1)
	go printLetters2(w2)
	// 何かが入るまで実行を中断
	<-w1
	<-w2
}
