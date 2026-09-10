package data

type File struct {
	Name     string
	Contents []byte
	Folder   *Folder
}

type Folder struct {
	Name   string
	Folder *Folder

	Folders []*Folder
	Files   []*File
}
