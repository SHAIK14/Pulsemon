package config

import (
	"os"
	"strings"
)

func LoadEnv(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))

	}
	return nil

}
