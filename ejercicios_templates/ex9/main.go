// Leer datos del archivo, convertirlos en un slices de un struct, enviarlo a un template.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./data/data.csv")
	if err != nil {
		log.Fatal("no se pudo leer el archivo")
	}
	// Conversion de []byte a string para poder trabajar con el archivo
	d := strings.Split(string(data), ",")

	// Ahora hay que procesar la informacion.
	// Tomar cabecera primeros. Guardamos un mapa con las cabeceras y el valor del mapa es un slice donde guardamos los datos.
	m := make(map[string][]string)
	// Seteando las cabeceras de la tabla
	m[d[0]] = []string{}
	m[d[1]] = []string{}
	m[d[2]] = []string{}
	m[d[3]] = []string{}
	m[d[4]] = []string{}
	m[d[5]] = []string{}
	// setando la ultima cabecera de la tabla.
	s := strings.Split(d[6], "\n")
	primera_fecha := s[1]
	m[s[0]] = []string{}
	fmt.Println(m)
	fmt.Println(primera_fecha)

	// Recorremos el fichero que es un slices de strings. Desde la posicion 7.
	for index, value := range d {
		// if math.Mod(float64(index), float64(7)) == 0 && index > 6 {
		// 	fmt.Printf("Fechas: %v\n", value)
		// }
		if index == 8 {
			fmt.Println(value)
		}
	}
}
