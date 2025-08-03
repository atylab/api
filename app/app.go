package app

import (
	"atylab-api/handlers"
	"atylab-api/routings"
	"atylab-api/services"

	"github.com/gin-gonic/gin"
)

type App struct {
	EncryptHandler *handlers.EncryptHandler
}

func NewApp() *App {
	return &App{
		EncryptHandler: handlers.NewEncryptHandler(&services.EncryptServiceImpl{}),
	}
}

func (a *App) InitRoutes(r *gin.Engine) {
	routings.EncryptRouting(r, a.EncryptHandler.Handle)
}
