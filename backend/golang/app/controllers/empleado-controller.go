package controllers

import (
	"net/http"
	"test-api/app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

//EmpleadoControllerI .
type EmpleadoControllerI interface {
	FindAll(ctx *gin.Context) []models.Empleado
}

//EmpleadoController .
type EmpleadoController struct{}

//FindAll empleados
func (c *EmpleadoController) FindAll(ctx *gin.Context) {
	db := ctx.MustGet("db").(*pg.DB)
	//service := services.EmpleadoService{}
	//var empleados []models.Empleado
	res, err := db.Query(&models.Empleado{}, "SELECT * FROM empleados")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, res)
	//	empleados, err := service.FindAll(db)
	//	if err != nil {
	//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	} else {
	//		ctx.JSON(200, empleados)
	//	}

}
