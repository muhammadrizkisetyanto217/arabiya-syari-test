package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}