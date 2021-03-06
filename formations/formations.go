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
	MinX        float64
	MinY        float64
	MaxX        float64
	MaxY        float64
	BallOffsetX float64
	BallOffsetY float64
}

type OffensePlayerAttributes struct {
	Position  string
	Thickness float64
	Radius    float64
	Color     color.RGBA
}

type OffensePlayer struct {
	Coordinates OffensePlayerCoordinates
	Attributes  OffensePlayerAttributes
}

type AllOffenseTeamFormations struct {
	Formations []OffenseTeamFormation
}

type AllDefenseTeamFormations struct {
	Formations []DefenseTeamFormation
}

type OffenseTeamFormation struct {
	Players       []OffensePlayer
	FormationName string
	SnapType      string
	Receivers     int
	RunningBacks  int
}

type DefensePlayerCoordinates struct {
	MinX        float64
	MinY        float64
	MaxX        float64
	MaxY        float64
	BallOffsetX float64
	BallOffsetY float64
}

type DefensePlayerAttributes struct {
	Position  string
	Thickness float64
	Radius    float64
	Color     color.RGBA
}

type DefensePlayer struct {
	Coordinates DefensePlayerCoordinates
	Attributes  DefensePlayerAttributes
}

type DefenseTeamFormation struct {
	Players       []DefensePlayer
	FormationName string
}

func DefineOffensivePlayerCoordinates(minX float64, minY float64, maxX float64, maxY float64) (playerCoordinates OffensePlayerCoordinates) {
	playerCoordinates.MinX = minX
	playerCoordinates.MinY = minY
	playerCoordinates.MaxX = maxX
	playerCoordinates.MaxY = maxY
	playerCoordinates.BallOffsetX = minX - 260.0
	playerCoordinates.BallOffsetY = minY - 145.0

	return playerCoordinates
}

func DefineDefensivePlayerCoordinates(minX float64, minY float64, maxX float64, maxY float64) (playerCoordinates DefensePlayerCoordinates) {
	playerCoordinates.MinX = minX
	playerCoordinates.MinY = minY
	playerCoordinates.MaxX = maxX
	playerCoordinates.MaxY = maxY
	playerCoordinates.BallOffsetX = minX - 260.0
	playerCoordinates.BallOffsetY = minY - 145.0

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
	AllFormations.Formations = append(AllFormations.Formations, SetOffensiveTeamFormationShotgunDoubleTwins())

	return AllFormations
}

func ReturnAllDefensiveTeamFormations() (AllFormations AllDefenseTeamFormations) {

	AllFormations.Formations = append(AllFormations.Formations, SetDefensiveTeamFormationZoneThreeFour())
	AllFormations.Formations = append(AllFormations.Formations, SetDefensiveTeamFormationZoneThreeThree())

	return AllFormations
}

func DefineOffensivePlayerAttributes(position string, thickness float64, radius float64, color color.RGBA) (PlayerAttributes OffensePlayerAttributes) {
	PlayerAttributes.Position = position
	PlayerAttributes.Thickness = thickness
	PlayerAttributes.Radius = radius
	PlayerAttributes.Color = color

	return PlayerAttributes
}

func DefineDefensivePlayerAttributes(position string, thickness float64, radius float64, color color.RGBA) (PlayerAttributes DefensePlayerAttributes) {
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

func DefineDefensivePlayer(playerAttributes *DefensePlayerAttributes, playerCoordinates *DefensePlayerCoordinates, player *DefensePlayer) {
	player.Coordinates = *playerCoordinates
	player.Attributes = *playerAttributes
}

func DefineOffensiveTeamFormation(Players []OffensePlayer) (OffensiveTeam OffenseTeamFormation) {

	OffensiveTeam.Players = Players

	return OffensiveTeam
}

func DefineDefensiveTeamFormation(Players []DefensePlayer) (DefensiveTeam DefenseTeamFormation) {

	DefensiveTeam.Players = Players

	return DefensiveTeam
}

func SetOffensiveTeamFormationShowRoutes() OffenseTeamFormation {

	var player OffensePlayer

	player.Attributes = DefineOffensivePlayerAttributes("Player", 2.0, 5.0, colornames.Black)

	player.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, player)

	offensiveTeam.FormationName = "Show Routes"
	offensiveTeam.SnapType = "Not Applicable"
	offensiveTeam.Receivers = 0
	offensiveTeam.RunningBacks = 0

	return offensiveTeam

}

