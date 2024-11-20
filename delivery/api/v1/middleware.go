package v1

import (
	"bytes"
	"crypto/hmac"
	"io"
	"net/http"

	"github.com/JamshedJ/WalletAPI/config"
	"github.com/JamshedJ/WalletAPI/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

func (ctrl *ControllerV1) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-UserId")
		digest := c.GetHeader("X-Digest")

		if userID == "" || digest == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication headers"})
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		computedDigest := utils.ComputeHMACSHA1(body, config.Get().App.SecretKey)
		if !hmac.Equal([]byte(digest), []byte(computedDigest)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid digest"})
			return
		}

		c.Next()
	}
}
