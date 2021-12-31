package field

import (
	"fmt"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
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

func DrawFootballFieldYardNumbers(imd *imdraw.IMDraw, win *pixelgl.Window) {

	//Right side of the field yard line numbers
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	basicTxt := text.New(pixel.V(490, 180), atlas)

	fmt.Fprintln(basicTxt, "1 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 280), atlas)

	fmt.Fprintln(basicTxt, "2 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 380), atlas)

	fmt.Fprintln(basicTxt, "3 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 480), atlas)

	fmt.Fprintln(basicTxt, "4 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 580), atlas)

	fmt.Fprintln(basicTxt, "5 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 680), atlas)

	fmt.Fprintln(basicTxt, "4 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 780), atlas)

	fmt.Fprintln(basicTxt, "3 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 880), atlas)

	fmt.Fprintln(basicTxt, "2 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	basicTxt = text.New(pixel.V(490, 980), atlas)

	fmt.Fprintln(basicTxt, "1 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, 1.55))

	//Left side of the field yard line numbers
	basicTxt = text.New(pixel.V(25, 220), atlas)

	fmt.Fprintln(basicTxt, "1 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 320), atlas)

	fmt.Fprintln(basicTxt, "2 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 420), atlas)

	fmt.Fprintln(basicTxt, "3 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 520), atlas)

	fmt.Fprintln(basicTxt, "4 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 620), atlas)

	fmt.Fprintln(basicTxt, "5 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 720), atlas)

	fmt.Fprintln(basicTxt, "4 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 820), atlas)

	fmt.Fprintln(basicTxt, "3 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 920), atlas)

	fmt.Fprintln(basicTxt, "2 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

	basicTxt = text.New(pixel.V(25, 1020), atlas)

	fmt.Fprintln(basicTxt, "1 0")

	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2).Rotated(basicTxt.Orig, -1.55))

}

func DrawFootballFieldLines(imd *imdraw.IMDraw, leftSideLinePixel int, rightSideLinePixel int) {

	//footballFieldLines *[]footballFieldLine, footballFieldOutsideLines *footballFieldLine, footballFieldHashLines *[]footballFieldLine, footballFieldEndZones *[]footballFieldLine,

	var footballFieldOutsideLines footballFieldLine
	var footballFieldLines []footballFieldLine
	var footballFieldHashLines []footballFieldLine
	var footballFieldEndZones []footballFieldLine

	defineFootballFieldLines(&footballFieldLines, &footballFieldOutsideLines,
		&footballFieldHashLines, &footballFieldEndZones, leftSideLinePixel, rightSideLinePixel)

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
}
