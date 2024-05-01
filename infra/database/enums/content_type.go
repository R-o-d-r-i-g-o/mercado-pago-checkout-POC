package enums

type ContentType uint

const (
	VIDEO_CONTENT ContentType = iota + 1
)

func (c ContentType) IsEqual(content ContentType) bool {
	return c == content
}
