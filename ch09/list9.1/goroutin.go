package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers1() {
	for i := 0; i < 100; i++ {
		// fmt.Printf("%d", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+100; i++ {
		// fmt.Printf("%c", i)
	}
}

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+100; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	wg.Done()
}

// func print1() {
// 	printNumbers1()
// 	printLetters1()
// }

// func print2() {
// 	printNumbers2()
// 	printLetters2()
// }

// func goPrint1() {
// 	go printNumbers1()
// 	go printLetters1()
// }

// func goPrint2() {
// 	go printNumbers2()
// 	go printLetters2()
// }

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()
}
