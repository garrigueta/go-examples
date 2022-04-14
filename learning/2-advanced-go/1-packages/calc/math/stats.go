package math

func Avg(values []float64) float64 {
	total := 0.0
	for _, val := range values {
		total += val
	}

	return total / float64(len(values))
}
