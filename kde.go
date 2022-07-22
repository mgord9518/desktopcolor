package desktopcolor

import (
	"os"
	"path/filepath"
	"image/color"
	"errors"
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

	d.Accent, err = strToRGBA(i.Section("General").Key("AccentColor").Value())
	if err != nil { return nil, err }
	
	d.Foreground, err = strToRGBA(i.Section("Colors:Window").Key("ForegroundActive").Value())
	if err != nil { return nil, err }
	
	d.Background, err = strToRGBA(i.Section("Colors:Window").Key("BackgroundNormal").Value())
	if err != nil { return nil, err }
	
	d.BackgroundAlternate, err = strToRGBA(i.Section("Colors:Window").Key("BackgroundAlternate").Value())
	if err != nil { return nil, err }

	return d, err
}

func strToRGBA(str string) (color.RGBA, error) {
	s := strings.Split(str, ",")

	if len(s) < 3 {
		return color.RGBA{}, errors.New("color value invalid")
	}
	
	var iArr []uint8

	for _, val := range s {
		num, _ := strconv.Atoi(val)
		iArr = append(iArr, uint8(num))
	}

	rgba := color.RGBA{
		R: iArr[0],
		G: iArr[1],
		B: iArr[2],
		A: 255,
	}

	return rgba
}

