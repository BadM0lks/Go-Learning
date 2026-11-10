package main

import "testing"

func TestValidarCorreo(t *testing.T) {
	casos := []struct {
		nombre       string
		destinatario string
		asunto       string
		esperado     bool
	}{
		{"campos completos", "correo@example.com", "Hola", true},
		{"destinatario vacío", "", "Hola", false},
		{"asunto vacío", "correo@example.com", "", false},
		{"solo espacios", "   ", "Hola", false},
		{"tabulaciones", "correo@example.com", "\t\n", false},
		{"espacios externos", "  correo@example.com  ", "  Hola  ", true},
		{"espacios internos", "nombre apellido", "Hola", true},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			obtenido := validarCorreo(caso.destinatario, caso.asunto)
			if obtenido != caso.esperado {
				t.Errorf("validarCorreo(%q, %q) = %t; se esperaba %t", caso.destinatario, caso.asunto, obtenido, caso.esperado)
			}
		})
	}
}
