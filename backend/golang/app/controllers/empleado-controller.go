package controllers

import (
	"net/http"
	"strconv"
	"test-api/app/models"
	"test-api/app/services"

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
	service := services.EmpleadoService{}
	empleados, err := service.FindAll(db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, empleados)
	}

}

//Create empleado
func (c *EmpleadoController) Create(ctx *gin.Context) {
	db := ctx.MustGet("db").(*pg.DB)
	var empleado models.Empleado
	ctx.ShouldBindJSON(&empleado)
	service := services.EmpleadoService{}
	if err := service.Create(empleado, db); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, "Se ha creado el empleado")
	}

}

//Update empleado
func (c *EmpleadoController) Update(ctx *gin.Context) {
	db := ctx.MustGet("db").(*pg.DB)
	var empleado models.Empleado
	err := ctx.ShouldBindJSON(&empleado)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	nss, err := strconv.ParseInt(ctx.Param("nss"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	empleado.Nss = nss
	service := services.EmpleadoService{}
	if err := service.Update(empleado, db); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Se ha actualizado el empleado"})
	}
}
