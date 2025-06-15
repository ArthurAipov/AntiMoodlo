package Handlers

import (
	"antimoodlo/db"
	"antimoodlo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Получить все курсы
// @Tags Courses
// @Produce json
// @Success 200 {array} Models.Course
// @Router /courses [get]
func GetCourses(c *gin.Context) {
	var courses []Models.Course
	DB.DB.Find(&courses)
	c.JSON(http.StatusOK, courses)
}

// @Summary Получить курс по ID
// @Tags Courses
// @Produce json
// @Param id path int true "ID курса"
// @Success 200 {object} Models.Course
// @Failure 404 {object} Models.ErrorResponse
// @Router /courses/{id} [get]
func GetCourseByID(c *gin.Context) {
	var course Models.Course
	if err := DB.DB.First(&course, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

// @Summary Создать курс
// @Tags Courses
// @Accept json
// @Produce json
// @Param course body Models.Course true "Курс"
// @Success 201 {object} Models.Course
// @Router /courses [post]
func CreateCourse(c *gin.Context) {
	var course Models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.DB.Create(&course)
	c.JSON(http.StatusCreated, course)
}

// @Summary Обновить курс
// @Tags Courses
// @Accept json
// @Produce json
// @Param id path int true "ID курса"
// @Param course body Models.Course true "Курс"
// @Success 200 {object} Models.Course
// @Failure 404 {object} Models.ErrorResponse
// @Router /courses/{id} [put]
func UpdateCourse(c *gin.Context) {
	var course Models.Course
	if err := DB.DB.First(&course, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	var input Models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.Title = input.Title
	DB.DB.Save(&course)
	c.JSON(http.StatusOK, course)
}

// @Summary Удалить курс
// @Tags Courses
// @Produce json
// @Param id path int true "ID курса"
// @Success 204
// @Failure 404 {object} Models.ErrorResponse
// @Router /courses/{id} [delete]
func DeleteCourse(c *gin.Context) {
	var course Models.Course
	if err := DB.DB.First(&course, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	DB.DB.Delete(&course)
	c.Status(http.StatusNoContent)
}
