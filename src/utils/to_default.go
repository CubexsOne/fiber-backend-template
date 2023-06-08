package utils

func ToDefault(valueToCheck string, defaultValue string) string {
	if valueToCheck == "" {
		return defaultValue
	}

	return valueToCheck
}
