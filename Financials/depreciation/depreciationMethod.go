package depreciation

func StraightLine(startValue float64, time int, predictedLastValue float64) float64 {
	salvageValue := (startValue - predictedLastValue) / float64(time)
	var resultValue = startValue
	for timeTik := 1; timeTik <= time; timeTik++ {
		resultValue = resultValue - salvageValue
	}
	return resultValue
}

func StraightLineArr(startValue float64, time int, predictedLastValue float64) []float64 {
	salvageValue := (startValue - predictedLastValue) / float64(time)
	var resultValue = startValue
	resultArray := make([]float64, 0)
	for timeTik := 1; timeTik <= time; timeTik++ {
		resultArray = append(resultArray, resultValue)
		resultValue = resultValue - salvageValue
	}
	return resultArray
}

func DoubleDeciliningBalance()  {
	
}
