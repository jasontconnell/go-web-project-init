package process

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"path/filepath"

	"github.com/jasontconnell/go-web-project-init/data"
)

func UnzipBytes(file []byte) (*data.Folder, error) {
	buf := bytes.NewReader(file)
	z, err := zip.NewReader(buf, int64(len(file)))
	if err != nil {
		return nil, err
	}

	dirmap := make(map[string]*data.Folder)
	root := &data.Folder{}
	dirmap[""] = root

	var allerrs error
	for _, f := range z.File {
		cleanName := filepath.Clean(f.Name)
		pardir := filepath.Dir(cleanName)
		parent := dirmap[pardir]
		if parent == nil {
			parent = &data.Folder{Folder: root, Name: "/"}
			root.Folders = append(root.Folders, parent)
			dirmap[pardir] = parent
		}

		if f.FileInfo().IsDir() {
			folder := &data.Folder{Name: f.FileInfo().Name(), Folder: parent}
			parent.Folders = append(parent.Folders, folder)
			dirmap[cleanName] = folder
		} else {
			fc, err := f.Open()
			if err != nil {
				allerrs = errors.Join(allerrs, err)
			}
			contents, err := io.ReadAll(fc)
			if err != nil {
				allerrs = errors.Join(allerrs, err)
			}
			ff := &data.File{Name: f.FileInfo().Name(), Contents: contents}
			parent.Files = append(parent.Files, ff)
		}
	}

	// first folder is like go-web-project-[branch]
	return root.Folders[0].Folders[0], allerrs
}
