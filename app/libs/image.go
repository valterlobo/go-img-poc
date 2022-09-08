package libs

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	fontBytes, err := ioutil.ReadFile("../img/Raleway-ExtraBold.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(col),
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    24,
			DPI:     72,
			Hinting: font.HintingNone,
		}),
		Dot: point,
	}
	d.DrawString(label)
}

func CreateImageLabel(label string, imgName string) {

	img := image.NewRGBA(image.Rect(0, 0, 300, 100))
	addLabel(img, 20, 30, label)

	f, err := os.Create(imgName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
