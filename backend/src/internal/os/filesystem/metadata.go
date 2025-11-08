package gestaltofs

import (
	"io/fs"
	"time"
)

type MetaData struct {
	FileName   string
	Path       string
	Size       uint64
	Mode       fs.FileMode
	IsDir      bool
	ModifiedAt time.Time
	CreatedAt  *time.Time
	Owner      string
	AccessTo   []string
	Group      []string
	Links      *uint64
}

type MacOSXMetaData struct{}

type WindowsMetaData struct{}

type GNULinuxMetaData struct {
	File     string
	Size     uint64
	FileType string
	Links    uint64
	Mode     string
	User     string
	Group    string
	Access   time.Time
	Modify   time.Time
	Change   time.Time
	Birth    time.Time
}

type AndroidMetaData struct{}

type iOSMetaData struct{}

type iPadOSMetaData struct{}
