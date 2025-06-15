package Handlers

import (
	"antimoodlo/db"
	"antimoodlo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Получить всех пользователей
// @Tags Users
// @Produce json
// @Success 200 {array} Models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []Models.User
	DB.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Получить пользователя по ID
// @Tags Users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} Models.User
// @Failure 404 {object} Models.ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user Models.User
	if err := DB.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Создать нового пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param user body Models.User true "Пользователь"
// @Success 201 {object} Models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// @Summary Обновить пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body Models.User true "Пользователь"
// @Success 200 {object} Models.User
// @Failure 404 {object} Models.ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user Models.User
	if err := DB.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input Models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.UserLogin = input.UserLogin
	user.UserPassword = input.UserPassword
	user.UserRole = input.UserRole

	DB.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary Удалить пользователя
// @Tags Users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 204
// @Failure 404 {object} Models.ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user Models.User
	if err := DB.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	DB.DB.Delete(&user)
	c.Status(http.StatusNoContent)
}
