package env

import (
	"os"
	"strings"
)

func GetEnv(key string, def string) string {
	result := os.Getenv(key)
	if result == "" {
		return def
	}
	return strings.TrimSpace(result)
}
