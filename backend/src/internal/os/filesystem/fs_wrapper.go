package gestaltofs

import (
	"io/fs"
	"log"
)

type GestaltoFS interface {
	fs.FS
	Stat(name string) (MetaData, error)
	Write(name string, data []byte, perms fs.FileMode) (bool, error)
	Remove(name string) (bool, error)
	MapTo(file fs.File) (fs.File, error)
}

func Stat(name string) (MetaData, error) {
	file, err := fs.File(name)
	if err != nil {
		log.Fatalf("unable to find file: %s", err)
	}
}

func Write(name string data []byte, perms fs.FileMode) (bool, error) {

}

func Remove(name string) (bool, error) {

}

func MapTO(file fs.File) (fs.File, error) {

}