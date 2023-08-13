/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package sock

import (
	"fmt"
	"go-tic-tac/game"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

const origin = "http://localhost/"
const endpoint = "ws://localhost:9876/ws"

// JoinServer of game and handle exchanges
func JoinServer(w *fyne.Window, notifChan chan string, replyChan chan game.Payload) {
	ws, err := websocket.Dial(endpoint, "", origin)
	if err != nil {
		close(notifChan)
		showErrorAndQuit(w, err)
		return
	}
	defer ws.Close()
	ws.MaxPayloadBytes = 1024

	//listen for updates from server
	go func(ws *websocket.Conn) {
	MatchLoop:
		for {
			var payload game.Payload
			if err = websocket.JSON.Receive(ws, &payload); err != nil {
				close(notifChan)
				showErrorAndQuit(w, err)
				break MatchLoop
			}
			// update UI with server message
			notifChan <- payload.Content
			log.Println("RAW payload", payload)
			switch payload.MessageType {
			case game.EXIT:
				close(notifChan)
				showErrorAndQuit(w, errors.New(payload.Content))
				break MatchLoop

			case game.WELCOME:
				game.SetMyPieceType(payload.FromUser)

			case game.START:
				game.IsReady.Store(true)

			case game.MOVE:
				if !game.IsMyTurn.Load() {
					//opponent played
					notifChan <- "OPPONENT PLAYED " + payload.Content + ". Your turn"
					targetIndex, _ := strconv.Atoi(payload.Content)
					game.PlaceOpponentMark(targetIndex, payload.FromUser)
					game.IsMyTurn.Swap(true)
				}

			case game.WIN, game.LOSE:
				game.Over.Store(true)
				close(notifChan)
				winningCells := payload.WinningCells
				iHaveWon := game.WIN == payload.MessageType
				game.HighlightBoxes(winningCells, iHaveWon)
				showLoserWinner(w, payload.Content)
				break MatchLoop

			case game.DRAW:
				game.Over.Store(true)
				close(notifChan)
				showLoserWinner(w, payload.Content)
				break MatchLoop
			}
		}
	}(ws)

uiLoop:
	for {
		select {
		// listen for UI messages, forward them to server
		case rr := <-replyChan:
			notifChan <- "YOU PLAYED " + rr.Content
			if err = websocket.JSON.Send(ws, rr); err != nil {
				close(notifChan)
				showErrorAndQuit(w, err)
				break uiLoop
			}
		default:
			//TODO add a countdown timer for 20 sec. if no move after expiry, force exit game
			// else reset the timer. repeat loop
			continue
		}

	}

}

// display error dialog and exit onClick "OK"
func showErrorAndQuit(w *fyne.Window, err error) {
	fmt.Printf("%+v\n", errors.WithStack(err))
	d := dialog.NewError(err, *w)
	d.Show()
	d.Resize(fyne.NewSize(300, 100))
	d.SetOnClosed(func() {
		fyne.CurrentApp().Quit()
	})
}

// display Match winner or loser, and exit
func showLoserWinner(w *fyne.Window, msg string) {
	d := dialog.NewInformation("Game over!", msg, *w)
	d.Show()
	d.Resize(fyne.NewSize(300, 100))
	d.SetOnClosed(func() {
		fyne.CurrentApp().Quit()
	})
}
