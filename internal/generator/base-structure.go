package generator

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	config2 "github.com/Dmitrijlin/go-skeleton/internal/config"
	"github.com/Dmitrijlin/go-skeleton/internal/file"
	"github.com/Dmitrijlin/go-skeleton/internal/project-struct"
	"os"
	"text/template"
)

//go:embed template/*
var templates embed.FS

func (g *Generator) generateStructure(
	_ context.Context,
	dir string,
) error {
	fmt.Println("Generating base structure into " + dir)

	config, err := config2.GetConfigFile(dir)
	if err != nil {
		return fmt.Errorf("generate: get config file: %w", err)
	}

	err = g.generateFromConfig(dir, config.ProjectStruct)
	if err != nil {
		return fmt.Errorf("generate: generate from config: %w", err)
	}

	return nil
}

//nolint:gocognit
func (g *Generator) generateFromConfig(baseDir string, config []projectstruct.ProjectStruct) error {
	for _, entity := range config {
		entityPath := fmt.Sprintf("%s/%s", baseDir, entity.Name)
		exists, err := file.Exists(entityPath)
		if err != nil {
			return fmt.Errorf("generate: error checking if file exists: %w", err)
		}

		switch entity.Type {
		case projectstruct.File:
			err = g.handleFile(exists, entityPath, entity)
		case projectstruct.Dir:
			err = g.handleDir(exists, entityPath, entity)
		}

		if err != nil {
			return fmt.Errorf("generate: handle file: %s %w", entityPath, err)
		}
	}

	return nil
}

func (g *Generator) handleFile(
	exists bool,
	entityPath string,
	entity projectstruct.ProjectStruct,
) error {
	if !exists {
		res, err := g.render(entity)
		if err != nil {
			return fmt.Errorf("generate: render: %w", err)
		}

		err = os.WriteFile(entityPath, res, 0755) //nolint:gosec
		if err != nil {
			return fmt.Errorf("generate: create file: %w", err)
		}
	} else {
		fmt.Println("Skipping existing file:", entityPath)
	}

	return nil
}

func (g *Generator) handleDir(exists bool,
	entityPath string,
	entity projectstruct.ProjectStruct,
) error {
	var err error

	if !exists {
		err = os.Mkdir(entityPath, 0755)
		if err != nil {
			return fmt.Errorf("generate: create dir: %w", err)
		}
	}

	if len(entity.Children) > 0 {
		err = g.generateFromConfig(entityPath, entity.Children)
		if err != nil {
			return fmt.Errorf("generate: generate from config: %w", err)
		}
	}

	return nil
}

func (g *Generator) render(file projectstruct.ProjectStruct) ([]byte, error) {
	if file.Type != projectstruct.File {
		return nil, nil
	}

	if v, ok := projectstruct.DefaultTagTemplates[file.Tag]; ok {
		tmpl := template.Must(template.ParseFS(templates, fmt.Sprintf("template/%s", v.Template)))
		buf := new(bytes.Buffer)

		if err := tmpl.Execute(buf, v.ParamsGenerator()); err != nil {
			return nil, fmt.Errorf("cannot render template: %w", err)
		}

		return buf.Bytes(), nil
	}

	return nil, nil
}
