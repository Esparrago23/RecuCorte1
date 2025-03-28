package polling

import (
	"demo/proyecto/entidad"
	"time"

	"github.com/gin-gonic/gin"
)

var LongPollingChan = make(chan int, 1)

func LongPollingHandler(c *gin.Context) {
	select {
	case count := <-LongPollingChan:
		c.JSON(200, gin.H{"cantidad": count})
	case <-time.After(30 * time.Second):
		c.JSON(408, gin.H{"error": "Timeout"})
	}
}

func NotifyLongPolling() {
	for {
		ProductosMutex.Lock()
		count := 0
		for _, producto := range Productos {
			if producto.Descuento {
				count++
			}
		}
		ProductosMutex.Unlock()

		LongPollingChan <- count
	}
}
func AddProductoHandler(c *gin.Context) {
	var nuevoProducto entidad.Producto
	if err := c.ShouldBindJSON(&nuevoProducto); err != nil {
		c.JSON(400, gin.H{"error": "Error al decodificar JSON"})
		return
	}

	ProductosMutex.Lock()
	Productos = append(Productos, nuevoProducto)
	ProductosMutex.Unlock()

	c.Status(201)
	if nuevoProducto.Descuento {
		count := 0
		for _, producto := range Productos {
			if producto.Descuento {
				count++
			}
		}
		LongPollingChan <- count
	}
}
