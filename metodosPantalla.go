package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
	fmt.Println("3- Ir al menú de exportaciones.")
	fmt.Println("4- Ir al menú de importaciones.")
	fmt.Println("5- Salir.")
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

func ImprimirMenuDeExportar() {
	fmt.Println("1- PDF.")
	fmt.Println("2- CSV.")
	fmt.Println("3- XLSX.")
	fmt.Println("4- JSON.")
	fmt.Println("5- XML.")
	fmt.Println("6- Volver al Menú Principal.")
	fmt.Print("\nElija una opcion => ")
}

func ImprimirMenuDeImportar() {
	fmt.Println("1- JSON.")
	fmt.Println("2- XML.")
	fmt.Println("3- Volver al Menú Principal.")
	fmt.Print("\nElija una opcion => ")
}

func LimpiarPantalla() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
