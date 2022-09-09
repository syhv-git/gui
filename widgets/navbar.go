package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gui/types"
)

func BuildNavBar(n *types.NavBar) fyne.CanvasObject {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MenuIcon(), func() {
			n.Menu.Show()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			n.Settings.Show()
		}),
	)
	return container.NewBorder(bar, nil, n.Menu, nil)
}
