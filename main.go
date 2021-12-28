package main

import (
	"fmt"
	"image"
	"image/color"
	formations "jbullfrog81/football-play-simulator/formations"
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

//rectangles to make the football field 5 yard lines
type footballFieldLine struct {
	minX  float64
	minY  float64
	maxX  float64
	maxY  float64
	color color.Color
}

type offensePlayerPosition struct {
	thickness float64
	radius    float64
	minX      float64
	minY      float64
	maxX      float64
	maxY      float64
	color     color.Color
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

func drawFootballFieldYardNumbers(imd *imdraw.IMDraw, win *pixelgl.Window) {

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

func drawOffensivePlayers(imd *imdraw.IMDraw, team *formations.OffenseTeamFormation) {

	imd.Color = team.Player1.Attributes.Color
	imd.Push(pixel.V(team.Player1.Coordinates.MinX, team.Player1.Coordinates.MinY))
	imd.Circle(team.Player1.Attributes.Radius, team.Player1.Attributes.Thickness)

	imd.Color = team.Player2.Attributes.Color
	imd.Push(pixel.V(team.Player2.Coordinates.MinX, team.Player2.Coordinates.MinY))
	imd.Circle(team.Player2.Attributes.Radius, team.Player2.Attributes.Thickness)

	imd.Color = team.Player3.Attributes.Color
	imd.Push(pixel.V(team.Player3.Coordinates.MinX, team.Player3.Coordinates.MinY))
	imd.Circle(team.Player3.Attributes.Radius, team.Player3.Attributes.Thickness)

	imd.Color = team.Player4.Attributes.Color
	imd.Push(pixel.V(team.Player4.Coordinates.MinX, team.Player4.Coordinates.MinY))
	imd.Circle(team.Player4.Attributes.Radius, team.Player4.Attributes.Thickness)

	imd.Color = team.Player5.Attributes.Color
	imd.Push(pixel.V(team.Player5.Coordinates.MinX, team.Player5.Coordinates.MinY))
	imd.Circle(team.Player5.Attributes.Radius, team.Player5.Attributes.Thickness)

	imd.Color = team.Player6.Attributes.Color
	imd.Push(pixel.V(team.Player6.Coordinates.MinX, team.Player6.Coordinates.MinY))
	imd.Circle(team.Player6.Attributes.Radius, team.Player6.Attributes.Thickness)

	imd.Color = team.Player7.Attributes.Color
	imd.Push(pixel.V(team.Player7.Coordinates.MinX, team.Player7.Coordinates.MinY))
	imd.Circle(team.Player7.Attributes.Radius, team.Player7.Attributes.Thickness)

}

func drawOffensePlayerRunPlay(imd *imdraw.IMDraw, route *routes.OffensePlayRoute, playerPosition *formations.OffensePlayer, iteration int) {

	println("starting draw offense run play")
	if iteration < len(route.MinX) {
		println("inside iteration loop")
		playerPosition.Coordinates.MinX += route.MinX[iteration]
		playerPosition.Coordinates.MinY += route.MinY[iteration]
		playerPosition.Coordinates.MaxX += route.MaxX[iteration]
		playerPosition.Coordinates.MaxY += route.MaxY[iteration]
	}

	imd.Color = playerPosition.Attributes.Color
	imd.Push(pixel.V(playerPosition.Coordinates.MinX, playerPosition.Coordinates.MinY))
	imd.Circle(playerPosition.Attributes.Radius, playerPosition.Attributes.Thickness)

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

func BuildDefaultOffensivePlayBook(defaultPlaybook *playbook.PlayBook) {

	var setupPlayerRoutes []routes.OffensePlayRoute

	//List of all the routes:
	//-----------------------
	//block = routes.DefineBlockRoute()
	//fiveYardOut = routes.DefineLeftOutFiveYardRoute()
	//tenYardOut = routes.DefineLeftOutTenYardRoute()
	//sevenYardOutAndUp = routes.DefineLeftOutAndUpSevenYardRoute()
	//tenYardPost = routes.DefineLeftPostTenYardRoute()
	//fiveYardWhip = routes.DefineWhipFiveYardRoute()
	//sevenYardWhip = routes.DefineLeftWhipSevenYardRoute()

	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineGoRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftOutFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftOutAndUpSevenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftWhipFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftPostTenYardRoute())

	playbook.AddPlayBookPage(defaultPlaybook, "Bunch Left In Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes)

	setupPlayerRoutes = nil

	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineGoRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutAndUpSevenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightWhipFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightPostTenYardRoute())

	playbook.AddPlayBookPage(defaultPlaybook, "Bunch Right In Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes)

}

func drawSpecificOffensiveFormations(imd *imdraw.IMDraw, win *pixelgl.Window, iteration int) {

	availableOffensiveFormations := formations.ReturnAllOffensiveTeamFormations()
	currentFormation := availableOffensiveFormations.Formations[iteration]
	//for i, v := range availableOffensiveFormations.Formations {
	//	if i < 1 {

	drawOffensivePlayers(imd, &currentFormation)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(600, 400), atlas)

	fmt.Fprintln(basicTxt, "Name: "+currentFormation.FormationName)
	fmt.Fprintln(basicTxt, "Snap Type: "+currentFormation.SnapType)
	fmt.Fprintln(basicTxt, "Recievers: "+fmt.Sprint(currentFormation.Receivers))
	fmt.Fprintln(basicTxt, "Running Backs: "+fmt.Sprint(currentFormation.RunningBacks))

	basicTxt.Draw(win, pixel.IM)

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

	imd := imdraw.New(nil)
	imd2 := imdraw.New(nil)

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

	drawFootballFieldLines(&footballFieldLines, &footballFieldOutsideLines,
		&footballFieldHashLines, &footballFieldEndZones, imd)

	drawFootballFieldYardNumbers(imd, win)

	//Manual creation of a playbook
	var myTeamOffensivePlayBook playbook.PlayBook

	myTeamOffensivePlayBook.PlayBookName = "Default"

	BuildDefaultOffensivePlayBook(&myTeamOffensivePlayBook)

	playbook.SavePlayBookToFile(&myTeamOffensivePlayBook)

	//myTeamOffenseInitialFormation = formations.SetOffensiveTeamFormationShotgunTripsLeft()

	//drawOffensivePlayers(imd, &myTeamOffensivePlayBook.OffensivePlays[0].Formation)

	//Use this for moving the players during the play
	myTeamOffenseRunPlayFormation := myTeamOffensivePlayBook.OffensivePlays[0].Formation

	//var fiveYardOut routes.OffensePlayRoute
	//var tenYardOut routes.OffensePlayRoute
	//var sevenYardOutAndUp routes.OffensePlayRoute
	//var tenYardPost routes.OffensePlayRoute
	//var fiveYardWhip routes.OffensePlayRoute
	//var sevenYardWhip routes.OffensePlayRoute
	//var block routes.OffensePlayRoute

	//block = routes.DefineBlockRoute()
	//fiveYardOut = routes.DefineLeftOutFiveYardRoute()
	//tenYardOut = routes.DefineLeftOutTenYardRoute()
	//sevenYardOutAndUp = routes.DefineLeftOutAndUpSevenYardRoute()
	//tenYardPost = routes.DefineLeftPostTenYardRoute()
	//fiveYardWhip = routes.DefineWhipFiveYardRoute()
	//sevenYardWhip = routes.DefineLeftWhipSevenYardRoute()

	//var rightTwin offensePlayerPosition
	//var leftTwin offensePlayerPosition
	//var leftWideReceiver offensePlayerPosition

	//defineOffensivePlayerPosition(&rightTwin, 2.0, 5.0, 400.0, 145.0, 400.0, 145.0, colornames.Black)
	//defineOffensivePlayerPosition(&leftTwin, 2.0, 5.0, 415.0, 145.0, 415.0, 145.0, colornames.Black)
	//defineOffensivePlayerPosition(&leftWideReceiver, 2.0, 5.0, 180.0, 145.0, 180.0, 145.0, colornames.Black)

	iteration := 0

	windowState := "paused"

	OffenseFormationIteration := 0

	for !win.Closed() {

		if windowState == "running" {
			win.Clear(colornames.Darkolivegreen)

			// restart the play when pressing enter
			if win.JustPressed(pixelgl.KeyEnter) {
				//redraw the initial play formation
				drawOffensivePlayers(imd, &myTeamOffensivePlayBook.OffensivePlays[0].Formation)

				//reset the run play formation
				myTeamOffenseRunPlayFormation = myTeamOffensivePlayBook.OffensivePlays[0].Formation
				iteration = 0
			}

			if win.JustPressed(pixelgl.KeySpace) {
				windowState = "paused"
			}

			imd.Clear()

			drawFootballFieldLines(&footballFieldLines, &footballFieldOutsideLines,
				&footballFieldHashLines, &footballFieldEndZones, imd)

			drawFootballFieldYardNumbers(imd, win)

			if iteration == 0 {
				drawOffensivePlayers(imd, &myTeamOffensivePlayBook.OffensivePlays[0].Formation)
			}

			drawOffensePlayerRunPlay(imd, &myTeamOffensivePlayBook.OffensivePlays[0].PlayerRoutes[0], &myTeamOffenseRunPlayFormation.Player1, iteration)
			drawOffensePlayerRunPlay(imd, &myTeamOffensivePlayBook.OffensivePlays[0].PlayerRoutes[1], &myTeamOffenseRunPlayFormation.Player2, iteration)
			drawOffensePlayerRunPlay(imd, &myTeamOffensivePlayBook.OffensivePlays[0].PlayerRoutes[2], &myTeamOffenseRunPlayFormation.Player3, iteration)
			drawOffensePlayerRunPlay(imd, &myTeamOffensivePlayBook.OffensivePlays[0].PlayerRoutes[3], &myTeamOffenseRunPlayFormation.Player4, iteration)
			drawOffensePlayerRunPlay(imd, &myTeamOffensivePlayBook.OffensivePlays[0].PlayerRoutes[4], &myTeamOffenseRunPlayFormation.Player5, iteration)
			drawOffensePlayerRunPlay(imd, &myTeamOffensivePlayBook.OffensivePlays[0].PlayerRoutes[5], &myTeamOffenseRunPlayFormation.Player6, iteration)
			drawOffensePlayerRunPlay(imd, &myTeamOffensivePlayBook.OffensivePlays[0].PlayerRoutes[6], &myTeamOffenseRunPlayFormation.Player7, iteration)

			imd.Draw(win)

			//drawOffenseRunPlay(imd, increment)

			// Add football
			//pic, err := loadPicture("./images/football2.png")
			//if err != nil {
			//	panic(err)
			//}
			//sprite := pixel.NewSprite(pic, pic.Bounds())
			//sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

			win.Update()

			iteration += 1
			println("iteration is: ", iteration)
			println("the windowState is:", windowState)

			if frameTick != nil {
				<-frameTick.C
			}

		} else {
			//when paused we have to send signals to screen or the window will bomb out

			if win.JustPressed(pixelgl.KeyDown) && OffenseFormationIteration > 0 {
				OffenseFormationIteration -= 1
			}

			if win.JustPressed(pixelgl.KeyUp) && OffenseFormationIteration < 9 {
				OffenseFormationIteration += 1
			}

			win.Clear(colornames.Darkolivegreen)

			imd2.Clear()
			drawSpecificOffensiveFormations(imd2, win, OffenseFormationIteration)

			drawFootballFieldYardNumbers(imd, win)

			imd.Draw(win)
			imd2.Draw(win)
			win.Update()

			if win.JustPressed(pixelgl.MouseButtonLeft) {
				windowState = "running"
			}

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
