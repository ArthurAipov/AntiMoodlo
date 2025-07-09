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
	Title        string     `gorm:"column:title" json:"title" binding:"required"`
	MaxGrade     int        `gorm:"column:maxgrade" json:"maxgrade"`
	Duration     int        `gorm:"column:duration" json:"duration"`
	StateID      uint       `gorm:"column:stateid" json:"stateid"`
	StartDate    time.Time  `gorm:"column:startdate" json:"startdate"`
	EndDate      time.Time  `gorm:"column:enddate" json:"enddate"`
	SubmitedDate *time.Time `gorm:"column:submiteddate" json:"submiteddate,omitempty"`
	CourseID     uint       `gorm:"column:courseid" json:"courseid"`
	Description  string     `gorm:"column:description" json:"description"`
	UserID       uint       `gorm:"column:userid" json:"userid"`
}

func (Quiz) TableName() string {
	return "quizzes"
}

type Question struct {
	ID             uint   `gorm:"primaryKey;column:id" json:"id"`
	QuizID         uint   `gorm:"column:quizid" json:"quizid"`
	QuestionText   string `gorm:"column:questiontext" json:"questiontext"`
	QuestionTypeID uint   `gorm:"column:questiontypeid" json:"questiontypeid"`
	ImageURL       string `gorm:"column:imageurl" json:"imageurl"`
	Points         int    `gorm:"column:points" json:"points"`
}

func (Question) TableName() string {
	return "questions"
}

type QuestionOption struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint   `gorm:"column:questionid" json:"questionid"`
	OptionText string `gorm:"column:optiontext" json:"optiontext"`
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

type AllAnswersResponse struct {
	Options        []QuestionOption `json:"options"`
	CorrectAnswers []CorrectAnswer  `json:"correctAnswers"`
	OpenAnswers    []OpenAnswer     `json:"openAnswers"`
	MatchPairs     []MatchPair      `json:"matchPairs"`
}

type Grades struct {
	ID         int     `gorm:"primaryKey;column:id" json:"id"`
	UserId     uint    `gorm:"column:userid" json:"userid"`
	QuizID     uint    `gorm:"column:quizid" json:"quizid"`
	Points     float32 `gorm:"column:points" json:"points"`
	GradeState int     `gorm:"column:completedstate" json:"gradestate"`
}

func (Grades) TableName() string {
	return "grades"
}

type GradeStateUpdate struct {
	GradeState int `json:"gradestate"`
}
