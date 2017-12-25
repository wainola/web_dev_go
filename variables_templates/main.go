// enviado datos al template y asignando esos datos a una variable
package main

import "text/template"
import "os"
import "log"
import "fmt"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mensaje := "Go es util puesto que incorpora un lenguaje sencillo como C pero pontente y listo para la programacion del siglo XXI"
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", mensaje)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}
