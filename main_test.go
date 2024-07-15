package main

import (
	"bytes"
	"encoding/json"
	"go-gin-productManagerPro/dto"
	"go-gin-productManagerPro/infra"
	"go-gin-productManagerPro/models"
	"go-gin-productManagerPro/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestDate(db *gorm.DB) {
	items := []models.Product{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト２", SoldOut: true, UserID: 1},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト３", SoldOut: false, UserID: 3},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, items := range items {
		db.Create(&items)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Product{}, &models.User{})

	setupTestDate(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)

	router.ServeHTTP(w, req)

	var res map[string][]models.Product
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["date"]))
}

func TestCreate(t *testing.T) {
	router := setup()
	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateProductInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}

	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+ *token)

	router.ServeHTTP(w, req)

	var res map[string]models.Product
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}

func TestCreateUnAuthorized(t *testing.T) {
	router := setup()

	createItemInput := dto.CreateProductInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}

	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	var res map[string]models.Product
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}