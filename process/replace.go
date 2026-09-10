package process

import (
	"strings"

	"github.com/jasontconnell/go-web-project-init/data"
)

func ReplaceAll(folder *data.Folder, replacements []data.KeyValue) {
	updateFolder(folder, replacements)
}

func updateFolder(folder *data.Folder, replacements []data.KeyValue) {
	for _, rep := range replacements {
		folder.Name = strings.ReplaceAll(folder.Name, rep.Key, rep.Value)
	}

	for _, f := range folder.Folders {
		updateFolder(f, replacements)
	}

	for _, f := range folder.Files {
		updateFile(f, replacements)
	}
}

func updateFile(file *data.File, replacements []data.KeyValue) {
	for _, rep := range replacements {
		file.Name = strings.ReplaceAll(file.Name, rep.Key, rep.Value)

		s := string(file.Contents)

		s = strings.ReplaceAll(s, rep.Key, rep.Value)
		file.Contents = []byte(s)
	}
}
