package game

import "golang.org/x/net/websocket"

type MessageType int

const (
	WELCOME MessageType = iota
	START
	EXIT
	MOVE
	WIN
	LOSE
)

func (r MessageType) String() string {
	switch r {
	case WELCOME:
		return "WELCOME"
	case START:
		return "START"
	case EXIT:
		return "EXIT"
	case MOVE:
		return "MOVE"
	case WIN:
		return "WIN"
	case LOSE:
		return "LOSE"
	}
	return "unknown"
}

// Payload sent to client in json
type Payload struct {
	MessageType MessageType `json:"messageType"`
	Content     string      `json:"content"`
	FromUser    string      `json:"fromUser"`
}

// SendMessage to server in JSON
func SendMessage(payload *Payload, ws *websocket.Conn) error {
	err := websocket.JSON.Send(ws, payload)
	return err
}
