package main

import (
	"bufio"
	"fmt"
	"os"
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

type Listado struct {
	Id          int
	Nombre      string
	Descripcion string
	Canciones   []Cancion
}

//Slice de canciones
var canciones = []Cancion{}

var listas = []Listado{}

//Metodos de pantalla
func imprimirCabecera() {
	fmt.Println("************************************")
	fmt.Println("*  Gestor de listado de Canciones  *")
	fmt.Println("************************************\n")
}

func imprimirMenuPrincipal() {
	fmt.Println("1- Ir al menú de Canciones.")
	fmt.Println("2- Ir al menú de Listas.")
	fmt.Println("3- Salir.")
	fmt.Print("\nElija una opcion => ")
}

func imprimirMenuDeListas() {
	fmt.Println("1- Ver Listas.")
	fmt.Println("2- Crear Lista.")
	fmt.Println("3- Añadir Cancion a Lista.")
	fmt.Println("4- Eliminar Cancion de Lista.")
	fmt.Println("5- Buscar por Id.")
	fmt.Println("6- Eliminar Lista.")
	fmt.Println("7- Volver al Menú Principal.")
	fmt.Print("\nElija una opcion => ")
}

func imprimirMenuDeCanciones() {
	fmt.Println("1- Ver Canciones.")
	fmt.Println("2- Añadir Cancion.")
	fmt.Println("3- Modificar Cancion.")
	fmt.Println("4- Eliminar Cancion.")
	fmt.Println("5- Buscar por Id.")
	fmt.Println("6- Buscar por Artista.")
	fmt.Println("7- Buscar por Genero.")
	fmt.Println("8- Volver al Menú Principal.")
	fmt.Print("\nElija una opcion => ")
}

//helpers
func limpiarPantalla() {
	//cmd := exec.Command("cmd", "/c", "cls")
	//cmd.Stdout = os.Stdout
	//cmd.Run()
	for i := 1; i < 10; i++ {
		fmt.Printf("\n\n\n\n\n\n\n\n\n")
	}
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
			fmt.Printf(strconv.Itoa(cancion.Id) + "- " + cancion.Nombre + ". By: " + cancion.Artista + "\n")
		}
	}
}

func imprimirListas(listas []Listado) {
	if len(listas) == 0 {
		fmt.Println("\nNo se encontraron listas.")
	} else {
		fmt.Println("\nListas de Reproducción")
		for _, lista := range listas {
			fmt.Printf(strconv.Itoa(lista.Id) + "- " + lista.Nombre + " (" + lista.Descripcion + ")\n")
		}
	}
}

func imprimirCancionesEnLista(canciones []Cancion) {
	if len(canciones) == 0 {
		fmt.Println("\nNo se encontraron listas.")
	} else {
		for _, cancion := range canciones {
			fmt.Printf(strconv.Itoa(cancion.Id) + "- " + cancion.Nombre + ". By: " + cancion.Artista + "\n")
		}
	}
}

func volverAlMenu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--Presione ENTER para regresar el menu principal--\n")
	reader.ReadString('\n')
	limpiarPantalla()
}

func volverAlMenuCanciones() {
	volverAlMenu()
	menuCanciones()
}

func volverAlMenuListas() {
	volverAlMenu()
	menuListas()
}

func verificarIdCancion(id int) bool {
	if id <= 0 || id > len(canciones) {
		return false
	} else {
		return true
	}
}

