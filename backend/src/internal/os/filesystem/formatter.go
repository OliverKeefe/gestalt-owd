package gestaltofs

import (
	"io/fs"
	"time"
)

type ISO8601Date time.Time

const isoLayout = "2006-01-02"

func (d ISO8601Date) String() string {
	t := time.Time(d)
	return t.Format(isoLayout)
}

// type time
type ISO8601Time time.Time

const isoTimeLayout = "09:54:16.067672290"
