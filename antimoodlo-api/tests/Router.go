package tests

import (
	"antimoodlo/Handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/quizzes", Handlers.CreateQuiz)
	router.GET("/quizzes/:id", Handlers.GetQuizByID)
	router.GET("/quizzes", Handlers.GetQuizzes)
	router.PUT("/quizzes/:id", Handlers.UpdateQuiz)

	return router
}
