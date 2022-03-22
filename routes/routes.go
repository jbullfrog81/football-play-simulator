package routes

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type OffensePlayRoute struct {
	RouteName string
	MinX      []float64
	MinY      []float64
	MaxX      []float64
	MaxY      []float64
}

type OffensePlayRoutes struct {
	Routes []OffensePlayRoute
}

func DrawOffensiveRoutesMenu(imd *imdraw.IMDraw, win *pixelgl.Window, route OffensePlayRoute) {

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	basicTxtMenu := text.New(pixel.V(600, 600), atlas)

	fmt.Fprintln(basicTxtMenu, "Routes Menu:")

	basicTxtMenu.Draw(win, pixel.IM.Scaled(basicTxtMenu.Orig, 2))

	basicTxt := text.New(pixel.V(600, 400), atlas)

	fmt.Fprintln(basicTxt, "Route Name: "+route.RouteName)

	basicTxt.Draw(win, pixel.IM)

}

func ReturnAllOffensePlayRoutes() (allRoutes OffensePlayRoutes) {

	var route OffensePlayRoute

	// TODO: figure out how to iterate through a slice of all route functions
	// to generate all of the routes
	//var sliceOfRoutes := {"","",}

	route = DefineNoneRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockLongRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockLeftRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockLongLeftRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockVeryLongLeftRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockRightRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockLongRightRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineBlockVeryLongRightRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightOutToSidelineRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightOutUpSidelineRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftOutToSidelineRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftOutUpSidelineRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineGoRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightSlantThreeYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightOutFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightOutTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightInFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightInTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightPostTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightWhipSevenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightWhipFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightOutAndUpSevenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightOutAndUpFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightCurlFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightCurlTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightFlatRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightQuickFlatRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightQuickRunRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightReceiverReverseRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineRightReceiverShortReverseRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftCurlFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftCurlTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftSlantThreeYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftOutFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftOutTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftInFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftInTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftPostTenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftWhipSevenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftWhipFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftOutAndUpSevenYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftOutAndUpFiveYardRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftFlatRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftQuickFlatRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftQuickRunRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftReceiverReverseRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	route = DefineLeftReceiverShortReverseRoute()
	allRoutes.Routes = append(allRoutes.Routes, route)

	return allRoutes

}

// Field Side Agnostic Routes

func DefineNoneRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.MinX = append(route.MinX, float64(0))
	route.MinY = append(route.MinY, float64(0))
	route.MaxX = append(route.MaxX, float64(0))
	route.MaxY = append(route.MaxY, float64(0))
	route.RouteName = "None"

	return route

}

func DefineBlockRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "Block"

	return route

}

func DefineBlockLongRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 41; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "BlockLong"

	return route

}

func DefineBlockRightRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "BlockRight"

	return route

}

func DefineBlockLeftRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "BlockLeft"

	return route

}

func DefineBlockLongRightRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 41; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "BlockLongRight"

	return route

}

func DefineBlockVeryLongRightRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 81; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "BlockVeryLongRight"

	return route

}

func DefineBlockLongLeftRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 41; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "BlockVeryLongLeft"

	return route

}

func DefineBlockVeryLongLeftRoute() OffensePlayRoute {

	var route OffensePlayRoute

	for i := 0; i < 81; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 10; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	route.RouteName = "BlockLongLeft"

	return route

}

func DefineGoRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Go"

	for i := 0; i < 151; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

// Right Side Field Routes

func DefineRightCurlFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right 5 yard Curl"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	return route

}

func DefineRightCurlTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right 10 yard Curl"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	return route

}

func DefineRightOutToSidelineRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right Out to Side Line"

	//var values OffensePlayRoute

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightOutUpSidelineRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right Out Up Side Line"

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightFlatRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right Flat"

	//var values OffensePlayRoute

	for i := 0; i < 81; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightQuickFlatRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right Quick Flat"

	//var values OffensePlayRoute

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightQuickRunRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right Quick Run"

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightReceiverReverseRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right Receiver Reverse Route"

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightReceiverShortReverseRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Right Receiver Short Reverse Route"

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightSlantThreeYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "3 yd Slant - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightOutFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd Out - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightOutTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "10 yd Out - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 151; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightInFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd In - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightInTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "10 yd In - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 151; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightPostTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "10 yd Post - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightWhipSevenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "7 yd Whip - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightWhipFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd Whip - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 41; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineRightOutAndUpFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd Out and Up - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineRightOutAndUpSevenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "7 yd Out and Up - Right Side"

	//var values OffensePlayRoute

	for i := 0; i < 71; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

// Left Side Field Routes

func DefineLeftCurlFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left 5 yard Curl"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	return route

}

func DefineLeftCurlTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left 10 yard Curl"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	return route

}

func DefineLeftOutToSidelineRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left Out to Side Line"

	//var values OffensePlayRoute

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftOutUpSidelineRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left Out Up Side Line"

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftFlatRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left Flat"

	//var values OffensePlayRoute

	for i := 0; i < 81; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftQuickFlatRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left Quick Flat"

	//var values OffensePlayRoute

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftQuickRunRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left Quick Run"

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftReceiverReverseRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left Receiver Reverse Route"

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 61; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftReceiverShortReverseRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Left Receiver Short Reverse Route"

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(-1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(-1))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftSlantThreeYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "3 yd Slant - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftOutFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd Out - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftOutTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "10 yd Out - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 151; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftInFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd In - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftInTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "10 yd In - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 151; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftPostTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "10 yd Post - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftWhipSevenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "7 yd Whip - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 21; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftWhipFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd Whip - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 41; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 11; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 5; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	return route

}

func DefineLeftOutAndUpFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "5 yd Out and Up - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

func DefineLeftOutAndUpSevenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "7 yd Out and Up - Left Side"

	//var values OffensePlayRoute

	for i := 0; i < 71; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
	}

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}
