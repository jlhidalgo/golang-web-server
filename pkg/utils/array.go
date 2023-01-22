package utils

// Check if a given string is contained in an array of strings
func StringArrayContains(strArray []string, str string) bool {
	for _, elem := range strArray {
		if elem == str {
			return true
		}
	}
	return false
}
