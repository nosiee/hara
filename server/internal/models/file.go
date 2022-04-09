package models

type File struct {
	Filename        string
	Fullpath        string
	Deletetimestamp int64
}

type FileRepository interface {
	Add(file File) error
	IsExists(filename string) bool
}

func NewFile(fname, fullpath string, ts int64) File {
	return File{
		fname,
		fullpath,
		ts,
	}
}
