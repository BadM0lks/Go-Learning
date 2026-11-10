package main

import (
	"fmt"
	"strings"
)

// Correo agrupa los datos que necesitamos para representar un mensaje.
type Correo struct {
	Destinatario string
	Asunto       string
	Mensaje      string
}

// validarCorreo recibe una estructura Correo y comprueba sus tres campos.
func validarCorreo(correo Correo) bool {
	if strings.TrimSpace(correo.Destinatario) == "" {
		return false
	}
	if strings.TrimSpace(correo.Asunto) == "" {
		return false
	}
	if strings.TrimSpace(correo.Mensaje) == "" {
		return false
	}
	return true
}

func main() {
	correoPrueba := Correo{
		Destinatario: "correo@example.com",
		Asunto:       "Confirmación de pedido",
		Mensaje:      "Tu compra se ha recibido correctamente.",
	}

	fmt.Println(validarCorreo(correoPrueba))
}
