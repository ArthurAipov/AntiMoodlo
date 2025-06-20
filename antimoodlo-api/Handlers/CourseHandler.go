package Handlers

import (
	"antimoodlo/Models"
	"antimoodlo/db"
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
	id := parseUint(c.Param("id"))
	var course Models.Course

	if err := DB.DB.First(&course, id).Error; err != nil {
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

	if err := DB.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

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
	id := parseUint(c.Param("id"))
	var course Models.Course

	if err := DB.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	var input Models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.Title = input.Title

	if err := DB.DB.Save(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

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
	id := parseUint(c.Param("id"))
	var course Models.Course

	if err := DB.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	if err := DB.DB.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.Status(http.StatusNoContent)
}
