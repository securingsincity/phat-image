package imagebuilder

import (
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
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

func style3(dc *gg.Context, boldFace font.Face, boldItalicFace font.Face, title string, artist string) *gg.Context {

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.Rotate(gg.Radians(0))

	dc.SetRGB255(255, 0, 0)
	dc.DrawRectangle(0, 52, 300, 100)
	dc.Fill()

	dc.SetFontFace(boldFace)
	dc.SetRGB255(255, 255, 255)
	uppercaseArtist := truncateString(strings.ToUpper(artist), 40)
	dc.DrawStringWrapped(uppercaseArtist, 0, 55, 0, 0, 210, 0.8, gg.AlignCenter)

	dc.SetFontFace(boldItalicFace)
	dc.SetRGB255(255, 0, 0)
	titleTruncated := truncateString(title, 40)
	// dc.DrawStringAnchored(text2, 100, 10, 0.5, 0.5)
	dc.DrawStringWrapped(titleTruncated, 0, 10, 0, 0, 212, 1, gg.AlignCenter)

	dc.SetRGB255(0, 0, 0)
	textBY := "BY"
	dc.DrawStringAnchored(textBY, 110, 50, 0.5, 0.5)
	return dc
}

func style4(dc *gg.Context, boldFace font.Face, boldItalicFace font.Face, title string, artist string) *gg.Context {

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.Rotate(gg.Radians(-20))

	dc.SetRGB255(255, 0, 0)
	dc.DrawRectangle(-100, 80, 300, 100)
	dc.Fill()

	dc.SetFontFace(boldFace)
	dc.SetRGB255(255, 255, 255)
	uppercaseArtist := truncateString(strings.ToUpper(artist), 40)
	dc.DrawStringWrapped(uppercaseArtist, -30, 85, 0, 0, 210, 0.8, gg.AlignCenter)

	dc.SetFontFace(boldItalicFace)
	dc.SetRGB255(255, 0, 0)
	titleTruncated := truncateString(title, 40)
	// dc.DrawStringAnchored(text2, 100, 10, 0.5, 0.5)
	dc.DrawStringWrapped(titleTruncated, -50, 40, 0, 0, 212, 1, gg.AlignCenter)

	dc.SetRGB255(0, 0, 0)
	textBY := "BY"
	dc.DrawStringAnchored(textBY, 180, 80, 0.5, 0.5)
	return dc
}

func GenerateImage(style int, boldFace font.Face, boldItalicFace font.Face, title string, artist string) *gg.Context {
	dc := gg.NewContext(212, 104)
	switch style {
	case 1:
		dc = style1(dc, boldFace, boldItalicFace, title, artist)
		break
	case 2:
		dc = style2(dc, boldFace, boldItalicFace, title, artist)
		break
	case 3:
		dc = style3(dc, boldFace, boldItalicFace, title, artist)
		break
	case 4:
		dc = style4(dc, boldFace, boldItalicFace, title, artist)
		break
	}
	return dc
}
