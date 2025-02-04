package file

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, fmt.Errorf("exists: %w", err)
}

func DeleteIfExists(path string) error {
	exists, err := Exists(path)
	if err != nil {
		return fmt.Errorf("delete if exists: %w", err)
	}

	if !exists {
		return nil
	}

	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("delete if exists: %w", err)
	}

	return nil
}

func ReadFile(path string) ([]byte, error) {
	exists, err := Exists(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	if !exists {
		return nil, fmt.Errorf("read file: %w", fs.ErrNotExist)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return data, nil
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0755)
}

func CreateDirIfNotExist(path string) error {
	exists, err := Exists(path)
	if err != nil {
		return fmt.Errorf("could not check if project exists: %w", err)
	}

	if !exists {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("could not create project directory: %w", err)
		}
	}

	return nil
}
