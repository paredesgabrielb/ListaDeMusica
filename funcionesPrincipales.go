package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//funciones principales
//Inicializa data en las canciones
func InicializarCanciones() {
	fileCanciones, err := readFileLines("files/canciones.txt")
	if(err != nil){
		fmt.Printf("No se ha podido encontrar el archivo de lectura de canciones.\n")
		os.Exit(1)
	}
	var cancionDetail []string
	if len(fileCanciones) > 0 {
		for _, line := range fileCanciones {
			cancionDetail = strings.Split(line, "|")
			id, err := strconv.Atoi(cancionDetail[0])
			if(err != nil){
				fmt.Printf("Ha ocurrido un error en la asignacion de IDs de canciones.\n")
				os.Exit(1)
			}
			duracion, err := strconv.Atoi(cancionDetail[3])
			Canciones = append(Canciones, Cancion{
				id,
				cancionDetail[1],
				cancionDetail[2],
				duracion,
				cancionDetail[4]})
		}

	}
}

//Inicializa ListaCanciones
func InicializarListaCancion() {
	fileListaCancion, err := readFileLines("files/listasCanciones.txt")
	if(err != nil){
		fmt.Printf("No se ha podido encontrar el archivo de lectura de listas.\n")
		os.Exit(1)
	}
	var listaCancionDetail []string
	if len(fileListaCancion) > 0 {
		for _, line := range fileListaCancion {
			listaCancionDetail = strings.Split(line, "|")
			idLista, err := strconv.Atoi(listaCancionDetail[0])
			if(err != nil){
				fmt.Printf("Ha ocurrido un error en la asignacion de IDs de lista de canciones.\n")
				os.Exit(1)
			}
			idCancion, err := strconv.Atoi(listaCancionDetail[1])
			if(err != nil){
				fmt.Printf("Ha ocurrido un error en la asignacion de IDs de Lista de Canciones.\n")
				os.Exit(1)
			}
			ListasCanciones = append(ListasCanciones, ListaCancion{
				idLista,
				idCancion})
		}

	}
}

func InicializarListas() {
	fileListas, err := readFileLines("files/listas.txt")
	if(err != nil){
		fmt.Printf("Ha ocurrido un error en la lectura del archivo de listas.\n")
		os.Exit(1)
	}
	var listaDetail []string
	if len(fileListas) > 0 {
		for _, line := range fileListas {
			listaDetail = strings.Split(line, "|")
			id, err := strconv.Atoi(listaDetail[0])
			if(err != nil){
				fmt.Printf("Ha ocurrido un error en la asignacion de IDs de Listas.\n")
				os.Exit(1)
			}
			Listas = append(Listas, Listado{
				id,
				listaDetail[1],
				listaDetail[2]})
		}
	}
}

// Funciones con canciones

func AnadirCancion(cancion Cancion) {
	Canciones = append(Canciones, cancion)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	writeFileCanciones(Canciones, "files/canciones.txt")
}

func ModificarCancion(id int) {
	reader := bufio.NewReader(os.Stdin)
	nombre := GetInput("Nombre: ", *reader)
	artista := GetInput("Artista: ", *reader)
	Canciones[id].Nombre = nombre
	Canciones[id].Artista = artista
	writeFileCanciones(Canciones, "files/canciones.txt")
}

func BuscarCancionPorGenero(genero string) []Cancion {
	var busqueda = []Cancion{}
	for _, cancion := range Canciones {
		if strings.Contains(strings.ToLower(cancion.Genero), strings.ToLower(genero)) {
			busqueda = append(busqueda, cancion)
		}
	}
	return busqueda
}

func BuscarCancionPorArtista(artista string) []Cancion {
	var busqueda = []Cancion{}
	for _, cancion := range Canciones {
		if strings.Contains(strings.ToLower(cancion.Artista), strings.ToLower(artista)) {
			busqueda = append(busqueda, cancion)
		}
	}
	return busqueda
}

func GenerarCancion() Cancion {
	reader := bufio.NewReader(os.Stdin)
	nombre := GetInput("Nombre: ", *reader)
	artista := GetInput("Artista: ", *reader)
	genero := GetInput("Género: ", *reader)
	var duracion int
	fmt.Print("Duración: ")
	fmt.Scan(&duracion)
	writeFileCanciones(Canciones, "files/canciones.txt")
	return Cancion{
		len(Canciones) + 1,
		nombre,
		artista,
		duracion,
		genero}

}

func GenerarLista() Listado {
	reader := bufio.NewReader(os.Stdin)
	nombre := GetInput("Nombre: ", *reader)
	descripcion := GetInput("Descripción: ", *reader)
	writeFileListas(Listas, "files/listas.txt")
	return Listado{
		len(Listas) + 1,
		nombre,
		descripcion}

}

func AnadirLista(lista Listado) {
	Listas = append(Listas, lista)
	writeFileListas(Listas, "files/listas.txt")
}

func EliminarLista(id int) {
	Listas = append(Listas[:id], Listas[id+1:]...)
	writeFileListas(Listas, "files/listas.txt")
}

func AnadirCancionALista(id int, cancion Cancion) {
	ListasCanciones = append(ListasCanciones, ListaCancion{id, cancion.Id})
	writeFileListaCancion(ListasCanciones, "files/listasCanciones.txt")
}

func EliminarCancion(id int) {
	Canciones = append(Canciones[:id], Canciones[id+1:]...)
	writeFileCanciones(Canciones, "files/canciones.txt")
}

func EliminarCancionDeLista(idListaCancion int) {
	ListasCanciones = append(ListasCanciones[:idListaCancion], ListasCanciones[idListaCancion+1:]...)
	writeFileListaCancion(ListasCanciones, "files/listasCanciones.txt")
}
