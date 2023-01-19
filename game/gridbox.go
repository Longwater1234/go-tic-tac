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
	"go-tic-tac/player"
	"image/color"
	"log"
)

var _ fyne.Tappable = (*gridBox)(nil)

var gameRecord map[int]player.SymbolGame  //keeps record of the game (cellIndex -> symbol)
var playerState map[string]*player.Player //keeps record of the player (playerName -> []indexes)

// Single cell inside the 3x3 grid.
// Custom widget. See https://developer.fyne.io/extend/custom-widget
type gridBox struct {
	widget.BaseWidget
	Index     int               //cell index
	rectangle *canvas.Rectangle //background of cell
	textVal   *canvas.Text      //text box
	container *fyne.Container   //hosts textVal and rectangle
	window    *fyne.Window      // master window
}

// default starts with X
var isPlayerXTurn = true

func (g *gridBox) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(g.container)
}

// Tapped overrides onClick listener
func (g *gridBox) Tapped(*fyne.PointEvent) {
	if isPlayerXTurn {
		g.textVal.Text = player.X.String()
		isPlayerXTurn = false
		gameRecord[g.Index] = player.X
	} else {
		g.textVal.Text = player.O.String()
		isPlayerXTurn = true
		gameRecord[g.Index] = player.O
	}

	if g.getWinner() != "" {
		g.displayWinner(g.getWinner())
		return
	}
	if g.allBoxFilled() {
		return
	}
	g.Refresh()
}

// NewGridBox creates a new single cell for grid
func NewGridBox(rectangle *canvas.Rectangle, Index int, window *fyne.Window) *gridBox {
	tv := &canvas.Text{
		Text:      "",
		Alignment: fyne.TextAlignCenter,
		Color:     color.Black,
		TextSize:  float32(80),
	}

	g := &gridBox{}
	g.ExtendBaseWidget(g)
	g.rectangle = rectangle
	g.Index = Index
	g.textVal = tv
	g.window = window
	g.container = container.NewMax(g.rectangle, tv)
	return g
}

// checks if all cells filled. If true, game over
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

// Evaluate who won the match
func (g *gridBox) getWinner() string {
	var p = playerState[g.textVal.Text]
	p.Vals = append(p.Vals, g.Index)

	//log.Printf("Game scoreboard %v", gameRecord)
	if p.HasWon() {
		return fmt.Sprintf("Player %v has Won!", p.Name)
	}
	return ""
}

// shows Winner and exit game
func (g *gridBox) displayWinner(msg string) {
	d := dialog.NewInformation("Game Over!", msg, *g.window)
	d.SetOnClosed(func() {
		fyne.CurrentApp().Quit()
	})
	d.Resize(fyne.NewSize(300, 100))
	d.Show()
}

// InitializeRecord (scoreboard) for the game
func InitializeRecord() {
	gameRecord = make(map[int]player.SymbolGame)
}

// InitializePlayers of the game, must be exactly 2 players
func InitializePlayers(p []player.Player) {
	if len(p) != 2 {
		log.Fatalf("players must be exactly 2, provided %d", len(p))
	}
	playerState = map[string]*player.Player{
		p[0].Name: &p[0],
		p[1].Name: &p[1],
	}
}
