package output

type (
	ContentOutputDTO struct {
		Hash         string         `json:"hash"`
		TypeID       uint           `json:"type_id"`
		VideoContent VideoOutputDTO `json:"video_content,omitempty"`
	}

	VideoOutputDTO struct {
		URL         string `json:"url"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)
