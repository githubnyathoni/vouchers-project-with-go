package brand

import (
	"net/http"
	"otto/vouchers-project/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *Handler) CreateBrand(c *gin.Context) {
	var input struct {
		Name string `json:"name" validate:"required,min=3,max=255"`
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

	brand := models.Brand{
		Name: input.Name,
	}

	if err := h.service.CreateBrand(&brand); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:   500,
			Status: "Failed to create brand",
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Code:   201,
		Status: "Success",
		Data: map[string]interface{}{
			"id":   brand.ID.String(),
			"name": brand.Name,
		},
	})
}

func (h *Handler) formatValidationErrors(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}
	return errors
}
