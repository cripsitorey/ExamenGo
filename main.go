package main

import "fmt"

var productosNombre []string
var productosPrecios []float64

// productosNombre = make([]string, 3)

func main() {
	opt := 0
	fmt.Println("1 para registrar venta, 2 para mostrar estadisticas y 3 para salir")
	fmt.Scan(&opt)
	switch opt {
	case 1:
		RegistrarVenta()
		main()
	case 2:
		main()
	case 3:
		break
	default:
		fmt.Println("Ingresa unicamente un numero valido")
		main()
	}
}

func RegistrarVenta() {
	productosNombre = make([]string, 3)
	elected := 0
	precio := 0.0
	fmt.Println("Selecciona 1 para arroz, 2 para leche y 3 para pan")
	fmt.Scan(&elected)
	productosNombre[0] = "Arroz"
	productosNombre[1] = "Leche"
	productosNombre[2] = "Pan"
	fmt.Println("Elige el producto que quieres")
	switch elected {
	case 1:
		precio = 1.25
		fmt.Println("El precio es de: 1.25")
	case 2:
		precio = 0.95
		fmt.Println("El precio es de: 0.95")
	case 3:
		precio = 0.5
		fmt.Println("El precio es de: 0.50")
	}
	fmt.Println(productosNombre)
	fmt.Println("Ingresa la cantidad")
	cantidad := 0
	fmt.Scan(&cantidad)

	subtotal := float64(cantidad) * precio
	fmt.Println("El subtotal es: ", subtotal, "Dolares")
}
