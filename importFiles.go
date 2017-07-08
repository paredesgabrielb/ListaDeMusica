package main

import (
	"encoding/json"
	"fmt"
)

func importToJson() {
	cancionesJson, _ := readFileLines("./exportedFiles/Canciones.json")
	var cancionData []Cancion
	b := []byte(cancionesJson[0])
	json.Unmarshal(b, &cancionData)
	for _, cancion := range cancionData {
		Canciones = append(Canciones, Cancion{
			cancion.Id,
			cancion.Nombre,
			cancion.Artista,
			cancion.Duracion,
			cancion.Genero})
	}
	/*listasJson, _ := readFileLines("./exportedFiles/Canciones.json")
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
	fmt.Printf("Import exitoso.")
}

func importFromXML(path string) {
	xmlFile, err := os.Open(path + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var q QueryXml
	xml.Unmarshal(b, &q)
	for _, cancion := range q.Canciones {
		Canciones = append(Canciones, Cancion{
			cancion.Id,
			cancion.Nombre,
			cancion.Artista,
			cancion.Duracion,
			cancion.Genero})
	}

	for _, lista := range q.Listados {
		Listas = append(Listas, Listado{
			lista.Id,
			lista.Nombre,
			lista.Descripcion})
	}
}
