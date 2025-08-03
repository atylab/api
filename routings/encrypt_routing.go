package routings

import (
	"github.com/gin-gonic/gin"
)

func EncryptRouting(r *gin.Engine, handler gin.HandlerFunc) {
	encryptRouter := r.Group("/encrypt")
	encryptRouter.POST("/shift", handler)
}
