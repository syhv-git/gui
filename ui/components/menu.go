package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gui/types"
)

func menu(g *types.GUI) (list fyne.CanvasObject) {
	l1 := widget.NewLabel("Test")
	l2 := widget.NewLabel("Apps")
	l3 := widget.NewLabel("Goes")
	l4 := widget.NewLabel("Here")
	list = container.NewVBox(l1, l2, l3, l4)
	return
}
