package models

type File struct {
	Filename string
	Filetype string
	Lifetime uint
}

type FileRepository interface {
	Add(file File) error
	IsExists(filename string) bool
}

func NewFile(fname, filetype string, lifetime uint) File {
	return File{
		fname,
		filetype,
		lifetime,
	}
}
