package animatedgif

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
)

var palette = []color.Color{color.Black,
	color.RGBA{
		R: 0,
		G: 100,
		B: 0,
		A: 1,
	},
	color.RGBA{
		R: 100,
		G: 90,
		B: 10,
		A: 1,
	}, color.RGBA{
		R: 70,
		G: 100,
		B: 50,
		A: 1,
	},
	color.RGBA{
		R: 0,
		G: 100,
		B: 150,
		A: 1,
	},
}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func Lissajous(out io.Writer, c *int) {
	var cycles = 5.0

	if c != nil {
		cycles = float64(*c)
	}

	const (
		//cycles = 5// number of complete x oscillator revolutions
		res     = 0.005 // angular resolution
		size    = 200   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 10    // delay between frames in 10ms units
	)

	freq := rand.Float64() * 1.5 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				//blackIndex,
				uint8(rand.Intn(5)),
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	err := gif.EncodeAll(out, &anim)

	if err != nil {
		log.Println(fmt.Errorf("error: %s", err))
		return
	} // NOTE: ignoring encoding errors
}
