package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
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
	"github.com/gen2brain/dlgs"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

//Global variables
var (
	frameTick *time.Ticker
	fps       float64
	// TODO add the following:
	// - View Offensive Routes
	// - Defense
	mainMenuLookup = map[string]string{
		"Offensive Playbook":        "OffensivePlaybook",
		"View Offensive Formations": "ViewOffensiveFormations",
		"View Offensive Routes":     "ViewOffensiveRoutes",
		"Exit":                      "Exit",
	}
	mainMenuOptions []string

	// Available Window States:
	// - paused
	// - running
	windowState = "paused"

	// Available Window Menus:
	// - MainMenu
	// - RunOffensivePlaybook
	// - RunPlay
	// - OffensiveFormations
	// - BuildOffensivePlaybook
	windowMenu = "MainMenu"

	// This is to hold where the user was prior to going to the main menu
	windowMenuPrevious = "MainMenu"

	//Manual creation of a playbook
	myTeamOffensivePlayBook playbook.PlayBook

	// Build a new playbook
	buildOffensivePlayBook playbook.PlayBook

	//Loaded playbook from file
	loadedTeamOffensivePlayBook playbook.PlayBook

	loadedTeamOffensivePlayBookRun playbook.PlayBook

	loadedPlaybookFileName string
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

func mainMenu() (wMenu string) {

	selectedMenuItem, okSelected, err := dlgs.List("Main Menu", "Program Options:", mainMenuOptions)

	if err != nil {
		panic(err)
	}

	if !okSelected {
		return windowMenuPrevious
	} else {
		fmt.Println("okSelected is:")
		fmt.Println(okSelected)
		fmt.Println("Menu item selected is:")
		fmt.Println(selectedMenuItem)
		return mainMenuLookup[selectedMenuItem]
	}
}

func loadPlaybookFromFile(fileName string, playbook *playbook.PlayBook) {
	file, err := ioutil.ReadFile(fileName)
	//file, err := os.Open(fileName)
	//defer file.Close()

	if err != nil {
		//log.Fatalf("failed to open file")
		_, err := dlgs.Info("Info", "Unable to open file")
		if err != nil {
			panic(err)
		}
	}

	err = json.Unmarshal([]byte(file), playbook)
	if err != nil {
		panic(err)
	}

	//loadedTeamOffensivePlayBookRun = loadedTeamOffensivePlayBook
	//util.DeepCopy(loadedTeamOffensivePlayBook, loadedTeamOffensivePlayBookRun)

	dlgs.MessageBox("File loaded", "The playbook was successfully loaded")
}

func loadPlaybookMenu() (wMenu string, fileName string) {

	fileName, selectedOk, err := dlgs.File("Select file", "*.playbook", false)
	if err != nil {
		panic(err)
	}

	fmt.Println("The filename is:")
	fmt.Println(fileName)
	fmt.Println("The selected ok is:")
	fmt.Println(selectedOk)

	if !selectedOk {
		return windowMenuPrevious, ""
	} else {

		loadPlaybookFromFile(fileName, &loadedTeamOffensivePlayBook)

		loadPlaybookFromFile(fileName, &loadedTeamOffensivePlayBookRun)

		return "OffensivePlaybookLoaded", fileName

	}
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
	imdOffensePlaybookLoadedRunPlay := imdraw.New(nil)
	imdFootballField := imdraw.New(nil)
	imdOffensiveFormations := imdraw.New(nil)
	imdViewOffensiveRoutes := imdraw.New(nil)
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

	myTeamOffensivePlayBook = playbook.BuildDefaultOffensivePlayBook()

	playbook.SavePlayBookToFile(myTeamOffensivePlayBook)

	//Use this for moving the players during the play
	iteration := 0

	// Generate Main Menu items
	for key, _ := range mainMenuLookup {
		mainMenuOptions = append(mainMenuOptions, key)
	}

	// Generate Offensive Playbook Menu items
	offensivePlaybookMenuLookup := map[string]string{
		"Load Playbook":        "LoadPlaybook",
		"Create New Playbook":  "CreateOffensivePlaybook",
		"Use Default Playbook": "UseDefaultPlaybook",
		"Back to Main Menu":    "MainMenu",
	}

	var offensivePlaybookMenuOptions []string

	for key, _ := range offensivePlaybookMenuLookup {
		offensivePlaybookMenuOptions = append(offensivePlaybookMenuOptions, key)
	}

	// Generate Offensive Playbook Loaded Menu items
	offensivePlaybookLoadedMenuLookup := map[string]string{
		"Edit Playbook":                   "EditPlaybook",
		"Run Plays in Playbook":           "LoadedOffensivePlaybookRunPlays",
		"View Plays in Playbook":          "ViewLoadedOffensivePlaybook",
		"Back to Offensive Playbook Menu": "OffensivePlaybook",
		"Back to Main Menu":               "MainMenu",
	}

	var offensivePlaybookLoadedMenuOptions []string

	for key, _ := range offensivePlaybookLoadedMenuLookup {
		offensivePlaybookLoadedMenuOptions = append(offensivePlaybookLoadedMenuOptions, key)
	}

	ViewOffensiveRoutesIteration := 0

	OffenseFormationIteration := 0

	OffensePlaybookPageNumber := 0

	OffensePlaybookLoadedPageNumber := 0

	OffenseRunPlayPlaybookPageNumber := 0

	OffensePlaybookLoadedRunPlayPageNumber := 0

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
	//var buildOffensivePlay playbook.OffensivePlay
	buildOffensivePlay := playbook.ReturnEmptyOffensivePlay()

	allOffensiveFormations := formations.ReturnAllOffensiveTeamFormations()
	allOffensiveRoutes := routes.ReturnAllOffensePlayRoutes()
	// Used for the view all routes menu
	singlePlayerOffensiveFormation := formations.SetOffensiveTeamFormationShowRoutes()

	for !win.Closed() {

		if windowMenu == "MainMenu" {

			windowMenu = mainMenu()
			fmt.Println("The window menu is:")
			fmt.Println(windowMenu)

		} else if windowMenu == "Exit" {

			os.Exit(0)

		} else if windowMenu == "LoadPlaybook" {

			windowMenu, loadedPlaybookFileName = loadPlaybookMenu()

		} else if windowMenu == "OffensivePlaybookLoaded" {

			offensivePlaybookLoadedMenuSelection, okSelected, err := dlgs.List("Offensive Playbook menu", "Offensive Playbook Options:", offensivePlaybookLoadedMenuOptions)
			if err != nil {
				panic(err)
			}

			if !okSelected {
				windowMenu = windowMenuPrevious
			} else {
				fmt.Println("okSelected is:")
				fmt.Println(okSelected)
				fmt.Println("Item selected is:")
				fmt.Println(offensivePlaybookLoadedMenuSelection)
				windowMenu = offensivePlaybookLoadedMenuLookup[offensivePlaybookLoadedMenuSelection]
			}

		} else if windowMenu == "OffensivePlaybook" {

			playbookMenuSelection, okSelected, err := dlgs.List("Offensive Playbook menu", "Offensive Playbook Options:", offensivePlaybookMenuOptions)
			if err != nil {
				panic(err)
			}

			if !okSelected {
				windowMenu = windowMenuPrevious
			} else {
				fmt.Println("okSelected is:")
				fmt.Println(okSelected)
				fmt.Println("Item selected is:")
				fmt.Println(playbookMenuSelection)
				windowMenu = offensivePlaybookMenuLookup[playbookMenuSelection]
			}

		} else if windowMenu == "RunPlay" {

			if win.JustPressed(pixelgl.KeyEscape) {

				//Reset the run play formation
				imdOffenseRunPlay.Clear()
				myTeamOffensePlayBookRun = playbook.BuildDefaultOffensivePlayBook()
				iteration = 0
				windowMenuPrevious = "RunPlay"
				windowMenu = "MainMenu"
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

		} else if windowMenu == "LoadedOffensivePlaybookRunPlays" {

			if win.JustPressed(pixelgl.KeyEscape) {

				//Reset the run play formation
				//- must do a deep copy as there are slices
				//- we don't want to share pointers
				imdOffensePlaybookLoadedRunPlay.Clear()
				loadPlaybookFromFile(loadedPlaybookFileName, &loadedTeamOffensivePlayBookRun)
				//util.DeepCopy(loadedTeamOffensivePlayBook, loadedTeamOffensivePlayBookRun)
				iteration = 0
				windowMenuPrevious = "LoadedOffensivePlaybookRunPlays"
				windowMenu = "OffensivePlaybookLoaded"
			}

			// restart the play when pressing enter
			if win.JustPressed(pixelgl.KeyEnter) {

				//Reset the run play formation
				//- must do a deep copy as there are slices
				//- we don't want to share pointers
				//util.DeepCopy(loadedTeamOffensivePlayBook, loadedTeamOffensivePlayBookRun)
				loadPlaybookFromFile(loadedPlaybookFileName, &loadedTeamOffensivePlayBookRun)
				iteration = 0
			}

			if win.JustPressed(pixelgl.KeySpace) {
				if windowState == "paused" {
					windowState = "running"
				} else {
					windowState = "paused"
				}
			}

			if win.JustPressed(pixelgl.KeyRight) && OffensePlaybookLoadedRunPlayPageNumber < (len(loadedTeamOffensivePlayBook.OffensivePlays)-1) {
				OffensePlaybookLoadedRunPlayPageNumber += 1
				iteration = 0
				//loadedTeamOffensivePlayBookRun = loadedTeamOffensivePlayBook
			}

			if win.JustPressed(pixelgl.KeyLeft) && OffensePlaybookLoadedRunPlayPageNumber > 0 {
				OffensePlaybookLoadedRunPlayPageNumber -= 1
				iteration = 0
				//loadedTeamOffensivePlayBookRun = loadedTeamOffensivePlayBook
			}

			if windowState == "paused" {

				win.Clear(colornames.Darkolivegreen)
				imdOffensePlaybookLoadedRunPlay.Clear()
				imdFootballField.Clear()

				field.DrawFootballFieldLines(imdFootballField, leftSideLinePixel, rightSideLinePixel)
				field.DrawFootballFieldYardNumbers(imdFootballField, win)
				imdFootballField.Draw(win)

				for i, _ := range loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].Formation.Players {

					formations.DrawOffensePlayerRunPlay(imdOffensePlaybookLoadedRunPlay, loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].PlayerRoutes[i], loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].Formation.Players[i], iteration)

				}

				imdOffensePlaybookLoadedRunPlay.Draw(win)
				playbook.DrawOffensiveRunPlayMenu(imdOffensePlaybookLoadedRunPlay, win, loadedTeamOffensivePlayBookRun, OffensePlaybookLoadedRunPlayPageNumber)

				win.Update()

			} else {

				win.Clear(colornames.Darkolivegreen)

				imdOffensePlaybookLoadedRunPlay.Clear()
				imdFootballField.Clear()

				field.DrawFootballFieldLines(imdFootballField, leftSideLinePixel, rightSideLinePixel)
				field.DrawFootballFieldYardNumbers(imdFootballField, win)
				imdFootballField.Draw(win)

				if iteration == 0 {
					formations.DrawOffensivePlayers(imdOffensePlaybookLoadedRunPlay, loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].Formation)
				} else {
					for i, _ := range loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].Formation.Players {

						loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].Formation.Players[i] = formations.DrawOffensePlayerRunPlay(imdOffensePlaybookLoadedRunPlay, loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].PlayerRoutes[i], loadedTeamOffensivePlayBookRun.OffensivePlays[OffensePlaybookLoadedRunPlayPageNumber].Formation.Players[i], iteration)

					}
				}

				playbook.DrawOffensiveRunPlayMenu(imdOffensePlaybookLoadedRunPlay, win, loadedTeamOffensivePlayBookRun, OffensePlaybookLoadedRunPlayPageNumber)

				imdOffensePlaybookLoadedRunPlay.Draw(win)

				win.Update()

				iteration += 1
				println("iteration is: ", iteration)
				println("the windowState is:", windowState)
			}

			if frameTick != nil {
				<-frameTick.C
			}

		} else if windowMenu == "RunOffensivePlaybook" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenuPrevious = "RunOffensivePlaybook"
				windowMenu = "MainMenu"
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

		} else if windowMenu == "ViewLoadedOffensivePlaybook" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenuPrevious = "ViewLoadedOffensivePlaybook"
				windowMenu = "OffensivePlaybookLoaded"
			}

			// need to utilize the max number of plays instead of 1
			if win.JustPressed(pixelgl.KeyRight) && OffensePlaybookLoadedPageNumber < (len(loadedTeamOffensivePlayBook.OffensivePlays)-1) {
				OffensePlaybookLoadedPageNumber += 1
			}

			if win.JustPressed(pixelgl.KeyLeft) && OffensePlaybookLoadedPageNumber > 0 {
				OffensePlaybookLoadedPageNumber -= 1
			}

			win.Clear(colornames.Darkolivegreen)

			imdFootballField.Draw(win)

			field.DrawFootballFieldYardNumbers(imdFootballField, win)

			imdOffensivePlayBook.Clear()

			playbook.DrawOffensivePlayBookPage(imdOffensivePlayBook, win, OffensePlaybookLoadedPageNumber, loadedTeamOffensivePlayBook)

			playbook.DrawOffensivePlayBookMenu(imdOffensivePlayBook, win, loadedTeamOffensivePlayBook, OffensePlaybookLoadedPageNumber)

			imdOffensivePlayBook.Draw(win)

			win.Update()

			if frameTick != nil {
				<-frameTick.C
			}

		} else if windowMenu == "ViewOffensiveRoutes" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenuPrevious = "ViewOffensiveRoutes"
				windowMenu = "MainMenu"
			}

			if win.JustPressed(pixelgl.KeyLeft) && ViewOffensiveRoutesIteration > 0 {
				ViewOffensiveRoutesIteration -= 1
			}

			if win.JustPressed(pixelgl.KeyRight) && ViewOffensiveRoutesIteration < len(allOffensiveRoutes.Routes)-1 {
				ViewOffensiveRoutesIteration += 1
			}

			win.Clear(colornames.Darkolivegreen)

			imdViewOffensiveRoutes.Clear()
			formations.DrawOffensivePlayers(imdViewOffensiveRoutes, singlePlayerOffensiveFormation)
			playbook.DrawOffensivePlayerPlayRoute(imdViewOffensiveRoutes, singlePlayerOffensiveFormation.Players[0].Coordinates, allOffensiveRoutes.Routes[ViewOffensiveRoutesIteration], colornames.Gold)
			routes.DrawOffensiveRoutesMenu(imdViewOffensiveRoutes, win, allOffensiveRoutes.Routes[ViewOffensiveRoutesIteration])

			field.DrawFootballFieldYardNumbers(imdFootballField, win)

			imdFootballField.Draw(win)
			imdViewOffensiveRoutes.Draw(win)
			win.Update()

			if frameTick != nil {
				<-frameTick.C
			}

		} else if windowMenu == "ViewOffensiveFormations" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenuPrevious = "ViewOffensiveFormations"
				windowMenu = "MainMenu"
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
		} else if windowMenu == "CreateOffensivePlaybook" {

			if win.JustPressed(pixelgl.KeyEscape) {
				windowMenuPrevious = "CreateOffensivePlaybook"
				windowMenu = "MainMenu"
			}

			if buildOffensivePlayBook.PlayBookName == "" {

				playbookName, okSelected, err := dlgs.Entry("Enter Playbook Name", "Name of the new playbook:", "default")
				if err != nil {
					panic(err)
				}

				if okSelected {
					buildOffensivePlayBook.PlayBookName = playbookName
				}

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

					playName, okSelected, err := dlgs.Entry("Enter Play Name", "Name of the play:", "")
					if err != nil {
						panic(err)
					}

					if okSelected {
						buildOffensivePlay.PlayName = playName
						buildOffensivePlay.Formation = allOffensiveFormations.Formations[selectedFormation]
						//buildOffensivePlayBook.OffensivePlays = append(buildOffensivePlayBook.OffensivePlays, 	allOffensiveFormations.Formations[selectedFormation])
						//buildOffensivePlayBook.OffensivePlays[0].Formation = allOffensiveFormations.Formations	[selectedFormation]
						fmt.Println("The selected formation is: ", selectedFormation)
						fmt.Println("the build playbook is:")
						fmt.Println(buildOffensivePlayBook)
						BuildOffensivePlaybookMenuSelection = "Route"
					}
				}

				if win.JustPressed(pixelgl.KeyUp) {
					if drawSelectFormationIteration > 0 {
						drawSelectFormationIteration += -1
					}
				} else if win.JustPressed(pixelgl.KeyDown) {
					if drawSelectFormationIteration < len(allOffensiveFormations.Formations)-1 {
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

				hasConfirmed, err := dlgs.Question("Done creating the play?", "Are you sure you are done creating this play?", true)
				if err != nil {
					panic(err)
				}

				if hasConfirmed {
					BuildOffensivePlaybookMenuSelection = "Done"
				} else {
					BuildOffensivePlaybookMenuSelection = "Route"
				}

				// Let the user know to use Enter to confirm adding the Play to the playbook
				// delete to return to editing the play.
				//if win.JustPressed(pixelgl.KeyEnter) {
				//	BuildOffensivePlaybookMenuSelection = "Done"
				//} else if win.JustPressed(pixelgl.KeyTab) {
				//	BuildOffensivePlaybookMenuSelection = "Route"
				//}

				//playbook.DrawBuildOffensivePlaybookMenuDoneConfirmation(imdBuildOffensivePlaybook, win)

			} else if BuildOffensivePlaybookMenuSelection == "Done" {

				//TODO build a success screen:
				// Let the user know the play successfully added to the playbook
				// save playbook
				// reset everything to add another play to the playbook
				buildOffensivePlayBook.OffensivePlays = append(buildOffensivePlayBook.OffensivePlays, playbook.AddPlayBookPage(buildOffensivePlay.PlayName, buildOffensivePlay.Formation, selectedPlayerRoutes.Routes))

				//Save the playbook to disk
				playbook.SavePlayBookToFile(buildOffensivePlayBook)

				//reset all the settings for editing/building a new play
				for i := 0; i < 7; i++ {
					selectedPlayerRoutes.Routes[i] = routes.DefineBlockRoute()
				}
				//reset the play back to null
				buildOffensivePlay = playbook.ReturnEmptyOffensivePlay()
				drawSelectPlayerIteration = 0
				drawSelectRouteIteration = 0

				hasConfirmed, err := dlgs.Question("Create another play?", "Do you want to create another play?", true)
				if err != nil {
					panic(err)
				}

				if hasConfirmed {
					BuildOffensivePlaybookMenuSelection = "Formation"
				} else {
					windowMenu = "OffensivePlaybook"
				}

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
