package Models

type State struct {
	ID   uint   `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (State) TableName() string {
	return "states"
}

type Course struct {
	CourseID uint   `gorm:"primaryKey;column:courseid" json:"courseid"`
	Title    string `gorm:"column:title" json:"title"`
}

func (Course) TableName() string {
	return "course"
}

type Block struct {
	BlockID   uint   `gorm:"primaryKey;column:blockid" json:"blockid"`
	BlockName string `gorm:"column:blockname" json:"blockname"`
	CourseID  uint   `gorm:"column:courseid" json:"courseid"`
}

func (Block) TableName() string {
	return "block"
}

type UserCourse struct {
	UserCourseID uint    `gorm:"primaryKey;column:usercourseid" json:"usercourseid"`
	UserID       uint    `gorm:"column:userid" json:"userid"`
	CourseID     uint    `gorm:"column:courseid" json:"courseid"`
	Grade        float64 `gorm:"column:grade" json:"grade"`
	LetterGrade  string  `gorm:"column:lettergrade" json:"lettergrade"`
}

func (UserCourse) TableName() string {
	return "usercourse"
}
