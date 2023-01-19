/*
 * Copyright (c) 2023, Davis Tibbz, MIT License.
 */

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"go-tic-tac/game"
	"go-tic-tac/player"
	"image/color"
	"log"
)

func init() {
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
	myApp := app.New()
	w := myApp.NewWindow("Tic-Tac-Tiba")
	grid := container.New(layout.NewGridLayout(3))
	for i := 0; i < 9; i++ {
		rect := canvas.NewRectangle(color.RGBA{
			R: 255,
			G: 165,
			B: 0,
			A: 255,
		})

		gridBox := game.NewGridBox(rect, i, &w)
		grid.Add(gridBox)
	}

	w.SetContent(grid)
	w.Resize(fyne.NewSize(900, 600))
	r, e := fyne.LoadResourceFromPath("game_icon.png")
	if e != nil {
		log.Fatalln(e)
	}
	w.Show()
	w.SetIcon(r)
	w.SetFixedSize(true)
	myApp.Run()
}
