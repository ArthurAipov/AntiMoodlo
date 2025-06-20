package Models

import "time"

// --- Тип вопроса ---
type QuestionType struct {
	ID   uint   `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (QuestionType) TableName() string {
	return "questiontypes"
}

// --- Тест (Quiz) ---
type Quiz struct {
	ID           uint       `gorm:"primaryKey;column:id" json:"id"`
	Title        string     `gorm:"column:title" json:"title"`
	MaxGrade     int        `gorm:"column:maxgrade" json:"maxgrade"`
	Duration     int        `gorm:"column:duration" json:"duration"`
	StateID      uint       `gorm:"column:stateid" json:"stateid"`
	StartDate    time.Time  `gorm:"column:startdate" json:"startdate"`
	EndDate      time.Time  `gorm:"column:enddate" json:"enddate"`
	SubmitedDate *time.Time `gorm:"column:submiteddate" json:"submiteddate,omitempty"`
	CourseID     uint       `gorm:"column:courseid" json:"courseid"`
}

func (Quiz) TableName() string {
	return "quizzes"
}

// --- Вопрос ---
type Question struct {
	ID             uint   `gorm:"primaryKey;column:id" json:"id"`
	QuizID         uint   `gorm:"column:quizid" json:"quizid"`
	QuestionText   string `gorm:"column:questiontext" json:"questiontext"`
	QuestionTypeID uint   `gorm:"column:questiontypeid" json:"questiontypeid"`
}

func (Question) TableName() string {
	return "questions"
}

// --- Варианты ответа ---
type QuestionOption struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint   `gorm:"column:questionid" json:"questionid"`
	OptionText string `gorm:"column:optiontext" json:"optiontext"`
}

func (QuestionOption) TableName() string {
	return "questionoptions"
}

// --- Правильный ответ ---
type CorrectAnswer struct {
	ID         uint `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint `gorm:"column:questionid" json:"questionid"`
	OptionID   uint `gorm:"column:optionid" json:"optionid"`
}

func (CorrectAnswer) TableName() string {
	return "correctanswers"
}

// --- Открытый ответ ---
type OpenAnswer struct {
	QuestionID uint   `gorm:"primaryKey;column:questionid" json:"questionid"`
	AnswerText string `gorm:"column:answertext" json:"answertext"`
}

func (OpenAnswer) TableName() string {
	return "openanswers"
}

// --- Сопоставление пар ---
type MatchPair struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	QuestionID uint   `gorm:"column:questionid" json:"questionid"`
	LeftText   string `gorm:"column:lefttext" json:"lefttext"`
	RightText  string `gorm:"column:righttext" json:"righttext"`
}

func (MatchPair) TableName() string {
	return "matchpairs"
}
