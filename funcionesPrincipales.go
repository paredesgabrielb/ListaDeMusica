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
	check(err)
	var cancionDetail []string
	if len(fileCanciones) > 0 {
		for _, line := range fileCanciones {
			cancionDetail = strings.Split(line, "|")
			id, err := strconv.Atoi(cancionDetail[0])
			check(err)
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
	check(err)
	var listaCancionDetail []string
	if len(fileListaCancion) > 0 {
		for _, line := range fileListaCancion {
			listaCancionDetail = strings.Split(line, "|")
			idLista, err := strconv.Atoi(listaCancionDetail[0])
			check(err)
			idCancion, err := strconv.Atoi(listaCancionDetail[1])
			check(err)
			ListasCanciones = append(ListasCanciones, ListaCancion{
				idLista,
				idCancion})
		}

	}
}


//Inicializa data en la lista
func InicializarListas() {
	fileListas, err := readFileLines("files/listas.txt")
	check(err)
	var listaDetail []string
	if len(fileListas) > 0 {
		for _, line := range fileListas {
			listaDetail = strings.Split(line, "|")
			id, err := strconv.Atoi(listaDetail[0])
			check(err)
			Listas = append(Listas, Listado{
				id,
				listaDetail[1],
				listaDetail[2]})
		}
	}
}

//graba los archivos
func saveDataToFile() {
	writeFileCanciones(Canciones, "files/canciones.txt")
	writeFileListas(Listas, "files/listas.txt")
	writeFileListaCancion(ListasCanciones, "files/listasCanciones.txt")
}

func AnadirCancion(cancion Cancion) {
	Canciones = append(Canciones, cancion)
}

func ModificarCancion(id int) {
	reader := bufio.NewReader(os.Stdin)
	nombre := GetInput("Nombre: ", *reader)
	artista := GetInput("Artista: ", *reader)
	Canciones[id].Nombre = nombre
	Canciones[id].Artista = artista
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
	return Listado{
		len(Listas) + 1,
		nombre,
		descripcion}
}

func AnadirLista(lista Listado) {
	Listas = append(Listas, lista)
}

func EliminarLista(id int) {
	Listas = append(Listas[:id], Listas[id+1:]...)
}

func AnadirCancionALista(id int, cancion Cancion) {
	ListasCanciones = append(ListasCanciones, ListaCancion{id, cancion.Id})
}

func EliminarCancion(id int) {
	Canciones = append(Canciones[:id], Canciones[id+1:]...)
}

func EliminarCancionDeLista(idListaCancion int) {
	ListasCanciones = append(ListasCanciones[:idListaCancion], ListasCanciones[idListaCancion+1:]...)
}
