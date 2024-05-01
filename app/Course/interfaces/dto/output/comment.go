package output

import "time"

type CommentOutputDTO struct {
	Hash        string             `json:"hash"`
	Content     string             `json:"content"`
	Reviews     uint               `json:"reviews,omitempty"`
	PublishedAt time.Time          `json:"published_at"`
	Author      AuthorOutputDTO    `json:"author"`
	Children    []CommentOutputDTO `json:"nested_comments,omitempty"`
}

type AuthorOutputDTO struct {
	Name string `json:"name"`
}
