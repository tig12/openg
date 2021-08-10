package tiglib

// Returns the min and max values of slice of strings
func MinMaxString(s []string) (min, max string) {
	max = s[0]
	min = s[0]
	for _, value := range s {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
