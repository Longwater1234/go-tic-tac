/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package game

type MessageType int

const (
	WELCOME MessageType = iota
	START
	EXIT
	MOVE
	WIN
	LOSE
	DRAW
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
	case DRAW:
		return "DRAW"
	}
	return "unknown"
}

// Payload sent to client in json
type Payload struct {
	MessageType  MessageType `json:"messageType"`  // what's the message about
	Content      string      `json:"content"`      // the main content
	FromUser     string      `json:"fromUser"`     // source user of message
	WinningCells []int       `json:"winningCells"` // grid cells which caused win
}
