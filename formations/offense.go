package offense

import (
	"image/color"

	"golang.org/x/image/colornames"
)

type offensePlayerCoordinates struct {
	minX float64
	minY float64
	maxX float64
	maxY float64
}

type offensePlayerAttributes struct {
	position  string
	thickness float64
	radius    float64
	color     color.Color
}

type offensePlayer struct {
	coordinates offensePlayerCoordinates
	attributes  offensePlayerAttributes
}

type offenseTeamFormation struct {
	player1 offensePlayer
	player2 offensePlayer
	player3 offensePlayer
	player4 offensePlayer
	player5 offensePlayer
	player6 offensePlayer
	player7 offensePlayer
}

func defineOffensivePlayerCoordinates(minX float64, minY float64, maxX float64, maxY float64) (playerCoordinates offensePlayerCoordinates) {
	playerCoordinates.minX = minX
	playerCoordinates.minY = minY
	playerCoordinates.maxX = maxX
	playerCoordinates.maxY = maxY

	return playerCoordinates
}

func defineOffensivePlayerAttributes(position string, thickness float64, radius float64, color color.Color) (playerAttributes offensePlayerAttributes) {
	playerAttributes.position = position
	playerAttributes.thickness = thickness
	playerAttributes.radius = radius
	playerAttributes.color = color

	return playerAttributes
}

func defineOffensivePlayer(playerAttributes *offensePlayerAttributes, playerCoordinates *offensePlayerCoordinates, player *offensePlayer) {
	player.coordinates = *playerCoordinates
	player.attributes = *playerAttributes
}

func defineOffensiveTeamFormation(player1 *offensePlayer, player2 *offensePlayer, player3 *offensePlayer, player4 *offensePlayer, player5 *offensePlayer, player6 *offensePlayer, player7 *offensePlayer) (offensiveTeam offenseTeamFormation) {

	offensiveTeam.player1 = *player1
	offensiveTeam.player2 = *player2
	offensiveTeam.player3 = *player3
	offensiveTeam.player4 = *player4
	offensiveTeam.player5 = *player5
	offensiveTeam.player6 = *player6
	offensiveTeam.player7 = *player7

	return offensiveTeam
}

func setOffensiveTeamFormationBunchRight() offenseTeamFormation {

	var leftGuard offensePlayer
	var center offensePlayer
	var rightGuard offensePlayer
	var quarterBack offensePlayer
	var receiver1 offensePlayer
	var receiver2 offensePlayer
	var receiver3 offensePlayer

	leftGuard.attributes = defineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.attributes = defineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.attributes = defineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.attributes = defineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.attributes = defineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.attributes = defineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver3.attributes = defineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.coordinates = defineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.coordinates = defineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.coordinates = defineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.coordinates = defineOffensivePlayerCoordinates(260.0, 130.0, 260.0, 130.0)
	receiver1.coordinates = defineOffensivePlayerCoordinates(360.0, 145.0, 360.0, 145.0)
	receiver2.coordinates = defineOffensivePlayerCoordinates(370.0, 130.0, 370.0, 145.0)
	receiver3.coordinates = defineOffensivePlayerCoordinates(380.0, 145.0, 380.0, 145.0)

	var offensiveTeam offenseTeamFormation

	offensiveTeam = defineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	return offensiveTeam

}
