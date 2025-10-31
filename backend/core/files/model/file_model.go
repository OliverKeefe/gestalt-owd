package model

type File struct {
	Metadata MetaData
	FileData []byte
}

type MetaData struct {
	Filename     string
	Size         uint64
	Permissions  fs.FileMode
	LastModified time.Time
	IsDirectory  bool
}

type MacOSXMetaData struct{}

type WindowsMetaData struct{}

type GNULinuxMetaData struct{}

type AndroidMetaData struct{}

type iPadOSMetaData struct{}
