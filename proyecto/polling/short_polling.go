package polling

import (
	"demo/proyecto/entidad"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortPolling(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"newProductAdded": entidad.NewProductAdded})
	entidad.NewProductAdded = false
}
func ConfigurarShortPolling(r *gin.Engine){
	r.GET("/isNewProductAdded", ShortPolling)
}