package DB

import (
	"antimoodlo/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1111 dbname=AntiMoodlo port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(
		&Models.User{},
		&Models.UserRole{},
		&Models.Course{},
		&Models.Block{},
		&Models.UserCourse{},
		&Models.State{},
		&Models.QuestionType{},
		&Models.Quiz{},
		&Models.Question{},
		&Models.QuestionOption{},
		&Models.CorrectAnswer{},
		&Models.OpenAnswer{},
		&Models.MatchPair{},
		&Models.Grades{},
	)
}
