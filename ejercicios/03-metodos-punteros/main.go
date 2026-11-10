package main

import (
	"fmt"
	"strings"
)

// Correo agrupa los datos del correo; en Go podemos asociar métodos a tipos.
type Correo struct {
	Destinatario string
	Asunto       string
	Mensaje      string
	Estado       EstadoCorreo
}

// EstadoCorreo es un tipo definido cuya base es string.
type EstadoCorreo string

// Las constantes permiten reutilizar los estados sin repetir sus textos.
const (
	EstadoPendiente EstadoCorreo = "Pendiente"
	EstadoEnviado   EstadoCorreo = "Enviado"
	EstadoError     EstadoCorreo = "Error"
)

func (correo Correo) validar() bool {
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

// enviar simula el envío y modifica el estado del correo original.
func (correo *Correo) enviar() bool {
	if !correo.validar() {
		correo.Estado = EstadoError
		return false
	}
	correo.Estado = EstadoEnviado
	return true
}

// marcarEnviado conserva el método de práctica de receptores puntero.
// Asigna el estado directamente, sin validar ni realizar un envío.
func (correo *Correo) marcarEnviado() {
	correo.Estado = EstadoEnviado
}

func main() {
	correoPrueba := Correo{
		Destinatario: "javier@example.com",
		Asunto:       "Prueba de GoMail",
		Mensaje:      "Su compra ha sido recibida",
		Estado:       EstadoPendiente,
	}

	enviado := correoPrueba.enviar()

	fmt.Println("Resultado:", enviado)
	fmt.Println("Estado:", correoPrueba.Estado)
}
