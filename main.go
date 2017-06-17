package main

import (
	"fmt"
	"os"
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
		
		break
	case 2: // 2- Ir al menú de Listas.
		LimpiarPantalla()
		ImprimirCabecera()
		ImprimirMenuDeListas()
		MenuListas()
		LimpiarPantalla()
		break
	case 3: // 3- Salir.
		os.Exit(3)
		break
	default:
		LimpiarPantalla()
	}
}

func main() {
	InicializarCanciones()
	InicializarListas()
	InicializarListaCancion()
	for{
		iniciar()
	}
}
