// Pasando tipos compuestos en go.
package main

import "text/template"
import "os"
import "log"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html", "templates/slices_struct.html", "templates/map.html", "templates/slices_struct_embbedd.html"))
}

type Juego struct {
	Nombre string
	Year   int
}

// Structs para parsear structs compuestos.
type Alumno struct {
	Nombre      string
	Edad        int
	Asignaturas []Asignatura
}

type Asignatura struct {
	Id       int
	Nombre   string
	Creditos float64
}

func main() {
	data := []string{"Nicolas Riquelme", "Camilo Riquelme", "Francisca Romero", "Jorge Aliaga", "Sylvia Guzman"}
	//err := tpl.Execute(os.Stdout, data)
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", data)
	if err != nil {
		log.Fatal(err)
	}

	data_structs := []Juego{
		Juego{
			Nombre: "El Señor de los Anillos",
			Year:   1999,
		},
		Juego{
			Nombre: "Command and Conquer Red Alert",
			Year:   2001,
		},
		Juego{
			Nombre: "Startcraft",
			Year:   2002,
		},
	}

	slicesStructs := tpl.ExecuteTemplate(os.Stdout, "slices_struct.html", data_structs)
	if slicesStructs != nil {
		log.Fatal(slicesStructs)
	}

	// Pasando un map a los templates.
	alumnos := map[int]string{
		1: "Nicolas Riquelme",
		2: "Francisca Romero",
		3: "Jorge Aliaga",
		4: "Camilo Riquelme",
		5: "Sylvia Guzman"}

	mapa := tpl.ExecuteTemplate(os.Stdout, "map.html", alumnos)
	if mapa != nil {
		log.Fatal(mapa)
	}

	// Pasamos ahora un slices de structs embebidos.
	colegio := []Alumno{
		Alumno{
			Nombre: "Nicolas Riquelme",
			Edad:   29,
			Asignaturas: []Asignatura{
				Asignatura{
					Id:       1,
					Nombre:   "Calculo",
					Creditos: 2.5,
				},
				Asignatura{
					Id:       2,
					Nombre:   "Programacion",
					Creditos: 4.5,
				},
			},
		},
		Alumno{
			Nombre: "Francisca Romero",
			Edad:   28,
			Asignaturas: []Asignatura{
				Asignatura{
					Id:       1,
					Nombre:   "Calculo",
					Creditos: 2.5,
				},
				Asignatura{
					Id:       3,
					Nombre:   "Diseño de Investigacion",
					Creditos: 5.7,
				},
			},
		},
	}

	compositeStructs := tpl.ExecuteTemplate(os.Stdout, "slices_struct_embbedd.html", colegio)
	if compositeStructs != nil {
		log.Fatal(compositeStructs)
	}
}