func SetDefensiveTeamFormationZoneThreeFour() DefenseTeamFormation {

	var leftGuard, noseguard, rightGuard, linebacker1, linebacker2, cornerback1, cornerback2 DefensePlayer

	leftGuard.Attributes = DefineDefensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	noseguard.Attributes = DefineDefensivePlayerAttributes("Nose Guard", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineDefensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	linebacker1.Attributes = DefineDefensivePlayerAttributes("Middle Linebacker", 2.0, 5.0, colornames.Black)
	linebacker2.Attributes = DefineDefensivePlayerAttributes("Middle Linebacker", 2.0, 5.0, colornames.Forestgreen)
	cornerback1.Attributes = DefineDefensivePlayerAttributes("Left Corner", 2.0, 5.0, colornames.Red)
	cornerback2.Attributes = DefineDefensivePlayerAttributes("Right Corner", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineDefensivePlayerCoordinates(200.0, 155.0, 200.0, 165.0)
	noseguard.Coordinates = DefineDefensivePlayerCoordinates(260.0, 155.0, 260.0, 165.0)
	rightGuard.Coordinates = DefineDefensivePlayerCoordinates(320.0, 155.0, 320.0, 165.0)
	linebacker1.Coordinates = DefineDefensivePlayerCoordinates(240.0, 205.0, 240.0, 205.0)
	linebacker2.Coordinates = DefineDefensivePlayerCoordinates(280.0, 205.0, 280.0, 205.0)
	cornerback1.Coordinates = DefineDefensivePlayerCoordinates(180.0, 205.0, 180.0, 205.0)
	cornerback2.Coordinates = DefineDefensivePlayerCoordinates(350.0, 205.0, 350.0, 205.0)

	var defensiveTeam DefenseTeamFormation

	defensiveTeam.Players = append(defensiveTeam.Players, leftGuard)
	defensiveTeam.Players = append(defensiveTeam.Players, noseguard)
	defensiveTeam.Players = append(defensiveTeam.Players, rightGuard)
	defensiveTeam.Players = append(defensiveTeam.Players, linebacker1)
	defensiveTeam.Players = append(defensiveTeam.Players, linebacker2)
	defensiveTeam.Players = append(defensiveTeam.Players, cornerback1)
	defensiveTeam.Players = append(defensiveTeam.Players, cornerback2)

	defensiveTeam.FormationName = "Zone - Three Four"

	return defensiveTeam

}

func SetDefensiveTeamFormationZoneThreeThree() DefenseTeamFormation {

	var leftGuard, noseguard, rightGuard, linebacker, safety, cornerback1, cornerback2 DefensePlayer

	leftGuard.Attributes = DefineDefensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	noseguard.Attributes = DefineDefensivePlayerAttributes("Nose Guard", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineDefensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	linebacker.Attributes = DefineDefensivePlayerAttributes("Middle Linebacker", 2.0, 5.0, colornames.Black)
	safety.Attributes = DefineDefensivePlayerAttributes("Safety", 2.0, 5.0, colornames.Forestgreen)
	cornerback1.Attributes = DefineDefensivePlayerAttributes("Left Corner", 2.0, 5.0, colornames.Red)
	cornerback2.Attributes = DefineDefensivePlayerAttributes("Right Corner", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineDefensivePlayerCoordinates(200.0, 155.0, 200.0, 165.0)
	noseguard.Coordinates = DefineDefensivePlayerCoordinates(260.0, 155.0, 260.0, 165.0)
	rightGuard.Coordinates = DefineDefensivePlayerCoordinates(320.0, 155.0, 320.0, 165.0)
	linebacker.Coordinates = DefineDefensivePlayerCoordinates(260.0, 205.0, 260.0, 205.0)
	safety.Coordinates = DefineDefensivePlayerCoordinates(260.0, 225.0, 260.0, 225.0)
	cornerback1.Coordinates = DefineDefensivePlayerCoordinates(180.0, 205.0, 180.0, 205.0)
	cornerback2.Coordinates = DefineDefensivePlayerCoordinates(350.0, 205.0, 350.0, 205.0)

	var defensiveTeam DefenseTeamFormation

	defensiveTeam.Players = append(defensiveTeam.Players, leftGuard)
	defensiveTeam.Players = append(defensiveTeam.Players, noseguard)
	defensiveTeam.Players = append(defensiveTeam.Players, rightGuard)
	defensiveTeam.Players = append(defensiveTeam.Players, linebacker)
	defensiveTeam.Players = append(defensiveTeam.Players, safety)
	defensiveTeam.Players = append(defensiveTeam.Players, cornerback1)
	defensiveTeam.Players = append(defensiveTeam.Players, cornerback2)

	defensiveTeam.FormationName = "Zone - Three Three"

	return defensiveTeam

}

func SetOffensiveTeamFormationShotgunBunchRight() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Red)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(330.0, 145.0, 330.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(340.0, 130.0, 340.0, 130.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(350.0, 145.0, 350.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver3)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	offensiveTeam.FormationName = "Bunch Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunBunchLeft() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Red)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(190.0, 145.0, 190.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(180.0, 130.0, 180.0, 130.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(170.0, 145.0, 170.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, receiver3)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	offensiveTeam.FormationName = "Bunch Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTripsLeft() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Red)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(190.0, 145.0, 190.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(180.0, 145.0, 180.0, 145.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(170.0, 145.0, 170.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, receiver3)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	offensiveTeam.FormationName = "Trips Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTripsRight() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, receiver3 OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Red)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(330.0, 145.0, 330.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(340.0, 145.0, 340.0, 145.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(350.0, 145.0, 350.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver3)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &receiver1, &receiver2, &receiver3)

	offensiveTeam.FormationName = "Trips Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 3
	offensiveTeam.RunningBacks = 0

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsRightBackRight() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Red)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(280.0, 75.0, 280.0, 75.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(330.0, 145.0, 330.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(340.0, 145.0, 340.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBack)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	offensiveTeam.FormationName = "Twins Right Running Back Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsRightBackLeft() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Red)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(240.0, 75.0, 280.0, 75.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(330.0, 145.0, 330.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(340.0, 145.0, 340.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBack)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	offensiveTeam.FormationName = "Twins Right Running Back Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsLeftBackRight() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Red)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(280.0, 75.0, 280.0, 75.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(190.0, 145.0, 190.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(200.0, 145.0, 200.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)
	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBack)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	offensiveTeam.FormationName = "Twins Left Running Back Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunTwinsLeftBackLeft() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, receiver2, runningBack OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBack.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Red)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	runningBack.Coordinates = DefineOffensivePlayerCoordinates(240.0, 75.0, 280.0, 75.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(200.0, 145.0, 200.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(210.0, 145.0, 210.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)
	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBack)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBack, &receiver1, &receiver2)

	offensiveTeam.FormationName = "Twins Left Running Back Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 2
	offensiveTeam.RunningBacks = 1

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunSplitBackRecieverRight() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, runningBackRight, runningBackLeft OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBackLeft.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Red)
	runningBackRight.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Blue)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	runningBackLeft.Coordinates = DefineOffensivePlayerCoordinates(240.0, 75.0, 240.0, 75.0)
	runningBackRight.Coordinates = DefineOffensivePlayerCoordinates(280.0, 75.0, 280.0, 75.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(330.0, 145.0, 330.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBackLeft)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBackRight)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBackRight, &receiver1, &runningBackLeft)

	offensiveTeam.FormationName = "Split Backs Reciever Right"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 1
	offensiveTeam.RunningBacks = 2

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunSplitBackRecieverLeft() OffenseTeamFormation {

	var leftGuard, center, rightGuard, quarterBack, receiver1, runningBackRight, runningBackLeft OffensePlayer

	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	rightGuard.Attributes = DefineOffensivePlayerAttributes("Right Guard", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)
	runningBackLeft.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Red)
	runningBackRight.Attributes = DefineOffensivePlayerAttributes("Running Back", 2.0, 5.0, colornames.Blue)
	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)

	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	rightGuard.Coordinates = DefineOffensivePlayerCoordinates(280.0, 145.0, 280.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)
	runningBackLeft.Coordinates = DefineOffensivePlayerCoordinates(240.0, 75.0, 240.0, 75.0)
	runningBackRight.Coordinates = DefineOffensivePlayerCoordinates(280.0, 75.0, 280.0, 75.0)
	receiver1.Coordinates = DefineOffensivePlayerCoordinates(190.0, 145.0, 190.0, 145.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, rightGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBackLeft)
	offensiveTeam.Players = append(offensiveTeam.Players, runningBackRight)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBackRight, &receiver1, &runningBackLeft)

	offensiveTeam.FormationName = "Split Backs Reciever Left"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 1
	offensiveTeam.RunningBacks = 2

	return offensiveTeam

}

