package desktopcolor

import (
	"os"
	"errors"
)

func GetColors() (*DesktopColor, error) {
	desktop, present := os.LookupEnv("XDG_SESSION_DESKTOP")

	// TODO: implement fallback
	if !present {
		desktop, present = os.LookupEnv("XDG_CURRENT_DESKTOP")
		if !present {
			return &DesktopColor{}, errors.New("XDG_SESSION_DESKTOP not set, desktop unknown")
		}
	}
	
	if desktop != "KDE" &&
	desktop != "LXQt" {
		return &DesktopColor{}, errors.New("unknown desktop enviornment, can't find colors")
	}

	if desktop == "KDE" {
		return GetColorsFromKDE()
	} else if desktop == "LXQt" {
		return GetColorsFromLXQT()
	}

	return &DesktopColor{}, nil
}
