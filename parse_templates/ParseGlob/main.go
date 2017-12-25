// Parseando templates con ParseGlob
package main

import "html/template"
import "os"
import "log"
import "fmt"

// declaracion de variable global template para parseo correspondiente
// la variable corresponde a un puntero del tipo template.Template
var tpl *template.Template

// init inicializa los template utilizando la funcion must y ParseGlob
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Con esto podemos ver que hemos llamado al primer template.
	fmt.Println(err)

	fmt.Println("Ejecutando el segundo template ahora!!")
	// Ahora podemos llamar templates por su nombre de archivo.
	err2 := tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal(err2)
	}

	fmt.Println(err2)
}
