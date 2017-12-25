// String to html
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	nombre := "Nicolas Riquelme"
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hola Mundo</title>
		</head>
		<body>
		<h1>` + nombre +
		`</h1></body></html>`)

	new_file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file ", err)
	}
	defer new_file.Close()

	io.Copy(new_file, strings.NewReader(str))
}
