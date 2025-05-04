package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectLineComment intenta detectar un comentario de línea (//).
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectLineComment(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)

	if i < textLen-1 && chars[i] == '/' && chars[i+1] == '/' {
		startPos := i
		i += 2 // Saltar //

		// Avanzar hasta el final de la línea
		for i < textLen && chars[i] != '\n' {
			i++
		}

		word := string(chars[startPos:i])
		tokenFinded := token.Token{
			Word:     word,
			Category: constants.COMENTARIO_LINEA,
			Index:    fmt.Sprintf("%d-%d", startPos, i),
		}

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}
