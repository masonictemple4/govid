package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"gocv.io/x/gocv"
)

func main() {
	wc, _ := gocv.VideoCaptureDevice(0)
	defer wc.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("test.xml") {
		log.Fatalf("Error reading cascade file: %s\n", "test.xml")
	}

	println("Reading camera device")
	lastVal := true
	awaycount := 0
	f, _ := os.Open("test.mp3")
	defer f.Close()
	d, _ := mp3.NewDecoder(f)
	c, _ := oto.NewContext(d.SampleRate(), 2, 2, 1024)
	defer c.Close()
	p := c.NewPlayer()
	defer p.Close()
	for {
		if ok := wc.Read(&img); !ok {
			log.Fatalf("Cannot read device")
		}
		if img.Empty() {
			continue
		}

		gocv.Flip(img, &img, 1)
		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("Found %d faces Uninterupted zeros: %d\n", len(rects), awaycount)
		val := len(rects) > 0
		if val == lastVal && val == false {
			awaycount++
		} else {
			awaycount = 0
		}
		lastVal = val

		if awaycount >= 10 {
			io.Copy(p, d)
			os.Exit(1)
		}

		// draw a rectangle around each face on the original image
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}
	}

}
