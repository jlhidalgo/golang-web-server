package utils

func StringArrayContains(strArray []string, str string) bool {
	for _, elem := range strArray {
		if elem == str {
			return true
		}
	}
	return false
}
