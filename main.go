/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package main

import (
	_ "embed"
	"fmt"
	"go-tic-tac/game"
	"go-tic-tac/sock"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

//go:embed Icon.png
var icon []byte

func initGame() {
	game.InitializeRecord()
}

func main() {
	initGame()
	myApp := app.New()
	w := myApp.NewWindow("Tic-Tac-Tiba")

	textPanel := canvas.NewText("Connecting...", color.White)
	txtContainer := container.NewMax(canvas.NewRectangle(color.Black), textPanel)

	grid := container.New(layout.NewGridLayout(3))
	replyChan := make(chan game.Payload) //from client UI -> SERVER
	notifChan := make(chan string, 1)    //from server -> client UI

	for i := 0; i < 9; i++ {
		//orange cell
		rect := canvas.NewRectangle(color.RGBA{
			R: 255,
			G: 165,
			B: 0,
			A: 255,
		})
		gridCell := game.NewGridCell(rect, i, &w, replyChan)
		grid.Add(gridCell)
	}

	mainWindow := container.NewVSplit(grid, txtContainer)
	mainWindow.SetOffset(0.8)
	w.SetContent(mainWindow)
	w.Resize(fyne.NewSize(900, 800))
	r := fyne.NewStaticResource("Icon.png", icon)
	w.Show()
	w.SetIcon(r)
	w.SetFixedSize(true)

	go sock.JoinServer(&w, notifChan, replyChan)

	go func() {
		//updates the notification box
		for msg := range notifChan {
			textPanel.Text = fmt.Sprintf("[%s]: %s", time.Now().Format("15:04:05"), msg)
			textPanel.Refresh()
		}
	}()
	myApp.Run()
}
