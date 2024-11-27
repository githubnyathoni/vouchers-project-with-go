package transaction

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

func (h *Handler) CreateTransaction(c *gin.Context) {
	var input struct {
		VoucherID   		uuid.UUID `json:"voucher_id" validate:"required"`
		Quantity				int				`json:"quantity" validate:"required,min=1"`	
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

	transaction, err := h.service.CreateTransaction(input.VoucherID, input.Quantity)
	if err != nil {
		if err.Error() == "Voucher not found" {
			c.JSON(http.StatusNotFound, models.Response{
				Code:   404,
				Status: "Voucher not found",
				Data:   nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.Response{
			Code:   500,
			Status: "Failed",
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Code:   201,
		Status: "Success",
		Data:   transaction,
	})
}

func (h *Handler) GetTransactionByID(c *gin.Context) {
	transactionID := c.Query("transactionId")
	transaction, err := h.service.GetTransactionByID(transactionID)

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:   404,
			Status: "Transaction not found",
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:   200,
		Status: "Success",
		Data:   transaction,
	})
}

func (h *Handler) formatValidationErrors(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}
	return errors
}