package polling

import (
	"demo/proyecto/entidad"
	"sync"
	"github.com/gin-gonic/gin"
)

var Productos []entidad.Producto
var ProductosMutex sync.Mutex
var UltimoProductoIndex int
var ProductosMostrados []entidad.Producto

func ShortPollingHandler(c *gin.Context) {
	ProductosMutex.Lock()
	defer ProductosMutex.Unlock()

	if UltimoProductoIndex == 0 {
		ProductosMostrados = nil
	}

	if UltimoProductoIndex < len(Productos) {
		nuevosProductos := Productos[UltimoProductoIndex:]
		ProductosMostrados = append(ProductosMostrados, nuevosProductos...)
		UltimoProductoIndex = len(Productos)
	}

	c.JSON(200, ProductosMostrados)
}

func ResetProductosMostrados() {
	ProductosMostrados = nil
	UltimoProductoIndex = 0
}
