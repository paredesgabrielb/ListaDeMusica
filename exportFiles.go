package main
//noInspection
import(
	"fmt"
	"os"
	"github.com/Luxurioust/excelize"
	"strconv"
	"encoding/json"
	"log"
  	"github.com/signintech/gopdf"
)

//funciones exportar

//Xlsx tiene un bug en el que si cambias el nombre de la hoja antes de escribir en ella por alguna razon no escribe en la hoja.
//El fix mas rapido es cambiar el nombre de la hoja luego de escribir en ella
func exportXlsx(){
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
	for i:=0; i < len(Canciones); i++ {
		row = strconv.Itoa((i+2))
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
		row = strconv.Itoa((i+2))
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
	PauseConsole()
}

func exportJson(){
	c, _ := json.Marshal(Canciones)
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
	fmt.Fprintf(fileListasCanciones, string(ListasCanciones))
	PauseConsole()
}

func exportPdf(){
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	/*err := pdf.AddTTFFont("HDZB_5", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}*/

	/*err = pdf.SetFont("HDZB_5", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}*/

	for i:=0; i < len(Canciones); i++ {
		pdf.Cell(nil, Canciones[i].Id + "|" +Canciones[i].Nombre+ "|" + Canciones[i].Artista+"|" + Canciones[i].Duracion + "|" + Canciones[i].Genero)
	    pdf.Br(20)
	}

	xlsx.SetCellValue("Sheet2", "A1", "Id")
	xlsx.SetCellValue("Sheet2", "B1", "Nombre")
	xlsx.SetCellValue("Sheet1", "C1", "Artista")
	for i := 0; i < len(Listas); i++ {
		pdf.Cell(nil, Listas[i].Id + "|" + Listas[i].Nombre + row + "|" +Listas[i].Descripcion)
		pdf.Br(20)
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("hello.pdf")
	fmt.Printf("Archivo exportado")
	PauseConsole()
}