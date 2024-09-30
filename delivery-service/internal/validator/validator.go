package validator

import (
	"delivery-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func RegisterCustomValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("deliverystatus", validateDeliveryStatus)
	}
}

func validateDeliveryStatus(fl validator.FieldLevel) bool {
	status, ok := fl.Field().Interface().(models.DeliveryStatus)
	if !ok {
		return false
	}
	return status.IsValid()
}

func ValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("validator", validate)
		c.Next()
	}
}

func Validate(obj interface{}) error {
	return validate.Struct(obj)
}

func HandleValidationErrors(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var validationErrors []string
		for _, e := range errs {
			validationErrors = append(validationErrors, e.Field()+" "+e.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
