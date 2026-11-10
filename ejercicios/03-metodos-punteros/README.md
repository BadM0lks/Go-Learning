# Ejercicio 03: métodos, punteros y estados

**Completado:** 10 de noviembre de 2026

## Objetivo

Asociar la validación al tipo Correo y simular un envío que devuelva un booleano y actualice el estado del correo original.

## Conceptos practicados

- Método validar() con receptor por valor: consulta los campos de una copia.
- Método enviar() con receptor *Correo: modifica el original.
- Tipo definido EstadoCorreo, basado en string, y constantes tipadas.
- Condiciones booleanas directas y negación con !.
- Retorno temprano: return termina el método, por lo que no hace falta un segundo if para el caso válido.

## Dudas resueltas

- **¿Por qué no usar dos if y un return false final?** Tras devolver false en el caso inválido, la ejecución restante ya corresponde al caso válido. Así se llama a validar() una sola vez y todos los caminos devuelven un valor.
- **¿El if supone true?** Ejecuta su bloque cuando la condición completa vale true. !correo.validar() invierte el resultado: equivale a correo.validar() == false. Una condición booleana directa también se puede usar en Dart.
- **¿Qué diferencia hay entre enviado y Estado?** enviado recibe el bool del return; Estado guarda la constante asignada en la estructura.
- **¿Por qué *Correo?** Con un receptor por valor se cambiaría una copia y el estado del original no se conservaría.

## Ejecutar y comprobar

Desde la raíz del repositorio:

```bash
go run ./ejercicios/03-metodos-punteros
```

| Caso | Resultado | Estado |
| --- | --- | --- |
| Destinatario, asunto y mensaje con contenido | true | Enviado |
| Asunto compuesto solo por espacios | false | Error |

El usuario compartió la salida del caso inválido y confirmó la salida del caso válido en el chat. La revisión final del código fue manual; el asistente no dispuso de un compilador de Go en esta sesión y no ejecutó pruebas automatizadas del ejercicio 03.

Para repetir el caso inválido, cambia el asunto del ejemplo a "   " y ejecuta de nuevo.

## Alcance

El envío es simulado: no hay conexión con un servicio de correo. La validación solo comprueba que los tres campos tengan contenido útil; no comprueba el formato de la dirección.

marcarEnviado() se conserva como práctica del receptor puntero. Asigna el estado directamente y no se utiliza en main(), que emplea enviar() para validar primero.

El código conserva la lógica entregada por el usuario, con comentarios aclarados y datos de ejemplo ficticios.
