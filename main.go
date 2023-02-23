/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/tevino/abool"
	"go-tic-tac/game"
	"go-tic-tac/player"
	"go-tic-tac/sock"
	"image/color"
)

//go:embed Icon.png
var icon []byte

var IsMyTurn abool.AtomicBool = 1

func initGame() {
	game.InitializeRecord()
	var players = []player.Player{{
		Name: player.X.String(),
		Vals: []int{},
	}, {
		Name: player.O.String(),
		Vals: []int{},
	}}
	game.InitializePlayers(players)
}

func main() {
	initGame()
	myApp := app.New()
	w := myApp.NewWindow("Tic-Tac-Tiba")

	textPanel := canvas.NewText("Connecting", color.White)
	grid := container.New(layout.NewGridLayout(3))
	payloadChan := make(chan game.Payload, 1) //for full responses from server
	notifChan := make(chan string, 1)         // for notifications

	for i := 0; i < 9; i++ {
		rect := canvas.NewRectangle(color.RGBA{
			R: 255,
			G: 165,
			B: 0,
			A: 255,
		})

		gridBox := game.NewGridBox(rect, i, &w, payloadChan)
		grid.Add(gridBox)
	}

	mainWindow := container.NewVSplit(grid, textPanel)
	mainWindow.SetOffset(0.8)
	w.SetContent(mainWindow)
	w.Resize(fyne.NewSize(900, 800))
	r := fyne.NewStaticResource("Icon.png", icon)
	w.Show()
	w.SetIcon(r)
	w.SetFixedSize(true)

	go sock.JoinServer(payloadChan, &w, notifChan)
	//updates the notification box
	go func() {
		for msg := range notifChan {
			textPanel.Text = msg
			textPanel.Refresh()
		}
	}()
	myApp.Run()
}
