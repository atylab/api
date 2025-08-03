package tests

import (
	"atylab-api/handlers"
	"atylab-api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEncryptShiftRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// モックサービスを使ってハンドラを生成
	mockService := &mocks.MockEncryptService{ReturnValue: "def"}
	encryptHandler := handlers.NewEncryptHandler(mockService)

	// テスト用にルーティングだけ直接定義（routings.InitRoutes に依存しない）
	encryptRouter := r.Group("/encrypt")
	encryptRouter.POST("/shift", encryptHandler.Handle)

	// リクエストを作成して送信
	body := `{"vector": 3, "text": "abc"}`
	req := httptest.NewRequest("POST", "/encrypt/shift", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 結果を検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "def") // モックで返した値
}
