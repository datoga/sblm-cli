package data

//IsValidLevel checks if the provided level is valid or not
func IsValidLevel(level string) bool {
	for _, v := range PredefinedLevels {
		if level == v {
			return true
		}
	}

	return false
}
