package v1

import (
	"errors"
	"net/http"

	"github.com/JamshedJ/WalletAPI/domain/errs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func errorMsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	default:
		return fe.Error()
	}
}

func handleError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]APIValidationError, len(ve))
		for i, v := range ve {
			out[i] = APIValidationError{Field: v.Field(), Message: errorMsgForTag(v)}
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}

	var errWithCode *errs.ErrorWithCode
	if errors.As(err, &errWithCode) {
		c.JSON(errWithCode.Code, gin.H{"error": errWithCode.Error()})
		return
	}

	c.JSON(400, gin.H{"error": err.Error()})
}
