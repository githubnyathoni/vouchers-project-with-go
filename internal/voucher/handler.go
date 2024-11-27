package voucher

import (
	"net/http"
	"otto/vouchers-project/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Handler struct {
	service 	Service
	validator *validator.Validate
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *Handler) CreateVoucher(c *gin.Context) {
	var input struct {
		Name        string    `json:"name" validate:"required,min=3,max=255"`
		CostInPoint int       `json:"cost_in_point" validate:"required,min=1"`
		BrandID     uuid.UUID `json:"brand_id" validate:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:   400,
			Status: "Invalid request",
			Data:   nil,
		})
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:   400,
			Status: "Validation error",
			Data:   gin.H{"errors": h.formatValidationErrors(err)},
		})
		return
	}

	voucher, err := h.service.CreateVoucher(input.Name, input.CostInPoint, input.BrandID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:   500,
			Status: "Failed to create voucher",
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusAccepted, models.Response{
		Code:   201,
		Status: "Success",
		Data:   voucher,
	})
}

func (h *Handler) GetVoucherByID(c *gin.Context) {
	id := c.Query("id")
	voucher, err := h.service.GetVoucherByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:   404,
			Status: "Voucher not found",
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:   200,
		Status: "Success",
		Data:   voucher,
	})
}

func (h *Handler) GetAllVoucherByBrand(c *gin.Context) {
	brandID := c.Query("id")
	vouchers, err := h.service.GetAllVoucherByBrand(brandID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:   500,
			Status: "Failed to fetch vouchers",
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:   200,
		Status: "Success",
		Data:   vouchers,
	})
}

func (h *Handler) formatValidationErrors(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}
	return errors
}