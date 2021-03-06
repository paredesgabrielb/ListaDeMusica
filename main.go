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
	case 3: // 3- Ir al menú de exportar.
		LimpiarPantalla()
		ImprimirCabecera()
		ImprimirMenuDeExportar()
		MenuExportar()
		LimpiarPantalla()
		break
	case 4: // 4- Ir al menú de importar.
		LimpiarPantalla()
		ImprimirCabecera()
		ImprimirMenuDeImportar()
		MenuImportar()
		LimpiarPantalla()
		break
	case 5: // 5- Salir.
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
	for {
		iniciar()
	}
}
