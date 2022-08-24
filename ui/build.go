package ui

import (
	"gui/types"
	"gui/ui/components"
)

func BuildUI(g *types.GUI) {
	nav := components.BuildNavBar(g)

	g.Main.SetContent(nav)
}
