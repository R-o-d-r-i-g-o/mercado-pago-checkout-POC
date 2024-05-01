package models

const CONTENT_TYPE_TABLE_NAME string = "content_types"

type ContentType struct {
	ID      uint    `gorm:"primary_key"`
	Name    string  `gorm:"unique; not null"`
	Content Content `gorm:"foreignkey:TypeID"`
}

func (ContentType) TableName() string {
	return CONTENT_TYPE_TABLE_NAME
}
