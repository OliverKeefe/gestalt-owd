package gestaltofs

import (
	"io/fs"
	"time"
)

type File struct {
	Data FileData
	Meta MetaData
}

// FileData is an array of bytes (uint8) and encompasses the actual content of a file.
type FileData []byte

// MetaData is a struct encompassing a file's information.
//
// It is OS-agnostic and used for the GestaltoFileSystem GestaltoFS only.
type MetaData struct {
	FileName   string
	Path       string
	Size       uint64
	Mode       fs.FileMode
	IsDir      bool
	ModifiedAt time.Time
	UploadedAt *time.Time
	Owner      string
	AccessTo   []string
	Group      []string
	Links      *uint64
}

// MacOSXMetaData contains metadata for files on Mac OS X.
type MacOSXMetaData struct{}

// WindowsMetaData contains metadata for files on Windows.
type WindowsMetaData struct{}

// GNULinuxMetaData contains metadata for files on Linux.
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

// AndroidMetaData contains metadata for files on Android OS.
type AndroidMetaData struct{}

// iOSMetaData contains metadata for files on iOS devices.
type iOSMetaData struct{}

// iPadOSMetaData contains metadata for files on iPad OS devices.
type iPadOSMetaData struct{}
