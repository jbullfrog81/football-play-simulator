package main

import (
	"fmt"
	"image"
	"image/color"
	"jbullfrog81/football-play-simulator/field"
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/playbook"
	"jbullfrog81/football-play-simulator/routes"
	"os"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
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

func DrawBuildPlaybookMenuSavedPlayerRoutes(imd *imdraw.IMDraw, win *pixelgl.Window, savedRoutes routes.OffensePlayRoutes, formationIteration int) {
	availableOffensiveFormations := formations.ReturnAllOffensiveTeamFormations()

	for i, _ := range savedRoutes.Routes {
		playbook.DrawOffensivePlayerPlayRoute(imd, availableOffensiveFormations.Formations[formationIteration].Players[i].Coordinates, savedRoutes.Routes[i], colornames.Gold)
	}

	imd.Draw(win)
}

func DrawSpecificOffensiveFormationHighlightOnePlayer(imd *imdraw.IMDraw, win *pixelgl.Window, formationIteration int, routeIteration int, playerHighlight int) {

	availableOffensiveFormations := formations.ReturnAllOffensiveTeamFormations()
	availableOffensiveRoutes := routes.ReturnAllOffensePlayRoutes()
	//currentFormation := availableOffensiveFormations.Formations[iteration]
	//for i, v := range availableOffensiveFormations.Formations {
	//	if i < 1 {

	DrawOffensivePlayersHighlightOnePlayer(imd, availableOffensiveFormations.Formations[formationIteration], playerHighlight, availableOffensiveRoutes.Routes[routeIteration])

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	playerTxt := text.New(pixel.V(600, 450), atlas)
	fmt.Fprintln(playerTxt, "Player Information: ")
	fmt.Fprintln(playerTxt, "Position - "+availableOffensiveFormations.Formations[formationIteration].Players[playerHighlight].Attributes.Position)

	playerTxt.Draw(win, pixel.IM)

	routeTxt := text.New(pixel.V(600, 400), atlas)
	fmt.Fprintln(routeTxt, "Route Information: ")
	fmt.Fprintln(routeTxt, "Route Name - "+availableOffensiveRoutes.Routes[routeIteration].RouteName)

	routeTxt.Draw(win, pixel.IM)

	basicTxt := text.New(pixel.V(600, 200), atlas)

	fmt.Fprintln(basicTxt, "Formation Information: ")
	fmt.Fprintln(basicTxt, "Name - "+availableOffensiveFormations.Formations[formationIteration].FormationName)
	fmt.Fprintln(basicTxt, "Snap Type - "+availableOffensiveFormations.Formations[formationIteration].SnapType)
	fmt.Fprintln(basicTxt, "Recievers - "+fmt.Sprint(availableOffensiveFormations.Formations[formationIteration].Receivers))
	fmt.Fprintln(basicTxt, "Running Backs - "+fmt.Sprint(availableOffensiveFormations.Formations[formationIteration].RunningBacks))

	basicTxt.Draw(win, pixel.IM)

}

// This function will draw the offensive team formation with a specific player highlighted in
// red and the route they are assigned
func DrawOffensivePlayersHighlightOnePlayer(imd *imdraw.IMDraw, team formations.OffenseTeamFormation, playerHighlight int, playerHighlightRoute routes.OffensePlayRoute) {

	for i, v := range team.Players {

		if i == playerHighlight {
			imd.Color = colornames.Red
			playbook.DrawOffensivePlayerPlayRoute(imd, v.Coordinates, playerHighlightRoute, colornames.Red)
		} else {
			imd.Color = v.Attributes.Color
		}
		imd.Push(pixel.V(v.Coordinates.MinX, v.Coordinates.MinY))
		imd.Circle(v.Attributes.Radius, v.Attributes.Thickness)
	}
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
	imdBuildOffensivePlaybook := imdraw.New(nil)

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

	// Build a new playbook
	var buildOffensivePlayBook playbook.PlayBook

	myTeamOffensivePlayBook = playbook.BuildDefaultOffensivePlayBook()

	playbook.SavePlayBookToFile(myTeamOffensivePlayBook)

	//Use this for moving the players during the play
	iteration := 0

	// Available Window States:
	// - paused
	// - running
	windowState := "paused"

	// Available Window Menus:
	// - RunOffensivePlaybook
	// - RunPlay
	// - OffensiveFormations
	// - BuildOffensivePlaybook
	windowMenu := "RunOffensivePlaybook"

	OffenseFormationIteration := 0

	OffensePlaybookPageNumber := 0

	OffenseRunPlayPlaybookPageNumber := 0

	myTeamOffensePlayBookRun := playbook.BuildDefaultOffensivePlayBook()

	drawSelectFormationIteration := 0

	// Available Build Offensive Playbook menu items:
	// - Formation
	// - Route
	BuildOffensivePlaybookMenuSelection := "Formation"

	var selectedFormation int
	//var selectedRoute int
	var selectedPlayerRoute routes.OffensePlayRoute

	//These are the temp player routes to build a page in our playbood
	var selectedPlayerRoutes routes.OffensePlayRoutes
	//set all the routes to a default of block
	for i := 0; i < 7; i++ {
		selectedPlayerRoutes.Routes = append(selectedPlayerRoutes.Routes, routes.DefineBlockRoute())
	}

	drawSelectPlayerIteration := 0
	drawSelectRouteIteration := 0
	var buildOffensivePlay playbook.OffensivePlay

	allOffensiveFormations := formations.ReturnAllOffensiveTeamFormations()
	allOffensiveRoutes := routes.ReturnAllOffensePlayRoutes()

	for !win.Closed() {

		if windowMenu == "RunPlay" {

			if win.JustPressed(pixelgl.KeyEscape) {

				//Reset the run play formation
				imdOffenseRunPlay.Clear()
				myTeamOffensePlayBookRun = playbook.BuildDefaultOffensivePlayBook()
				iteration = 0

				windowMenu = "RunOffensivePlaybook"
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
				playbook.DrawOffensiveRunPlayMenu(imdOffenseRunPlay, win, myTeamOffensePlayBookRun, OffenseRunPlayPlaybookPageNumber)

				win.Update()

			} else {

				win.Clear(colornames.Darkolivegreen)

				imdOffenseRunPlay.Clear()
				imdFootballField.Clear()

				field.DrawFootballFieldLines(imdFootballField, leftSideLinePixel, rightSideLinePixel)
				field.DrawFootballFieldYardNumbers(imdFootballField, win)
				imdFootballField.Draw(win)

				if iteration == 0 {
					formations.DrawOffensivePlayers(imdOffenseRunPlay, myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation)
				} else {
					for i, _ := range myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players {

						myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players[i] = formations.DrawOffensePlayerRunPlay(imdOffenseRunPlay, myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].PlayerRoutes[i], myTeamOffensePlayBookRun.OffensivePlays[OffenseRunPlayPlaybookPageNumber].Formation.Players[i], iteration)

					}
				}

				playbook.DrawOffensiveRunPlayMenu(imdOffenseRunPlay, win, myTeamOffensePlayBookRun, OffenseRunPlayPlaybookPageNumber)

				imdOffenseRunPlay.Draw(win)

				win.Update()

				iteration += 1
				//println("iteration is: ", iteration)
				//println("the windowState is:", windowState)
			}

			if frameTick != nil {
				<-frameTick.C
			}

		} else if windowMenu == "RunOffensivePlaybook" {

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
				windowMenu = "BuildOffensivePlaybook"
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
		} else if windowMenu == "BuildOffensivePlaybook" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenu = "RunPlay"
			}

			win.Clear(colornames.Darkolivegreen)

			imdBuildOffensivePlaybook.Clear()

			field.DrawFootballFieldYardNumbers(imdFootballField, win)

			imdFootballField.Draw(win)

			if BuildOffensivePlaybookMenuSelection == "Formation" {

				playbook.DrawBuildOffensivePlaybookMenuSelectFormation(imdBuildOffensivePlaybook, win, drawSelectFormationIteration)

				imdBuildOffensivePlaybook.Draw(win)

				if win.JustPressed(pixelgl.KeyEnter) {
					selectedFormation = drawSelectFormationIteration
					buildOffensivePlay.PlayName = "play 1"
					buildOffensivePlay.Formation = allOffensiveFormations.Formations[selectedFormation]
					//buildOffensivePlayBook.OffensivePlays = append(buildOffensivePlayBook.OffensivePlays, allOffensiveFormations.Formations[selectedFormation])
					//buildOffensivePlayBook.OffensivePlays[0].Formation = allOffensiveFormations.Formations[selectedFormation]
					fmt.Println("The selected formation is: ", selectedFormation)
					fmt.Println("the build playbook is:")
					fmt.Println(buildOffensivePlayBook)
					BuildOffensivePlaybookMenuSelection = "Route"
				}

				if win.JustPressed(pixelgl.KeyUp) {
					if drawSelectFormationIteration > 0 {
						drawSelectFormationIteration += -1
					}
				} else if win.JustPressed(pixelgl.KeyDown) {
					if drawSelectFormationIteration < len(allOffensiveFormations.Formations) {
						drawSelectFormationIteration += 1
					}
				}

			} else if BuildOffensivePlaybookMenuSelection == "Route" {

				playbook.DrawBuildOffensivePlaybookMenuSelectRoute(imdBuildOffensivePlaybook, win, drawSelectFormationIteration, drawSelectRouteIteration, drawSelectPlayerIteration)

				//Draw the routes that have alredy been selected
				DrawBuildPlaybookMenuSavedPlayerRoutes(imdBuildOffensivePlaybook, win, selectedPlayerRoutes, drawSelectFormationIteration)

				//Draw the currently selected player to select a route
				DrawSpecificOffensiveFormationHighlightOnePlayer(imdBuildOffensivePlaybook, win, drawSelectFormationIteration, drawSelectRouteIteration, drawSelectPlayerIteration)

				imdBuildOffensivePlaybook.Draw(win)

				if win.JustPressed(pixelgl.KeyEnter) {

					selectedPlayerRoute = allOffensiveRoutes.Routes[drawSelectRouteIteration]
					selectedPlayerRoutes.Routes[drawSelectPlayerIteration] = selectedPlayerRoute
					//buildOffensivePlayBook.OffensivePlays[0].PlayName = "Play 1"
					//buildOffensivePlayBook.OffensivePlays[0].PlayerRoutes[drawSelectPlayerIteration] = allOffensiveRoutes.Routes[drawSelectRouteIteration]

					fmt.Println("The selected player and route is: ", selectedPlayerRoute)
					fmt.Println("the build playbook is:")
					fmt.Println(buildOffensivePlayBook)
				}

				if win.JustPressed(pixelgl.KeyLeft) {
					if drawSelectPlayerIteration > 0 {
						drawSelectPlayerIteration += -1
					}
				} else if win.JustPressed(pixelgl.KeyRight) {
					if drawSelectPlayerIteration < 6 {
						drawSelectPlayerIteration += 1
					}
				}

				if win.JustPressed(pixelgl.KeyUp) {
					if drawSelectRouteIteration > 0 {
						drawSelectRouteIteration += -1
					}
				} else if win.JustPressed(pixelgl.KeyDown) {
					if drawSelectRouteIteration < len(allOffensiveRoutes.Routes)-1 {
						drawSelectRouteIteration += 1
					}
				}

				if win.JustPressed(pixelgl.KeyTab) {
					BuildOffensivePlaybookMenuSelection = "Confirmation"
				}

			} else if BuildOffensivePlaybookMenuSelection == "Confirmation" {

				//TODO build a confirmation screen:
				// Let the user know to use Enter to confirm adding the Play to the playbook
				// delete to return to editing the play.
				if win.JustPressed(pixelgl.KeyEnter) {
					BuildOffensivePlaybookMenuSelection = "Done"
				} else if win.JustPressed(pixelgl.KeyDelete) {
					BuildOffensivePlaybookMenuSelection = "Route"
				}

			} else if BuildOffensivePlaybookMenuSelection == "Done" {

				//TODO build a success screen:
				// Let the user know the play successfully added to the playbook
				// save playbook
				// reset everything to add another play to the playbook
				buildOffensivePlayBook.OffensivePlays = append(buildOffensivePlayBook.OffensivePlays, playbook.AddPlayBookPage(buildOffensivePlay.PlayName, buildOffensivePlay.Formation, selectedPlayerRoutes.Routes))

			}

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
