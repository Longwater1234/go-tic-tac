/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package game

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go-tic-tac/player"
	"image/color"
	"log"
	"sync"
	"sync/atomic"
)

// force implementation
var _ fyne.Tappable = (*gridCell)(nil)

var gameRecord map[int]string // keeps record of the game (cellIndex -> symbol)
var gridMap map[int]*gridCell // maps cellIndex to gridCell
var isMyTurn atomic.Bool      // player turn,  default starts with X (player 1)
var IsReady atomic.Bool       // whether match is ready to start
var Over atomic.Bool          // whether game is over
var mu sync.Mutex

var myPieceType player.SymbolGame //can be either `X` or `O`

// SetMyPieceType for current match
func SetMyPieceType(val string) {
	mu.Lock()
	defer mu.Unlock()
	if val == player.X.String() {
		myPieceType = player.X
		log.Println("i am player X")
		isMyTurn.Store(true)
	} else {
		myPieceType = player.O
		log.Println("i am player O")
		isMyTurn.Store(false)
	}
}

// ToggleMyTurn switches current player's turn
func ToggleMyTurn() {
	var old = isMyTurn.Load()
	isMyTurn.Store(!old)
}

// GetMyTurn returns true if it's my turn
func GetMyTurn() bool {
	return isMyTurn.Load()
}

// Single cell inside the 3x3 grid.
// Custom widget. See https://developer.fyne.io/extend/custom-widget
type gridCell struct {
	widget.BaseWidget
	Index     int               //cell index
	rectangle *canvas.Rectangle //background of cell
	textBox   *canvas.Text      //text box
	container *fyne.Container   //hosts textBox and rectangle
	window    *fyne.Window      //master window
	replyChan chan Payload      //for messages from client UI to server
}

// CreateRenderer for custom widgets
func (g *gridCell) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(g.container)
}

// Tapped overrides onClick listener
func (g *gridCell) Tapped(_ *fyne.PointEvent) {
	if g.textBox.Text != "" || Over.Load() || !IsReady.Load() {
		//cannot use this cell
		return
	}

	log.Printf("I tapped gridIndex %d", g.Index)

	if isMyTurn.Load() {
		g.textBox.Text = myPieceType.String()
		gameRecord[g.Index] = myPieceType.String()
		pp := Payload{
			MessageType: MOVE,
			Content:     fmt.Sprintf("%d", g.Index),
			FromUser:    myPieceType.String(),
		}
		isMyTurn.Swap(false)
		g.replyChan <- pp
	}
	g.Refresh()
}

// NewGridCell creates a new single cell of 3x3 grid
func NewGridCell(rectangle *canvas.Rectangle, index int, window *fyne.Window, replyChan chan Payload) *gridCell {
	tv := &canvas.Text{
		Text:      "",
		Alignment: fyne.TextAlignCenter,
		Color:     color.Black,
		TextSize:  float32(80),
	}

	g := &gridCell{
		Index:     index,
		rectangle: rectangle,
		textBox:   tv,
		container: container.NewMax(rectangle, tv),
		window:    window,
		replyChan: replyChan,
	}
	g.ExtendBaseWidget(g)
	gridMap[index] = g
	return g
}

// HighlightBoxes with either GREEN (if I won) or RED (if I lost)
func HighlightBoxes(arr []int, won bool) {
	//RED
	fillColor := color.NRGBA{
		R: 255,
		A: 255,
	}
	if won {
		//GREEN
		fillColor = color.NRGBA{
			G: 255,
			A: 255,
		}
	}
	for _, v := range arr {
		if g, exists := gridMap[v]; exists {
			g.rectangle.FillColor = fillColor
			g.Refresh()
		}
	}
}

// PlaceOpponentPiece at given index with symbol (X or O)
func PlaceOpponentPiece(targetIndex int, symbolChar string) {
	for i, cell := range gridMap {
		if i == targetIndex {
			cell.textBox.Text = symbolChar
			gameRecord[targetIndex] = symbolChar
			cell.Refresh()
			break
		}
	}
}

// InitializeRecord for the game
func InitializeRecord() {
	gameRecord = make(map[int]string)
	gridMap = make(map[int]*gridCell)
	Over.Swap(false)
}
