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

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// --- Users ---
	r.GET("/users", Handlers.GetUsers)
	r.GET("/users/id/:id", Handlers.GetUserByID)
	r.GET("/users/login/:login", Handlers.GetUserByLogin)
	r.POST("/users", Handlers.CreateUser)
	r.PUT("/users/:id", Handlers.UpdateUser)
	r.DELETE("/users/:id", Handlers.DeleteUser)

	// --- Courses ---
	r.GET("/courses", Handlers.GetCourses)
	r.GET("/courses/:id", Handlers.GetCourseByID)
	r.POST("/courses", Handlers.CreateCourse)
	r.PUT("/courses/:id", Handlers.UpdateCourse)
	r.DELETE("/courses/:id", Handlers.DeleteCourse)

	// --- Quizzes ---
	r.GET("/quizzes", Handlers.GetQuizzes)
	r.GET("/quizzes/:id", Handlers.GetQuizByID)
	r.POST("/quizzes", Handlers.CreateQuiz)
	r.PUT("/quizzes/:id", Handlers.UpdateQuiz)
	r.DELETE("/quizzes/:id", Handlers.DeleteQuiz)

	// --- Questions ---
	r.GET("/questions", Handlers.GetAllQuestions)
	r.GET("/quizzes/:id/questions", Handlers.GetQuestionsByQuiz)
	r.POST("/quizzes/:id/questions", Handlers.CreateQuestion)
	r.PUT("/questions/:id", Handlers.UpdateQuestion)
	r.DELETE("/questions/:id", Handlers.DeleteQuestion)
	r.GET("/questions/:id/answers", Handlers.GetAllAnswers)

	// --- Blocks ---
	r.GET("/blocks", Handlers.GetBlocks)
	r.POST("/blocks", Handlers.CreateBlock)
	r.PUT("/blocks/:id", Handlers.UpdateBlock)
	r.DELETE("/blocks/:id", Handlers.DeleteBlock)

	// --- Answers ---
	r.POST("/questions/:id/options", Handlers.AddQuestionOption)
	r.POST("/questions/:id/answers/correct", Handlers.AddCorrectAnswer)
	r.POST("/questions/:id/answers/open", Handlers.AddOpenAnswer)
	r.POST("/questions/:id/answers/match", Handlers.AddMatchPair)

	// Start server
	r.Run(":8080")
}
