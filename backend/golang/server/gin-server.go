package server

import (
	"test-api/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

//StartServer .
func StartServer(db *pg.DB) {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})
	api := r.Group("/api")
	{
		routes.SetEmpleadoRoutes(api)
	}
	r.Run()
}
