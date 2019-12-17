package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	go func() {
		ch <- "hello, golang"
	}()
	resp := <-ch
	fmt.Printf("resp = %+v\n", resp)
}
