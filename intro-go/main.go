// Definicion del paquete
package main

// Dependencias
import (
	"encoding/json"
	"fmt"
)

// Para instalar paquetes externos se debe crear un go.mod y un go.sum.
//   Para ello debemos correr los comandos:
//  go mod init 'nombre del proyecto'
//  go mod tidy

// Estructuras
type Producto struct {
	Titulo      string  `json:"titulo"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Disponible  bool    `json:"disponible"`
}

// Funciones
func main() {
	fmt.Println("Intro Go")

	producto1 := Producto{
		Titulo:      "TV 50 pulgadas",
		Descripcion: "TV Samsung 50 pulgadas Oferta",
		Precio:      1000000,
		Disponible:  true,
	}

	// Muestra producto
	bytes, err := json.Marshal(producto1)
	fmt.Println(string(bytes))
	fmt.Println(err)

	// Aplicar descuento
	// productoConDescuento := aplicarDescuento(producto1)
	aplicarDescuento2(&producto1)

	// Muestra producto con descuento
	bytes, err = json.Marshal(producto1)
	fmt.Println(string(bytes))
	fmt.Println(err)
}

// Ejercicio: Implementar función aplicarDescuento
// Que recibe un producto
// y rebaja su precio al 50%
// Mostrar el producto por la terminal

// Opcion 1
func aplicarDescuento(producto Producto) Producto {
	producto.Precio = producto.Precio / 2
	return producto
	// return Producto{
	// 	Titulo:      producto.Titulo,
	// 	Descripcion: producto.Descripcion,
	// 	Precio:      producto.Precio / 2,
	// 	Disponible:  producto.Disponible,
	// }
}

// Opcion 2
func aplicarDescuento2(producto *Producto) {
	producto.Precio = producto.Precio / 2
}
