package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
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

// Learning Go K&D Book
func main() {
	//fmt.Println("Hello World")

	//files := os.Args[1:]
	////dup1(files)
	//dup1(files)

	file, err := os.Create("lissajous.gif")

	if err != nil {
		log.Fatalf("%v", err)
		//log.Println(fmt.Errorf("error: %s", err))
		//file = os.Stdout
	}

	//fmt.Printf("File: %v\n,", file)

	lissajous(file)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
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
				uint8(rand.Intn(5)))
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

// Unix uniq
// dup1 Find duplicate lines in given files -v1
func dup1(files []string) {
	stats := make(map[string]map[string]int)

	for _, file := range files {
		c := make(map[string]int)

		f, err := os.Open(file)
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			c[scanner.Text()]++
		}

		stats[file] = c

		f.Close()
	}

	printResults(stats)
}

// dup2 Find duplicate lines in given files -v2
func dup2(files []string) {
	stats := make(map[string]map[string]int)

	for _, file := range files {
		c := make(map[string]int)

		bytes, err := os.ReadFile(file)

		if err != nil {
			continue
		}

		lines := strings.Split(string(bytes), "\n")

		for _, line := range lines {
			c[line]++
		}

		stats[file] = c
	}

	printResults(stats)
}

func printResults(stats map[string]map[string]int) {
	fmt.Println("------------ Duplicates Summary -------------")

	for name, counter := range stats {
		fmt.Printf("Filename: \t%s\n", name)
		fmt.Println("Count: \t Line")
		fmt.Println(strings.Repeat("---", 10))
		for line, c := range counter {
			if c > 1 {
				fmt.Printf("%d\t%s\n", c, line)
			}
		}
	}
}
