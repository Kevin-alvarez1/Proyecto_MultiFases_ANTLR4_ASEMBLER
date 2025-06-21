package traducciones

import (
	"fmt"
	"strings"
)

var contadorEtiqueta int = 0
var DataBuilder strings.Builder
var TextBuilder strings.Builder
var mensajesUnicos = make(map[string]string)

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

func EmitirCodigoCompleto(principal string) string {
	encabezado := fmt.Sprintf(`.global _start
.text
_start:
    bl fn_%s
    mov x0, #0
    mov x8, #93
    svc #0
`, principal)

	return ".data\n" + DataBuilder.String() + "\n" + encabezado + "\n" + TextBuilder.String()
}

func ResetearCodigoASM() {
	contadorEtiqueta = 0
	DataBuilder.Reset()
	TextBuilder.Reset()
	mensajesUnicos = make(map[string]string)
}
