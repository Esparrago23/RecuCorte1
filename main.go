package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()

	r.POST("/addProduct", addProduct)
	r.GET("/isNewProductAdded", isNewProductAdded)
	r.GET("/countProductInDiscount", countProductInDiscount)

	r.Run(":8080")
}
