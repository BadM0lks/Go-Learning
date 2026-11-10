# Ejercicio 02: estructuras (`structs`) en Go

**Trabajado:** 1 de noviembre de 2026  
**Estado:** completado en el itinerario de aprendizaje.

## Objetivo

Representar un correo mediante un `struct` con destinatario, asunto y mensaje, y crear una función que valide que ninguno de estos campos esté vacío o contenga únicamente espacios en blanco.

## Conceptos practicados

- Declaración de `type Correo struct { ... }`.
- Creación de valores usando literales de estructura, como `correoPrueba := Correo{...}`.
- Acceso a los campos con `correo.Destinatario`, `correo.Asunto` y `correo.Mensaje`.
- Funciones que aceptan un valor `Correo` como argumento.
- Reutilización de `strings.TrimSpace`.

## Dudas y aprendizajes

1. **¿Qué significa `correoPrueba := Correo{...}`?** Construye un valor del tipo `Correo` y declara una variable cuyo tipo se infiere. En Go un `struct` no equivale exactamente a una clase de Dart.
2. **¿Qué ocurre con `Correo{}` vacío?** Sus campos `string` quedan con su valor cero `""`.
3. **¿Qué cambia respecto al primer ejercicio?** En lugar de pasar dos cadenas por separado, los datos relacionados viajan agrupados en una estructura y la función revisa tres campos.

## Ejecución y pruebas

Desde la raíz del repositorio:

```bash
go run ./ejercicios/02-structs
go test ./ejercicios/02-structs
```

Las pruebas incluyen una estructura válida, una vacía y variantes con destinatario, asunto o mensaje sin contenido útil.

## Límite del ejercicio

Aquí usamos una **función independiente** `validarCorreo(correo Correo)`. Convertirla en un método con receptor es materia del ejercicio 03 y, por tanto, no se adelanta en este commit. El ejemplo público utiliza datos ficticios.
