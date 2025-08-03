package routings

import (
	"atylab-api/handlers"
	"atylab-api/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEncryptRouting(t *testing.T) {
	expected := map[string]string{
		"/encrypt/shift": "POST",
	}

	r := gin.Default()

	// ハンドラのDIをしてからルーティング登録
	svc := &services.EncryptServiceImpl{}
	handler := handlers.NewEncryptHandler(svc)
	EncryptRouting(r, handler.Handle)

	for path, method := range expected {
		found := false
		for _, route := range r.Routes() {
			if route.Path == path && route.Method == method {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected route %s [%s] to be registered", path, method)
	}
}
