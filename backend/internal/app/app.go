package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/harshh-2/slink/internal/config"
	"github.com/harshh-2/slink/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
    Config *config.Config
    DB     *pgxpool.Pool
    Router *gin.Engine
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() error {
	var err error
	a.Config, err = config.Load()
	if err != nil {
		return err
	}
	a.DB, err = db.NewPostgresPool(a.Config)
	if err != nil {
		return err
	}
	a.Router = gin.Default()
	a.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	return a.Router.Run(":" + a.Config.Port)
}