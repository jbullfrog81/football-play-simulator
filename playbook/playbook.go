package playbook

import (
	"fmt"
	"image/color"
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/routes"
	"strconv"

	"encoding/json"
	"io/ioutil"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"

	"github.com/jung-kurt/gofpdf"
)

type PlayBook struct {
	PlayBookName   string
	OffensivePlays []OffensivePlay
}

type OffensivePlay struct {
	PlayName     string
	Formation    formations.OffenseTeamFormation
	PlayerRoutes []routes.OffensePlayRoute
}

func AddPlayBookPage(playName string, playFormation formations.OffenseTeamFormation, playRoutes []routes.OffensePlayRoute) (playBookPage OffensivePlay) {

	playBookPage.PlayName = playName
	playBookPage.Formation = playFormation
	playBookPage.PlayerRoutes = playRoutes

	//playBook.OffensivePlays = append(playBook.OffensivePlays, playBookPage)

	return playBookPage

}

func SavePlayBookToFile(playBook PlayBook, filename string) {

	file, _ := json.MarshalIndent(playBook, "", " ")
	if filename == "" {
		_ = ioutil.WriteFile(playBook.PlayBookName+".playbook", file, 0644)
	} else {
		_ = ioutil.WriteFile(filename, file, 0644)
	}

}

