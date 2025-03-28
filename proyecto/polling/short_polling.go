package polling

import (
	"demo/proyecto/entidad"
	"sync"
	"github.com/gin-gonic/gin"
)

var Productos []entidad.Producto
var ProductosMutex sync.Mutex

func ShortPollingHandler(c *gin.Context) {
	ProductosMutex.Lock()
	defer ProductosMutex.Unlock()
	c.JSON(200, Productos)
}
