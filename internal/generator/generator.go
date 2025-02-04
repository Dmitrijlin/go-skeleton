package generator

import (
	"context"
	"fmt"
	"github.com/Dmitrijlin/go-skeleton/internal/dialog"
	"github.com/Dmitrijlin/go-skeleton/internal/file"
	"os"
)

type Generator struct {
	dialog *dialog.Dialog
}

func NewGenerator() *Generator {
	return &Generator{
		dialog: dialog.NewDialog(),
	}
}

func (g *Generator) Generate(
	ctx context.Context,
	dir string,
) error {
	var err error

	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current directory: %w", err)
		}
	}

	dirExists, err := file.Exists(dir)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	if !dirExists {
		fmt.Println("Directory not exists. Creating new directory", dir)
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("generate: create directory: %w", err)
		}
	}

	if err = g.generateStructure(ctx, dir); err != nil {
		return fmt.Errorf("generate base structure: %w", err)
	}

	return nil
}
