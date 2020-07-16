package routes

import (
	"test-api/app/controllers"

	"github.com/gin-gonic/gin"
)

//SetEmpleadoRoutes .
func SetEmpleadoRoutes(r *gin.RouterGroup) {
	controller := controllers.EmpleadoController{}
	r.GET("/empleados", func(ctx *gin.Context) {
		controller.FindAll(ctx)
	})
}
