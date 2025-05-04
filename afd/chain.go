package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectChain intenta detectar una cadena entre comillas (") con manejo de escapes.
// Retorna true si detecta una cadena, el nuevo valor de `i` y el Token generado.
func DetectChain(chars []rune, i int) (bool, int, token.Token) {

	textLen := len(chars)

	if i >= textLen || chars[i] != '"' {
		return false, i, token.Token{}
	}

	startPos := i
	i++ // Saltar la comilla inicial
	escaped := false

	for i < textLen {
		if escaped {
			escaped = false
		} else if chars[i] == '\\' {
			escaped = true
		} else if chars[i] == '"' {
			i++ // Incluir la comilla final
			break
		}
		i++
	}

	word := string(chars[startPos:i])
	tokenFinded := token.Token{
		Word:     word,
		Category: constants.CADENA,
		Index:    fmt.Sprintf("%d-%d", startPos, i),
	}

	return true, i, tokenFinded
}
