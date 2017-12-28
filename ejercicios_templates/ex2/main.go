// Tarea => crear una estructura de datos para pasarsela al template
// Estructura de datos debe contener informacion sobre hoteles en y las regiones del hotel. Incluir tambien codigo postal
package main

import "text/template"
import "os"
import "log"

type Hotel struct {
	Nombre    string
	Direccion string
	Comuna    string
	Ubicacion Ubicacion
}

type Ubicacion struct {
	Nombre  string
	ZipCode int
	Coords  Coordenadas
}

type Coordenadas struct {
	Lat string
	Lng string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func main() {
	// Slice de Hoteles
	h := []Hotel{
		Hotel{
			"Hotel 1",
			"Calle 25",
			"Ñuñoa",
			Ubicacion{
				"Santiago",
				123456,
				Coordenadas{
					"-33.12323",
					"-77.23432",
				},
			},
		},
		Hotel{
			"Hotel 2",
			"Calle 26",
			"Providencia",
			Ubicacion{
				"Santiago",
				98765,
				Coordenadas{
					"-33.9889",
					"-77.4534",
				},
			},
		},
		Hotel{
			"Hotel 3",
			"Calle 27",
			"Santiago Centro",
			Ubicacion{
				"Santiago",
				12321321,
				Coordenadas{
					"-33.123",
					"-77.321",
				},
			},
		},
	}

	// Ahora pasamos los datos al template
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", h)
	if err != nil {
		log.Fatal(err)
	}
}
