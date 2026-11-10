package main

import "testing"

func TestValidarCorreo(t *testing.T) {
	base := Correo{
		Destinatario: "correo@example.com",
		Asunto:       "Hola",
		Mensaje:      "Texto del mensaje",
	}

	casos := []struct {
		nombre   string
		correo   Correo
		esperado bool
	}{
		{"todos los campos completos", base, true},
		{"estructura sin inicializar", Correo{}, false},
		{"destinatario vacío", Correo{"", base.Asunto, base.Mensaje}, false},
		{"asunto en blanco", Correo{base.Destinatario, "  ", base.Mensaje}, false},
		{"mensaje en blanco", Correo{base.Destinatario, base.Asunto, "\t"}, false},
		{"espacios externos", Correo{" correo@example.com ", " Hola ", " Mensaje "}, true},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			obtenido := validarCorreo(caso.correo)
			if obtenido != caso.esperado {
				t.Errorf("validarCorreo(%+v) = %t; se esperaba %t", caso.correo, obtenido, caso.esperado)
			}
		})
	}
}
