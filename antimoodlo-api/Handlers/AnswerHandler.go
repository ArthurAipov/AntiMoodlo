package Handlers

import (
	"antimoodlo/Models"
	"antimoodlo/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddQuestionOption godoc
// @Summary      Добавить вариант ответа
// @Description  Добавляет один из вариантов ответа для вопроса (Multiple/Single Choice)
// @Tags         answers
// @Accept       json
// @Produce      json
// @Param        id      path      int                     true  "ID вопроса"
// @Param        option  body      Models.QuestionOption   true  "Вариант ответа"
// @Success      201     {object}  Models.QuestionOption
// @Failure      400     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /questions/{id}/options [post]
func AddQuestionOption(c *gin.Context) {
	var option Models.QuestionOption
	questionID := parseUint(c.Param("id"))

	if err := DB.DB.First(&Models.Question{}, questionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if err := c.ShouldBindJSON(&option); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	option.QuestionID = questionID
	if err := DB.DB.Create(&option).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create option"})
		return
	}

	c.JSON(http.StatusCreated, option)
}

// AddCorrectAnswer godoc
// @Summary      Добавить правильный ответ
// @Description  Добавляет правильный вариант ответа к вопросу (по ID)
// @Tags         answers
// @Accept       json
// @Produce      json
// @Param        id      path      int                    true  "ID вопроса"
// @Param        answer  body      Models.CorrectAnswer   true  "Правильный ответ"
// @Success      201     {object}  Models.CorrectAnswer
// @Failure      400     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /questions/{id}/answers/correct [post]
func AddCorrectAnswer(c *gin.Context) {
	var answer Models.CorrectAnswer
	questionID := parseUint(c.Param("id"))

	if err := DB.DB.First(&Models.Question{}, questionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer.QuestionID = questionID
	if err := DB.DB.Create(&answer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create correct answer"})
		return
	}

	c.JSON(http.StatusCreated, answer)
}

// AddOpenAnswer godoc
// @Summary      Добавить открытый ответ
// @Description  Добавляет открытый ответ для вопроса с открытой формой
// @Tags         answers
// @Accept       json
// @Produce      json
// @Param        id      path      int                 true  "ID вопроса"
// @Param        answer  body      Models.OpenAnswer   true  "Открытый ответ"
// @Success      201     {object}  Models.OpenAnswer
// @Failure      400     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /questions/{id}/answers/open [post]
func AddOpenAnswer(c *gin.Context) {
	var answer Models.OpenAnswer
	questionID := parseUint(c.Param("id"))

	if err := DB.DB.First(&Models.Question{}, questionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer.QuestionID = questionID
	if err := DB.DB.Create(&answer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create open answer"})
		return
	}

	c.JSON(http.StatusCreated, answer)
}

// AddMatchPair godoc
// @Summary      Добавить пару для сопоставления
// @Description  Добавляет пару "лево-право" для вопросов типа Match
// @Tags         answers
// @Accept       json
// @Produce      json
// @Param        id    path      int              true  "ID вопроса"
// @Param        pair  body      Models.MatchPair  true  "Пара для сопоставления"
// @Success      201   {object}  Models.MatchPair
// @Failure      400   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /questions/{id}/answers/match [post]
func AddMatchPair(c *gin.Context) {
	var pair Models.MatchPair
	questionID := parseUint(c.Param("id"))

	if err := DB.DB.First(&Models.Question{}, questionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if err := c.ShouldBindJSON(&pair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pair.QuestionID = questionID
	if err := DB.DB.Create(&pair).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create match pair"})
		return
	}

	c.JSON(http.StatusCreated, pair)
}

// GetAllAnswers godoc
// @Summary      Получить все ответы на вопрос
// @Description  Возвращает все типы ответов (варианты, правильные, открытые, пары)
// @Tags         answers
// @Produce      json
// @Param        id   path      int  true  "ID вопроса"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /questions/{id}/answers [get]
func GetAllAnswers(c *gin.Context) {
	questionID := parseUint(c.Param("id"))

	if err := DB.DB.First(&Models.Question{}, questionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	var (
		options []Models.QuestionOption
		correct []Models.CorrectAnswer
		open    []Models.OpenAnswer
		match   []Models.MatchPair
	)

	DB.DB.Where("questionid = ?", questionID).Find(&options)
	DB.DB.Where("questionid = ?", questionID).Find(&correct)
	DB.DB.Where("questionid = ?", questionID).Find(&open)
	DB.DB.Where("questionid = ?", questionID).Find(&match)

	if options == nil {
		options = []Models.QuestionOption{}
	}
	if correct == nil {
		correct = []Models.CorrectAnswer{}
	}
	if open == nil {
		open = []Models.OpenAnswer{}
	}
	if match == nil {
		match = []Models.MatchPair{}
	}

	c.JSON(http.StatusOK, gin.H{
		"options":        options,
		"correctAnswers": correct,
		"openAnswers":    open,
		"matchPairs":     match,
	})
}

// UpdateMatchPair godoc
// @Summary      Обновить пару сопоставления
// @Description  Обновляет пару "лево-право" по ID
// @Tags         answers
// @Accept       json
// @Produce      json
// @Param        id    path      int              true  "ID пары сопоставления"
// @Param        pair  body      Models.MatchPair  true  "Обновленная пара сопоставления"
// @Success      200   {object}  Models.MatchPair
// @Failure      400   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /answers/match/{id} [put]
func UpdateMatchPair(c *gin.Context) {
	var pair Models.MatchPair
	pairID := parseUint(c.Param("id"))

	if err := DB.DB.First(&pair, pairID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match pair not found"})
		return
	}

	var input Models.MatchPair
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pair.LeftText = input.LeftText
	pair.RightText = input.RightText
	if err := DB.DB.Save(&pair).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update match pair"})
		return
	}

	c.JSON(http.StatusOK, pair)
}

// UpdateOpenAnswer godoc
// @Summary      Обновить открытый ответ
// @Description  Обновляет открытый ответ по ID
// @Tags         answers
// @Accept       json
// @Produce      json
// @Param        id      path      int                 true  "ID открытого ответа"
// @Param        answer  body      Models.OpenAnswer   true  "Обновленный открытый ответ"
// @Success      200     {object}  Models.OpenAnswer
// @Failure      400     {object}  map[string]interface{}
// @Failure      404     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /answers/open/{id} [put]
func UpdateOpenAnswer(c *gin.Context) {
	var answer Models.OpenAnswer
	answerID := parseUint(c.Param("id"))

	if err := DB.DB.First(&answer, answerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Open answer not found"})
		return
	}

	var input Models.OpenAnswer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer.AnswerText = input.AnswerText
	if err := DB.DB.Save(&answer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update open answer"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// UpdateQuestionOption godoc
// @Summary      Обновить вариант ответа
// @Description  Обновляет текст варианта ответа по ID
// @Tags         answers
// @Accept       json
// @Produce      json
// @Param        id        path      int                     true  "ID варианта ответа"
// @Param        option    body      Models.QuestionOption   true  "Обновленный вариант ответа"
// @Success      200       {object}  Models.QuestionOption
// @Failure      400       {object}  map[string]interface{}
// @Failure      404       {object}  map[string]interface{}
// @Failure      500       {object}  map[string]interface{}
// @Router       /answers/options/{id} [put]
func UpdateQuestionOption(c *gin.Context) {
	var option Models.QuestionOption
	optionID := parseUint(c.Param("id"))

	if err := DB.DB.First(&option, optionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
		return
	}

	var input Models.QuestionOption
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	option.OptionText = input.OptionText
	if err := DB.DB.Save(&option).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update option"})
		return
	}

	c.JSON(http.StatusOK, option)
}
