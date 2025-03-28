package main

import (
	_"demo/proyecto/entidad"
	"demo/proyecto/polling"
	"demo/proyecto/rutas"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	

	go polling.NotifyLongPolling()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	rutas.RegistrarRutas(r)

	r.Run(":8080")
}
