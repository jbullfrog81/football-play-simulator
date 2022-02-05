package playbook

import (
	"fmt"
	"image/color"
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/routes"

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

	footballLocationX = 100.0
	footballLocationY = 100.0

	var playOutlinesX float64
	var playOutlinesY float64

	playOutlinesX = 30
	playOutlinesY = 30

	var playOutlinesHeight float64
	var playOutlinesWidth float64

	playOutlinesHeight = 90
	playOutlinesWidth = 80

	for i := 1; i <= 24; i++ {

		pdf.Rect(playOutlinesX, playOutlinesY, playOutlinesWidth, playOutlinesHeight, "D")

		playOutlinesX += playOutlinesWidth

		if i%4 == 0 {
			playOutlinesY += playOutlinesHeight
			playOutlinesX = 30
		}

	}
	//Rect(x, y, w, h float64, styleStr string)
	//pdf.Rect(10, 10, 400, 200, "D")
	//pdf.Rect(210, 410, 400, 200, "D")

	for playNumber := range offensivePlaybook.OffensivePlays {

		if (playNumber+1)%4 == 0 {
			footballLocationY += 200.0
			footballLocationX = 100.0
		} else {
			if playNumber != 0 {
				footballLocationX += 100.0
			}
		}

		for playerRouteNumber, playerValue := range offensivePlaybook.OffensivePlays[playNumber].Formation.Players {
			fmt.Println("The player is:" + playerValue.Attributes.Position)
			//xCurrent = playerValue.Coordinates.MinX
			//yCurrent = playerValue.Coordinates.MinY

			playerCurrentLocationX = playerValue.Coordinates.BallOffsetX*scaleFactor + footballLocationX
			playerCurrentLocationY = playerValue.Coordinates.BallOffsetY*-1*scaleFactor + footballLocationY

			//Being lazy as I don't want to retype as playerCurrentLocationX and playerCurrentLocationY
			xCurrent = playerCurrentLocationX
			yCurrent = playerCurrentLocationY

			//Player color - R,G,B inputs for the fill color
			pdf.SetFillColor(int(playerValue.Attributes.Color.R), int(playerValue.Attributes.Color.G), int(playerValue.Attributes.Color.B))

			//Draw the player as a cirlce with fill no outline (FD for fill and outline)
			pdf.Circle(playerCurrentLocationX, playerCurrentLocationY, playerCircleRadius, "FD")

			//fmt.Sprintf("The player coordinates are: %f", playerValue.Attributes.Position)
			//fmt.Sprintf("xCurrent: %f", xCurrent)
			//fmt.Sprintf("yCurrent: %f", xCurrent)
			fmt.Println("The player initial coordinates are:")
			fmt.Println(playerValue.Attributes.Position)
			fmt.Println("Initial xCurrent:")
			fmt.Println(xCurrent)
			fmt.Println("Initial yCurrent:")
			fmt.Println(yCurrent)

			xNew = xCurrent
			yNew = yCurrent

			scaledRouteX = nil
			scaledRouteY = nil
			copyRouteX = nil
			copyRouteY = nil

			for i, _ := range offensivePlaybook.OffensivePlays[playNumber].PlayerRoutes[playerRouteNumber].MinX {
				copyRouteX = append(copyRouteX, offensivePlaybook.OffensivePlays[playNumber].PlayerRoutes[playerRouteNumber].MinX[i])
				// we need to do the * -1 to reverse as pdf's coordinates start at the top left of the page
				// instead of the bottom left of the page
				copyRouteY = append(copyRouteY, offensivePlaybook.OffensivePlays[playNumber].PlayerRoutes[playerRouteNumber].MinY[i]*-1.0)
			}

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
			fmt.Println("length of original route:")
			fmt.Println(len(offensivePlaybook.OffensivePlays[playNumber].PlayerRoutes[playerRouteNumber].MinX))
			fmt.Println("length of copy route:")
			fmt.Println(len(copyRouteX))
			fmt.Println("length of scaled route:")
			fmt.Println(len(scaledRouteX))

			for i, _ := range scaledRouteX {
				xNew += scaledRouteX[i]
				yNew += scaledRouteY[i]
				if i < len(scaledRouteX) {
					pdf.Line(xCurrent, yCurrent, xNew, yNew)
				}
				xCurrent = xNew
				yCurrent = yNew
			}
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
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftOutFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftOutAndUpSevenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftWhipFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineLeftPostTenYardRoute())

	defaultPlaybook.OffensivePlays = append(defaultPlaybook.OffensivePlays, AddPlayBookPage("Bunch Left In Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes))

	setupPlayerRoutes = nil

	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineGoRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutAndUpSevenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightWhipFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightPostTenYardRoute())

	defaultPlaybook.OffensivePlays = append(defaultPlaybook.OffensivePlays, AddPlayBookPage("Bunch Right In Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes))

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

func DrawBuildOffensivePlaybookMenuSelectRoute(imd *imdraw.IMDraw, win *pixelgl.Window, formationIteration int, routeIteration int, selectedPlayer int) {

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
