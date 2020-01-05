package main

import "fmt"

type Cmd struct {
	id   int
	data string
}

func main() {
	ch := make(chan string, 1)
	go func() {
		ch <- "gogo"
	}()
	resp := <-ch
	fmt.Println(resp)

	var msgChan = make(chan *Cmd, 1)
	go func() {
		msgChan <- &Cmd{
			id:   1,
			data: "hello golang",
		}
	}()
	msgResp := <-msgChan
	fmt.Println(msgResp)
}
