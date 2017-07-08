package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//helpers

func GetInput(message string, reader bufio.Reader) string {
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

func ImprimirCanciones(canciones []Cancion) {
	if len(canciones) == 0 {
		fmt.Println("\nNo se encontraron canciones.")
	} else {
		fmt.Println("\nLista de canciones")
		for _, cancion := range canciones {
			fmt.Printf(strconv.Itoa(cancion.Id) + "- " + cancion.Nombre + ". By: " + cancion.Artista + "\n")
		}
	}
}

func GetListaCancionesByListaId(idLista int) []int {
	var IdCanciones = []int{}
	for i := 0; i < len(ListasCanciones); i++ {
		if ListasCanciones[i].IdLista == idLista {
			IdCanciones = append(IdCanciones, i)
		}
	}
	return IdCanciones
}

func GetCancionEnLista(idLista int, idCancion int) int {
	for i := 0; i < len(ListasCanciones); i++ {
		if ListasCanciones[i].IdLista == idLista && ListasCanciones[i].IdCancion == idCancion {
			return i
		}
	}
	return -1
}

func ImprimirCancionesByLista(id int) {
	var idCanciones = GetListaCancionesByListaId(id)
	if len(idCanciones) == 0 {
		fmt.Println("\nNo se encontraron canciones.")
	} else {
		for _, index := range GetListaCancionesByListaId(id) {
			fmt.Printf("\t" + strconv.Itoa(Canciones[index].Id) + "- " + Canciones[index].Nombre + ". By: " + Canciones[index].Artista + "\n")
		}
	}
}

func ImprimirListas(listas []Listado) {
	if len(listas) == 0 {
		fmt.Println("\nNo se encontraron listas.")
	} else {
		fmt.Println("\nListas de Reproducción")
		for _, lista := range listas {
			fmt.Printf(strconv.Itoa(lista.Id) + "- " + lista.Nombre + " (" + lista.Descripcion + ")\n")
		}
	}
}

func ImprimirCancionesEnLista(canciones []Cancion) {
	if len(canciones) == 0 {
		fmt.Println("\nNo se encontraron listas.")
	} else {
		for _, cancion := range canciones {
			fmt.Printf(strconv.Itoa(cancion.Id) + "- " + cancion.Nombre + ". By: " + cancion.Artista + "\n")
		}
	}
}

func PauseConsole() {
	if runtime.GOOS == "windows" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\n--Presione ENTER para regresar el menu principal--\n")
		reader.ReadString('\r')
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\n--Presione ENTER para regresar el menu principal--\n")
		reader.ReadString('\n')
	}
}
func VolverAlMenu() {
	PauseConsole()
	LimpiarPantalla()
}

func VolverAlMenuCanciones() {
	VolverAlMenu()
	ImprimirCabecera()
	ImprimirMenuDeCanciones()
	MenuCanciones()
}

func VolverAlMenuListas() {
	VolverAlMenu()
	ImprimirMenuDeExportar()
	ImprimirMenuDeListas()
	MenuListas()
}


func VolverAlMenuExportar() {
	VolverAlMenu()
	ImprimirCabecera()
	ImprimirMenuDeExportar()
	MenuExportar()
}

func VolverAlMenuImportar() {
	VolverAlMenu()
	ImprimirCabecera()
	ImprimirMenuDeImportar()
	MenuImportar()
}

func VerificarIdCancion(id int) bool {
	if id <= 0 || id > len(Canciones) {
		return false
	} else {
		return true
	}
}

func VerificarIdCancionesEnLista(idLista, idCancion int) bool {
	if idCancion <= 0 || idCancion > len(Canciones) {
		return false
	} else {
		return true
	}
}

func VerificarIdListas(id int) bool {
	if id <= 0 || id > len(Listas) {
		return false
	} else {
		return true
	}
}

func ReorganizarIds() {
	for i := range Canciones {
		Canciones[i].Id = i + 1
	}
}

func ReorganizarIdsListas() {
	for i := range Listas {
		Listas[i].Id = i + 1
	}
}

