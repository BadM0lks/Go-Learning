package main

import (
	"fmt"
	"strings"
)

// validarCorreo comprueba que el destinatario y el asunto no estén vacíos.
// Este ejercicio no comprueba todavía el formato real de una dirección de email.
func validarCorreo(destinatario string, asunto string) bool {
	if strings.TrimSpace(destinatario) == "" {
		return false
	}
	if strings.TrimSpace(asunto) == "" {
		return false
	}
	return true
}

func main() {
	fmt.Println(validarCorreo("correo@example.com", "Bienvenida"))
}
