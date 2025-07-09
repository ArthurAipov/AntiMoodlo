package tests

import (
	"antimoodlo/Models"
	"antimoodlo/db"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	DB.InitDB()
	return SetupRouter()
}

func TestCreateQuiz(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Math Quiz",
		"description": "Test your math skills",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response Models.Quiz
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Math Quiz", response.Title)
	assert.Equal(t, "Test your math skills", response.Description)
	assert.NotZero(t, response.ID)

	t.Cleanup(func() {
		DB.DB.Delete(&Models.Quiz{}, response.ID)
	})
}

func TestGetQuizByID(t *testing.T) {
	router := setupTestRouter()

	state := Models.State{Name: "Active"}
	DB.DB.Create(&state)

	course := Models.Course{Title: "Test Course"}
	DB.DB.Create(&course)

	user := Models.User{
		UserLogin: "TestUser",
		UserRole:  3,
	}

	DB.DB.Create(&user)
	quiz := Models.Quiz{
		Title:       "Science Quiz",
		Description: "Test your science knowledge",
		StateID:     state.ID,
		CourseID:    course.CourseID,
		UserID:      user.UserID,
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
	}
	DB.DB.Create(&quiz)

	req, _ := http.NewRequest("GET", "/quizzes/"+fmt.Sprint(quiz.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response Models.Quiz
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, quiz.Title, response.Title)
	assert.Equal(t, quiz.Description, response.Description)
	assert.Equal(t, quiz.ID, response.ID)

	t.Cleanup(func() {
		DB.DB.Delete(&Models.Quiz{}, quiz.ID)
		DB.DB.Delete(&Models.State{}, state.ID)
		DB.DB.Delete(&Models.Course{}, course.CourseID)
		DB.DB.Delete(&Models.User{}, user.UserID)
	})
}

func TestGetQuizByID_NotFound(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/quizzes/99999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateQuiz_MissingTitle(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"description": "Test missing title",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateQuiz_InvalidDate(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Invalid Date Quiz",
		"description": "Test invalid date",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "invalid-date",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateQuiz_MissingBody(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("POST", "/quizzes", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIntegration_CreateAndDeleteQuiz(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Temp Quiz",
		"description": "To be deleted",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdQuiz Models.Quiz
	_ = json.Unmarshal(w.Body.Bytes(), &createdQuiz)

	DB.DB.Delete(&Models.Quiz{}, createdQuiz.ID)

	getReq, _ := http.NewRequest("GET", fmt.Sprintf("/quizzes/%d", createdQuiz.ID), nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	assert.Equal(t, http.StatusNotFound, getW.Code)
}

func TestIntegration_CreateAndRetrieveQuiz(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Integration Quiz",
		"description": "Full flow test",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdQuiz Models.Quiz
	_ = json.Unmarshal(w.Body.Bytes(), &createdQuiz)

	getReq, _ := http.NewRequest("GET", fmt.Sprintf("/quizzes/%d", createdQuiz.ID), nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	assert.Equal(t, http.StatusOK, getW.Code)

	var retrievedQuiz Models.Quiz
	_ = json.Unmarshal(getW.Body.Bytes(), &retrievedQuiz)

	assert.Equal(t, createdQuiz.Title, retrievedQuiz.Title)
	assert.Equal(t, createdQuiz.Description, retrievedQuiz.Description)

	t.Cleanup(func() {
		DB.DB.Delete(&Models.Quiz{}, createdQuiz.ID)
	})
}

func TestCreateQuiz_MissingUserID(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Quiz Without User",
		"description": "Test missing user ID",
		"stateid":     1,
		"courseid":    1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateQuiz_EndDateBeforeStartDate(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Date Error Quiz",
		"description": "End date before start date",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-30T00:00:00Z",
		"enddate":     "2025-06-29T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetQuizByID_InvalidID(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/quizzes/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIntegration_CreateUpdateRetrieveQuiz(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Initial Quiz",
		"description": "Initial description",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdQuiz Models.Quiz
	_ = json.Unmarshal(w.Body.Bytes(), &createdQuiz)

	updateData := map[string]interface{}{
		"title":       "Updated Quiz",
		"description": "Updated Description",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	updateJson, _ := json.Marshal(updateData)
	updateReq, _ := http.NewRequest("PUT", fmt.Sprintf("/quizzes/%d", createdQuiz.ID), bytes.NewBuffer(updateJson))
	updateReq.Header.Set("Content-Type", "application/json")

	updateW := httptest.NewRecorder()
	router.ServeHTTP(updateW, updateReq)

	assert.Equal(t, http.StatusOK, updateW.Code)

	getReq, _ := http.NewRequest("GET", fmt.Sprintf("/quizzes/%d", createdQuiz.ID), nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	assert.Equal(t, http.StatusOK, getW.Code)

	var updatedQuiz Models.Quiz
	_ = json.Unmarshal(getW.Body.Bytes(), &updatedQuiz)

	assert.Equal(t, "Updated Quiz", updatedQuiz.Title)

	t.Cleanup(func() {
		DB.DB.Delete(&Models.Quiz{}, createdQuiz.ID)
	})
}

func TestIntegration_CreateQuiz_InvalidStateID(t *testing.T) {
	router := setupTestRouter()

	quizData := map[string]interface{}{
		"title":       "Invalid State Quiz",
		"description": "StateID does not exist",
		"stateid":     99999,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	jsonValue, _ := json.Marshal(quizData)
	req, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIntegration_CreateMultipleAndListQuizzes(t *testing.T) {
	router := setupTestRouter()

	quiz1 := map[string]interface{}{
		"title":       "Quiz One",
		"description": "First quiz",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-06-29T00:00:00Z",
		"enddate":     "2025-06-30T00:00:00Z",
	}

	quiz2 := map[string]interface{}{
		"title":       "Quiz Two",
		"description": "Second quiz",
		"stateid":     1,
		"courseid":    1,
		"userid":      1,
		"startdate":   "2025-07-01T00:00:00Z",
		"enddate":     "2025-07-02T00:00:00Z",
	}

	jsonValue1, _ := json.Marshal(quiz1)
	req1, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue1))
	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)
	assert.Equal(t, http.StatusCreated, w1.Code)

	var createdQuiz1 Models.Quiz
	_ = json.Unmarshal(w1.Body.Bytes(), &createdQuiz1)

	jsonValue2, _ := json.Marshal(quiz2)
	req2, _ := http.NewRequest("POST", "/quizzes", bytes.NewBuffer(jsonValue2))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusCreated, w2.Code)

	var createdQuiz2 Models.Quiz
	_ = json.Unmarshal(w2.Body.Bytes(), &createdQuiz2)

	getReq, _ := http.NewRequest("GET", "/quizzes", nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	assert.Equal(t, http.StatusOK, getW.Code)

	var quizzes []Models.Quiz
	err := json.Unmarshal(getW.Body.Bytes(), &quizzes)
	assert.NoError(t, err)

	foundQuiz1 := false
	foundQuiz2 := false
	for _, q := range quizzes {
		if q.ID == createdQuiz1.ID {
			foundQuiz1 = true
		}
		if q.ID == createdQuiz2.ID {
			foundQuiz2 = true
		}
	}

	assert.True(t, foundQuiz1, "Quiz One not found in the list")
	assert.True(t, foundQuiz2, "Quiz Two not found in the list")

	t.Cleanup(func() {
		DB.DB.Delete(&Models.Quiz{}, createdQuiz1.ID)
		DB.DB.Delete(&Models.Quiz{}, createdQuiz2.ID)
	})
}
