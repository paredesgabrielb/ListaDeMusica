package main

//noInspection
import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"log"
	"github.com/Luxurioust/excelize"
	"github.com/signintech/gopdf"
)

//funciones exportar

//Xlsx tiene un bug en el que si cambias el nombre de la hoja antes de escribir en ella por alguna razon no escribe en la hoja.
//El fix mas rapido es cambiar el nombre de la hoja luego de escribir en ella
func exportToXlsx() {
	var row = ""
	xlsx := excelize.NewFile()

	// Create a new sheet.
	xlsx.NewSheet(2, "Sheet2")

	//Sets the value of the cells
	xlsx.SetCellValue("Sheet1", "A1", "Id")
	xlsx.SetCellValue("Sheet1", "B1", "Nombre")
	xlsx.SetCellValue("Sheet1", "C1", "Artista")
	xlsx.SetCellValue("Sheet1", "D1", "Duracion")
	xlsx.SetCellValue("Sheet1", "E1", "Genero")
	for i := 0; i < len(Canciones); i++ {
		row = strconv.Itoa((i + 2))
		xlsx.SetCellValue("Sheet1", "A"+row, Canciones[i].Id)
		xlsx.SetCellValue("Sheet1", "B"+row, Canciones[i].Nombre)
		xlsx.SetCellValue("Sheet1", "C"+row, Canciones[i].Artista)
		xlsx.SetCellValue("Sheet1", "D"+row, Canciones[i].Duracion)
		xlsx.SetCellValue("Sheet1", "E"+row, Canciones[i].Genero)
	}

	xlsx.SetCellValue("Sheet2", "A1", "Id")
	xlsx.SetCellValue("Sheet2", "B1", "Nombre")
	xlsx.SetCellValue("Sheet1", "C1", "Artista")
	for i := 0; i < len(Listas); i++ {
		row = strconv.Itoa((i + 2))
		xlsx.SetCellValue("Sheet2", "A"+row, Listas[i].Id)
		xlsx.SetCellValue("Sheet2", "B"+row, Listas[i].Nombre)
		xlsx.SetCellValue("Sheet2", "C"+row, Listas[i].Descripcion)
	}

	// Set a new Name to the Given Sheet
	xlsx.SetSheetName("Sheet1", "Canciones")
	xlsx.SetSheetName("Sheet2", "Listas")

	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./exportedFiles/ListaDeMusica.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Archivo exportado")
}

func exportToCsv(canciones []Cancion, listas []Listado) {
	file, err := os.Create("exportedFiles/export.csv")
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
	fmt.Printf("Archivo exportado")
}

func exportToJson(){
	var jExport JsonHandler
	for _, cancion := range Canciones{
		jExport.Cancion = append(jExport.Cancion, struct {
			Artista  string `json:"artista"`
			Duracion string `json:"duracion"`
			Genero   string `json:"genero"`
			ID       int    `json:"id"`
			Nombre   string `json:"nombre"`
		}{
			cancion.Artista,
			strconv.Itoa(cancion.Duracion),
			cancion.Genero,
			cancion.Id,
			cancion.Nombre})
	}
	for _, lista := range Listas{
		var cancionesInLista []int
		cancionesInLista = GetListaCancionesByListaId(lista.Id)
		jExport.Listacanciones = append(jExport.Listacanciones, struct {
			Canciones []int `json:"canciones"`
			ID        int `json:"id"`
			Nombre    string `json:"nombre"`
		}{
			cancionesInLista,
			lista.Id,
			lista.Nombre})
		fmt.Print(cancionesInLista)
	}

	e, _ := json.Marshal(jExport)
	fileExport, _ := os.Create("./exportedFiles/listacanciones.json")
	defer fileExport.Close()
	fmt.Fprintf(fileExport, string(e))
	/*c, _ := json.Marshal(Canciones)
	listas, _ := json.Marshal(Listas)
	ListasCanciones, _ := json.Marshal(ListasCanciones)
	fileCanciones, _ := os.Create("./exportedFiles/Canciones.json")
	defer fileCanciones.Close()
	fmt.Fprintf(fileCanciones, string(c))
	fileListas, _ := os.Create("./exportedFiles/Listas.json")
	defer fileListas.Close()
	fmt.Fprintf(fileListas, string(listas))
	fileListasCanciones, _ := os.Create("./exportedFiles/ListasCanciones.json")
	defer fileListasCanciones.Close()
	fmt.Fprintf(fileListasCanciones, string(ListasCanciones))*/
	fmt.Printf("Archivo exportado")
}

func exportToXML(canciones []Cancion, listas []Listado) {
	file, err := os.Create("exportedFiles/export.xml")
	if err != nil {
		panic(err)
	} // panic if error
	defer file.Close()

	fmt.Fprintf(file, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n")
	fmt.Fprintf(file, "<libreriacanciones>\r\n")

	fmt.Fprintf(file, "<canciones>\r\n")
	for _, cancion := range canciones {
		id := strconv.Itoa(cancion.Id)
		duracion := strconv.Itoa(cancion.Duracion)
		line := "\t<cancion>\r\n\t\t<id>" + id + "</id>\r\n\t\t<nombre>" + cancion.Nombre + "</nombre>\r\n\t\t<artista>" + cancion.Artista + "</artista>\r\n\t\t<duracion>" + duracion + "</duracion>\r\n\t\t<genero>" + cancion.Genero + "</genero>\r\n\t</cancion>\r\n"
		fmt.Fprintf(file, line)
	}
	fmt.Fprintf(file, "</canciones>\r\n\n")
	fmt.Fprintf(file, "<listacanciones>\r\n")
	for _, lista := range listas {
		id := strconv.Itoa(lista.Id)
		line := "\r\n\t<id>" + id + "</id>\r\n\t<nombre>" + lista.Nombre + "</nombre>\r\n\t<descripcion>" + lista.Descripcion + "</descripcion>\r\n"
		fmt.Fprintf(file, line)
		fmt.Fprintf(file, "\t<canciones>\r\n")

		for _, index := range GetListaCancionesByListaId(lista.Id) {
			id := strconv.Itoa(Canciones[index].Id)
			line := "\t\t<id>" + id + "</id>\r\n"
			fmt.Fprintf(file, line)
		}
		fmt.Fprintf(file, "\t</canciones>\r\n")

		fmt.Fprintf(file, "</listacanciones>\r\n")
	}
	fmt.Fprintf(file, "</libreriacanciones>\r\n\n")
}

func exportToPdf() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("time", "./res/times.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("time", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	for i := 0; i < len(Canciones); i++ {
		pdf.Cell(nil, Canciones[i].Nombre) //Canciones[i].Id + "|" +Canciones[i].Nombre+ "|" + Canciones[i].Artista+"|" + Canciones[i].Duracion + "|" + Canciones[i].Genero)
		pdf.Br(20)
	}

	for i := 0; i < len(Listas); i++ {
		pdf.Cell(nil, Listas[i].Nombre) //Listas[i].Id + "|" + Listas[i].Nombre + "|" +Listas[i].Descripcion)
		pdf.Br(20)
	}
	pdf.WritePdf("exportedFiles/Listas.pdf")
	fmt.Printf("Archivo exportado")
}

