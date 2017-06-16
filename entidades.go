package main

//Entidad con sus propiedades
type Cancion struct {
	Id       int
	Nombre   string
	Artista  string
	Duracion int
	Genero   string
}

type Listado struct {
	Id          int
	Nombre      string
	Descripcion string
}

type ListaCancion struct {
	IdLista   int
	IdCancion int
}

//Slice de canciones
var Canciones = []Cancion{}

var Listas = []Listado{}

var ListasCanciones = []ListasCanciones{}
