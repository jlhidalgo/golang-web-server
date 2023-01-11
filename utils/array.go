package utils

func Contains(options []string, str string) bool {
	for _, option := range options {
		if option == str {
			return true
		}
	}
	return false
}
