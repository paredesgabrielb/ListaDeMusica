package main

import(
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error){
	if e != nil {
		panic(e);
	}
}

//Recibe la direccion del archivo a leer, retorna un slice de string con lo leido
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

//Recibe un slice de strings y crea un archivo y escribe en el la data
//TIENE BUG Y NO SEPARA LAS LINEAS, NO ACEPTA EL ESCAPE \N
func writeFile(data []string){
	file, err := os.Create("canciones.txt")
	check(err)
	defer file.Close()
	for _, line :=range data {
		fmt.Fprintf(file, line+"\t")
	}
}


