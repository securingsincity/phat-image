package imagebuilder
import (
	"image"
	"log"
	"github.com/fogleman/gg"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/experimental/devices/inky"
	"periph.io/x/periph/host"
)

// WriteToEink writes to the phat eeink
func WriteToEink(dc *gg.Context) {
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
}