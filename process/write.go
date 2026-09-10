package process

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jasontconnell/go-web-project-init/data"
)

func WriteProject(folder *data.Folder, destination string) error {
	stat, err := os.Stat(destination)
	if err != nil || !stat.IsDir() {
		return err
	}

	return writeFolder(folder, destination)
}

func writeFolder(folder *data.Folder, destination string) error {
	path := filepath.Join(destination, folder.Name)

	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create dir %s %w", path, err)
	}

	var ferrs error
	for _, f := range folder.Files {
		err = writeFile(f, path)
		if err != nil {
			ferrs = errors.Join(ferrs, err)
		}
	}

	var derrs error
	for _, f := range folder.Folders {
		err = writeFolder(f, path)
		if err != nil {
			derrs = errors.Join(derrs, err)
		}
	}

	return ferrs
}

func writeFile(file *data.File, destination string) error {
	path := filepath.Join(destination, file.Name)
	err := os.WriteFile(path, file.Contents, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create file %s %w", path, err)
	}
	return nil
}
