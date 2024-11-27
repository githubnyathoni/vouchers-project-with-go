package voucher

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

func (m *MockService) CreateVoucher(name string, costInPoint int, brandID uuid.UUID) (*models.Voucher, error) {
	args := m.Called(name, costInPoint, brandID)
	return args.Get(0).(*models.Voucher), args.Error(1)
}

func (m *MockService) GetVoucherByID(id string) (*models.Voucher, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Voucher), args.Error(1)
}

func (m *MockService) GetAllVoucherByBrand(brandID string) ([]models.Voucher, error) {
	args := m.Called(brandID)
	return args.Get(0).([]models.Voucher), args.Error(1)
}

func TestCreateVoucher(t *testing.T) {
	gin.SetMode(gin.TestMode)

	input := models.Voucher{
		Name:        "Voucher 1",
		CostInPoint: 5000,
		BrandID:     uuid.New(),
	}
	jsonInput, _ := json.Marshal(input)

	mockService := new(MockService)
	mockService.On("CreateVoucher", input.Name, input.CostInPoint, input.BrandID).Return(&input, nil)

	handler := NewHandler(mockService)

	req, _ := http.NewRequest(http.MethodPost, "/v1/api/voucher", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/v1/api/voucher", handler.CreateVoucher)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 201, response.Code)
	assert.Equal(t, "Success", response.Status)

	data := response.Data.(map[string]interface{})
	assert.NotEmpty(t, data["id"])
	assert.Equal(t, input.BrandID.String(), data["brand_id"])
	assert.Equal(t, input.Name, data["name"])
	assert.Equal(t, float64(input.CostInPoint), data["cost_in_point"])
}

func TestGetVoucherByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	voucherID := uuid.New().String()
	voucherData := models.Voucher{
		ID:					 uuid.Must(uuid.Parse(voucherID)),
		Name:        "Voucher 1",
		CostInPoint: 5000,
		BrandID:     uuid.New(),
	}

	mockService := new(MockService)
	mockService.On("GetVoucherByID", voucherID).Return(&voucherData, nil)

	handler := NewHandler(mockService)

	req, _ := http.NewRequest(http.MethodGet, "/v1/api/voucher?id="+voucherID, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/v1/api/voucher", handler.GetVoucherByID)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Success", response.Status)

	data := response.Data.(map[string]interface{})
	assert.NotEmpty(t,  voucherID, data["id"])
	assert.Equal(t, voucherData.BrandID.String(), data["brand_id"])
	assert.Equal(t, voucherData.Name, data["name"])
	assert.Equal(t, float64(voucherData.CostInPoint), data["cost_in_point"])
}

func TestGetAllVoucherByBrand(t *testing.T) {
	gin.SetMode(gin.TestMode)

	brandID := uuid.New().String()
	voucherData := []models.Voucher{
		{
			ID:					 uuid.New(),
			Name:        "Voucher 1",
			CostInPoint: 5000,
			BrandID:     uuid.Must(uuid.Parse(brandID)),
		},
	}

	mockService := new(MockService)
	mockService.On("GetAllVoucherByBrand", brandID).Return(voucherData, nil)

	handler := NewHandler(mockService)

	req, _ := http.NewRequest(http.MethodGet, "/v1/api/voucher/brand?id="+brandID, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/v1/api/voucher/brand", handler.GetAllVoucherByBrand)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Success", response.Status)

	data, ok := response.Data.([]interface{})
	assert.True(t, ok, "Response data is not an array")

	assert.Len(t, data, len(voucherData), "Voucher data length does not match")

	voucher := data[0].(map[string]interface{})
	assert.NotEmpty(t,  voucherData[0].ID, voucher["id"])
	assert.Equal(t, voucherData[0].BrandID.String(), voucher["brand_id"])
	assert.Equal(t, voucherData[0].Name, voucher["name"])
	assert.Equal(t, float64(voucherData[0].CostInPoint), voucher["cost_in_point"])
}