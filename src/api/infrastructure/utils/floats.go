package utils

func NullableInt32ToFloat64(param *int32) float64 {
	var result float64
	if param != nil {
		result = float64(*param)
	}

	return result
}