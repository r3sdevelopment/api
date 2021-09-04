package utils

func Contains(arr []string, str string) bool {
	for _, i := range arr {
		if i == str {
			return true
		}
	}
	return false
}
