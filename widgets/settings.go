package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gui/types"
)

func BuildSettingsMenu(g *types.GUI) *fyne.Container {
	t := themeButton(g)

	e := widget.NewButtonWithIcon("Exit", theme.LogoutIcon(), func() {
		g.App.Quit()
	})
	s := widget.NewSeparator()
	v := layout.NewVBoxLayout()
	c := container.New(v, t, s, e)
	return c
}

func themeButton(g *types.GUI) fyne.CanvasObject {
	var (
		b  bool
		lt func()
	)
	if g.App.Settings().ThemeVariant() == theme.VariantDark {
		b = true
	} else {
		b = false
	}

	tb := new(widget.Button)
	l, _ := fyne.LoadResourceFromPath("images/light_mode.svg")
	d, _ := fyne.LoadResourceFromPath("images/dark_mode.svg")
	dt := func() {
		tb.Icon = l
		tb.Text = "Light Mode"
		tb.Refresh()
		tb.OnTapped = lt
		b = false
	}
	lt = func() {
		tb.Icon = d
		tb.Text = "Dark Mode"
		tb.Refresh()
		tb.OnTapped = dt
		b = true
	}

	if b {
		tb.Icon = d
		tb.Text = "Dark Mode"
		tb.OnTapped = dt
		return tb
	}
	tb.Icon = l
	tb.Text = "Light Mode"
	tb.OnTapped = lt
	return tb
}