func SetOffensiveTeamFormationShotgunDoubleTwins() OffenseTeamFormation {

	var leftGuard, center, quarterBack, receiver1, receiver2, receiver3, receiver4 OffensePlayer

	receiver1.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Forestgreen)
	receiver2.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Red)
	leftGuard.Attributes = DefineOffensivePlayerAttributes("Left Guard", 2.0, 5.0, colornames.Orange)
	center.Attributes = DefineOffensivePlayerAttributes("Center", 2.0, 5.0, colornames.Lightskyblue)
	receiver3.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Blue)
	receiver4.Attributes = DefineOffensivePlayerAttributes("Wide Receiver", 2.0, 5.0, colornames.Orange)
	quarterBack.Attributes = DefineOffensivePlayerAttributes("Quarterback", 2.0, 5.0, colornames.Black)

	receiver1.Coordinates = DefineOffensivePlayerCoordinates(190.0, 145.0, 190.0, 145.0)
	receiver2.Coordinates = DefineOffensivePlayerCoordinates(200.0, 145.0, 200.0, 145.0)
	leftGuard.Coordinates = DefineOffensivePlayerCoordinates(240.0, 145.0, 240.0, 145.0)
	center.Coordinates = DefineOffensivePlayerCoordinates(260.0, 145.0, 260.0, 145.0)
	receiver3.Coordinates = DefineOffensivePlayerCoordinates(330.0, 145.0, 330.0, 145.0)
	receiver4.Coordinates = DefineOffensivePlayerCoordinates(340.0, 145.0, 340.0, 145.0)
	quarterBack.Coordinates = DefineOffensivePlayerCoordinates(260.0, 115.0, 260.0, 115.0)

	var offensiveTeam OffenseTeamFormation

	offensiveTeam.Players = append(offensiveTeam.Players, receiver1)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver2)
	offensiveTeam.Players = append(offensiveTeam.Players, leftGuard)
	offensiveTeam.Players = append(offensiveTeam.Players, center)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver3)
	offensiveTeam.Players = append(offensiveTeam.Players, receiver4)
	offensiveTeam.Players = append(offensiveTeam.Players, quarterBack)

	//offensiveTeam = DefineOffensiveTeamFormation(&leftGuard, &center, &rightGuard, &quarterBack, &runningBackRight, &receiver1, &runningBackLeft)

	offensiveTeam.FormationName = "Double Twins"
	offensiveTeam.SnapType = "Shotgun"
	offensiveTeam.Receivers = 4
	offensiveTeam.RunningBacks = 0

	return offensiveTeam

}

