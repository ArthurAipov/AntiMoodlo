package Handlers

import (
	"antimoodlo/db"
	"antimoodlo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /questions/:id/options

// @Summary Добавить вариант ответа
// @Tags Answers
// @Accept json
// @Produce json
// @Param id path int true "ID вопроса"
// @Param option body Models.QuestionOption true "Вариант"
// @Success 201 {object} Models.QuestionOption
// @Router /questions/{id}/options [post]
func AddQuestionOption(c *gin.Context) {
	var option Models.QuestionOption
	questionID := c.Param("id")

	if err := c.ShouldBindJSON(&option); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	option.QuestionID = parseUint(questionID)
	DB.DB.Create(&option)
	c.JSON(http.StatusCreated, option)
}

// POST /questions/:id/answers/correct

// @Summary Добавить правильный ответ
// @Tags Answers
// @Accept json
// @Produce json
// @Param id path int true "ID вопроса"
// @Param answer body Models.CorrectAnswer true "Ответ"
// @Success 201 {object} Models.CorrectAnswer
// @Router /questions/{id}/answers/correct [post]
func AddCorrectAnswer(c *gin.Context) {
	var answer Models.CorrectAnswer
	questionID := c.Param("id")

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer.QuestionID = parseUint(questionID)
	DB.DB.Create(&answer)
	c.JSON(http.StatusCreated, answer)
}

// POST /questions/:id/answers/open

// @Summary Добавить открытый ответ
// @Tags Answers
// @Accept json
// @Produce json
// @Param id path int true "ID вопроса"
// @Param answer body Models.OpenAnswer true "Ответ"
// @Success 201 {object} Models.OpenAnswer
// @Router /questions/{id}/answers/open [post]
func AddOpenAnswer(c *gin.Context) {
	var answer Models.OpenAnswer
	questionID := c.Param("id")

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer.QuestionID = parseUint(questionID)
	DB.DB.Create(&answer)
	c.JSON(http.StatusCreated, answer)
}

// POST /questions/:id/answers/match

// @Summary Добавить пару для сопоставления
// @Tags Answers
// @Accept json
// @Produce json
// @Param id path int true "ID вопроса"
// @Param pair body Models.MatchPair true "Пара"
// @Success 201 {object} Models.MatchPair
// @Router /questions/{id}/answers/match [post]
func AddMatchPair(c *gin.Context) {
	var pair Models.MatchPair
	questionID := c.Param("id")

	if err := c.ShouldBindJSON(&pair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pair.QuestionID = parseUint(questionID)
	DB.DB.Create(&pair)
	c.JSON(http.StatusCreated, pair)
}
