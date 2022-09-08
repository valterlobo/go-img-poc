package libs

import (
	"github.com/fogleman/gg"
	"log"
)

func CreateImageFogleman(path string, label string) {

	im, err := gg.LoadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	const S = 256

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	/*if err := dc.LoadFontFace("/Library/Fonts/Arial.ttf", 96); err != nil {
		panic(err)
	}*/
	//dc.DrawStringAnchored(label, S/2, S/2, 0.5, 0.5)

	dc.DrawRoundedRectangle(0, 0, 256, 256, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored(label, S, S-S/12, 1, 0)
	dc.Clip()
	dc.SavePNG("out.png")
}
