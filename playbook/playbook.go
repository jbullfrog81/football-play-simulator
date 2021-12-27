package playbook

import (
	"jbullfrog81/football-play-simulator/formations"
	"jbullfrog81/football-play-simulator/routes"
)

type PlayBook struct {
	OffensivePlays []OffensivePlay
}

type OffensivePlay struct {
	PlayName     string
	Formation    formations.OffenseTeamFormation
	PlayerRoutes []routes.OffensePlayRoute
}
