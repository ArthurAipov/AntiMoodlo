package Handlers

import (
	"antimoodlo/db"
	"antimoodlo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GET /quizzes/:quizid/questions

// @Summary Получить вопросы теста
// @Tags Questions
// @Produce json
// @Param id path int true "ID теста"
// @Success 200 {array} Models.Question
// @Router /quizzes/{id}/questions [get]
func GetQuestionsByQuiz(c *gin.Context) {
	quizID := c.Param("id")
	var questions []Models.Question
	DB.DB.Where("quiz_id = ?", quizID).Find(&questions)
	c.JSON(http.StatusOK, questions)
}

// POST /quizzes/:quizid/questions

// @Summary Добавить вопрос в тест
// @Tags Questions
// @Accept json
// @Produce json
// @Param id path int true "ID теста"
// @Param question body Models.Question true "Вопрос"
// @Success 201 {object} Models.Question
// @Router /quizzes/{id}/questions [post]
func CreateQuestion(c *gin.Context) {
	quizID := c.Param("id")
	var question Models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	question.QuizID = parseUint(quizID)
	DB.DB.Create(&question)
	c.JSON(http.StatusCreated, question)
}

// DELETE /questions/:id

// @Summary Удалить вопрос
// @Tags Questions
// @Produce json
// @Param id path int true "ID вопроса"
// @Success 204
// @Router /questions/{id} [delete]
func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	DB.DB.Delete(&Models.Question{}, id)
	c.Status(http.StatusNoContent)
}

func parseUint(s string) uint {
	id, _ := strconv.ParseUint(s, 10, 32)
	return uint(id)
}
