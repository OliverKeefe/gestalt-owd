package filesystem

type MacOSXMetaData struct{}

type WindowsMetaData struct{}

type GNULinuxMetaData struct {
	File     string
	Size     uint64
	Blocks   uint16
	IOBlock  uint64
	FileType string
	Device   string
	Inode    uint64
	Links    uint8
	Access
}

type AndroidMetaData struct{}

type iOSMetaData struct{}

type iPadOSMetaData struct{}
