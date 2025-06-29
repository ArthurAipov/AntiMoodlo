package Handlers

import (
	"antimoodlo/Models"
	"antimoodlo/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateGrade godoc
// @Summary Создать оценку
// @Description Создает новую запись в таблице grades
// @Tags grades
// @Accept json
// @Produce json
// @Param grade body Models.Grades true "Grade object"
// @Success 201 {object} Models.Grades
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /grades [post]
func CreateGrade(c *gin.Context) {
	var grade Models.Grades
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.DB.Create(&grade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grade)
}

// GetGrades godoc
// @Summary Получить все оценки
// @Description Возвращает список всех записей из таблицы grades
// @Tags grades
// @Produce json
// @Success 200 {array} Models.Grades
// @Failure 500 {object} map[string]interface{}
// @Router /grades [get]
func GetGrades(c *gin.Context) {
	var grades []Models.Grades
	if err := DB.DB.Find(&grades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grades)
}

// GetGradeByID godoc
// @Summary Получить оценку по ID
// @Description Возвращает одну оценку по ID
// @Tags grades
// @Produce json
// @Param id path int true "Grade ID"
// @Success 200 {object} Models.Grades
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /grades/{id} [get]
func GetGradeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var grade Models.Grades
	if err := DB.DB.First(&grade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
		return
	}

	c.JSON(http.StatusOK, grade)
}

// UpdateGrade godoc
// @Summary Обновить оценку
// @Description Обновляет запись в таблице grades по ID
// @Tags grades
// @Accept json
// @Produce json
// @Param id path int true "Grade ID"
// @Param grade body Models.Grades true "Grade object"
// @Success 200 {object} Models.Grades
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /grades/{id} [put]
func UpdateGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var grade Models.Grades
	if err := DB.DB.First(&grade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
		return
	}

	var input Models.Grades
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade.UserId = input.UserId
	grade.QuizID = input.QuizID
	grade.Points = input.Points

	if err := DB.DB.Save(&grade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grade)
}

// DeleteGrade godoc
// @Summary Удалить оценку
// @Description Удаляет запись в таблице grades по ID
// @Tags grades
// @Produce json
// @Param id path int true "Grade ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /grades/{id} [delete]
func DeleteGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := DB.DB.Delete(&Models.Grades{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade deleted successfully"})
}
