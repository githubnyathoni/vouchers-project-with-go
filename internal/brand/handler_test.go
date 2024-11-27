package brand

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"otto/vouchers-project/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateBrand(brand *models.Brand) error {
	args := m.Called(brand)
	return args.Error(0)
}

func TestCreateBrand(t *testing.T) {
	gin.SetMode(gin.TestMode)

	input := map[string]string{
		"name": "Test Brand",
	}
	jsonInput, _ := json.Marshal(input)

	mockService := new(MockService)
	mockService.On("CreateBrand", mock.Anything).Return(nil)

	handler := NewHandler(mockService)

	req, _ := http.NewRequest(http.MethodPost, "/v1/api/brand", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/v1/api/brand", handler.CreateBrand)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 201, response.Code)
	assert.Equal(t, "Success", response.Status)

	data := response.Data.(map[string]interface{})
	assert.NotEmpty(t, data["id"])
	assert.Equal(t, "Test Brand", data["name"])

	mockService.AssertExpectations(t)
}