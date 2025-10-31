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

type iPadOSMetaData struct{}
