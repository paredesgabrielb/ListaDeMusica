package main

import (
	"fmt"
)

func iniciar() {
	var opcion int

	ImprimirCabecera()
	ImprimirMenuPrincipal()

	fmt.Scan(&opcion)

	switch opcion {
	case 1: // 1- Ir al menú de Canciones.
		LimpiarPantalla()
		ImprimirCabecera()
		ImprimirMenuDeCanciones()
		MenuCanciones()
		LimpiarPantalla()
		iniciar()
		
		break
	case 2: // 2- Ir al menú de Listas.
		LimpiarPantalla()
		ImprimirCabecera()
		ImprimirMenuDeListas()
		MenuListas()
		LimpiarPantalla()
		iniciar()
		break
	case 3: // 3- Salir.
		break
	default:
		LimpiarPantalla()
		iniciar()
	}
}

func main() {
	iniciar()
}