func verificarIdListas(id int) bool {
	if id <= 0 || id > len(listas) {
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

func reorganizarIdsListas() {
	for i := range listas {
		listas[i].Id = i + 1
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
	genero := getInput("Género: ", *reader)
	var duracion int
	fmt.Print("Duración: ")
	fmt.Scan(&duracion)
	return Cancion{
		len(canciones) + 1,
		nombre,
		artista,
		duracion,
		genero}
}

func generarLista() Listado {
	reader := bufio.NewReader(os.Stdin)
	nombre := getInput("Nombre: ", *reader)
	descripcion := getInput("Descripción: ", *reader)
	return Listado{
		len(listas) + 1,
		nombre,
		descripcion,
		[]Cancion{}}
}

func anadirLista(lista Listado) {
	listas = append(listas, lista)
}

func eliminarLista(id int) {
	listas = append(listas[:id], listas[id+1:]...)
}

func anadirCancionALista(id int, cancion Cancion) {
	cancion.Id = len(listas[id].Canciones) + 1
	listas[id].Canciones = append(listas[id].Canciones, cancion)
}

func eliminarCancion(id int) {
	canciones = append(canciones[:id], canciones[id+1:]...)
}

func eliminarCancionDeLista(id int, cancion Cancion) {
	listas[id].Canciones = append(listas[id].Canciones[:cancion.Id], listas[id].Canciones[cancion.Id+1:]...)
}

func menuListas() {
	var opcion int

	limpiarPantalla()
	imprimirCabecera()
	imprimirMenuDeListas()

	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		imprimirListas(listas)
		volverAlMenuListas()
		break
	case 2:
		fmt.Println("\nDigite los detalles de la lista:")
		anadirLista(generarLista())
		fmt.Println("\nLista creada satisfactoriamente!")
		volverAlMenuListas()
		break
	case 3:
		var idCancion, idLista int
		imprimirListas(listas)
		fmt.Print("\nDigite el ID de la lista: ")
		fmt.Scan(&idLista)
		imprimirCanciones(canciones)
		fmt.Print("\nDigite el ID de la cancion a agregar: ")
		fmt.Scan(&idCancion)
		if verificarIdCancion(idCancion) {
			anadirCancionALista(idLista, canciones[idCancion-1])
			fmt.Println("\nCancion anadida satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenuListas()
		break
	case 4:
		//TODO
		volverAlMenuListas()
		break
	case 5:
		fmt.Print("\nDigite el ID de la lista a buscar: ")
		var id int
		fmt.Scan(&id)
		if verificarIdListas(id) {
			imprimirListas([]Listado{listas[id-1]})
			imprimirCancionesEnLista(listas[id-1].Canciones)
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenuListas()
		break
	case 6:
		imprimirListas(listas)
		fmt.Print("\nDigite el ID de la lista a eliminar: ")
		var id int
		fmt.Scan(&id)
		if verificarIdListas(id) {
			eliminarLista(id - 1)
			reorganizarIdsListas()
			fmt.Println("\nLista eliminada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenuListas()
		break
	case 7:
		break
	default:
		volverAlMenuListas()
	}
}

func menuCanciones() {
	var opcion int

	limpiarPantalla()
	imprimirCabecera()
	imprimirMenuDeCanciones()

	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		imprimirCanciones(canciones)
		volverAlMenuCanciones()
		break
	case 2:
		fmt.Println("\nDigite los detalles de la cancion:")
		anadirCancion(generarCancion())
		fmt.Println("\nCancion añadida satisfactoriamente!")
		volverAlMenuCanciones()
		break
	case 3:
		imprimirCanciones(canciones)
		fmt.Print("\nDigite el ID de la cancion a modificar: ")
		var id int
		fmt.Scan(&id)
		if verificarIdCancion(id) {
			modificarCancion(id - 1)
			fmt.Println("\nCancion modificada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenuCanciones()
		break
	case 4:
		imprimirCanciones(canciones)
		fmt.Print("\nDigite el ID de la cancion a eliminar: ")
		var id int
		fmt.Scan(&id)
		if verificarIdCancion(id) {
			eliminarCancion(id - 1)
			reorganizarIds()
			fmt.Println("\nCancion eliminada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenuCanciones()
		break
	case 5:
		fmt.Print("\nDigite el ID de la cancion a buscar: ")
		var id int
		fmt.Scan(&id)
		if verificarIdCancion(id) {
			imprimirCanciones([]Cancion{canciones[id-1]})
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		volverAlMenuCanciones()
		break
	case 6:
		fmt.Println("\nIngrese los criterios de busqueda.")
		reader := bufio.NewReader(os.Stdin)
		artista := getInput("Artista: ", *reader)
		imprimirCanciones(buscarCancionPorArtista(artista))
		volverAlMenuCanciones()
		break
	case 7:
		fmt.Println("\nIngrese los criterios de busqueda.")
		reader := bufio.NewReader(os.Stdin)
		genero := getInput("Genero: ", *reader)
		imprimirCanciones(buscarCancionPorGenero(genero))
		volverAlMenuCanciones()
		break
	case 8:
		break
	default:
		volverAlMenuCanciones()
	}
}

func iniciar() {
	var opcion int

	limpiarPantalla()
	imprimirCabecera()
	imprimirMenuPrincipal()

	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		menuCanciones()
		iniciar()
		break
	case 2:
		menuListas()
		iniciar()
		break
	case 3:
		break
	default:
		iniciar()
	}
}

func main() {
	iniciar()
}
