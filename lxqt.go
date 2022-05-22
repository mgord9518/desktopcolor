package desktopcolor

import (
	"os"
	"path/filepath"
	"image/color"
	"strconv"

	ini "github.com/go-ini/ini"
	xdg "github.com/adrg/xdg"
)

func GetColorsFromLXQT() (*DesktopColor, error) {
	d := &DesktopColor{}

	// Open `lxqt.conf` file (INI format)
	f, err := os.Open(filepath.Join(xdg.ConfigHome, "lxqt", "lxqt.conf"))
	if err != nil { return nil, err }

	i, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, f)
	if err != nil { return nil, err }

	d.Foreground, err          = hexStrToRGBA(i.Section("Palette").Key("window_text_color").Value())
	if err != nil { return nil, err }
	d.Background, err          = hexStrToRGBA(i.Section("Palette").Key("window_color").Value())
	if err != nil { return nil, err }
	d.BackgroundAlternate, err = hexStrToRGBA(i.Section("Palette").Key("window_color").Value())

	return d, err
}

func hexStrToRGBA(str string) (color.RGBA, error) {
	r, err := strconv.ParseInt(str[1:3], 16, 16)
	g, err := strconv.ParseInt(str[3:5], 16, 16)
	b, err := strconv.ParseInt(str[5:7], 16, 16)
	if err != nil { return color.RGBA{}, err }

	rgba := color.RGBA {
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}

	return rgba, nil
}