func CreateOffensivePlaybookPdf(pdf *gofpdf.Fpdf, offensivePlaybook PlayBook) {

	var xCurrent float64
	var yCurrent float64
	var playerCurrentLocationX float64
	var playerCurrentLocationY float64
	var playerStartingLocationX float64
	var playerStartingLocationY float64
	//var playerNextLocationX float64
	//var playerNextLocationY float64
	var footballLocationX float64
	var footballLocationY float64
	var xNew float64
	var yNew float64
	var copyRouteX []float64
	var copyRouteY []float64
	var scaledRouteX []float64
	var scaledRouteY []float64
	//var scaledRouteTempIteration int

	//how much should we scale down
	scaleFactor := .35

	playerCircleRadius := 3.0

	scaledRouteTempIteration := 0.0

	footballOrigLocationX := 70.0
	//footballOrigLocationY = 100.0
	//footballLocationX = footballOrigLocationX
	//footballLocationY = 90.0

	var playOutlinesX float64
	var playOutlinesY float64

	playOutlinesX = 30
	playOutlinesY = 30

	var playOutlinesHeight float64
	var playOutlinesWidth float64

	playOutlinesHeight = 90
	playOutlinesWidth = 80

	var playCounterBoxOutlinesHeight float64
	var playCounterBoxOutlinesWidth float64
	playCounterBoxOutlinesHeight = 15
	playCounterBoxOutlinesWidth = 20
	var playCounterBoxOutlinesX float64
	var playCounterBoxOutlinesY float64
	playCounterBoxOutlinesX = playOutlinesX + 30
	playCounterBoxOutlinesY = playOutlinesY

	//Create the gridlines for the plays and the play numbers
	for i := 1; i <= 24; i++ {

		//Rect(x, y, w, h float64, styleStr string)
		//Draw the outside box of the play
		pdf.Rect(playOutlinesX, playOutlinesY, playOutlinesWidth, playOutlinesHeight, "D")
		//Draw the outside box for the play number
		pdf.Rect(playCounterBoxOutlinesX, playCounterBoxOutlinesY, playCounterBoxOutlinesWidth, playCounterBoxOutlinesHeight, "D")

		//Write the play numbers
		pdf.SetFont("Arial", "B", 11)
		pdf.Text(playCounterBoxOutlinesX+5, playCounterBoxOutlinesY+11, strconv.Itoa(i))

		//pdf.Text(10, 10, "11234567890")

		playOutlinesX += playOutlinesWidth

		if i%4 == 0 {
			playOutlinesY += playOutlinesHeight
			playOutlinesX = 30
		}

		if i%8 == 0 {
			playOutlinesY += 30
		}

		playCounterBoxOutlinesX = playOutlinesX + 30
		playCounterBoxOutlinesY = playOutlinesY

	}

	for playNumber := range offensivePlaybook.OffensivePlays {
		if playNumber == 0 {
			footballLocationX = footballOrigLocationX
			footballLocationY = 90.0
		} else if playNumber%8 == 0 {
			footballLocationY += 120.0
			footballLocationX = footballOrigLocationX
		} else if playNumber%4 == 0 {
			footballLocationY += 90.0
			footballLocationX = footballOrigLocationX
		} else {
			if playNumber != 0 {
				footballLocationX += playOutlinesWidth
			}
		}

		for playerRouteNumber, playerValue := range offensivePlaybook.OffensivePlays[playNumber].Formation.Players {

			playerStartingLocationX = playerValue.Coordinates.BallOffsetX*scaleFactor + footballLocationX
			playerStartingLocationY = playerValue.Coordinates.BallOffsetY*-1*scaleFactor + footballLocationY

			playerCurrentLocationX = playerStartingLocationX
			playerCurrentLocationY = playerStartingLocationY

			//Being lazy as I don't want to retype as playerCurrentLocationX and playerCurrentLocationY
			xCurrent = playerCurrentLocationX
			yCurrent = playerCurrentLocationY

			xNew = xCurrent
			yNew = yCurrent

			scaledRouteX = nil
			scaledRouteY = nil
			copyRouteX = nil
			copyRouteY = nil

			//copy the route for scaling later
			for i, _ := range offensivePlaybook.OffensivePlays[playNumber].PlayerRoutes[playerRouteNumber].MinX {
				copyRouteX = append(copyRouteX, offensivePlaybook.OffensivePlays[playNumber].PlayerRoutes[playerRouteNumber].MinX[i])
				// we need to do the * -1 to reverse as pdf's coordinates start at the top left of the page
				// instead of the bottom left of the page
				copyRouteY = append(copyRouteY, offensivePlaybook.OffensivePlays[playNumber].PlayerRoutes[playerRouteNumber].MinY[i]*-1.0)
			}

			//scale the route for the PDF document
			for i, _ := range copyRouteX {

				if i < len(copyRouteX)-1 {
					if copyRouteX[i] == copyRouteX[i+1] && copyRouteY[i] == copyRouteY[i+1] {
						scaledRouteTempIteration += 1
					} else {
						scaledRouteTempIteration = scaledRouteTempIteration * scaleFactor
						for iter := 0.0; iter < scaledRouteTempIteration; iter += 1 {
							scaledRouteX = append(scaledRouteX, copyRouteX[i])
							scaledRouteY = append(scaledRouteY, copyRouteY[i])
						}
						scaledRouteTempIteration = 0
					}
				} else {
					scaledRouteTempIteration = scaledRouteTempIteration * scaleFactor
					for iter := 0.0; iter < scaledRouteTempIteration; iter += 1 {
						scaledRouteX = append(scaledRouteX, copyRouteX[i])
						scaledRouteY = append(scaledRouteY, copyRouteY[i])
					}
					scaledRouteTempIteration = 0
				}

			}

			//draw the route for the player
			for i, _ := range scaledRouteX {
				xNew += scaledRouteX[i]
				yNew += scaledRouteY[i]
				//this will ensure that the playroute doesn't print outside the play boxes
				if i < len(scaledRouteX) && (xNew < footballLocationX+(playOutlinesWidth/2) && xNew > footballLocationX-(playOutlinesWidth/2)) && (yNew < footballLocationY+(playOutlinesHeight/2) && yNew > footballLocationY-(playOutlinesHeight/2)) {
					// set the width of the route line for easier viewing
					pdf.SetLineWidth(1)
					// set the color of the route for the player
					pdf.SetDrawColor(int(playerValue.Attributes.Color.R), int(playerValue.Attributes.Color.G), int(playerValue.Attributes.Color.B))
					// draw the route for the player
					pdf.Line(xCurrent, yCurrent, xNew, yNew)
				}
				xCurrent = xNew
				yCurrent = yNew
			}

			//Player color - R,G,B inputs for the fill color
			pdf.SetFillColor(int(playerValue.Attributes.Color.R), int(playerValue.Attributes.Color.G), int(playerValue.Attributes.Color.B))

			//Draw the player as a cirlce with fill no outline (FD for fill and outline)
			//do this here so it will be on top of the route
			pdf.Circle(playerStartingLocationX, playerStartingLocationY, playerCircleRadius, "FD")
		}
	}
}

