package main-footballfield2

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type footballFieldLine struct {
	minX  float64
	minY  float64
	maxX  float64
	maxY  float64
	color color.Color
}

func (p *footballFieldLine) draw(imd *imdraw.IMDraw) {
	imd.Color = p.color
	imd.Push(pixel.V(p.minX, p.minY))
	imd.Push(pixel.V(p.maxX, p.maxY))
	imd.Rectangle(2)
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

	//rect := pixel.R(1, 1, 7, 7)

	// choose your football field background color from
	// https://pkg.go.dev/golang.org/x/image/colornames#pkg-variables
	win.Clear(colornames.Darkolivegreen)

	//rect.Draw()

	//fmt.Println(rect.W(), rect.H()) // 6 6
	//fmt.Println(rect.Size())        // Vec(6, 6)
	//
	//fmt.Println(rect.Center())              // Vec(4, 4)
	//fmt.Println(rect.Moved(pixel.V(4, 10))) // Rect(5, 11, 11, 17)

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

	var footballFieldLines []footballFieldLine

	//footballFieldLines := []footballFieldLine{
	//	{: 10, : 100, : 500, : 105},
	//	{10, 200, 500, 205},
	//	{10, 300, 500, 305},
	//	{10, 400, 500, 405},
	//	{10, 500, 500, 505},
	//	{10, 600, 500, 605},
	//	{10, 700, 500, 705},
	//	{10, 800, 500, 805},
	//	{10, 900, 500, 905},
	//	{10, 1000, 500, 1005},
	//	{10, 1100, 500, 1105},
	//}

	yardLine := 100

	for i := 0; i < 11; i++ {
		footballFieldLines[i].minX = 10.0
		footballFieldLines[i].minY = yardLine
		footballFieldLines[i].maxX = 500.0
		footballFieldLines[i].maxY = yardLine + 5.0
		footballFieldLines[i].color = pixel.RGB(256, 256, 256)
		yardLine += 100.0
	}

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
