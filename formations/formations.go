package formations

import (
	"fmt"
	"image/color"

	"jbullfrog81/football-play-simulator/routes"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
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

type AllOffenseTeamFormations struct {
	Formations []OffenseTeamFormation
}

type OffenseTeamFormation struct {
	Player1       OffensePlayer
	Player2       OffensePlayer
	Player3       OffensePlayer
	Player4       OffensePlayer
	Player5       OffensePlayer
	Player6       OffensePlayer
	Player7       OffensePlayer
	FormationName string
	SnapType      string
	Receivers     int
	RunningBacks  int
}

func DefineOffensivePlayerCoordinates(minX float64, minY float64, maxX float64, maxY float64) (playerCoordinates OffensePlayerCoordinates) {
	playerCoordinates.MinX = minX
	playerCoordinates.MinY = minY
	playerCoordinates.MaxX = maxX
	playerCoordinates.MaxY = maxY

	return playerCoordinates
}

func ReturnAllOffensiveTeamFormations() (AllFormations AllOffenseTeamFormations) {

	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunBunchRight())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunBunchLeft())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunSplitBackRecieverLeft())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunSplitBackRecieverRight())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunTripsLeft())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunTripsRight())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunTwinsLeftBackLeft())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunTwinsLeftBackRight())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunTwinsRightBackLeft())
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunTwinsRightBackRight())

	return AllFormations
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

	offensiveTeam.FormationName = "Bunch Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

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

	offensiveTeam.FormationName = "Bunch Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

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

	offensiveTeam.FormationName = "Trips Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

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

	offensiveTeam.FormationName = "Trips Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

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

	offensiveTeam.FormationName = "Twins Right Running Back Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

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

	offensiveTeam.FormationName = "Twins Right Running Back Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

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

	offensiveTeam.FormationName = "Twins Left Running Back Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

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

	offensiveTeam.FormationName = "Twins Left Running Back Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

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

	offensiveTeam.FormationName = "Split Backs Reciever Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 1
	offensiveTeam.RunningBacks = 2

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

	offensiveTeam.FormationName = "Split Backs Reciever Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 1
	offensiveTeam.RunningBacks = 2

	return offensiveTeam

}

func DrawOffensivePlayers(imd *imdraw.IMDraw, team *OffenseTeamFormation) {

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

func DrawOffensePlayerRunPlay(imd *imdraw.IMDraw, route *routes.OffensePlayRoute, playerPosition *OffensePlayer, iteration int) {

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

func DrawSpecificOffensiveFormation(imd *imdraw.IMDraw, win *pixelgl.Window, iteration int) {

	availableOffensiveFormations := ReturnAllOffensiveTeamFormations()
	currentFormation := availableOffensiveFormations.Formations[iteration]
	//for i, v := range availableOffensiveFormations.Formations {
	//	if i < 1 {

	DrawOffensivePlayers(imd, &currentFormation)

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(600, 400), atlas)

	fmt.Fprintln(basicTxt, "Name: "+currentFormation.FormationName)
	fmt.Fprintln(basicTxt, "Snap Type: "+currentFormation.SnapType)
	fmt.Fprintln(basicTxt, "Recievers: "+fmt.Sprint(currentFormation.Receivers))
	fmt.Fprintln(basicTxt, "Running Backs: "+fmt.Sprint(currentFormation.RunningBacks))

	basicTxt.Draw(win, pixel.IM)

}
