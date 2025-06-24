package traducciones

import (
	"fmt"
	"strings"
)

var contadorEtiqueta int = 0
var DataBuilder strings.Builder
var TextBuilder strings.Builder
var mensajesUnicos = make(map[string]string)
var FuncionesBuilder strings.Builder
var contadorSuma int = 0

func GenerarCodigoPrint(msg string, addNewline bool) {
	if addNewline {
		msg += "\n"
	}
	if etiqueta, existe := mensajesUnicos[msg]; existe {
		// Ya fue agregado, simplemente usar etiqueta
		TextBuilder.WriteString(fmt.Sprintf(`mov x0, #1
adr x1, %s
mov x2, #%d
mov w8, #64
svc #0

`, etiqueta, len(msg)))
		return
	}

	etiqueta := fmt.Sprintf("msg_%d", contadorEtiqueta)
	contadorEtiqueta++
	mensajesUnicos[msg] = etiqueta

	DataBuilder.WriteString(fmt.Sprintf("%s: .ascii \"%s\"\n", etiqueta, escape(msg)))

	TextBuilder.WriteString(fmt.Sprintf(`mov x0, #1
adr x1, %s
mov x2, #%d
mov w8, #64
svc #0

`, etiqueta, len(msg)))
}

func escape(input string) string {
	input = strings.ReplaceAll(input, "\n", `\n`)
	input = strings.ReplaceAll(input, `"`, `\"`)
	return input
}

func ResetearCodigoASM() {
	contadorEtiqueta = 0
	DataBuilder.Reset()
	TextBuilder.Reset()
	mensajesUnicos = make(map[string]string)
	FuncionesBuilder.Reset()
	variablesReservadas = make(map[string]bool)
	contadorSuma = 0
}
