package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

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

	win.Clear(colornames.Darkolivegreen)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
