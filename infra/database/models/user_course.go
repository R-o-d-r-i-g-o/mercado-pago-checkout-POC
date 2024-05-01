package models

const USER_COURSE_TABLE_NAME string = "user_courses"

type UserCourse struct {
	UserID   uint `gorm:"primaryKey;autoIncrement:false"`
	CourseID uint `gorm:"primaryKey;autoIncrement:false"`
	Grade    uint
	Review   string
}

func (UserCourse) TableName() string {
	return USER_COURSE_TABLE_NAME
}
