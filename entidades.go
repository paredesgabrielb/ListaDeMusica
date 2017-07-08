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

type QueryXml struct {
	Canciones []Cancion 'xml:cancion'
	Listados []Listado 'xml:listacanciones'
}

//Slice de canciones
var Canciones = []Cancion{}

var Listas = []Listado{}

var ListasCanciones = []ListaCancion{}
