package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectComma detecta una coma ',' como separador.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectComma(chars []rune, i int) (bool, int, token.Token) {
	if chars[i] == ',' {
		tokenFinded := token.Token{
			Word:     ",",
			Category: constants.SEPARADOR,
			Index:    fmt.Sprintf("%d-%d", i, i+1),
		}

		i++ // Avanzar después de la coma

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}
