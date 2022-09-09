package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gui/types"
)

func BuildMenu(g *types.GUI) fyne.CanvasObject {
	l1 := widget.NewLabel("Main Page")
	l2 := widget.NewLabel("Apps")
	l3 := widget.NewLabel("Go")
	l4 := widget.NewLabel("Here")
	list := container.NewVBox(l1, l2, l3, l4)
	m := widget.NewModalPopUp(list, g.Main.Canvas())
	s := fyne.Size{Width: 200, Height: 400}
	m.Resize(s)
	p := fyne.Position{X: 10, Y: 100}
	m.Move(p)
	m.Hide()
	return m
}
