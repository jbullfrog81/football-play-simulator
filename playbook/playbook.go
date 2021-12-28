package playbook

import (
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/routes"

	"encoding/json"
	"io/ioutil"
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
