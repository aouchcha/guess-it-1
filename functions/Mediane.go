package functions

import "math"

func Median(Numbers []int, average float64) float64 {
	for i := 0; i < len(Numbers); i++ {
		for j := i + 1; j < len(Numbers); j++ {
			var temp int
			if Numbers[i] > Numbers[j] {
				temp = Numbers[i]
				Numbers[i] = Numbers[j]
				Numbers[j] = temp
			}
		}
	}

	if len(Numbers)%2 != 0 {
		return float64(Numbers[len(Numbers)/2])
	}
	return math.Round(float64(Numbers[len(Numbers)/2]+Numbers[len(Numbers)/2]) / 2)
}