func DrawOffensivePlayers(imd *imdraw.IMDraw, team OffenseTeamFormation) {

	for _, v := range team.Players {
		imd.Color = v.Attributes.Color
		imd.Push(pixel.V(v.Coordinates.MinX, v.Coordinates.MinY))
		imd.Circle(v.Attributes.Radius, v.Attributes.Thickness)
	}

}

func DrawDefensivePlayers(imd *imdraw.IMDraw, team DefenseTeamFormation) {

	for _, v := range team.Players {
		imd.Color = v.Attributes.Color
		imd.Push(pixel.V(v.Coordinates.MinX, v.Coordinates.MinY))
		imd.Circle(v.Attributes.Radius, v.Attributes.Thickness)
	}

}

func DrawOffensePlayerRunPlay(imd *imdraw.IMDraw, route routes.OffensePlayRoute, playerPosition OffensePlayer, iteration int) OffensePlayer {

	//println("starting draw offense run play")
	if iteration < len(route.MinX) {
		//println("inside iteration loop")
		playerPosition.Coordinates.MinX += route.MinX[iteration]
		playerPosition.Coordinates.MinY += route.MinY[iteration]
		playerPosition.Coordinates.MaxX += route.MaxX[iteration]
		playerPosition.Coordinates.MaxY += route.MaxY[iteration]
	}

	imd.Color = playerPosition.Attributes.Color
	imd.Push(pixel.V(playerPosition.Coordinates.MinX, playerPosition.Coordinates.MinY))
	imd.Circle(playerPosition.Attributes.Radius, playerPosition.Attributes.Thickness)

	return playerPosition

}

