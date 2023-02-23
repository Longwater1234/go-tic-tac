package sock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go-tic-tac/game"
	"golang.org/x/net/websocket"
)

const origin = "http://localhost/"
const endpoint = "ws://localhost:9876/ws"

// JoinServer of game and handle responses
func JoinServer(payloadChan chan game.Payload, w *fyne.Window, notifChan chan string) {
	var payload game.Payload
	ws, err := websocket.Dial(endpoint, "", origin)
	if err != nil {
		showErrorAndQuit(payloadChan, w, err)
		return
	}
	defer ws.Close()

	err = websocket.JSON.Receive(ws, &payload)
	if err != nil {
		showErrorAndQuit(payloadChan, w, err)
		return
	}
	payloadChan <- payload
	notifChan <- payload.Content
}

func showErrorAndQuit(msgChan chan game.Payload, w *fyne.Window, err error) {
	d := dialog.NewInformation("Error", err.Error(), *w)
	d.Show()
	close(msgChan)
	d.Resize(fyne.NewSize(300, 100))
	d.SetOnClosed(func() {
		fyne.CurrentApp().Quit()
	})
}
