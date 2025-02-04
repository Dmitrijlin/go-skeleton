package initializer

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"github.com/Dmitrijlin/go-skeleton/internal/file"
	projectstruct "github.com/Dmitrijlin/go-skeleton/internal/project-struct"
	"os"
	"path/filepath"
	"text/template"
)

type Initializer struct {
	Global      bool
	ProjectPath string
	ConfigPath  string
}

var ErrCannotBeUsedTogether = errors.New("cannot be used together")

//go:embed template/*
var templates embed.FS

func NewInitializer(
	global bool,
	projectPath string,
	configPath string,
) *Initializer {
	return &Initializer{
		Global:      global,
		ProjectPath: projectPath,
		ConfigPath:  configPath,
	}
}

func (i *Initializer) Initialize(_ context.Context) error {
	if i.Global && i.ProjectPath != "" {
		return fmt.Errorf("flags global and config: %w", ErrCannotBeUsedTogether)
	}

	path, err := i.getPath()
	if err != nil {
		return fmt.Errorf("get path: %w", err)
	}

	configPath := filepath.Join(path, projectstruct.ConfigFileName)
	configLockPath := filepath.Join(path, projectstruct.ConfigLockFileName)

	err = file.DeleteIfExists(configPath)
	if err != nil {
		return fmt.Errorf("could not delete config file: %w", err)
	}

	err = file.DeleteIfExists(configLockPath)
	if err != nil {
		return fmt.Errorf("could not delete config file: %w", err)
	}

	configFileData, err := i.getConfigFileData(configPath)
	if err != nil {
		return fmt.Errorf("get config file: %w", err)
	}

	err = file.WriteFile(configPath, configFileData)
	if err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}

func (i *Initializer) getConfigFromTemplate() ([]byte, error) {
	tmpl := template.Must(template.ParseFS(templates, "template/skeleton.json.tpl"))
	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, map[string]any{}); err != nil {
		return nil, fmt.Errorf("cannot render template: %w", err)
	}

	return buf.Bytes(), nil
}

func (i *Initializer) getPath() (string, error) {
	var path string
	var err error

	if i.Global {
		path, err = os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("could not get home directory: %w", err)
		}

		path = filepath.Join(path, projectstruct.ConfigDirInHome)
	} else if i.ProjectPath != "" {
		path = i.ProjectPath
	} else {
		path, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("could not get current working directory: %w", err)
		}
	}

	err = file.CreateDirIfNotExist(path)
	if err != nil {
		return "", fmt.Errorf("could not check if project exists: %w", err)
	}

	return path, nil
}

func (i *Initializer) getConfigFileData(configPath string) ([]byte, error) {
	var configFileData []byte
	var err error

	if i.ConfigPath != "" {
		configFileData, err = file.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("could not read config file: %w", err)
		}
	} else {
		configFileData, err = i.getConfigFromTemplate()
		if err != nil {
			return nil, fmt.Errorf("could not read config file: %w", err)
		}
	}

	return configFileData, nil
}
