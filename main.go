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
	"go-tic-tac/game"
	"go-tic-tac/sock"
	"image/color"
)

//go:embed Icon.png
var icon []byte

func initGame() {
	game.InitializeRecord()
	//game.InitializePlayers(players)
}

func main() {
	initGame()
	myApp := app.New()
	w := myApp.NewWindow("Tic-Tac-Tiba")

	textPanel := canvas.NewText("Connecting", color.White)
	grid := container.New(layout.NewGridLayout(3))
	serverChan := make(chan game.Payload)   //for full responses from server
	clientChan := make(chan game.Payload)   //for full responses to server
	replyChan := make(chan game.Payload, 1) //for full responses to server
	notifChan := make(chan string, 1)       // for on-screen notifications

	for i := 0; i < 9; i++ {
		rect := canvas.NewRectangle(color.RGBA{
			R: 255,
			G: 165,
			B: 0,
			A: 255,
		})

		commChannel := game.NewCommChannel(serverChan, clientChan, replyChan)
		gridBox := game.NewGridBox(rect, i, &w, commChannel)
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

	go sock.JoinServer(&w, notifChan, serverChan, clientChan, replyChan)
	//updates the notification box
	go func() {
		for msg := range notifChan {
			textPanel.Text = msg
			textPanel.Refresh()
		}
	}()
	myApp.Run()
}
