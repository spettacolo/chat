package chat

type Message struct {
	Sender      string `json:"sender"`  // Username dell'utente che invia il messaggio
	RoomID      string `json:"room_id"` // ID della stanza (privata, gruppo o canale)
	Content     string `json:"content"` // Contenuto del messaggio
	MessageType string `json:"type"`    // Tipo di messaggio ("private", "group", "channel")
}
