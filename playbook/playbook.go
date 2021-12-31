package playbook

import (
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/routes"

	"encoding/json"
	"io/ioutil"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

func AddPlayBookPage(playBook *PlayBook, playName string, playFormation formations.OffenseTeamFormation, playRoutes []routes.OffensePlayRoute) {

	var playBookPage OffensivePlay

	playBookPage.PlayName = playName
	playBookPage.Formation = playFormation
	playBookPage.PlayerRoutes = playRoutes

	playBook.OffensivePlays = append(playBook.OffensivePlays, playBookPage)

}

func SavePlayBookToFile(playBook *PlayBook) {

	file, _ := json.MarshalIndent(playBook, "", " ")
	_ = ioutil.WriteFile(playBook.PlayBookName+".playbook", file, 0644)

}

func BuildDefaultOffensivePlayBook(defaultPlaybook *PlayBook) {

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

	AddPlayBookPage(defaultPlaybook, "Bunch Left In Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes)

	setupPlayerRoutes = nil

	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineGoRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineBlockRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightOutAndUpSevenYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightWhipFiveYardRoute())
	setupPlayerRoutes = append(setupPlayerRoutes, routes.DefineRightPostTenYardRoute())

	AddPlayBookPage(defaultPlaybook, "Bunch Right In Whipper", formations.SetOffensiveTeamFormationShotgunBunchLeft(), setupPlayerRoutes)

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

func DrawOffensivePlayBookPage(imd *imdraw.IMDraw, win *pixelgl.Window, pageNumber int, offensivePlayBook *PlayBook) {

	DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player1.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[0])

	DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player2.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[1])

	DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player3.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[2])

	DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player4.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[3])

	DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player5.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[4])

	DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player6.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[5])

	DrawOffensivePlayerPlayRoute(imd, offensivePlayBook.OffensivePlays[pageNumber].Formation.Player7.Coordinates, offensivePlayBook.OffensivePlays[pageNumber].PlayerRoutes[6])

	formations.DrawOffensivePlayers(imd, &offensivePlayBook.OffensivePlays[pageNumber].Formation)

}
