package libs

import (
	gim "github.com/ozankasikci/go-image-merge"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"os"
)

func CreateMergeImage(label string) {

	grids := []*gim.Grid{
		{
			ImageFilePath:   "../img/wolf-04.jpg",
			BackgroundColor: color.Black,
			// these grids will be drawn on top of the first grid
			Grids: []*gim.Grid{
				{
					ImageFilePath: "../img/wolf-logo.png",
					OffsetX:       10, OffsetY: 10,
				},
			},
		},
	}
	rgba, err := gim.New(grids, 2, 1).Merge()

	if err != nil {
		panic(err)
	}

	addLabelMerge(rgba, 1200, 300, "WOLF - 102080160")
	//CreateImageLabel("WOLF - 102080160" , "grid-rgb-go-new.png")
	f, errCreate := os.Create("grid-rgb-go.png")
	if errCreate != nil {
		panic(errCreate)
	}
	defer f.Close()
	if err := png.Encode(f, rgba); err != nil {
		panic(err)
	}
}

func addLabelMerge(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
