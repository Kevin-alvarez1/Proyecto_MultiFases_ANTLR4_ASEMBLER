package operaciones

import "strings"

func SecuenciaDeEscape(s string) string {
	processedString := s
	processedString = strings.ReplaceAll(processedString, "\\n", "\n")  // salto de línea
	processedString = strings.ReplaceAll(processedString, "\\t", "\t")  // tabulación
	processedString = strings.ReplaceAll(processedString, "\\\"", "\"") // comilla doble
	processedString = strings.ReplaceAll(processedString, "\\r", "\r")  // retorno de carro
	processedString = strings.ReplaceAll(processedString, "\\\\", "\\") // barra invertida
	processedString = strings.ReplaceAll(processedString, "\\'", "'")   // comilla simple
	return processedString
}
