// Leer datos del archivo, convertirlos en un slices de un struct, enviarlo a un template.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// Abrimos el archivo con os.Open
	data, err := os.Open("./data/data.csv")
	if err != nil {
		log.Fatal("no se pudo leer el archivo")
	}
	defer data.Close()

	// Parseamos a csv.
	reader := csv.NewReader(data)

	// Generamos mapa para setear nuestr estructura de datos.
	// Mapa con llaves de string y slice de strings.
	datos := make(map[string][]string)

	cabeceras, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range cabeceras {
		// Seteando cabeceras del map.
		datos[value] = []string{}
	}

	//fmt.Println(datos)

	d, err := reader.ReadAll()
	for _, value := range d {
		// Con eso procesamos y seteamos correctamente el mapa
		datos["Date"] = append(datos["Date"], value[0])
		datos["Open"] = append(datos["Open"], value[1])
		datos["High"] = append(datos["High"], value[2])
		datos["Low"] = append(datos["Low"], value[3])
		datos["Close"] = append(datos["Close"], value[4])
		datos["Volume"] = append(datos["Volume"], value[5])
		datos["Adj Close"] = append(datos["Adj Close"], value[6])
	}

	// Iteracion sobre los datos como prueba
	for index := range datos["Date"] {
		fmt.Println(datos["Date"][index], " ", datos["Volume"][index])
	}
}
