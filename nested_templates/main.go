// Trabajando con templates nesteados.
package main

import "text/template"
import "os"
import "log"

var tpl *template.Template

type Perro struct {
	Nombre string
	Edad   int
}

func (p Perro) GetNombre() string {
	return p.Nombre
}

func (p Perro) GetEdad() int {
	return p.Edad
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html", "templates/template.html", "templates/dataTemplate.html"))
}

func main() {

	// Pasando datos a un template
	p := Perro{
		Nombre: "Lucas",
		Edad:   10,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", p)
	if err != nil {
		log.Fatal(err)
	}
}
