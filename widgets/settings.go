package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gui/types"
)

func settingsMenu(g *types.GUI) *fyne.Container {
	t := setTheme(g)
	s := widget.NewSeparator()
	e := widget.NewButtonWithIcon("Exit", theme.LogoutIcon(), func() {
		g.App.Quit()
	})

	l := layout.NewVBoxLayout()
	return container.New(l, t, s, e)
}

func setTheme(g *types.GUI) (change fyne.CanvasObject) {
	light, _ := fyne.LoadResourceFromPath("images/light_mode.svg")
	dark, _ := fyne.LoadResourceFromPath("images/dark_mode.svg")
	c := g.App.Settings().Theme()
	v := g.App.Settings().ThemeVariant()

	l := widget.NewButtonWithIcon("Dark Mode", dark, func() {
		c.Color(theme.ColorGray, theme.VariantLight)
		g.App.Settings().SetTheme(c)
	})

	d := widget.NewButtonWithIcon("Light Mode", light, func() {
		c.Color(theme.ColorGray, theme.VariantDark)
		g.App.Settings().SetTheme(c)
	})
	if v == theme.VariantDark {
		return l
	}
	return d
}