// Draw a specific offensive formation from the available offensive formations
func DrawSpecificOffensiveFormation(imd *imdraw.IMDraw, win *pixelgl.Window, iteration int) {

	availableOffensiveFormations := ReturnAllOffensiveTeamFormations()
	//currentFormation := availableOffensiveFormations.Formations[iteration]
	//for i, v := range availableOffensiveFormations.Formations {
	//	if i < 1 {

	DrawOffensivePlayers(imd, availableOffensiveFormations.Formations[iteration])

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(600, 400), atlas)

	fmt.Fprintln(basicTxt, "Name: "+availableOffensiveFormations.Formations[iteration].FormationName)
	fmt.Fprintln(basicTxt, "Snap Type: "+availableOffensiveFormations.Formations[iteration].SnapType)
	fmt.Fprintln(basicTxt, "Recievers: "+fmt.Sprint(availableOffensiveFormations.Formations[iteration].Receivers))
	fmt.Fprintln(basicTxt, "Running Backs: "+fmt.Sprint(availableOffensiveFormations.Formations[iteration].RunningBacks))

	basicTxt.Draw(win, pixel.IM)

}

func DrawSpecificDefensiveFormation(imd *imdraw.IMDraw, win *pixelgl.Window, iteration int) {

	availableDefensiveFormations := ReturnAllDefensiveTeamFormations()

	DrawDefensivePlayers(imd, availableDefensiveFormations.Formations[iteration])

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(600, 400), atlas)

	fmt.Fprintln(basicTxt, "Name: "+availableDefensiveFormations.Formations[iteration].FormationName)

	basicTxt.Draw(win, pixel.IM)

}

// Draw an offensive formation that is provided as an input
func DrawProvidedOffensiveFormation(imd *imdraw.IMDraw, win *pixelgl.Window, offensiveFormation OffenseTeamFormation) {

	DrawOffensivePlayers(imd, offensiveFormation)

}

func DrawOffensiveFormatonsMenu(imd *imdraw.IMDraw, win *pixelgl.Window) {

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	basicTxtMenu := text.New(pixel.V(600, 600), atlas)

	fmt.Fprintln(basicTxtMenu, "Formations Menu:")

	basicTxtMenu.Draw(win, pixel.IM.Scaled(basicTxtMenu.Orig, 2))

}

func DrawDefensiveFormatonsMenu(imd *imdraw.IMDraw, win *pixelgl.Window) {

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	basicTxtMenu := text.New(pixel.V(600, 600), atlas)

	fmt.Fprintln(basicTxtMenu, "Formations Menu:")

	basicTxtMenu.Draw(win, pixel.IM.Scaled(basicTxtMenu.Orig, 2))

}
