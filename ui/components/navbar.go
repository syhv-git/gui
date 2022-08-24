package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gui/types"
	"image/color"
)

func BuildNavBar(g *types.GUI) fyne.CanvasObject {
	sm := settingsMenu(g)
	sm.Hide()
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MenuIcon(), func() {
			_ = menu(g)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			sm.Move(fyne.Position{X: 600, Y: 40})
			sm.Resize(fyne.Size{Width: 140})
			sm.Show()
		}),
	)
	c := color.RGBA{A: 42}
	r := canvas.NewRectangle(c)
	return container.NewBorder(bar, widget.NewSeparator(), nil, nil, r, sm)
}
