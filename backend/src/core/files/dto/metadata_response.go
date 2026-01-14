package dto

// MetaDataDTO is a Data Transfer Object for file Metadata.
// When the frontend sends a multipart form, metadata is stored
// as raw json,
type MetaDataDTO struct {
	Path             string `json:"path"`
	RelativePath     string `json:"relativePath"`
	LastModified     int64  `json:"lastModified"`
	LastModifiedDate string `json:"lastModifiedDate"`
	Size             uint64 `json:"size"`
	FileType         string `json:"fileType"`

	ID       string `json:"id"`
	OwnerID  string `json:"ownerId"`
	CheckSum string `json:"checkSum"`
}
