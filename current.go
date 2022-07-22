package desktopcolor

import (
	"os"
	"errors"
)

func GetColors() (*DesktopColor, error) {
	desktop := os.Getenv("XDG_SESSION_DESKTOP")

	// TODO: implement fallback
	if len(desktop) < 1 {
		desktop = os.Getenv("XDG_CURRENT_DESKTOP")
		if len(desktop) < 1 {
			return &DesktopColor{}, errors.New("XDG_SESSION_DESKTOP not set, desktop unknown")
		}
	}

	if desktop == "KDE" {
		return GetColorsFromKDE()
	} else if desktop == "LXQt" {
		return GetColorsFromLXQT()
	}

	return nil, errors.New("unknown desktop enviornment, can't find colors")
}
