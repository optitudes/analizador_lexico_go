package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectSemicolon detecta el terminal ';'.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectSemicolon(chars []rune, i int) (bool, int, token.Token) {
	if chars[i] == ';' {
		tokenFinded := token.Token{
			Word:     ";",
			Category: constants.TERMINAL,
			Index:    fmt.Sprintf("%d-%d", i, i+1),
		}

		// Incrementar i después de procesar el punto y coma
		i++

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}
