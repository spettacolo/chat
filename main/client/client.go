package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Set up a channel to listen for interrupt signals to close the connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	// Send "hello world" message
	err = conn.WriteMessage(websocket.TextMessage, []byte("hello world"))
	if err != nil {
		log.Println("write:", err)
		return
	}

	// Close the connection by sending a close message and then
	// waiting (with timeout) for the server to close the connection.
	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
		return
	}
	select {
	case <-done:
	case <-time.After(time.Second):
	}
}
