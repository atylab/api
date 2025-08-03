package handlers

import (
	"atylab-api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEncryptHandler_WithMock_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"vector": 3, "text": "abc"}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	mockService := &mocks.MockEncryptService{ReturnValue: "mocked!"}
	handler := NewEncryptHandler(mockService)
	handler.Handle(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "mocked!")
}

func TestEncryptShiftHandlerBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectContains string
	}{
		{
			name:           "missing text",
			body:           `{"vector": 3}`,
			expectedStatus: http.StatusBadRequest,
			expectContains: "error",
		},
		{
			name:           "missing vector",
			body:           `{"text": "abc"}`,
			expectedStatus: http.StatusBadRequest,
			expectContains: "error",
		},
		{
			name:           "vector out of range",
			body:           `{"vector": 999, "text": "abc"}`,
			expectedStatus: http.StatusBadRequest,
			expectContains: "error",
		},
		{
			name:           "invalid json",
			body:           `not a json`,
			expectedStatus: http.StatusBadRequest,
			expectContains: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("POST", "/", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			mockService := &mocks.MockEncryptService{ReturnValue: "mocked!"}
			handler := NewEncryptHandler(mockService)
			handler.Handle(c)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectContains)
		})
	}
}
