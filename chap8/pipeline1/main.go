package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Square
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // チャネルは閉じられて空になっている
			}
			squares <- x * x
		}
	}()

	// Printer (main go routine)
	for {
		fmt.Println(<-squares)
	}
}
