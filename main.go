package main

import (
	"fmt"
)

func iniciar() {
	var opcion int

	LimpiarPantalla()
	ImprimirCabecera()
	ImprimirMenuPrincipal()

	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		MenuCanciones()
		iniciar()
		break
	case 2:
		MenuListas()
		iniciar()
		break
	case 3:
		break
	default:
		iniciar()
	}
}

func main() {
	iniciar()
}
