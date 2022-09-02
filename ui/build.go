package ui

import (
	"gui/types"
	"gui/widgets"
)

func BuildUI(g *types.GUI) {
	nav := widgets.BuildNavBar(g)

	g.Main.SetContent(nav)
}
