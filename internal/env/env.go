package env

import (
	"os"
	"strconv"
)

func GetString(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return ""
}

func GetInt(key string) int {
	value := os.Getenv(key)
	if value != "" {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
	}

	return 0
}
