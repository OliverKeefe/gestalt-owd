package gestaltofs

import "io/fs"

type GestaltoFS interface {
	fs.FS
	Stat(name string) (MetaData, error)
	Write(name string, data []byte, perms fs.FileMode) error
	Remove(name string) error
}
