package filesystem

import "time"

type ISO8601Date time.Time

const isoLayout = "2006-01-02"

func (d ISO8601Date) String() string {
	t := time.Time(d)
	return t.Format(isoLayout)
}

// type time

type MacOSXMetaData struct{}

type WindowsMetaData struct{}

type Birth struct {
}
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