func ReturnEmptyOffensivePlay() (blankPlay OffensivePlay) {
	return blankPlay
}

func BuildDefaultOffensivePlayBook() PlayBook {

	var setupPlayerRoutes []routes.OffensePlayRoute
	var defaultPlaybook PlayBook

	defaultPlaybook.PlayBookName = "Default"

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
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftOutTenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftWhipFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftInFiveYardRoute())

	defaultPlaybook.OffensivePlays = append(defaultPlaybook.OffensivePlays, AddPlayBookPage("Bunch Left In Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes))

	setupPlayerRoutes = nil

	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftInTenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftOutTenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftWhipFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineGoRoute())

	defaultPlaybook.OffensivePlays = append(defaultPlaybook.OffensivePlays, AddPlayBookPage("Bunch Left Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes))

	return defaultPlaybook

}

func DrawOffensivePlayerPlayRoute(imd *imdraw.IMDraw, playerCoordinates formations.OffensePlayerCoordinates, playerRoutes routes.OffensePlayRoute, routeColor color.RGBA) {

	imd.Color = routeColor

	for i, _ := range playerRoutes.MinX {
		playerCoordinates.MinX += playerRoutes.MinX[i]
		playerCoordinates.MinY += playerRoutes.MinY[i]
		imd.Push(pixel.V(playerCoordinates.MinX, playerCoordinates.MinY))
	}

	imd.Line(1)

}

func DrawOffensivePlayBookPage(imd *imdraw.IMDraw, win *pixelgl.Window, pageNumber int, offensivePlayBook PlayBook) {

	for i, v := range offensivePlayBook.OffensivePlays[pageNumber].Formation.Players {

		DrawOffensivePlayerPlayRoute(imd, v.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[i], colornames.Gold)

	}

	formations.DrawOffensivePlayers(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation)

}

func DrawOffensivePlayBookMenu(imd *imdraw.IMDraw, win *pixelgl.Window, offensivePlayBook PlayBook, pageNumber int) {

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	basicTxtMenu := text.New(pixel.V(600, 600), atlas)

	fmt.Fprintln(basicTxtMenu, "Playbook Menu:")

	basicTxtMenu.Draw(win, pixel.IM.Scaled(basicTxtMenu.Orig, 2))

	basicTxt := text.New(pixel.V(600, 400), atlas)

	fmt.Fprintln(basicTxt, "Playbook Name: "+offensivePlayBook.PlayBookName)
	fmt.Fprintln(basicTxt, "Play Name: "+offensivePlayBook.OffensivePlays[pageNumber].PlayName)
	fmt.Fprintln(basicTxt, "Formation Name: "+offensivePlayBook.OffensivePlays[pageNumber].Formation.FormationName)
	fmt.Fprintln(basicTxt, "Snap Type: "+offensivePlayBook.OffensivePlays[pageNumber].Formation.SnapType)
	fmt.Fprintln(basicTxt, "Recievers: "+fmt.Sprint(offensivePlayBook.OffensivePlays[pageNumber].Formation.Receivers))
	fmt.Fprintln(basicTxt, "Running Backs: "+fmt.Sprint(offensivePlayBook.OffensivePlays[pageNumber].Formation.RunningBacks))

	basicTxt.Draw(win, pixel.IM)
}

func DrawOffensiveRunPlayMenu(imd *imdraw.IMDraw, win *pixelgl.Window, offensivePlayBook PlayBook, pageNumber int) {

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	basicTxtMenu := text.New(pixel.V(600, 600), atlas)

	fmt.Fprintln(basicTxtMenu, "Play Run Menu:")

	basicTxtMenu.Draw(win, pixel.IM.Scaled(basicTxtMenu.Orig, 2))

	basicTxt := text.New(pixel.V(600, 400), atlas)

	fmt.Fprintln(basicTxt, "Playbook Name: "+offensivePlayBook.PlayBookName)
	fmt.Fprintln(basicTxt, "Play Name: "+offensivePlayBook.OffensivePlays[pageNumber].PlayName)
	fmt.Fprintln(basicTxt, "Formation Name: "+offensivePlayBook.OffensivePlays[pageNumber].Formation.FormationName)
	fmt.Fprintln(basicTxt, "Snap Type: "+offensivePlayBook.OffensivePlays[pageNumber].Formation.SnapType)
	fmt.Fprintln(basicTxt, "Recievers: "+fmt.Sprint(offensivePlayBook.OffensivePlays[pageNumber].Formation.Receivers))
	fmt.Fprintln(basicTxt, "Running Backs: "+fmt.Sprint(offensivePlayBook.OffensivePlays[pageNumber].Formation.RunningBacks))

	basicTxt.Draw(win, pixel.IM)
}

func DrawBuildOffensivePlaybookMenu(imd *imdraw.IMDraw, win *pixelgl.Window) {

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxtMenu := text.New(pixel.V(600, 600), atlas)

	fmt.Fprintln(basicTxtMenu, "Build Offensive Playbook Menu:")
	basicTxtMenu.Draw(win, pixel.IM.Scaled(basicTxtMenu.Orig, 2))

}

func DrawBuildOffensivePlaybookMenuSelectFormation(imd *imdraw.IMDraw, win *pixelgl.Window, formationIteration int) {

	//formationIteration = 0

	DrawBuildOffensivePlaybookMenu(imd, win)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxtMenuSelectFormation := text.New(pixel.V(600, 500), atlas)

	fmt.Fprintln(basicTxtMenuSelectFormation, "Select your formation:")
	basicTxtMenuSelectFormation.Draw(win, pixel.IM)

	formations.DrawSpecificOffensiveFormation(imd, win, formationIteration)

}

//func DrawBuildOffensivePlaybookMenuSelectRoute(imd *imdraw.IMDraw, win *pixelgl.Window, formationIteration int, routeIteration int, selectedPlayer int) {
func DrawBuildOffensivePlaybookMenuSelectRoute(imd *imdraw.IMDraw, win *pixelgl.Window) {
	DrawBuildOffensivePlaybookMenu(imd, win)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxtMenuSelectFormation := text.New(pixel.V(600, 500), atlas)

	fmt.Fprintln(basicTxtMenuSelectFormation, "Select your player route:")
	basicTxtMenuSelectFormation.Draw(win, pixel.IM)

}

func DrawBuildOffensivePlaybookMenuDoneConfirmation(imd *imdraw.IMDraw, win *pixelgl.Window) {

	DrawBuildOffensivePlaybookMenu(imd, win)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxtMenuSelectFormation := text.New(pixel.V(600, 500), atlas)

	fmt.Fprintln(basicTxtMenuSelectFormation, "Are you sure you're done with your play?")
	fmt.Fprintln(basicTxtMenuSelectFormation, "Press enter to save")
	fmt.Fprintln(basicTxtMenuSelectFormation, "Press tab to go back to edit")
	basicTxtMenuSelectFormation.Draw(win, pixel.IM)

}
