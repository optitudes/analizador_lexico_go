package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectOpenParenthesis detecta un paréntesis de apertura '('.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectOpenParenthesis(chars []rune, i int) (bool, int, token.Token) {
	if chars[i] == '(' {
		tokenFinded := token.Token{
			Word:     "(",
			Category: constants.PARENTESIS_ABRE,
			Index:    fmt.Sprintf("%d-%d", i, i+1),
		}

		// Incrementar i después de procesar el paréntesis
		i++

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}

// DetectCloseParenthesis detecta un paréntesis de cierre ')'.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectCloseParenthesis(chars []rune, i int) (bool, int, token.Token) {
	if chars[i] == ')' {
		tokenFinded := token.Token{
			Word:     ")",
			Category: constants.PARENTESIS_CIERRA,
			Index:    fmt.Sprintf("%d-%d", i, i+1),
		}

		// Incrementar i después de procesar el paréntesis
		i++

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}
