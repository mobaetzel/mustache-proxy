package src

import (
	"os"
	"strings"
)

func readAllowedTargets() []string {
	allowedTargets := os.Getenv("ALLOWED_TARGETS")
	return strings.Split(allowedTargets, ",")
}
