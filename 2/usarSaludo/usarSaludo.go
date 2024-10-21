package main

import (
	"fmt"

	"b2/saludos"
)

func main () {

	mensaje := saludos.SaludoFormal("Bumbo")
	fmt.Println(mensaje)

}