func MenuListas() {
	var opcion int

	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		ImprimirListas(Listas)
		VolverAlMenuListas()
		break
	case 2:
		fmt.Println("\nDigite los detalles de la lista:")
		AnadirLista(GenerarLista())
		fmt.Println("\nLista creada satisfactoriamente!")
		VolverAlMenuListas()
		break
	case 3:
		var idCancion, idLista int
		ImprimirListas(Listas)
		fmt.Print("\nDigite el ID de la lista: ")
		fmt.Scan(&idLista)
		ImprimirCanciones(Canciones)
		fmt.Print("\nDigite el ID de la cancion a agregar: ")
		fmt.Scan(&idCancion)
		if VerificarIdCancion(idCancion) {
			AnadirCancionALista(idLista-1, Canciones[idCancion-1])
			fmt.Println("\nCancion añadida satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		VolverAlMenuListas()
		break
	case 4:
		var idCancion, idLista int
		ImprimirListas(Listas)
		fmt.Print("\nDigite el ID de la lista: ")
		fmt.Scan(&idLista)
		ImprimirCancionesByLista(idLista - 1)
		fmt.Print("\nDigite el ID de la cancion a eliminar: ")
		fmt.Scan(&idCancion)
		if VerificarIdCancionesEnLista(idLista, idCancion) {
			EliminarCancionDeLista(GetCancionEnLista(idLista-1, idCancion-1))
			fmt.Println("\nCancion eliminada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		VolverAlMenuListas()
		break
	case 5:
		ImprimirListas(Listas)
		fmt.Print("\nDigite el ID de la lista a buscar: ")
		var id int
		fmt.Scan(&id)
		if VerificarIdListas(id) {
			ImprimirListas([]Listado{Listas[id-1]})
			ImprimirCancionesByLista(id - 1)
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		VolverAlMenuListas()
		break
	case 6:
		ImprimirListas(Listas)
		fmt.Print("\nDigite el ID de la lista a eliminar: ")
		var id int
		fmt.Scan(&id)
		if VerificarIdListas(id) {
			EliminarLista(id - 1)
			ReorganizarIdsListas()
			fmt.Println("\nLista eliminada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		VolverAlMenuListas()
		break
	case 7:
		break
	default:
		VolverAlMenuListas()
	}
}

func MenuCanciones() {
	var opcion int

	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		ImprimirCanciones(Canciones)
		VolverAlMenuCanciones()
		break
	case 2:
		fmt.Println("\nDigite los detalles de la cancion:")
		AnadirCancion(GenerarCancion())
		fmt.Println("\nCancion añadida satisfactoriamente!")
		VolverAlMenuCanciones()
		break
	case 3:
		ImprimirCanciones(Canciones)
		fmt.Print("\nDigite el ID de la cancion a modificar: ")
		var id int
		fmt.Scan(&id)
		if VerificarIdCancion(id) {
			ModificarCancion(id - 1)
			fmt.Println("\nCancion modificada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		VolverAlMenuCanciones()
		break
	case 4:
		ImprimirCanciones(Canciones)
		fmt.Print("\nDigite el ID de la cancion a eliminar: ")
		var id int
		fmt.Scan(&id)
		if VerificarIdCancion(id) {
			EliminarCancion(id - 1)
			ReorganizarIds()
			fmt.Println("\nCancion eliminada satisfactoriamente!")
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		VolverAlMenuCanciones()
		break
	case 5:
		fmt.Print("\nDigite el ID de la cancion a buscar: ")
		var id int
		fmt.Scan(&id)
		if VerificarIdCancion(id) {
			ImprimirCanciones([]Cancion{Canciones[id-1]})
		} else {
			fmt.Print("\nEl ID especificado no existe.")
		}
		VolverAlMenuCanciones()
		break
	case 6:
		fmt.Println("\nIngrese los criterios de busqueda.")
		reader := bufio.NewReader(os.Stdin)
		artista := GetInput("Artista: ", *reader)
		ImprimirCanciones(BuscarCancionPorArtista(artista))
		VolverAlMenuCanciones()
		break
	case 7:
		fmt.Println("\nIngrese los criterios de busqueda.")
		reader := bufio.NewReader(os.Stdin)
		genero := GetInput("Genero: ", *reader)
		ImprimirCanciones(BuscarCancionPorGenero(genero))
		VolverAlMenuCanciones()
		break
	case 8:
		break
	default:
		VolverAlMenuCanciones()
	}
}


func MenuExportar() {
	var opcion int

	fmt.Scan(&opcion)

	switch opcion {
	case 1: // PDF 
		//TODO
		exportToPdf()
		VolverAlMenuExportar()
		break
	case 2: // CSV
		exportToCsv(Canciones, Listas)
		fmt.Println("\nSe ha exportado los datos!")
		VolverAlMenuExportar()
		break
	case 3: //XLSX
		//TODO
		exportToXlsx()
		VolverAlMenuExportar()
		break
	case 4: // JSON
		//TODO
		exportToJson()
		VolverAlMenuExportar()
		break
	case 5: // XML
		//TODO
		//exportToXml()
		VolverAlMenuExportar()
		break
	case 6: // Volver al menu principal
		//TODO
		break
	default:
		VolverAlMenuExportar()
	}
}

func MenuImportar() {
	var opcion int

	fmt.Scan(&opcion)

	switch opcion {
	case 1: // JSON
		//TODO
		importToJson()
		VolverAlMenuImportar()
		break
	case 2: //XML
		//TODO
		VolverAlMenuImportar()
		break
	case 3: // Volver al Menu principal
		break
	default:
		VolverAlMenuImportar()
	}
}
