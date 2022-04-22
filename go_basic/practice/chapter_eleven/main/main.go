package main

import "fmt"

func main() {
	average := Average{
		Sum:         12356,
		Num:         35,
		SumRiseRate: 0.2 / 100,
		NumRiseRate: 0.5 / 100,
	}
	a, b, c := GetAverage(average)
	fmt.Printf("%f, %f, %f", a, b, c)
}

func GetAverage(average Average) (basicAverage float64, averageChange float64, averageChangeRate float64) {
	basicAverage = (average.Sum / average.Num) * ((1 + average.NumRiseRate) / (1 + average.SumRiseRate))
	averageChange = (average.Sum / average.Num) * ((average.SumRiseRate - average.NumRiseRate) / (1 + average.SumRiseRate))
	averageChangeRate = (average.SumRiseRate - average.NumRiseRate) / (1 + average.NumRiseRate)
	return basicAverage, averageChange, averageChangeRate
}

type Average struct {
	Sum         float64 `json:"sum"`
	Num         float64 `json:"num"`
	SumRiseRate float64 `json:"sum_rise_rate"`
	NumRiseRate float64 `json:"num_rise_rate"`
}
