package main

import "encoding/json"
import "fmt"

type Msg struct {
	MsgId int    `json:"msg_id"`
	Data  string `json:"data"`
}

func main() {
	ch := make(chan string, 1)
	go func() {
		ch <- "gogo"
	}()
	resp := <-ch
	fmt.Println(resp)

	var msgChan = make(chan *Msg, 1)
	go func() {
		msgChan <- &Msg{
			MsgId: 1,
			Data:  "hello golang",
		}
	}()
	fmt.Println(<-msgChan)

	var bytes, err = json.Marshal(&Msg{
		MsgId: 2,
		Data:  "welcome to golang",
	})
	if err != nil {
		panic("json Marshal fail")
	}
	var bCh = make(chan []byte, 1)
	go func() {
		bCh <- bytes
	}()
	var msg Msg
	var wrong = json.Unmarshal(<-bCh, &msg)
	if wrong != nil {
		panic("Unmarshal fail")
	}
	fmt.Printf("msg = %+v\n", msg)
}
