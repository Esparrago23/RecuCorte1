package rutas

import (
	"demo/proyecto/polling"

	"github.com/gin-gonic/gin"
)

func RegistrarRutas(r *gin.Engine) {
	r.GET("/isNewProductAdded", polling.ShortPollingHandler)
	r.GET("/countProductInDiscount", polling.LongPollingHandler)
	r.POST("/addProduct", polling.AddProductoHandler)
}
