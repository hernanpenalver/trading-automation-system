package slices

func GetStringSliceIndex (array []string, value string) int {
	for i, s := range array {
		if s == value {
			return i
		}
	}
	return -1
}
