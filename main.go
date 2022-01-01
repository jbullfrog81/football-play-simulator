package main

import (
	"image"
	"image/color"
	"jbullfrog81/football-play-simulator/field"
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/playbook"
	"os"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//Global variables
var (
	frameTick *time.Ticker
	fps       float64
)

type offensePlayerPosition struct {
	thickness float64
	radius    float64
	minX      float64
	minY      float64
	maxX      float64
	maxY      float64
	color     color.Color
}

func defineOffensivePlayerPosition(playerPosition *offensePlayerPosition, thickness float64,
	radius float64, minX float64, minY float64, maxX float64, maxY float64, color color.RGBA) {
	playerPosition.thickness = thickness
	playerPosition.radius = radius
	playerPosition.minX = minX
	playerPosition.minY = minY
	playerPosition.maxX = maxX
	playerPosition.maxY = maxY
	playerPosition.color = color
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

	// Set the frames per second to be 60
	setFPS(60)

	// smooth out the graphics i.e. don't be pixelated
	win.SetSmooth(true)

	// choose your football field background color from
	// https://pkg.go.dev/golang.org/x/image/colornames#pkg-variables
	win.Clear(colornames.Darkolivegreen)

	//imdMenu := imdraw.New(nil)
	imdOffenseRunPlay := imdraw.New(nil)
	imdFootballField := imdraw.New(nil)
	imdOffensiveFormations := imdraw.New(nil)
	imdOffensivePlayBook := imdraw.New(nil)

	// The lines on the football field:
	// 1 pixel = 3.6 inches
	// 10 pixels = 1 yard
	// 500 pixels = 50 yard wide field
	// 1000 pixels = 100 yard long field

	leftSideLinePixel := 10
	rightSideLinePixel := 510

	field.DrawFootballFieldLines(imdFootballField, leftSideLinePixel, rightSideLinePixel)

	field.DrawFootballFieldYardNumbers(imdFootballField, win)

	//Manual creation of a playbook
	var myTeamOffensivePlayBook playbook.PlayBook

	myTeamOffensivePlayBook = playbook.BuildDefaultOffensivePlayBook()

	playbook.SavePlayBookToFile(myTeamOffensivePlayBook)

	//Use this for moving the players during the play

	iteration := 0

	// Available Window States:
	// - paused
	// - running
	windowState := "running"

	// Available Window Menus:
	// - OffensivePlaybook
	// - RunPlay
	// - OffensiveFormations
	windowMenu := "OffensivePlaybook"

	OffenseFormationIteration := 0

	OffensePlaybookPageNumber := 0

	OffenseRunPlayPlaybookPageNumber := 0

	myTeamOffensePlayBookRun := playbook.BuildDefaultOffensivePlayBook()

	for !win.Closed() {

		if windowMenu == "RunPlay" {

			if win.JustPressed(pixelgl.KeyEscape) {

				//Reset the run play formation
				imdOffenseRunPlay.Clear()
				myTeamOffensePlayBookRun = playbook.BuildDefaultOffensivePlayBook()
				iteration = 0

				windowMenu = "OffensivePlaybook"
			}

			// restart the play when pressing enter
			if win.JustPressed(pixelgl.KeyEnter) {

				//reset the run play formation
				myTeamOffensePlayBookRun = playbook.BuildDefaultOffensivePlayBook()
				iteration = 0
			}

			if win.JustPressed(pixelgl.KeySpace) {
				if windowState == "paused" {
					windowState = "running"
				} else {
					windowState = "paused"
				}
			}

			if windowState == "paused" {

				win.Clear(colornames.Darkolivegreen)
				imdOffenseRunPlay.Clear()
				imdFootballField.Clear()

				field.DrawFootballFieldLines(imdFootballField, leftSideLinePixel, rightSideLinePixel)
				field.DrawFootballFieldYardNumbers(imdFootballField, win)
				imdFootballField.Draw(win)

				for i, _ := range myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players {

					formations.DrawOffensePlayerRunPlay(imdOffenseRunPlay, myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].PlayerRoutes[i], myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players[i], iteration)

				}

				imdOffenseRunPlay.Draw(win)
				playbook.DrawOffensiveRunPlayMenu(imdOffenseRunPlay, win, myTeamOffensivePlayBook, OffenseRunPlayPlaybookPageNumber)

				win.Update()

			} else {

				win.Clear(colornames.Darkolivegreen)

				imdOffenseRunPlay.Clear()
				imdFootballField.Clear()

				field.DrawFootballFieldLines(imdFootballField, leftSideLinePixel, rightSideLinePixel)
				field.DrawFootballFieldYardNumbers(imdFootballField, win)
				imdFootballField.Draw(win)

				if iteration == 0 {
					formations.DrawOffensivePlayers(imdOffenseRunPlay, myTeamOffensivePlayBook.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation)
				} else {
					for i, _ := range myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players {

						myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players[i] = formations.DrawOffensePlayerRunPlay(imdOffenseRunPlay, myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].PlayerRoutes[i], myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players[i], iteration)

					}
				}

				playbook.DrawOffensiveRunPlayMenu(imdOffenseRunPlay, win, myTeamOffensivePlayBook, OffenseRunPlayPlaybookPageNumber)

				imdOffenseRunPlay.Draw(win)

				win.Update()

				iteration += 1
				//println("iteration is: ", iteration)
				//println("the windowState is:", windowState)
			}

			if frameTick != nil {
				<-frameTick.C
			}

		} else if windowMenu == "OffensivePlaybook" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenu = "OffensiveFormations"
			}

			if win.JustPressed(pixelgl.KeyRight) && OffensePlaybookPageNumber < 1 {
				OffensePlaybookPageNumber += 1
			}

			if win.JustPressed(pixelgl.KeyLeft) && OffensePlaybookPageNumber > 0 {
				OffensePlaybookPageNumber -= 1
			}

			win.Clear(colornames.Darkolivegreen)

			imdFootballField.Draw(win)

			field.DrawFootballFieldYardNumbers(imdFootballField, win)

			imdOffensivePlayBook.Clear()

			playbook.DrawOffensivePlayBookPage(imdOffensivePlayBook, win, OffensePlaybookPageNumber, myTeamOffensivePlayBook)

			playbook.DrawOffensivePlayBookMenu(imdOffensivePlayBook, win, myTeamOffensivePlayBook, OffensePlaybookPageNumber)

			imdOffensivePlayBook.Draw(win)

			win.Update()

			if frameTick != nil {
				<-frameTick.C
			}

		} else if windowMenu == "OffensiveFormations" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenu = "RunPlay"
			}

			if win.JustPressed(pixelgl.KeyDown) && OffenseFormationIteration > 0 {
				OffenseFormationIteration -= 1
			}

			if win.JustPressed(pixelgl.KeyUp) && OffenseFormationIteration < 9 {
				OffenseFormationIteration += 1
			}

			win.Clear(colornames.Darkolivegreen)

			imdOffensiveFormations.Clear()
			formations.DrawSpecificOffensiveFormation(imdOffensiveFormations, win, OffenseFormationIteration)
			formations.DrawOffensiveFormatonsMenu(imdOffensiveFormations, win)

			field.DrawFootballFieldYardNumbers(imdFootballField, win)

			imdFootballField.Draw(win)
			imdOffensiveFormations.Draw(win)
			win.Update()

			if frameTick != nil {
				<-frameTick.C
			}
		}
	}
}

func setFPS(fps int) {
	if fps <= 0 {
		frameTick = nil
	} else {
		frameTick = time.NewTicker(time.Second / time.Duration(fps))
	}
}

func main() {
	pixelgl.Run(run)
}
