package main

import (
	"fmt"

	"github.com/spettacolo/chat/main/chat"
)

func main() {
	chat, err := chat.NewChat()
	if err != nil {
		fmt.Printf("Errore nell'inizializzazione della chat: %v\n", err)
		return
	}
	chat.Run()
}
