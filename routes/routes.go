package routes

type OffensePlayRoute struct {
	MinX []float64
	MinY []float64
	MaxX []float64
	MaxY []float64
}

func DefineOutFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

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

func DefineOutTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

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

func DefinePostTenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

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

func DefineWhipSevenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	//var values OffensePlayRoute

	for i := 0; i < 71; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
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

func DefineWhipFiveYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
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

func DefineOutAndUpSevenYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

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

func DefineBlockRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.MinX = append(route.MinX, float64(0))
	route.MinY = append(route.MinY, float64(0))
	route.MaxX = append(route.MaxX, float64(0))
	route.MaxY = append(route.MaxY, float64(0))

	return route

}
