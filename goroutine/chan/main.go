package main

import "fmt"

type Cmd struct {
	id   int
	data string
}

func main() {
	ch := make(chan string, 1)
	go func() {
		ch <- "hello, golang"
	}()
	resp := <-ch
	fmt.Println(resp)

	var cmdChan = make(chan *Cmd, 1)
	go func() {
		cmdChan <- &Cmd{
			id:   1,
			data: "hello golang",
		}
	}()
	cmdResp := <-cmdChan
	fmt.Println(cmdResp)
}
