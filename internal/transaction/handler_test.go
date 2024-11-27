package transaction

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"otto/vouchers-project/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateTransaction(voucherID uuid.UUID, quantity int) (*models.Transaction, error) {
	args := m.Called(voucherID, quantity)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *MockService) GetTransactionByID(id string) (*models.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	input := models.Transaction{
		VoucherID: uuid.New(),
		Quantity: 2,
	}
	jsonInput, _ := json.Marshal(input)

	mockService := new(MockService)
	mockService.On("CreateTransaction", input.VoucherID, input.Quantity).Return(&input, nil)

	handler := NewHandler(mockService)

	req, _ := http.NewRequest(http.MethodPost, "/v1/api/transaction/redemption", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/v1/api/transaction/redemption", handler.CreateTransaction)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 201, response.Code)
	assert.Equal(t, "Success", response.Status)

	data := response.Data.(map[string]interface{})
	assert.NotEmpty(t, data["id"])
	assert.Equal(t, input.VoucherID.String(), data["voucher_id"])
	assert.Equal(t, float64(input.Quantity), data["quantity"])
}

func TestGetTransactionByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	transactionID := uuid.New().String()
	transactionData := models.Transaction{
		ID:								uuid.Must(uuid.Parse(transactionID)),
		VoucherID:				uuid.New(),
		TotalPointsUsed: 	100,
		Quantity: 				10,
		Status: 					"completed",
	}

	mockService := new(MockService)
	mockService.On("GetTransactionByID", transactionID).Return(&transactionData, nil)

	handler := NewHandler(mockService)

	req, _ := http.NewRequest(http.MethodGet, "/v1/api/transaction/redemption?transactionId="+transactionID, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/v1/api/transaction/redemption", handler.GetTransactionByID)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Success", response.Status)

	data := response.Data.(map[string]interface{})
	assert.Equal(t, transactionID, data["id"])
	assert.Equal(t, transactionData.VoucherID.String(), data["voucher_id"])
	assert.Equal(t, float64(transactionData.TotalPointsUsed), data["total_points_used"])
	assert.Equal(t, float64(transactionData.Quantity), data["quantity"])
	assert.Equal(t, transactionData.Status, data["status"])
}