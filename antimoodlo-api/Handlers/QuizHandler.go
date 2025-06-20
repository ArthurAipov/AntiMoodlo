package Handlers

import (
	"antimoodlo/Models"
	"antimoodlo/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Получить все тесты
// @Tags Quizzes
// @Produce json
// @Success 200 {array} Models.Quiz
// @Router /quizzes [get]
func GetQuizzes(c *gin.Context) {
	var quizzes []Models.Quiz
	DB.DB.Find(&quizzes)
	c.JSON(http.StatusOK, quizzes)
}

// @Summary Получить тест по ID
// @Tags Quizzes
// @Produce json
// @Param id path int true "ID теста"
// @Success 200 {object} Models.Quiz
// @Failure 404 {object} Models.ErrorResponse
// @Router /quizzes/{id} [get]
func GetQuizByID(c *gin.Context) {
	id := parseUint(c.Param("id"))
	var quiz Models.Quiz

	if err := DB.DB.First(&quiz, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// @Summary Создать тест
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param quiz body Models.Quiz true "Тест"
// @Success 201 {object} Models.Quiz
// @Router /quizzes [post]
func CreateQuiz(c *gin.Context) {
	var quiz Models.Quiz

	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка внешних ключей
	if err := DB.DB.First(&Models.Course{}, quiz.CourseID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}
	if err := DB.DB.First(&Models.State{}, quiz.StateID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State not found"})
		return
	}

	if err := DB.DB.Create(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
		return
	}

	c.JSON(http.StatusCreated, quiz)
}

// @Summary Обновить тест
// @Tags Quizzes
// @Accept json
// @Produce json
// @Param id path int true "ID теста"
// @Param quiz body Models.Quiz true "Тест"
// @Success 200 {object} Models.Quiz
// @Failure 404 {object} Models.ErrorResponse
// @Router /quizzes/{id} [put]
func UpdateQuiz(c *gin.Context) {
	id := parseUint(c.Param("id"))
	var quiz Models.Quiz

	if err := DB.DB.First(&quiz, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	var input Models.Quiz
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.DB.First(&Models.Course{}, input.CourseID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}
	if err := DB.DB.First(&Models.State{}, input.StateID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State not found"})
		return
	}

	quiz.Title = input.Title
	quiz.MaxGrade = input.MaxGrade
	quiz.Duration = input.Duration
	quiz.StateID = input.StateID
	quiz.StartDate = input.StartDate
	quiz.EndDate = input.EndDate
	quiz.SubmitedDate = input.SubmitedDate
	quiz.CourseID = input.CourseID

	if err := DB.DB.Save(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quiz"})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// @Summary Удалить тест
// @Tags Quizzes
// @Produce json
// @Param id path int true "ID теста"
// @Success 204
// @Failure 404 {object} Models.ErrorResponse
// @Router /quizzes/{id} [delete]
func DeleteQuiz(c *gin.Context) {
	id := parseUint(c.Param("id"))
	var quiz Models.Quiz

	if err := DB.DB.First(&quiz, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	if err := DB.DB.Delete(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete quiz"})
		return
	}

	c.Status(http.StatusNoContent)
}
