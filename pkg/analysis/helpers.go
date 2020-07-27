package analysis

import (
	"math"
)

func ShiftValueSlice(src []float64, lastValue float64) {
	srcLastIndex := len(src) - 1
	for i := range src {
		if i == srcLastIndex {
			src[i] = lastValue
			break
		}
		src[i] = src[i+1]
	}
}

func GetSliceVariance(src []float64, until int, mean float64) float64 {
	variance := float64(0)
	for i, value := range src {
		variance += math.Pow(value-mean, 2)
		if i == until {
			variance = variance / float64(i)
			break
		}
	}
	return variance
}
