package laba8

func SearchingString(filename, valueForSearch string) bool {
	if valueForSearch == "" {
		return false
	}
	values := ReadDataFromFile(filename)

	for _, value := range values {
		if value == valueForSearch {
			return true
		}
	}
	return false
}
