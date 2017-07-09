package main

//Entidad con sus propiedades
type Cancion struct {
	Id       int `json:"id"`
	Nombre   string `json:"nombre"`
	Artista  string `json:"artista"`
	Duracion int `json:"duracion"`
	Genero   string `json:"genero"`
}

type Listado struct {
	Id          int `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type ListaCancion struct {
	IdLista   int
	IdCancion int
}
//Json Entity Handler
type JsonHandler struct {
	Cancion        []struct {
		Artista  string `json:"artista"`
		Duracion string `json:"duracion"`
		Genero   string `json:"genero"`
		ID       int    `json:"id"`
		Nombre   string `json:"nombre"`
	} `json:"canciones"`
	Listacanciones []struct {
		Canciones []int `json:"canciones"`
		ID        int `json:"id"`
		Nombre    string `json:"nombre"`
	}`json:"listacanciones"`
}
type QueryXml struct {
	Canciones []Cancion `xml:Cancion`
	Listados []Listado `xml:Listacanciones`
}



//Slice de canciones
var Canciones = []Cancion{}

var Listas = []Listado{}

var ListasCanciones = []ListaCancion{}
