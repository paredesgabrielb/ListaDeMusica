package main

import (
	"fmt"
)

//Metodos de pantalla
func ImprimirCabecera() {
	fmt.Println("************************************")
	fmt.Println("*  Gestor de listado de Canciones  *")
	fmt.Println("************************************\n")
}

func ImprimirMenuPrincipal() {
	fmt.Println("1- Ir al menú de Canciones.")
	fmt.Println("2- Ir al menú de Listas.")
	fmt.Println("3- Salir.")
	fmt.Print("\nElija una opcion => ")
}

func ImprimirMenuDeListas() {
	fmt.Println("1- Ver Listas.")
	fmt.Println("2- Crear Lista.")
	fmt.Println("3- Añadir Cancion a Lista.")
	fmt.Println("4- Eliminar Cancion de Lista.")
	fmt.Println("5- Buscar por Id.")
	fmt.Println("6- Eliminar Lista.")
	fmt.Println("7- Volver al Menú Principal.")
	fmt.Print("\nElija una opcion => ")
}

func ImprimirMenuDeCanciones() {
	fmt.Println("1- Ver Canciones.")
	fmt.Println("2- Añadir Cancion.")
	fmt.Println("3- Modificar Cancion.")
	fmt.Println("4- Eliminar Cancion.")
	fmt.Println("5- Buscar por Id.")
	fmt.Println("6- Buscar por Artista.")
	fmt.Println("7- Buscar por Genero.")
	fmt.Println("8- Volver al Menú Principal.")
	fmt.Print("\nElija una opcion => ")
}
