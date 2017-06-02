package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//funciones principales
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
		descripcion,
		[]Cancion{}}
}

func AnadirLista(lista Listado) {
	Listas = append(Listas, lista)
}

func EliminarLista(id int) {
	Listas = append(Listas[:id], Listas[id+1:]...)
}

func AnadirCancionALista(id int, cancion Cancion) {
	cancion.Id = len(Listas[id].Canciones) + 1
	Listas[id].Canciones = append(Listas[id].Canciones, cancion)
}

func EliminarCancion(id int) {
	Canciones = append(Canciones[:id], Canciones[id+1:]...)
}

func EliminarCancionDeLista(id, idCancion int) {
	Listas[id].Canciones = append(Listas[id].Canciones[:idCancion], Listas[id].Canciones[idCancion+1:]...)
}
