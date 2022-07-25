package desktopcolor

import (
	"image/color"
)

type DesktopColor struct {
	Accent              color.RGBA
	Background          color.RGBA
	BackgroundAlternate color.RGBA
	Foreground          color.RGBA
	HighlightBackground color.RGBA
	HighlightForeground color.RGBA
	Hyperlink           color.RGBA
	HyperlinkVisited    color.RGBA
}
