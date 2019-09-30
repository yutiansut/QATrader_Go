//package QATrader_Go

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"time"

	"github.com/gorilla/websocket"
)

type peek struct {
	Aid string `json:"aid"`
}

type login_msg struct {
	Aid      string `json:"aid"`
	Bid      string `json:"bid"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

var addr = flag.String("addr", "www.yutiansut.com:7988", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	login := login_msg{
		Aid:      "req_login",
		Bid:      "QUANTAXIS",
		Username: "133496",
		Password: "133496",
	}

	login1, err := json.Marshal(login)
	log.Printf(string(login1))

	u := url.URL{Scheme: "ws", Host: *addr, Path: ""}
	//log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}

			err = c.WriteMessage(websocket.TextMessage, login1)
			if err != nil {
				log.Println("write:", err)
				return
			}
			pk, _ := json.Marshal(peek{Aid: "peek_message"})
			log.Println(string(pk))
			err = c.WriteMessage(websocket.TextMessage, pk)
			if err != nil {
				log.Println("write:", err)
				return
			}

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
