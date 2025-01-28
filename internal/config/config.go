package config

import (
	"encoding/json"
	"fmt"
	"github.com/Dmitrijlin/go-skeleton/internal/helper"
	"github.com/Dmitrijlin/go-skeleton/internal/project-struct"
	"os"
)

func GetConfigFile(configPath string) ([]projectstruct.ProjectStruct, error) {
	exists, err := helper.Exists(configPath)
	if err != nil {
		return nil, fmt.Errorf("error checking if file exists: %w", err)
	}

	if !exists {
		return nil, nil
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", configPath, err)
	}

	var config []projectstruct.ProjectStruct
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling file %s: %w", configPath, err)
	}

	return config, nil
}
