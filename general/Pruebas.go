package general

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

var numeroLines = 365347
var bloques = 30000
var contadorArchivos = 0
var rutaExportacion = "/Users/jafajardo/Desktop/source/"
var nombreExcel = "/Users/jafajardo/Desktop/source/pruebaexcel.xlsx"
var fileActual *os.File

func Probar() {
	xlsx, _ := excelize.OpenFile(nombreExcel)
	rows, _ := xlsx.GetRows("hojaP")
	for index, row := range rows {

		if _, err := strconv.Atoi(row[15]); err == nil {
			CompletarArechivo(row[15], index)
		}

		//fmt.Println(row[15])
		//CompletarArechivo(row[15], index)
	}
}

func CompletarArechivo(fila string, indice int) {
	if (indice % bloques) == 0 {
		contadorArchivos++
		archivoActual := fmt.Sprintf("%s%s%d%s", rutaExportacion, "archivo_", contadorArchivos, ".txt")
		fileActual, _ = os.Create(archivoActual)
	} else {
		fmt.Fprintln(fileActual, fila)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
