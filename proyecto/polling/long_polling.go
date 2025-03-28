package polling

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"demo/proyecto/entidad"
)

func LongPolling(c *gin.Context){
	for (
		time.Sleep(1 * time.Second)

		c.JSON(http.StatusOK, gin.H{"ProductosConDescuento": entidad.ProductosEnDescuento})
		return
	)
}

func ConfigurarLongPolling(r *gin.Engine){
	r.GET("/countProductInDiscount", LongPolling)
}