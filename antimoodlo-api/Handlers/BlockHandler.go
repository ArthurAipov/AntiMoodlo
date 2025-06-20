package Handlers

import (
	"antimoodlo/Models"
	"antimoodlo/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBlocks godoc
// @Summary      Получить все блоки
// @Description  Возвращает список всех блоков
// @Tags         blocks
// @Produce      json
// @Success      200  {array}   Models.Block
// @Router       /blocks [get]
func GetBlocks(c *gin.Context) {
	var blocks []Models.Block
	DB.DB.Find(&blocks)
	c.JSON(http.StatusOK, blocks)
}

// CreateBlock godoc
// @Summary      Создать новый блок
// @Description  Создает блок, привязанный к курсу
// @Tags         blocks
// @Accept       json
// @Produce      json
// @Param        block  body      Models.Block  true  "Новый блок"
// @Success      201    {object}  Models.Block
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /blocks [post]
func CreateBlock(c *gin.Context) {
	var block Models.Block
	if err := c.ShouldBindJSON(&block); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.DB.First(&Models.Course{}, block.CourseID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	if err := DB.DB.Create(&block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create block"})
		return
	}

	c.JSON(http.StatusCreated, block)
}

// UpdateBlock godoc
// @Summary      Обновить блок
// @Description  Обновляет имя блока и ID курса
// @Tags         blocks
// @Accept       json
// @Produce      json
// @Param        id     path      int           true  "ID блока"
// @Param        block  body      Models.Block  true  "Данные для обновления"
// @Success      200    {object}  Models.Block
// @Failure      400    {object}  map[string]interface{}
// @Failure      404    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /blocks/{id} [put]
func UpdateBlock(c *gin.Context) {
	id := parseUint(c.Param("id"))
	var block Models.Block

	if err := DB.DB.First(&block, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}

	var input Models.Block
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.DB.First(&Models.Course{}, input.CourseID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	block.BlockName = input.BlockName
	block.CourseID = input.CourseID

	if err := DB.DB.Save(&block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update block"})
		return
	}

	c.JSON(http.StatusOK, block)
}

// DeleteBlock godoc
// @Summary      Удалить блок
// @Description  Удаляет блок по ID
// @Tags         blocks
// @Param        id   path      int  true  "ID блока"
// @Success      204  "No Content"
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /blocks/{id} [delete]
func DeleteBlock(c *gin.Context) {
	id := parseUint(c.Param("id"))
	var block Models.Block

	if err := DB.DB.First(&block, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
		return
	}

	if err := DB.DB.Delete(&block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete block"})
		return
	}

	c.Status(http.StatusNoContent)
}
