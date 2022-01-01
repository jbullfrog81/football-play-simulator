package playbook

import (
	"fmt"
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/routes"

	"encoding/json"
	"io/ioutil"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
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

func SavePlayBookToFile(playBook PlayBook) {

	file, _ := json.MarshalIndent(playBook, "", " ")
	_ = ioutil.WriteFile(playBook.PlayBookName+".playbook", file, 0644)

}

func BuildDefaultOffensivePlayBook() PlayBook {

	var setupPlayerRoutes []routes.OffensePlayRoute
	var defaultPlaybook PlayBook

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

func DrawOffensivePlayerPlayRoute(imd *imdraw.IMDraw, playerCoordinates formations.OffensePlayerCoordinates, playerRoutes routes.OffensePlayRoute) {

	imd.Color = colornames.Gold

	for i, _ := range playerRoutes.MinX {
		playerCoordinates.MinX += playerRoutes.MinX[i]
		playerCoordinates.MinY += playerRoutes.MinY[i]
		imd.Push(pixel.V(playerCoordinates.MinX, playerCoordinates.MinY))
	}

	imd.Line(1)

}

func DrawOffensivePlayBookPage(imd *imdraw.IMDraw, win *pixelgl.Window, pageNumber int, offensivePlayBook PlayBook) {

	for i, v := range offensivePlayBook.OffensivePlays[pageNumber].Formation.Players {

		DrawOffensivePlayerPlayRoute(imd, v.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[i])

	}
	//DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player1.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[0])

	//DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player2.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[1])

	//DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player3.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[2])

	//DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player4.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[3])

	//DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player5.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[4])

	//DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player6.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[5])

	//DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player7.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[6])

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
