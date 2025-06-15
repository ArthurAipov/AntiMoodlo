package main

import (
	"antimoodlo/Handlers"
	"antimoodlo/db"
	_ "antimoodlo/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	DB.InitDB()

	r := gin.Default()

	// Users
	r.GET("/users", Handlers.GetUsers)
	r.GET("/users/:id", Handlers.GetUserByID)
	r.POST("/users", Handlers.CreateUser)
	r.PUT("/users/:id", Handlers.UpdateUser)
	r.DELETE("/users/:id", Handlers.DeleteUser)

	// Courses
	r.GET("/courses", Handlers.GetCourses)
	r.GET("/courses/:id", Handlers.GetCourseByID)
	r.POST("/courses", Handlers.CreateCourse)
	r.PUT("/courses/:id", Handlers.UpdateCourse)
	r.DELETE("/courses/:id", Handlers.DeleteCourse)

	// Quizzes
	r.GET("/quizzes", Handlers.GetQuizzes)
	r.GET("/quizzes/:id", Handlers.GetQuizByID)
	r.POST("/quizzes", Handlers.CreateQuiz)
	r.PUT("/quizzes/:id", Handlers.UpdateQuiz)
	r.DELETE("/quizzes/:id", Handlers.DeleteQuiz)

	// Questions
	r.GET("/quizzes/:id/questions", Handlers.GetQuestionsByQuiz)
	r.POST("/quizzes/:id/questions", Handlers.CreateQuestion)
	r.DELETE("/questions/:id", Handlers.DeleteQuestion)

	// Answer submission
	r.POST("/questions/:id/options", Handlers.AddQuestionOption)
	r.POST("/questions/:id/answers/correct", Handlers.AddCorrectAnswer)
	r.POST("/questions/:id/answers/open", Handlers.AddOpenAnswer)
	r.POST("/questions/:id/answers/match", Handlers.AddMatchPair)

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080")
}
