package main

import (
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
	_ "image/png"

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

func defineFootballFieldLines(footballFieldLines *[]footballFieldLine, footballFieldOutsideLines *footballFieldLine, footballFieldHashLines *[]footballFieldLine, footballFieldEndZones *[]footballFieldLine, leftSideLinePixel int, rightSideLinePixel int) {
	//create the outside of the football field
	footballFieldOutsideLines.minX = float64(leftSideLinePixel)
	footballFieldOutsideLines.minY = float64(100)
	footballFieldOutsideLines.maxX = float64(rightSideLinePixel)
	footballFieldOutsideLines.maxY = float64(1000)
	footballFieldOutsideLines.color = pixel.RGB(256, 256, 256)

	var values footballFieldLine

	//define the starting y pixel
	yardLine := 100

	//create the 5 yard lines
	for i := 0; i < 21; i++ {
		values.minX = float64(leftSideLinePixel)
		values.minY = float64(yardLine)
		values.maxX = float64(rightSideLinePixel)
		values.maxY = float64(yardLine)
		values.color = pixel.RGB(256, 256, 256)

		*footballFieldLines = append(*footballFieldLines, values)
		yardLine += 50
	}
	//reset the starting y pixel
	yardLine = 100

	//Create the right side hashlines
	for i := 0; i < 81; i++ {
		values.minX = float64(rightSideLinePixel - 5)
		values.minY = float64(yardLine)
		values.maxX = float64(rightSideLinePixel)
		values.maxY = float64(yardLine)
		values.color = pixel.RGB(256, 256, 256)

		*footballFieldHashLines = append(*footballFieldHashLines, values)
		yardLine += 10
	}

	//reset the starting y pixel
	yardLine = 100

	//Create the left side hashlines
	for i := 0; i < 81; i++ {
		values.minX = float64(leftSideLinePixel)
		values.minY = float64(yardLine)
		values.maxX = float64(leftSideLinePixel + 5)
		values.maxY = float64(yardLine)
		values.color = pixel.RGB(256, 256, 256)

		*footballFieldHashLines = append(*footballFieldHashLines, values)
		yardLine += 10
	}

	//reset the starting y pixel
	yardLine = 100

	//Create the left center hashlines
	// center hash marks are about .442 from each side of the whole distance = 221 pixels
	for i := 0; i < 81; i++ {
		values.minX = float64(leftSideLinePixel + 221)
		values.minY = float64(yardLine)
		values.maxX = float64(leftSideLinePixel + 5 + 221 - 5 - 5)
		values.maxY = float64(yardLine)
		values.color = pixel.RGB(256, 256, 256)

		*footballFieldHashLines = append(*footballFieldHashLines, values)
		yardLine += 10
	}

	//reset the starting y pixel
	yardLine = 100

	//Create the right center hashlines
	// center hash marks are about .442 from each side of the whole distance = 221 pixels
	for i := 0; i < 81; i++ {
		values.minX = float64(rightSideLinePixel - 5 - 221)
		values.minY = float64(yardLine)
		values.maxX = float64(rightSideLinePixel - 221)
		values.maxY = float64(yardLine)
		values.color = pixel.RGB(256, 256, 256)

		*footballFieldHashLines = append(*footballFieldHashLines, values)
		yardLine += 10
	}

	yardLine = 1

	//create the end zones
	for i := 0; i < 21; i++ {
		values.minX = float64(leftSideLinePixel)
		values.minY = float64(yardLine)
		values.maxX = float64(rightSideLinePixel)
		values.maxY = float64(yardLine + 100)
		values.color = pixel.RGB(256, 256, 256)

		*footballFieldEndZones = append(*footballFieldEndZones, values)
		yardLine += 1000
	}
}

func drawFootballFieldLines(footballFieldLines *[]footballFieldLine, footballFieldOutsideLines *footballFieldLine, footballFieldHashLines *[]footballFieldLine, footballFieldEndZones *[]footballFieldLine, imd *imdraw.IMDraw) {

	for _, p := range *footballFieldLines {
		p.draw(imd)
	}
	for _, p := range *footballFieldHashLines {
		p.drawHashes(imd)
	}
	for _, p := range *footballFieldEndZones {
		p.draw(imd)
	}

	footballFieldOutsideLines.draw(imd)
}

func drawOffensivePlayersStartingPosition(imd *imdraw.IMDraw) {

	//Offensive Players
	// left wide receiver
	imd.Color = colornames.Black
	imd.Push(pixel.V(180, 145))
	imd.Circle(5, 2)

	// left guard
	imd.Color = colornames.Black
	imd.Push(pixel.V(240, 145))
	imd.Circle(5, 2)

	// center
	imd.Color = colornames.Black
	imd.Push(pixel.V(255, 140))
	imd.Push(pixel.V(265, 150))
	imd.Rectangle(2)

	// right guard
	imd.Color = colornames.Black
	imd.Push(pixel.V(280, 145))
	imd.Circle(5, 2)

	// QB
	imd.Color = colornames.Black
	imd.Push(pixel.V(260, 130))
	imd.Circle(5, 2)

	// Right Twins
	// inside twin
	imd.Color = colornames.Black
	imd.Push(pixel.V(400, 145))
	imd.Circle(5, 2)

	// outside twin
	imd.Color = colornames.Black
	imd.Push(pixel.V(415, 145))
	imd.Circle(5, 2)

}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
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

	// choose your football field background color from
	// https://pkg.go.dev/golang.org/x/image/colornames#pkg-variables
	win.Clear(colornames.Darkolivegreen)

	imd := imdraw.New(nil)

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
	rightSideLinePixel := 510

	defineFootballFieldLines(&footballFieldLines, &footballFieldOutsideLines,
		&footballFieldHashLines, &footballFieldEndZones, leftSideLinePixel, rightSideLinePixel)

	for !win.Closed() {
		win.Clear(colornames.Darkolivegreen)

		//imd.Clear()
		drawFootballFieldLines(&footballFieldLines, &footballFieldOutsideLines,
			&footballFieldHashLines, &footballFieldEndZones, imd)

		drawOffensivePlayersStartingPosition(imd)

		imd.Draw(win)

		// Add football
		//pic, err := loadPicture("./images/football2.png")
		//if err != nil {
		//	panic(err)
		//}
		//sprite := pixel.NewSprite(pic, pic.Bounds())
		//sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
