package routes

type OffensePlayRoute struct {
	MinX []float64
	MinY []float64
	MaxX []float64
	MaxY []float64
}

func DefineOutFiveYardRoute() OffensePlayRoute {

	var outFiveYard OffensePlayRoute

	//var values OffensePlayRoute

	for i := 0; i < 51; i++ {
		outFiveYard.MinX = append(outFiveYard.MinX, float64(0))
		outFiveYard.MinY = append(outFiveYard.MinY, float64(1))
		outFiveYard.MaxX = append(outFiveYard.MaxX, float64(0))
		outFiveYard.MaxY = append(outFiveYard.MaxY, float64(1))
	}

	for i := 0; i < 101; i++ {
		outFiveYard.MinX = append(outFiveYard.MinX, float64(-1))
		outFiveYard.MinY = append(outFiveYard.MinY, float64(0))
		outFiveYard.MaxX = append(outFiveYard.MaxX, float64(-1))
		outFiveYard.MaxY = append(outFiveYard.MaxY, float64(0))
	}

	return outFiveYard

}
