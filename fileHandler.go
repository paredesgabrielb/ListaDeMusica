package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

<<<<<<< HEAD
=======
func check(e error) {
	if e != nil {
		panic(e)
	}
}

>>>>>>> a1d69821ef1525d3b1e029abe8e4c7194e3a9304
//Recibe la direccion del archivo a leer, retorna un arreglo de string con lo leido
func readFileLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if(err != nil){
		fmt.Print("No existe el archivo de lectura.\n")
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func exportToCsv(canciones []Cancion, listas []Listado) {
	file, err := os.Create("export/export.csv")
	check(err)
	defer file.Close()
	for _, cancion := range canciones {
		id := strconv.Itoa(cancion.Id)
		duracion := strconv.Itoa(cancion.Duracion)
		line := id + "," + cancion.Nombre + "," + cancion.Artista + "," + duracion + "," + cancion.Genero + "\r\n"
		fmt.Fprintf(file, line)
	}
	for _, lista := range listas {
		id := strconv.Itoa(lista.Id)
		line := id + "," + lista.Nombre + "," + lista.Descripcion + "\r\n"
		fmt.Fprintf(file, line)
	}
}

//Recibe un arreglo de canciones, abre la direccion especificada y escribe en el archivo la data
func writeFileCanciones(canciones []Cancion, path string) {
	file, err := os.Create(path)
	if(err != nil){
		fmt.Printf("No se ha podido guardar el archivo de Canciones.\n")
		os.Exit(1)
	}
	defer file.Close()
	for _, cancion := range canciones {
		id := strconv.Itoa(cancion.Id)
		duracion := strconv.Itoa(cancion.Duracion)
		line := id + "|" + cancion.Nombre + "|" + cancion.Artista + "|" + duracion + "|" + cancion.Genero + "\r\n"
		fmt.Fprintf(file, line)
	}
}

//Recibe un arreglo de Listas, abre la direccion especificada y escribe en el archivo la data
func writeFileListas(listas []Listado, path string) {
	file, err := os.Create(path)
	if(err != nil){
		fmt.Printf("No se ha podido guardar el archivo de Listas.\n")
		os.Exit(1)
	}
	defer file.Close()
	for _, lista := range listas {
		id := strconv.Itoa(lista.Id)
		line := id + "|" + lista.Nombre + "|" + lista.Descripcion + "\r\n"
		fmt.Fprintf(file, line)
	}
}

//Recibe un arreglo de Listas, abre la direccion especificada y escribe en el archivo la data
func writeFileListaCancion(listaCancion []ListaCancion, path string) {
	file, err := os.Create(path)
	if(err != nil){
		fmt.Printf("No se ha podido guardar el archivo de Lista de Candiones.")
		os.Exit(1)
	}
	defer file.Close()
	for _, lista := range listaCancion {
		idCancion := strconv.Itoa(lista.IdCancion)
		idLista := strconv.Itoa(lista.IdLista)
		line := idLista + "|" + idCancion + "\r\n"
		fmt.Fprintf(file, line)
	}
}
