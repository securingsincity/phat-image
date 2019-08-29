package main

import (
	"flag"

	"github.com/golang/freetype/truetype"
	"github.com/securingsincity/phat-image/imagebuilder"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
)

func main() {
	const S = 400
	var writeOutImage = false
	var title = ""
	var artist = ""
	var style = 1
	flag.StringVar(&title, "title", "", "name of the song")
	flag.StringVar(&artist, "artist", "", "name of the song")
	flag.IntVar(&style, "style", 1, "name of the song")
	flag.BoolVar(&writeOutImage, "outto", false, "output images")
	flag.Parse()

	boldFont, err := truetype.Parse(gobold.TTF)
	if err != nil {
		panic("")
	}
	boldFace := truetype.NewFace(boldFont, &truetype.Options{
		Size: 18,
	})

	boldItalicfont, _ := truetype.Parse(gobolditalic.TTF)

	boldItalicFace := truetype.NewFace(boldItalicfont, &truetype.Options{
		Size: 16,
	})

	dc := imagebuilder.GenerateImage(style, boldFace, boldItalicFace, title, artist)
	if writeOutImage == true {
		imagebuilder.WriteToEink(dc)
	} else {
		dc.SavePNG("out.png")
	}

}
