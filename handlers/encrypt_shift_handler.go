package handlers

import (
	"atylab-api/models"
	"atylab-api/services"

	"github.com/gin-gonic/gin"
)

type EncryptHandler struct {
	Service services.EncryptService
}

func NewEncryptHandler(svc services.EncryptService) *EncryptHandler {
	return &EncryptHandler{Service: svc}
}

func (h *EncryptHandler) Handle(c *gin.Context) {
	var req models.EncryptShiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	encrypted := h.Service.EncryptCaesarShift(req.Text, req.Vector)
	c.JSON(200, gin.H{"encrypt_string": encrypted})
}
