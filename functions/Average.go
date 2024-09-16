package functions

import "math"

func Average(Numbers []int) float64 {
	var sum int
	for _, num := range Numbers {
		sum += num
	}
	return math.Round(float64(sum) / float64(len(Numbers)))
}
