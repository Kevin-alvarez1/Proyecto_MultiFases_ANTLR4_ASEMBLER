package traducciones

import (
	"fmt"
	"strings"
)

var contadorEtiqueta int = 0
var dataBuilder strings.Builder
var textBuilder strings.Builder

func GenerarCodigoPrint(msg string, addNewline bool) {
	if addNewline {
		msg += "\n"
	}
	etiqueta := fmt.Sprintf("msg_%d", contadorEtiqueta)
	contadorEtiqueta++

	dataBuilder.WriteString(fmt.Sprintf("%s: .ascii \"%s\"\n", etiqueta, escape(msg)))

	textBuilder.WriteString(fmt.Sprintf(`mov x0, #1
adr x1, %s
mov x2, #%d
mov w8, #64
svc #0

`, etiqueta, len(msg)))
}

func GenerarCodigoParseFloat(original string, valor float64, outputASM *strings.Builder) {
	etiqueta := fmt.Sprintf("parse_float_val_%d", contadorEtiqueta)
	contadorEtiqueta++

	dataBuilder.WriteString(fmt.Sprintf("%s: .float %f\n", etiqueta, valor))
	textBuilder.WriteString(fmt.Sprintf("ldr s0, =%s\n\n", etiqueta))
}

func GenerarCodigoTypeOf(tipo string) {
	etiqueta := fmt.Sprintf("typeof_result_%d", contadorEtiqueta)
	contadorEtiqueta++

	dataBuilder.WriteString(fmt.Sprintf("%s: .ascii \"%s\\n\"\n", etiqueta, tipo))

	textBuilder.WriteString(fmt.Sprintf(`mov x0, #1
adr x1, %s
mov x2, #%d
mov w8, #64
svc #0

`, etiqueta, len(tipo)+1))
}

func escape(input string) string {
	input = strings.ReplaceAll(input, "\n", `\n`)
	input = strings.ReplaceAll(input, `"`, `\"`)
	return input
}
func EmitirCodigoCompleto() string {
	return ".data\n" + dataBuilder.String() + "\n.text\n" + textBuilder.String()
}

func ResetearCodigoASM() {
	dataBuilder.Reset()
	textBuilder.Reset()
	contadorEtiqueta = 0
}
