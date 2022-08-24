package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"gui/types"
	"gui/ui"
)

func main() {
	g := new(types.GUI)
	g.App = app.New()
	g.Main = g.App.NewWindow("Home")
	g.Main.Resize(fyne.Size{Width: 750, Height: 500})
	g.Main.SetFixedSize(true)
	g.Main.SetMaster()
	g.Main.Show()

	go ui.BuildUI(g)

	g.App.Run()
}
