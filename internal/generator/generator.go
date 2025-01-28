package generator

import (
	"context"
	"fmt"
	"github.com/Dmitrijlin/go-skeleton/internal/dialog"
	"github.com/Dmitrijlin/go-skeleton/internal/helper"
	"os"
)

type Generator struct {
	dialog *dialog.Dialog
}

var (
	baseStructure = map[string]string{
		"main.go": "",
	}
)

func NewGenerator() *Generator {
	return &Generator{
		dialog: dialog.NewDialog(),
	}
}

func (g *Generator) Generate(
	ctx context.Context,
	dir, configPath string,
	interactMode bool,
) error {
	dirExists, err := helper.Exists(dir)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	if !dirExists {
		fmt.Println("Directory not exists. Creating new directory", dir)
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("generate: create directory: %w", err)
		}
	}

	if err = g.generateBaseStructure(ctx, dir, configPath); err != nil {
		return fmt.Errorf("generate base structure: %w", err)
	}

	return nil
}
