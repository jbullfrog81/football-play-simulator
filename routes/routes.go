package routes

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

func ReturnAllOffensePlayRoutes() (allRoutes OffensePlayRoutes) {

	var route OffensePlayRoute

	// TODO: figure out how to iterate through a slice of all route functions
	// to generate all of the routes
	//var sliceOfRoutes := {"","",}

	route = DefineBlockRoute()
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

	return allRoutes

}

// Field Side Agnostic Routes

func DefineBlockRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.MinX = append(route.MinX, float64(0))
	route.MinY = append(route.MinY, float64(0))
	route.MaxX = append(route.MaxX, float64(0))
	route.MaxY = append(route.MaxY, float64(0))
	route.RouteName = "Block"

	return route

}

func DefineGoRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "Go"

	for i := 0; i < 201; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

// Right Side Field Routes

func DefineRightSlantThreeYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.RouteName = "3 yd Right Slant"

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

	for i := 0; i < 151; i++ {
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

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
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

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}

// Left Side Field Routes

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

	for i := 0; i < 151; i++ {
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

	for i := 0; i < 71; i++ {
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

	for i := 0; i < 31; i++ {
		route.MinX = append(route.MinX, float64(1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(1))
		route.MaxY = append(route.MaxY, float64(0))
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

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(-1))
		route.MinY = append(route.MinY, float64(0))
		route.MaxX = append(route.MaxX, float64(-1))
		route.MaxY = append(route.MaxY, float64(0))
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

	for i := 0; i < 101; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
		route.MaxY = append(route.MaxY, float64(1))
	}

	return route

}
