package formations

import (
	"image/color"

	"golang.org/x/image/colornames"
)

type OffensePlayerCoordinates struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

type OffensePlayerAttributes struct {
	Position  string
	Thickness float64
	Radius    float64
	Color     color.Color
}

type OffensePlayer struct {
	Coordinates OffensePlayerCoordinates
	Attributes  OffensePlayerAttributes
}

type OffenseTeamFormation struct {
	Player1 OffensePlayer
	Player2 OffensePlayer
	Player3 OffensePlayer
	Player4 OffensePlayer
	Player5 OffensePlayer
	Player6 OffensePlayer
	Player7 OffensePlayer
}

func DefineOffensivePlayerCoordinates(minX float64, minY float64, maxX float64, maxY float64) (playerCoordinates OffensePlayerCoordinates) {
	playerCoordinates.MinX = minX
	playerCoordinates.MinY = minY
	playerCoordinates.MaxX = maxX
	playerCoordinates.MaxY = maxY

	return playerCoordinates
}

func DefineOffensivePlayerAttributes(position string, thickness float64, radius float64, color color.Color) (PlayerAttributes OffensePlayerAttributes) {
	PlayerAttributes.Position = position
	PlayerAttributes.Thickness = thickness
	PlayerAttributes.Radius = radius
	PlayerAttributes.Color = color

	return PlayerAttributes
}

func DefineOffensivePlayer(playerAttributes *OffensePlayerAttributes, playerCoordinates *OffensePlayerCoordinates, player *OffensePlayer) {
	player.Coordinates = *playerCoordinates
	player.Attributes = *playerAttributes
}

func DefineOffensiveTeamFormation(Player1 *OffensePlayer, player2 *OffensePlayer, player3 *OffensePlayer, player4 *OffensePlayer, player5 *OffensePlayer, player6 *OffensePlayer, player7 *OffensePlayer) (OffensiveTeam OffenseTeamFormation) {

	OffensiveTeam.Player1 = *Player1
	OffensiveTeam.Player2 = *player2
	OffensiveTeam.Player3 = *player3
	OffensiveTeam.Player4 = *player4
	OffensiveTeam.Player5 = *player5
	OffensiveTeam.Player6 = *player6
	OffensiveTeam.Player7 = *player7

	return OffensiveTeam
}

func SetOffensiveTeamFormationShotgunBunchRight() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(360.0, 145.0, 360.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(370.0, 130.0, 370.0, 130.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(380.0, 145.0, 380.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunBunchLeft() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(160.0, 145.0, 160.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(150.0, 130.0, 150.0, 130.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(140.0, 145.0, 140.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTripsLeft() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(160.0, 145.0, 160.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(140.0, 145.0, 140.0, 145.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(120.0, 145.0, 120.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTripsRight() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(360.0, 145.0, 360.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(380.0, 145.0, 380.0, 145.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(400.0, 145.0, 400.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsRightBackRight() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(280.0, 85.0, 280.0, 85.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(360.0, 145.0, 360.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(380.0, 145.0, 380.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsRightBackLeft() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(240.0, 85.0, 280.0, 85.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(360.0, 145.0, 360.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(380.0, 145.0, 380.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsLeftBackRight() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(280.0, 85.0, 280.0, 85.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(160.0, 145.0, 160.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(180.0, 145.0, 180.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsLeftBackLeft() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var receiver2 OffensePlayer
	var runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(240.0, 85.0, 280.0, 85.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(160.0, 145.0, 160.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(180.0, 145.0, 180.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunSplitBackRecieverRight() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var runningBackRight OffensePlayer
	var runningBackLeft OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBackLeft.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	runningBackRight.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	runningBackLeft.Coordinates = DefineOffensivePlayerCoordinates(240.0, 85.0, 240.0, 85.0)
	runningBackRight.Coordinates = DefineOffensivePlayerCoordinates(280.0, 85.0, 280.0, 85.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(360.0, 145.0, 360.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBackRight, &receiver1, &runningBackLeft)

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunSplitBackRecieverLeft() OffenseTeamFormation {

	var leftGuard OffensePlayer
	var center OffensePlayer
	var rightGuard OffensePlayer
	var quarterBack OffensePlayer
	var receiver1 OffensePlayer
	var runningBackRight OffensePlayer
	var runningBackLeft OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Black)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Black)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Black)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBackLeft.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	runningBackRight.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Black)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 105.0, 260.0, 105.0)
	runningBackLeft.Coordinates = DefineOffensivePlayerCoordinates(240.0, 85.0, 240.0, 85.0)
	runningBackRight.Coordinates = DefineOffensivePlayerCoordinates(280.0, 85.0, 280.0, 85.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(160.0, 145.0, 160.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBackRight, &receiver1, &runningBackLeft)

	return offensiveTeam

}
