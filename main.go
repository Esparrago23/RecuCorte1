package main

import (
	"demo/proyecto/polling"
	"demo/proyecto/rutas"
	"demo/proyecto/entidad"
"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	polling.Productos = []entidad.Producto{
		{"Laptop", 15000, "LAP123", true},
		{"Celular", 7000, "CEL456", false},
	}

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

