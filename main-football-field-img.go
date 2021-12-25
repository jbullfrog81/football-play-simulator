package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/golang/glog"
	"golang.org/x/image/colornames"
)

func logClose(file io.Closer) {
	err := file.Close()
	if err != nil {
		glog.Error(err)
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer logClose(file)
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Football Play Simulator",
		Bounds: pixel.R(0, 0, 1024, 700),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// smooth out the graphics i.e. don't be pixelated
	win.SetSmooth(true)

	rect := pixel.R(1, 1, 7, 7)

	// choose your football field background color from
	// https://pkg.go.dev/golang.org/x/image/colornames#pkg-variables
	win.Clear(colornames.Darkolivegreen)

	//rect.Draw()

	fmt.Println(rect.W(), rect.H()) // 6 6
	fmt.Println(rect.Size())        // Vec(6, 6)

	fmt.Println(rect.Center())              // Vec(4, 4)
	fmt.Println(rect.Moved(pixel.V(4, 10))) // Rect(5, 11, 11, 17)

	pic, err := loadPicture("./images/footballField.jpg")
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	//imd.Color = colornames.Blueviolet
	//imd.EndShape = imdraw.RoundEndShape
	//imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	//imd.EndShape = imdraw.SharpEndShape
	//imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	//imd.Line(30)

	// left wide receiver
	imd.Color = colornames.Black
	imd.Push(pixel.V(100, 100))
	imd.Circle(15, 2)

	// left guard
	imd.Color = colornames.Black
	imd.Push(pixel.V(300, 100))
	imd.Circle(15, 2)

	// center
	imd.Color = colornames.Black
	imd.Push(pixel.V(325, 85))
	imd.Push(pixel.V(355, 115))
	imd.Rectangle(2)

	// right guard
	imd.Color = colornames.Black
	imd.Push(pixel.V(380, 100))
	imd.Circle(15, 2)

	// QB
	imd.Color = colornames.Black
	imd.Push(pixel.V(340, 60))
	imd.Circle(15, 2)

	// Right Twins
	// inside twin
	imd.Color = colornames.Black
	imd.Push(pixel.V(580, 100))
	imd.Circle(15, 2)
	// outside twin
	imd.Color = colornames.Black
	imd.Push(pixel.V(620, 100))
	imd.Circle(15, 2)

	//for !win.Closed() {
	//	win.Update()
	//}
	for !win.Closed() {
		win.Clear(colornames.Darkolivegreen)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
