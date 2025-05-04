package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectBlockComment intenta detectar un comentario de bloque (/* ... */),
// incluyendo comentarios anidados.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectBlockComment(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)

	if i < textLen-1 && chars[i] == '/' && chars[i+1] == '*' {
		startPos := i
		i += 2 // Saltar /*

		nested := 1
		for i < textLen-1 && nested > 0 {
			if chars[i] == '*' && chars[i+1] == '/' {
				nested--
				i += 2
			} else if chars[i] == '/' && chars[i+1] == '*' {
				nested++
				i += 2
			} else {
				i++
			}
		}

		word := string(chars[startPos:i])
		tokenFinded := token.Token{
			Word:     word,
			Category: constants.COMENTARIO_BLOQUE,
			Index:    fmt.Sprintf("%d-%d", startPos, i),
		}

		return true, i, tokenFinded
	}

	return false, i, token.Token{}
}
