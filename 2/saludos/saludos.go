package saludos

import "fmt"

func SaludoFormal (nombre string) {
	mensaje := fmt.Sprintf("Hola, %s. ¡Bienvenido!", nombre)
	return mensaje
}