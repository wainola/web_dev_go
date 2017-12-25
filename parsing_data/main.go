// Parseando datos a los templates.
package main

import "html/template"
import "os"
import "log"
import "fmt"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// notar que le indicamos como primer parametro a la funcion, en donde o hacia donde debe enviar los datos, en este caso, a la consola
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", "Enviado datos a un template desde funcion en GO")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}
