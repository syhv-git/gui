package ui

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"gui/types"
	"gui/widgets"
	"image/color"
)

func BuildUI(g *types.GUI) {
	nav := new(types.NavBar)
	nav.Menu = widgets.BuildMenu(g)
	nav.Settings = widgets.BuildSettingsMenu(g)
	nav.Bar = widgets.BuildNavBar(nav)

	l := layout.NewMaxLayout()
	c := color.RGBA{A: 81}
	r := canvas.NewRectangle(c)
	page := container.New(l, r, nav.Menu, nav.Settings)
	main := container.NewVBox(nav.Bar, page)
	g.Main.SetContent(main)
}
