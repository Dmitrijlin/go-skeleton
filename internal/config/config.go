package config

import (
	"encoding/json"
	"fmt"
	file2 "github.com/Dmitrijlin/go-skeleton/internal/file"
	"github.com/Dmitrijlin/go-skeleton/internal/project-struct"
	"os"
	"path/filepath"
)

type Config struct {
	UsedTags      map[projectstruct.Tag]bool    `json:"usedTags"`
	ProjectStruct []projectstruct.ProjectStruct `json:"projectStruct"`
}

func GetConfigFile(projectPath string) (*Config, error) {
	var configuration Config

	configFileDir, err := getConfigFileDir(projectPath)
	if err != nil {
		return nil, fmt.Errorf("could not get config directory: %w", err)
	}

	configLockFileData, err := getConfigLockFileData(configFileDir)
	if err != nil {
		return nil, fmt.Errorf("could not get config lock file: %w", err)
	}

	if configLockFileData != nil {
		err = json.Unmarshal(configLockFileData, &configuration)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling file %s: %w", projectPath, err)
		}

		projectstruct.SetTags(configuration.UsedTags)

		return &configuration, nil
	}

	configFileData, err := getConfigFileData(projectPath)
	if err != nil {
		return nil, fmt.Errorf("could not get config file data: %w", err)
	}

	var config []projectstruct.ProjectStruct
	err = json.Unmarshal(configFileData, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling file %s: %w", projectPath, err)
	}

	configuration.ProjectStruct = config
	configuration.UsedTags = projectstruct.CollectTags(config)

	return &configuration, nil
}

func getConfigFileDir(projectPath string) (string, error) {
	configFileDir := projectPath
	configFilePath := filepath.Join(projectPath, projectstruct.ConfigFileName)
	exists, err := file2.Exists(configFilePath)
	if err != nil {
		return "", fmt.Errorf("error checking if file exists: %w", err)
	}

	if !exists {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user home directory: %w", err)
		}

		configFileDir = filepath.Join(homeDir, projectstruct.ConfigDirInHome)
		configFilePath = filepath.Join(configFileDir, projectstruct.ConfigFileName)
		exists, err = file2.Exists(configFilePath)
		if err != nil {
			return "", fmt.Errorf("error checking if file exists: %w", err)
		}

		if !exists {
			return "", fmt.Errorf("file does not exist. Please run `skeleton init`: %s", configFilePath)
		}
	}

	return configFileDir, nil
}

func getConfigLockFileData(configFileDir string) ([]byte, error) {
	return getFileData(configFileDir, projectstruct.ConfigLockFileName)
}

func getConfigFileData(configFileDir string) ([]byte, error) {
	return getFileData(configFileDir, projectstruct.ConfigFileName)
}

func getFileData(configFileDir, filename string) ([]byte, error) {
	configFilePath := filepath.Join(configFileDir, filename)
	exists, err := file2.Exists(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error checking if file exists: %w", err)
	}

	if !exists {
		return nil, nil
	}

	return os.ReadFile(configFilePath)
}
