package generator

import (
	"context"
	"fmt"
	config2 "github.com/Dmitrijlin/go-skeleton/internal/config"
	project_struct "github.com/Dmitrijlin/go-skeleton/internal/project-struct"
	"os"
)

func (g *Generator) generateBaseStructure(
	ctx context.Context,
	dir, configPath string,
) error {
	fmt.Println("Generating base structure into " + dir)

	config, err := config2.GetConfigFile(configPath)
	if err != nil {
		return fmt.Errorf("generate: get config file: %w", err)
	}

	err = g.generateFromConfig(dir, config)
	if err != nil {
		return fmt.Errorf("generate: generate from config: %w", err)
	}

	return nil
}

func (g *Generator) generateFromConfig(baseDir string, config []project_struct.ProjectStruct) error {
	for _, entity := range config {
		entityPath := fmt.Sprintf("%s/%s", baseDir, entity.Name)
		switch entity.Type {
		case project_struct.File:
			err := os.WriteFile(entityPath, []byte(""), 0600)
			if err != nil {
				return fmt.Errorf("generate: create file: %w", err)
			}
		case project_struct.Dir:
			err := os.Mkdir(entityPath, 0755)
			if err != nil {
				return fmt.Errorf("generate: create dir: %w", err)
			}

			if len(entity.Children) > 0 {
				err = g.generateFromConfig(entityPath, entity.Children)
				if err != nil {
					return fmt.Errorf("generate: generate from config: %w", err)
				}
			}
		}
	}

	return nil
}
