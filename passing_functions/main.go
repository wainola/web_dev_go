// Pasando funciones a los templates para ocuparlas
package main

import "text/template"

import "os"
import "log"

var tpl *template.Template

// Antes de la funcion init creamos un FuncMap para registrar las funciones que queremos usar en nuestros templates

var funciones = template.FuncMap{
	"imprimir": printData,
}

// inicializamos el template ahora.
func init() {
	tpl = template.Must(template.New("").Funcs(funciones).ParseFiles("templates/index.html"))
}

// Notar que las funciones debe retornar valores para poder ser usados en lo templates. No pueden haber llamados dinamicos
// como ocurriria si estuvieramos recorriendo un slice y estuvieramos en go llamando a ftm.Elemento
func printData(data []string) string {
	var primero string
	for index, value := range data {
		if index == 1 {
			primero = value
		}
	}
	return primero
}

func main() {
	// Creamos un slices de strings para pasarselo al template.
	data := []string{
		"El Señor de los Anillos",
		"The Matrix",
		"V de Venganza",
		"El Padrino",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", data)
	if err != nil {
		log.Fatal(err)
	}
}
