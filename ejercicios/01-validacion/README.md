# Ejercicio 01: validación básica de correos

**Trabajado:** 31 de octubre de 2026  
**Estado:** completado en el itinerario de aprendizaje.

## Objetivo

Crear una función en Go que reciba un destinatario y un asunto y devuelva `true` cuando ambos tengan contenido útil; si alguno está vacío o solo contiene espacios en blanco, devolver `false`.

## Conceptos practicados

- Funciones con parámetros de tipo `string` y retorno `bool`.
- Condicionales `if` y operadores de comparación.
- Importación del paquete estándar `strings`.
- Uso de `strings.TrimSpace` para detectar valores vacíos o con solo espacios en blanco.

## Dudas y aprendizajes

1. **¿Basta con `valor == "" || valor == " "`?** No: se escaparían las entradas con varios espacios o tabulaciones. La solución fue `strings.TrimSpace(valor) == ""`.
2. **¿`TrimSpace` une las palabras?** No. Elimina espacios al principio y al final, pero conserva los internos. Por ejemplo, `" Nombre Apellido "` se convierte en `"Nombre Apellido"`.
3. **Comparación con Dart:** en Go un `string` tiene como valor cero `""`; un `string` ordinario no puede ser `nil`. Esto no significa que Go no tenga valores `nil` para otros tipos.

## Ejecución y pruebas

Desde la raíz del repositorio:

```bash
go run ./ejercicios/01-validacion
go test ./ejercicios/01-validacion
```

Las pruebas contemplan campos completos, vacíos, espacios, tabulaciones y espacios interiores.

## Límite del ejercicio

La función **no valida que el destinatario tenga formato de correo electrónico**: por ahora solo comprueba que los textos no estén vacíos. La implementación y los tests están preparados a partir de la solución documentada durante el aprendizaje.
