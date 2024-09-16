package functions

import "math"

func Variance(Numbers []int, average float64) (float64, float64) {
	var sum, variance float64
	for _, num := range Numbers {
		sum += math.Pow(float64(num)-average, 2)
	}
	variance = math.Round(sum / float64(len(Numbers)))
	return variance, math.Round(math.Sqrt(variance))
}
