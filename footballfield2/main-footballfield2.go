package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//rectangles to make the football field 5 yard lines
type footballFieldLine struct {
	minX  float64
	minY  float64
	maxX  float64
	maxY  float64
	color color.Color
}

//create lines - yard hash marks on the football field
func (p *footballFieldLine) drawHashes(imd *imdraw.IMDraw) {
	imd.Color = p.color
	imd.Push(pixel.V(p.minX, p.minY))
	imd.Push(pixel.V(p.maxX, p.maxY))
	imd.Line(1)
}

//create the lines on the field that are rectangles
func (p *footballFieldLine) draw(imd *imdraw.IMDraw) {
	imd.Color = p.color
	imd.Push(pixel.V(p.minX, p.minY))
	imd.Push(pixel.V(p.maxX, p.maxY))
	imd.Rectangle(1)
}

func (p *footballFieldLine) defineHashLines() {

}

func createHashLines()

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Football Play Simulator",
		Bounds: pixel.R(0, 0, 700, 700),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// smooth out the graphics i.e. don't be pixelated
	win.SetSmooth(true)

	// choose your football field background color from
	// https://pkg.go.dev/golang.org/x/image/colornames#pkg-variables
	win.Clear(colornames.Darkolivegreen)

	imd := imdraw.New(nil)

	//Offensive Players
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

	// The lines on the football field:
	// 1 pixel = 3.6 inches
	// 10 pixels = 1 yard
	// 500 pixels = 50 yard wide field
	// 1000 pixels = 100 yard long field

	var footballFieldOutsideLines footballFieldLine
	var footballFieldLines []footballFieldLine
	var footballFieldHashLines []footballFieldLine
	var footballFieldEndZones []footballFieldLine

	leftSideLinePixel := 10
	rightSideLinePixel := 600

	//create the outside of the football field
	footballFieldOutsideLines.minX = float64(leftSideLinePixel)
	footballFieldOutsideLines.minY = float64(100)
	footballFieldOutsideLines.maxX = float64(rightSideLinePixel)
	footballFieldOutsideLines.maxY = float64(1000)
	footballFieldOutsideLines.color = pixel.RGB(256, 256, 256)

	var values footballFieldLine

	//create the 5 yard lines
	yardLine := 100
	for i := 0; i < 21; i++ {
		values.minX = float64(leftSideLinePixel)
		values.minY = float64(yardLine)
		values.maxX = float64(rightSideLinePixel)
		values.maxY = float64(yardLine)
		values.color = pixel.RGB(256, 256, 256)

		footballFieldLines = append(footballFieldLines, values)
		yardLine += 50
	}

	//create the end zones
	yardLine = 0
	for i := 0; i < 21; i++ {
		values.minX = float64(leftSideLinePixel)
		values.minY = float64(yardLine)
		values.maxX = float64(rightSideLinePixel)
		values.maxY = float64(yardLine + 100)
		values.color = pixel.RGB(256, 256, 256)

		footballFieldEndZones = append(footballFieldEndZones, values)
		yardLine += 1000
	}

	for !win.Closed() {
		win.Clear(colornames.Darkolivegreen)

		//imd.Clear()
		for _, p := range footballFieldLines {
			p.draw(imd)
		}
		for _, p := range footballFieldHashLines {
			p.drawHashes(imd)
		}
		for _, p := range footballFieldEndZones {
			p.draw(imd)
		}

		footballFieldOutsideLines.draw(imd)

		imd.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
