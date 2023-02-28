package sock

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/pkg/errors"
	"go-tic-tac/game"
	"go-tic-tac/player"
	"golang.org/x/net/websocket"
	"log"
)

const origin = "http://localhost/"
const endpoint = "ws://localhost:9876/ws"

// JoinServer of game and handle responses
func JoinServer(w *fyne.Window, notifChan chan string, serverChan, clientChan, replyChan chan game.Payload) {
	ws, err := websocket.Dial(endpoint, "", origin)
	if err != nil {
		showErrorAndQuit(w, err)
		return
	}
	defer ws.Close()

	go func(ws *websocket.Conn) {
		for {
			var payload game.Payload
			err = websocket.JSON.Receive(ws, &payload)
			if err != nil {
				showErrorAndQuit(w, err)
				return
			}
			serverChan <- payload
		}
	}(ws)

	for {
		select {
		case payload := <-serverChan:
			switch payload.MessageType {
			case game.START:
				game.IsReady.Set()
				if game.MyCurrentSymbol.ValString == player.X {
					log.Println("game ready")
					notifChan <- payload.Content
					game.IsMyTurn.Set()
				}
			case game.WELCOME:
				notifChan <- payload.Content
				game.UpdateSymbol(payload.FromUser)
			case game.MOVE:
				clientChan <- payload
			case game.EXIT:
				notifChan <- payload.Content
				ws.Close()
				return
			}
		case uiResponse := <-replyChan:
			err = websocket.JSON.Send(ws, uiResponse)
			if err != nil {
				showErrorAndQuit(w, err)
				return
			}
		}
	}

}

func showErrorAndQuit(w *fyne.Window, err error) {
	fmt.Printf("%+v\n", errors.WithStack(err))
	d := dialog.NewInformation("Error", err.Error(), *w)
	d.Show()
	d.Resize(fyne.NewSize(300, 100))
	d.SetOnClosed(func() {
		fyne.CurrentApp().Quit()
	})
}
