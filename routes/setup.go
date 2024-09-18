package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbal13/tugas3hactiv/controllers"
)

func SetupRouter() *gin.Engine {
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "API Tugas Ketiga Hactiv8")
	})
	app.POST("/updatedata", controllers.CreateWaterWind)
	return app
}
