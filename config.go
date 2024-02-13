package blog

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server      Server   `json:"server"`
	CustomPages []string `json:"customPages"`
}

type Server struct {
	Startup Startup `json:"startup"`
}

type Startup struct {
	Convert bool `json:"convert"`
}

func GetConfig() (*Config, error) {
	configPath := "config.json"

	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return nil, err
	}

	fmt.Printf("Parsed config: %+v\n", config)
	return &config, nil
}
