package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Entidad con sus propiedades
type Cancion struct {
	Id       int
	Nombre   string
	Artista  string
	Duracion int
	Genero   string
}

//Slice de canciones
var canciones = []Cancion{}

//Metodos de pantalla
func imprimirCabecera() {
	fmt.Println("************************************")
	fmt.Println("*  Gestor de listado de Canciones  *")
	fmt.Println("************************************\n")
}

func imprimirMenu() {
	fmt.Println("1- Listar Canciones.")
	fmt.Println("2- Añadir Cancion.")
	fmt.Println("3- Modificar Cancion.")
	fmt.Println("4- Eliminar Cancion.")
	fmt.Println("5- Buscar por Id.")
	fmt.Println("6- Buscar por Artista.")
	fmt.Println("7- Buscar por Genero.")
	fmt.Println("8- Salir.")
	fmt.Print("\nElija una opcion => ")
}

//helpers
func limpiarPantalla() {
	cmd := exec.Command("cmd", "/c", "clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
	//for i := 1; i < 10; i++ {
	//	fmt.Printf("\n\n\n\n\n\n\n\n\n")
	//}
}

func getInput(message string, reader bufio.Reader) string {
	fmt.Print(message)
	var value string
	for {
		input, err := reader.ReadString('\n')
		if strings.TrimSpace(input) == "" {
			continue
		}
		if err == nil {
			value = input
			break
		} else {
			fmt.Printf("Ocurrio un Error: %s\n", err.Error())
		}
	}
	return strings.TrimSpace(value)
}

func imprimirCanciones(canciones []Cancion) {
	if len(canciones) == 0 {
		fmt.Println("\nNo se encontraron canciones.")
	} else {
		fmt.Println("\nLista de canciones")
		for _, cancion := range canciones {
			fmt.Printf(strconv.Itoa(cancion.Id) + "- " + cancion.Nombre + " (" + cancion.Artista + ")\n")
		}
	}
}

func volverAlMenu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--Presione ENTER para regresar el menu principal--\n")
	reader.ReadString('\n')
	limpiarPantalla()
	iniciar()
}

func verificarId(id int) bool {
	if id <= 0 || id > len(canciones) {
		return false
	} else {
		return true
	}
}

func reorganizarIds() {
	for i := range canciones {
		canciones[i].Id = i + 1
	}
}

//funciones principales
func anadirCancion(cancion Cancion) {
	canciones = append(canciones, cancion)
}

func modificarCancion(id int) {
	reader := bufio.NewReader(os.Stdin)
	nombre := getInput("Nombre: ", *reader)
	artista := getInput("Artista: ", *reader)
	canciones[id].Nombre = nombre
	canciones[id].Artista = artista
}

func eliminarCancion(id int) {
	canciones = append(canciones[:id], canciones[id+1:]...)
}

func buscarCancionPorGenero(genero string) []Cancion {
	var busqueda = []Cancion{}

	for _, cancion := range canciones {
		if strings.Contains(strings.ToLower(cancion.Genero), strings.ToLower(genero)) {
			busqueda = append(busqueda, cancion)
		}
	}

	return busqueda
}

func buscarCancionPorArtista(artista string) []Cancion {
	var busqueda = []Cancion{}

	for _, cancion := range canciones {
		if strings.Contains(strings.ToLower(cancion.Artista), strings.ToLower(artista)) {
			busqueda = append(busqueda, cancion)
		}
	}

	return busqueda
}

func generarCancion() Cancion {
	reader := bufio.NewReader(os.Stdin)
	nombre := getInput("Nombre: ", *reader)
	artista := getInput("Artista: ", *reader)
	genero := getInput("Genero: ", *reader)
	var duracion int
	fmt.Print("Duracion: ")
	fmt.Scan(&duracion)
	return Cancion{
		len(canciones) + 1,
		nombre,
		artista,
		duracion,
		genero}
}

func iniciar() {
	var opcion int

	limpiarPantalla()
	imprimirCabecera()
	imprimirMenu()

	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		imprimirCanciones(canciones)
		volverAlMenu()
		break
	case 2:
		fmt.Println("\nDigite los detalles de la cancion:")
		anadirCancion(generarCancion())
		fmt.Println("\nCancion añadida satisfactoriamente!")
		volverAlMenu()
		break
	case 3:
		imprimirCanciones(canciones)
		fmt.Print("\nDigite el ID de la cancion a modificar: ")
		var id int
		fmt.Scan(&id)
		if verificarId(id) {
			modificarCancion(id - 1)
			fmt.Println("\nCancion modificada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenu()
		break
	case 4:
		imprimirCanciones(canciones)
		fmt.Print("\nDigite el ID de la cancion a eliminar: ")
		var id int
		fmt.Scan(&id)
		if verificarId(id) {
			eliminarCancion(id - 1)
			reorganizarIds()
			fmt.Println("\nCancion eliminada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenu()
		break
	case 5:
		fmt.Print("\nDigite el ID de la cancion a buscar: ")
		var id int
		fmt.Scan(&id)
		if verificarId(id) {
			imprimirCanciones([]Cancion{canciones[id-1]})
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenu()
		break
	case 6:
		fmt.Println("\nIngrese los criterios de busqueda.")
		reader := bufio.NewReader(os.Stdin)
		artista := getInput("Artista: ", *reader)
		imprimirCanciones(buscarCancionPorArtista(artista))
		volverAlMenu()
		break
	case 7:
		fmt.Println("\nIngrese los criterios de busqueda.")
		reader := bufio.NewReader(os.Stdin)
		genero := getInput("Genero: ", *reader)
		imprimirCanciones(buscarCancionPorGenero(genero))
		volverAlMenu()
		break
	case 8:
		break
	default:
		iniciar()
	}
}

func main() {
	iniciar()
}
