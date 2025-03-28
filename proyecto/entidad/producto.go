package entidad

type Producto struct {
	Nombre    string `json:"nombre"`
	Precio    int    `json:"precio"`
	Codigo    string `json:"codigo"`
	Descuento bool   `json:"descuento"`
}

var Productos []Producto
var ProductosEnDescuento int
var NewProductAdded bool
