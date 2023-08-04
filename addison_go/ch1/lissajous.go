// lissajous generates a lissajous curve on standard out
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallet = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	f, err := os.Create("a_multicolor.gif")
	pallet = []color.Color{
		color.Black,
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 6
		res     = 0.001
		size    = 300
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallet)
		th := 0.0
		for cnx := uint8(1); cnx <= 3; cnx++ {
			for dth := 0.0; dth < cycles*2*math.Pi/3; dth += res {
				x := math.Sin(th + dth)
				y := math.Sin((th+dth)*freq + phase)
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), cnx)
			}
			th += 2 * math.Pi / 3
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
