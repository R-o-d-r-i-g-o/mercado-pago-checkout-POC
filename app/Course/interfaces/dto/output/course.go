package output

type CourseOutputDTO struct {
	Name        string             `json:"name"`
	Hash        string             `json:"hash"`
	Description string             `json:"description"`
	Modules     []ChapterOutputDTO `json:"modules,omitempty"`
}
