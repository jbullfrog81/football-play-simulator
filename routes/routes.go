package routes

type OffensePlayRoute struct {
	MinX []float64
	MinY []float64
	MaxX []float64
	MaxY []float64
}

// Field Side Agnostic Routes

func DefineBlockRoute() OffensePlayRoute {

	var route OffensePlayRoute

	route.MinX = append(route.MinX, float64(0))
	route.MinY = append(route.MinY, float64(0))
	route.MaxX = append(route.MaxX, float64(0))
	route.MaxY = append(route.MaxY, float64(0))

	return route

}

func DefineGoRoute() OffensePlayRoute {

	var route OffensePlayRoute

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

func DefineRightWhipFiveYardRoute() OffensePlayRoute {

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

func DefineRightOutAndUpSevenYardRoute() OffensePlayRoute {

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

// Left Side Field Routes

func DefineLeftSlantThreeYardRoute() OffensePlayRoute {

	var route OffensePlayRoute

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

	//var values OffensePlayRoute

	for i := 0; i < 71; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
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

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		route.MinX = append(route.MinX, float64(0))
		route.MinY = append(route.MinY, float64(1))
		route.MaxX = append(route.MaxX, float64(0))
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
