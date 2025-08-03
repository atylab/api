package app_test

import (
	"atylab-api/app"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewAppAndInitRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// App生成とルート初期化
	a := app.NewApp()
	a.InitRoutes(r)

	// テスト用リクエスト
	body := `{"vector": 3, "text": "abc"}`
	req := httptest.NewRequest("POST", "/encrypt/shift", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// 結果検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "encrypt_string")
}
