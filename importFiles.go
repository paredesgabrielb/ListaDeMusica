package main

import (
	"strconv"
	"fmt"
	"encoding/json"
	"strings"
)

func BusquedaCancion(nombre string, artista string, genero string) Cancion{
	var busqueda Cancion
	for _, cancion := range Canciones {
		if (strings.Contains(strings.ToLower(cancion.Nombre), strings.ToLower(nombre)) && strings.Contains(strings.ToLower(cancion.Artista), strings.ToLower(artista)) && strings.Contains(strings.ToLower(cancion.Genero), strings.ToLower(genero))) {
			busqueda = cancion
			break
		}
	}
	return busqueda
}

func BusquedaDuplicadoCancion(nombre string, artista string, genero string) bool{
	for _, cancion := range Canciones {
		if (strings.Contains(strings.ToLower(cancion.Nombre), strings.ToLower(nombre)) && strings.Contains(strings.ToLower(cancion.Artista), strings.ToLower(artista)) && strings.Contains(strings.ToLower(cancion.Genero), strings.ToLower(genero))) {
			return false
		}
	}
	return true
}

func BusquedaDuplicadoLista(nombre string) bool{
	for _, lista := range Listas {
		if (strings.Contains(strings.ToLower(lista.Nombre), strings.ToLower(nombre))) {
			return false
		}
	}
	return true
}

func importToJson(){
	cancionesJson, _ := readFileLines("./exportedFiles/listacanciones.json")
	var cancionData JsonHandler
	var cancionesImportedJson []Cancion
	b := []byte(cancionesJson[0])
	json.Unmarshal(b,&cancionData)
	for _, cancion := range cancionData.Cancion{
		if(BusquedaDuplicadoCancion(cancion.Nombre, cancion.Artista, cancion.Genero)){
			duracion, _ := strconv.Atoi(cancion.Duracion)
			Canciones = append(Canciones, Cancion{
				len(Canciones)+1,
				cancion.Nombre,
				cancion.Artista,
				duracion,
				cancion.Genero})
			cancionesImportedJson = append(cancionesImportedJson, Cancion{
				cancion.ID,
				cancion.Nombre,
				cancion.Artista,
				duracion,
				cancion.Genero})
		}else{
			fmt.Printf("No se ha agregado la cancion "+cancion.Nombre+" del artista "+cancion.Artista+" debido a que esta duplicada\n")
		}
	}

	for _, lista := range cancionData.Listacanciones{
		if(BusquedaDuplicadoLista(lista.Nombre)){
			var idCancionImported int
			var idLista = len(Listas)+1
			if(len(lista.Canciones) > 0){
				for _, id := range lista.Canciones{
					for _, cancion := range cancionesImportedJson {
						if (cancion.Id == id) {
							idCancionImported = BusquedaCancion(cancion.Nombre, cancion.Artista, cancion.Genero).Id
							ListasCanciones = append(ListasCanciones, ListaCancion{
								idLista,
								idCancionImported})
						}
					}
				}

			}

			Listas = append(Listas, Listado{
				idLista,
				lista.Nombre,
				"Importado desde Json"})
		}else{
			fmt.Printf("No se ha agregado la lista "+lista.Nombre+" debido a que esta duplicada\n")
		}
	}
	/*cancionesJson, _ := readFileLines("./exportedFiles/Canciones.json")
	var cancionData []Cancion
	b := []byte(cancionesJson[0])
	json.Unmarshal(b,&cancionData)
	for _, cancion := range cancionData{
		Canciones = append(Canciones, Cancion{
			cancion.Id,
			cancion.Nombre,
			cancion.Artista,
			cancion.Duracion,
			cancion.Genero})
	}
	listasJson, _ := readFileLines("./exportedFiles/Canciones.json")
	var listasData []Listado
	b := []byte(listasJson[0])
	json.Unmarshal(b,&listasData)
	for _, listas := range listasData{

	}
	listasCancionesJson, _ := readFileLines("./exportedFiles/Canciones.json")
	var listaCancionesData []Cancion
	b := []byte(listasCancionesJson[0])
	json.Unmarshal(b,&listaCancionesData)
	for _, listaCancion := range listaCancionesData{


	}*/
	writeFileCanciones(Canciones, "files/canciones.txt")
	writeFileListas(Listas, "files/listas.txt")
	writeFileListaCancion(ListasCanciones, "files/listasCanciones.txt")
	fmt.Printf("Import exitoso.")
}