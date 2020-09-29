package src

import (
	"bufio"
	"os"
	"strings"
)

func readAllowedTargets(configFile *string) []string {
	_, err := os.Stat(*configFile)
	if err != nil {
		return make([]string, 0)
	}

	f, _ := os.Open(*configFile)

	scanner := bufio.NewScanner(f)
	result := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\t \n\r")
		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result
}
