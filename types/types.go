package types

import (
	"fyne.io/fyne/v2"
)

// GUI is an organized struct for the fyne API types
type GUI struct {
	App  fyne.App
	Main fyne.Window
}

type NavBar struct {
	Bar      fyne.CanvasObject
	Menu     fyne.CanvasObject
	Settings fyne.CanvasObject
}
