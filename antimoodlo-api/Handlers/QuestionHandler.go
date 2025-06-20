package Handlers

import (
	"antimoodlo/Models"
	"antimoodlo/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Получить все вопросы
// @Tags Questions
// @Produce json
// @Success 200 {array} Models.Question
// @Router /questions [get]
func GetAllQuestions(c *gin.Context) {
	var questions []Models.Question
	DB.DB.Find(&questions)
	c.JSON(http.StatusOK, questions)
}

// @Summary Получить вопросы теста
// @Tags Questions
// @Produce json
// @Param id path int true "ID теста"
// @Success 200 {array} Models.Question
// @Failure 404 {object} Models.ErrorResponse
// @Router /quizzes/{id}/questions [get]
func GetQuestionsByQuiz(c *gin.Context) {
	id := parseUint(c.Param("id"))

	if err := DB.DB.First(&Models.Quiz{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	var questions []Models.Question
	DB.DB.Where("quizid = ?", id).Find(&questions)
	c.JSON(http.StatusOK, questions)
}

// @Summary Добавить вопрос в тест
// @Tags Questions
// @Accept json
// @Produce json
// @Param id path int true "ID теста"
// @Param question body Models.Question true "Вопрос"
// @Success 201 {object} Models.Question
// @Failure 400 {object} Models.ErrorResponse
// @Router /quizzes/{id}/questions [post]
func CreateQuestion(c *gin.Context) {
	quizID := parseUint(c.Param("id"))
	var question Models.Question

	if err := DB.DB.First(&Models.Quiz{}, quizID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quiz not found"})
		return
	}

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.QuizID = quizID

	if err := DB.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	c.JSON(http.StatusCreated, question)
}

// @Summary Обновить вопрос
// @Tags Questions
// @Accept json
// @Produce json
// @Param id path int true "ID вопроса"
// @Param question body Models.Question true "Вопрос"
// @Success 200 {object} Models.Question
// @Failure 404 {object} Models.ErrorResponse
// @Router /questions/{id} [put]
func UpdateQuestion(c *gin.Context) {
	id := parseUint(c.Param("id"))
	var existing Models.Question

	if err := DB.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	var input Models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.DB.First(&Models.Quiz{}, input.QuizID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quiz not found"})
		return
	}

	existing.QuestionText = input.QuestionText
	existing.QuestionTypeID = input.QuestionTypeID
	existing.QuizID = input.QuizID

	if err := DB.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update question"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

// @Summary Удалить вопрос
// @Tags Questions
// @Produce json
// @Param id path int true "ID вопроса"
// @Success 204
// @Failure 404 {object} Models.ErrorResponse
// @Router /questions/{id} [delete]
func DeleteQuestion(c *gin.Context) {
	id := parseUint(c.Param("id"))

	var existing Models.Question
	if err := DB.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if err := DB.DB.Delete(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete question"})
		return
	}

	c.Status(http.StatusNoContent)
}
