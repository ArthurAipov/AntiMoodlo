package Models

import "time"

type QuestionType struct {
	ID   uint   `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (QuestionType) TableName() string {
	return "questiontypes"
}

type Quiz struct {
	ID           uint       `gorm:"primaryKey;column:id" json:"id"`
	Title        string     `gorm:"column:title" json:"title"`
	MaxGrade     int        `gorm:"column:maxgrade" json:"maxgrade"`
	Duration     int        `gorm:"column:duration" json:"duration"`
	StateID      uint       `gorm:"column:stateid" json:"stateid"`
	StartDate    time.Time  `gorm:"column:startdate" json:"startdate"`
	EndDate      time.Time  `gorm:"column:end_date" json:"end_date"`
	SubmitedDate *time.Time `gorm:"column:submiteddate" json:"submiteddate,omitempty"`
	CourseID     uint       `gorm:"column:course_id" json:"course_id"`
}

func (Quiz) TableName() string {
	return "quizzes"
}

type Question struct {
	ID             uint   `gorm:"primaryKey;column:id" json:"id"`
	QuizID         uint   `gorm:"column:quizid" json:"quizid"`
	QuestionText   string `gorm:"column:questiontext" json:"questiontext"`
	QuestionTypeID uint   `gorm:"column:question_typeid" json:"question_typeid"`
}

func (Question) TableName() string {
	return "questions"
}

type QuestionOption struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint   `gorm:"column:question_id" json:"question_id"`
	OptionText string `gorm:"column:option_text" json:"option_text"`
}

func (QuestionOption) TableName() string {
	return "questionoptions"
}

type CorrectAnswer struct {
	ID         uint `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint `gorm:"column:questionid" json:"questionid"`
	OptionID   uint `gorm:"column:optionid" json:"optionid"`
}

func (CorrectAnswer) TableName() string {
	return "correctanswers"
}

type OpenAnswer struct {
	QuestionID uint   `gorm:"primaryKey;column:questionid" json:"questionid"`
	AnswerText string `gorm:"column:answertext" json:"answertext"`
}

func (OpenAnswer) TableName() string {
	return "openanswers"
}

type MatchPair struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint   `gorm:"column:questionid" json:"questionid"`
	LeftText   string `gorm:"column:lefttext" json:"lefttext"`
	RightText  string `gorm:"column:righttext" json:"righttext"`
}

func (MatchPair) TableName() string {
	return "matchpairs"
}
