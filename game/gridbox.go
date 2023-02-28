/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package game

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/tevino/abool"
	"go-tic-tac/player"
	"image/color"
	"log"
	"strconv"
	"sync"
	"time"
)

var _ fyne.Tappable = (*gridBox)(nil)

var gameRecord map[int]string //keeps record of the game (cellIndex -> symbol)
// var playerState map[string]*player.Player //keeps record of the player (playerName -> []indexes)
var gridMap map[int]*gridBox      //maps cellIndex to gridBox
var IsMyTurn abool.AtomicBool = 1 // player turn,  default starts with X (player 1)

type PieceType struct {
	ValString player.SymbolGame
	sync.Mutex
}

func (t *PieceType) SetPieceType(val string) {
	t.Lock()
	defer t.Unlock()
	if val == player.X.String() {
		t.ValString = player.X
		log.Println("i am player x")
		IsMyTurn.Set()
	} else {
		t.ValString = player.O
		log.Println("i am player O")
		IsMyTurn.UnSet()
	}
}

var MyCurrentSymbol PieceType //default X

func UpdateSymbol(val string) {
	MyCurrentSymbol.SetPieceType(val)
}

// Single cell inside the 3x3 grid.
// Custom widget. See https://developer.fyne.io/extend/custom-widget
type gridBox struct {
	widget.BaseWidget
	Index     int               //cell index
	rectangle *canvas.Rectangle //background of cell
	textVal   *canvas.Text      //text box
	container *fyne.Container   //hosts textVal and rectangle
	window    *fyne.Window      //master window
	commChannel
}

// For communicating with game server
type commChannel struct {
	serverChan chan Payload //recieves messages from server
	clientChan chan Payload // sends messages to Server
	replyChan  chan Payload
}

// NewCommChannel constructor
func NewCommChannel(serverChan chan Payload, clientChan chan Payload, replyChan chan Payload) *commChannel {
	return &commChannel{serverChan: serverChan, clientChan: clientChan, replyChan: replyChan}
}

var gameOver = false

// CreateRenderer for custom widgets
func (g *gridBox) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(g.container)
}

// Tapped overrides onClick listener
func (g *gridBox) Tapped(_ *fyne.PointEvent) {
	if g.textVal.Text != "" || gameOver || IsMyTurn.IsNotSet() {
		//already filled
		return
	}
	log.Println("i am here")

	for payload := range g.commChannel.clientChan {
		switch payload.MessageType {
		case MOVE:
			if IsMyTurn.IsSet() {
				IsMyTurn.UnSet()
				g.textVal.Text = MyCurrentSymbol.ValString.String()
				g.Refresh()
				gameRecord[g.Index] = MyCurrentSymbol.ValString.String()
				pp := Payload{
					MessageType: MOVE,
					Content:     fmt.Sprintf("%d", g.Index),
					FromUser:    MyCurrentSymbol.ValString.String(),
				}
				g.replyChan <- pp

			} else {
				IsMyTurn.UnSet()
				targetIndex, _ := strconv.Atoi(payload.Content)
				placeOpponentMark(targetIndex, payload.FromUser)
			}

		case WIN:
		case LOSE:
			gameOver = true
			go func() {
				time.Sleep(500 * time.Millisecond)
				g.displayWinner(payload.Content)
			}()
			close(g.clientChan)
			return
		default:
			log.Println("Unknown command sent")
		}
	}
	if g.allBoxFilled() {
		return
	}

	g.Refresh()
}

// NewGridBox creates a new single cell of grid
func NewGridBox(rectangle *canvas.Rectangle, Index int, window *fyne.Window, commChannel *commChannel) *gridBox {
	tv := &canvas.Text{
		Text:      "",
		Alignment: fyne.TextAlignCenter,
		Color:     color.Black,
		TextSize:  float32(80),
	}

	g := &gridBox{
		Index:       Index,
		rectangle:   rectangle,
		textVal:     tv,
		container:   container.NewMax(rectangle, tv),
		window:      window,
		commChannel: *commChannel,
	}
	g.ExtendBaseWidget(g)
	gridMap[Index] = g
	return g
}

// checks if all cells filled. If true, it's a draw
func (g *gridBox) allBoxFilled() bool {
	if len(gameRecord) == 9 {
		d := dialog.NewInformation("Game Over", "It's a draw", *g.window)
		d.SetOnClosed(func() {
			fyne.CurrentApp().Quit()
		})
		d.Resize(fyne.NewSize(300, 100))
		d.Show()
		return true
	}
	return false
}

// getWinner of the m

// highlightBoxes green color (winning cells)
func highlightBoxes(arr []int) {
	for _, v := range arr {
		if g, exists := gridMap[v]; exists {
			g.rectangle.FillColor = color.RGBA{
				R: 0,
				G: 255,
				B: 0,
				A: 255,
			}
			g.Refresh()
		}
	}

}

// placeOpponentMark at given index with symbol (X or O)
func placeOpponentMark(targetIndex int, symbolChar string) {
	for i, box := range gridMap {
		if i == targetIndex {
			box.textVal.Text = symbolChar
			gameRecord[targetIndex] = symbolChar
			box.Refresh()
			break
		}
	}
}

// displayWinner and exit game
func (g *gridBox) displayWinner(msg string) {
	d := dialog.NewInformation("Game Over!", msg, *g.window)
	d.SetOnClosed(func() {
		fyne.CurrentApp().Quit()
	})
	d.Resize(fyne.NewSize(300, 100))
	d.Show()
}

// InitializeRecord for the game
func InitializeRecord() {
	gameRecord = make(map[int]string)
	gridMap = make(map[int]*gridBox)
	gameOver = false
}
