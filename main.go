package main

import (
	"flag"
	"image"
	"log"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/inky"
	"periph.io/x/periph/host"
)

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}

func style2(dc *gg.Context, boldFace font.Face, boldItalicFace font.Face, title string, artist string) *gg.Context {
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.Rotate(gg.Radians(20))
	dc.DrawRectangle(0, -80, 300, 100)
	dc.Fill()

	dc.SetRGB255(255, 0, 0)
	dc.DrawRectangle(0, 20, 300, 100)
	dc.Fill()

	dc.SetFontFace(boldFace)
	dc.SetRGB255(0, 0, 0)
	uppercaseArtist := truncateString(strings.ToUpper(artist), 40)
	dc.DrawStringWrapped(uppercaseArtist, 30, 25, 0, 0, 150, 0.8, gg.AlignCenter)

	dc.SetFontFace(boldItalicFace)
	dc.SetRGB255(255, 0, 0)
	titleTruncated := truncateString(title, 40)
	// dc.DrawStringAnchored(text2, 100, 10, 0.5, 0.5)
	dc.DrawStringWrapped(titleTruncated, 50, -35, 0, 0, 152, 1, gg.AlignCenter)

	dc.SetRGB255(255, 255, 255)
	textBY := "BY"
	dc.DrawStringAnchored(textBY, 110, 18, 0.5, 0.5)
	return dc
}

func style1(dc *gg.Context, boldFace font.Face, boldItalicFace font.Face, title string, artist string) *gg.Context {

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.Rotate(gg.Radians(0))
	dc.DrawRectangle(0, 0, 300, 100)
	dc.Fill()

	dc.SetRGB255(255, 0, 0)
	dc.DrawRectangle(0, 52, 300, 100)
	dc.Fill()

	dc.SetFontFace(boldFace)
	dc.SetRGB255(0, 0, 0)
	uppercaseArtist := truncateString(strings.ToUpper(artist), 40)
	dc.DrawStringWrapped(uppercaseArtist, 0, 55, 0, 0, 210, 0.8, gg.AlignCenter)

	dc.SetFontFace(boldItalicFace)
	dc.SetRGB255(255, 0, 0)
	titleTruncated := truncateString(title, 40)
	// dc.DrawStringAnchored(text2, 100, 10, 0.5, 0.5)
	dc.DrawStringWrapped(titleTruncated, 0, 10, 0, 0, 212, 1, gg.AlignCenter)

	dc.SetRGB255(255, 255, 255)
	textBY := "BY"
	dc.DrawStringAnchored(textBY, 110, 50, 0.5, 0.5)
	return dc
}

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

	dc := gg.NewContext(212, 104)
	switch style {
	case 1:
		dc = style1(dc, boldFace, boldItalicFace, title, artist)
		break
	case 2:
		dc = style2(dc, boldFace, boldItalicFace, title, artist)
		break
	}

	if writeOutImage == true {
		img := dc.Image()
		if _, err := host.Init(); err != nil {
			log.Fatal(err)
		}

		b, err := spireg.Open("SPI0.0")
		if err != nil {
			log.Fatal(err)
		}

		dcPin := gpioreg.ByName("22")
		reset := gpioreg.ByName("27")
		busy := gpioreg.ByName("17")

		dev, err := inky.New(b, dcPin, reset, busy, &inky.Opts{
			Model:       inky.PHAT,
			ModelColor:  inky.Red,
			BorderColor: inky.Black,
		})
		if err != nil {
			log.Fatal(err)
		}

		if err := dev.Draw(img.Bounds(), img, image.ZP); err != nil {
			log.Fatal(err)
		}
	} else {
		dc.SavePNG("out.png")
	}

}
