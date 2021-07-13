package types

type PathInfo struct {
	Path  string         `json:"path,omitempty"`
	Items []PathItemInfo `json:"items,omitempty"`
}

type PathItemInfo struct {
	Name     string `json:"name,omitempty"`
	Path     string `json:"path,omitempty"`
	IsDir    bool   `json:"isDir,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}
