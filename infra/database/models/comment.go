package models

import "gorm.io/gorm"

const COMMENT_TABLE_NAME string = "comments"

type Comment struct {
	gorm.Model
	Text      string `gorm:"not null"`
	Hash      string `gorm:"unique; not null"`
	AuthorID  uint   `gorm:"not null"`
	Author    User
	ParentID  *uint `gorm:"index; check:chk_different_parent,parent_id <> id"`
	ContentID uint
	Children  []Comment `gorm:"foreignKey:ParentID"`
}

func (Comment) TableName() string {
	return COMMENT_TABLE_NAME
}
