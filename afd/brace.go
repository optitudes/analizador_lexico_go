package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectOpenBrace detecta una llave de apertura '{'.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectOpenBrace(chars []rune, i int) (bool, int, token.Token) {
	if chars[i] == '{' {
		tokenFinded := token.Token{
			Word:     "{",
			Category: constants.LLAVE_ABRE,
			Index:    fmt.Sprintf("%d-%d", i, i+1),
		}

		// Incrementar i después de procesar la llave
		i++

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}

// DetectCloseBrace detecta una llave de cierre '}'.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectCloseBrace(chars []rune, i int) (bool, int, token.Token) {
	if chars[i] == '}' {
		tokenFinded := token.Token{
			Word:     "}",
			Category: constants.LLAVE_CIERRA,
			Index:    fmt.Sprintf("%d-%d", i, i+1),
		}

		// Incrementar i después de procesar la llave
		i++

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}
