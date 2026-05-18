package config

import (
	"encoding/json"
	"os"
)

type Service struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Services []Service `json:"services"`
}

func Load(path string) (Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var data Config
	err = json.Unmarshal(bytes, &data)
	if err != nil {

		return Config{}, err
	}
	return data, nil
}
