package models

type EncryptShiftRequest struct {
	Vector int    `json:"vector" binding:"required,min=-100,max=100"`
	Text   string `json:"text" binding:"required"`
}
