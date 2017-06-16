package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error){
	if e != nil {
		panic(e);
	}
}

//Recibe la direccion del archivo a leer, retorna un arreglo de string con lo leido
func readFileLines(path string)([]string, error){
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	return lines,scanner.Err()
}

//Recibe un arreglo de canciones, abre la direccion especificada y escribe en el archivo la data
func writeFileCanciones(canciones []Cancion, path string){
	file, err := os.Create(path)
	check(err)
	defer file.Close()
<<<<<<< HEAD
	for _, cancion :=range canciones {
		id := strconv.Itoa(cancion.Id)
		duracion := strconv.Itoa(cancion.Duracion)
		line := id+"|"+cancion.Nombre+"|"+cancion.Artista+"|"+duracion+"|"+cancion.Genero+"\r\n"
		fmt.Fprintf(file, line)
=======
	for _, line :=range data {
		fmt.Fprintf(file, "%s\n",line)
>>>>>>> 0643ad52209b4772a693be37142eb6f70f393338
	}

	// Revisar este codigo
	//content := []byte(data) // convertir el []string a []byte 
	//err = ioutil.WriteFile(filename:"canciones.txt",content, perm: 0644)
	//if err != nil{
	//	fmt.Println("Exploto");
	//}
}

//Recibe un arreglo de Listas, abre la direccion especificada y escribe en el archivo la data
func writeFileListas(listas []Listado, path string){
	file, err := os.Create(path)
	check(err)
	defer file.Close()
	for _, cancion :=range canciones {
		id := strconv.Itoa(cancion.Id)
		duracion := strconv.Itoa(cancion.Duracion)
		line := id+"|"+cancion.Nombre+"|"+cancion.Artista+"|"+duracion+"|"+cancion.Genero+"\r\n"
		fmt.Fprintf(file, line)
	}
}



