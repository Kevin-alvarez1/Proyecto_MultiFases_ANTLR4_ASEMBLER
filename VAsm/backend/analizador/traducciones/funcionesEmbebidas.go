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

func generarEtiquetaUnica() string {
	etiqueta := fmt.Sprintf("msg_%d", contadorEtiqueta)
	contadorEtiqueta++
	return etiqueta
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
