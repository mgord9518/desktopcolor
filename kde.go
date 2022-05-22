package desktopcolor

import (
	"os"
	"path/filepath"
	"image/color"
	"strings"
	"strconv"

	ini "gopkg.in/ini.v1"
	xdg "github.com/adrg/xdg"
)

func GetColorsFromKDE() (*DesktopColor, error) {
	d := &DesktopColor{}

	// Open `kdeglobals` file (INI format)
	f, err := os.Open(filepath.Join(xdg.ConfigHome, "kdeglobals"))
	if err != nil { return nil, err }

	i, err := ini.Load(f)
	if err != nil { return nil, err }

	d.Accent              = strToRGBA(i.Section("General").Key("AccentColor").Value(), ",")
	d.Foreground          = strToRGBA(i.Section("Colors:Window").Key("ForegroundActive").Value(), ",")
	d.Background          = strToRGBA(i.Section("Colors:Window").Key("BackgroundNormal").Value(), ",")
	d.BackgroundAlternate = strToRGBA(i.Section("Colors:Window").Key("BackgroundAlternate").Value(), ",")

	return d, err
}

func strToRGBA(str string, d string) color.RGBA {
	s := strings.Split(str, d)

	var iArr []uint8

	for _, val := range s {
		num, _ := strconv.Atoi(val)
		iArr = append(iArr, uint8(num))
	}

	rgba := color.RGBA {
		R: iArr[0],
		G: iArr[1],
		B: iArr[2],
		A: 255,
	}

	return rgba
}

