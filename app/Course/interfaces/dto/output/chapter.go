package output

type ChapterOutputDTO struct {
	Name     string             `json:"name"`
	Hash     string             `json:"hash"`
	Contents []ContentOutputDTO `json:"content,omitempty"`
}
