package v1

import (
	"bytes"
	"crypto/hmac"
	"io"
	"net/http"

	"github.com/JamshedJ/WalletAPI/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *ControllerV1) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		partnerIDStr := c.GetHeader("X-UserId")
		partnerID, err := uuid.Parse(partnerIDStr)
		if err != nil {
			handleError(c, err)
			return
		}
		digest := c.GetHeader("X-Digest")

		if partnerID == uuid.Nil || digest == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication headers"})
			return
		}

		partner, err := ctrl.Services.Wallet.GetPartnerByID(c, partnerID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid partner"})
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		computedDigest := utils.ComputeHMACSHA1(body, partner.SecretKey)
		if !hmac.Equal([]byte(digest), []byte(computedDigest)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid digest"})
			return
		}

		c.Set("partnerID", partnerIDStr)
		c.Next()
	}
}
