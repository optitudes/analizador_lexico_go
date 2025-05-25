package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

func DetectStaticAccess(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)
	startPos := i

	if i+1 >= textLen || chars[i] != ':' || chars[i+1] != ':' {
		return false, i, token.Token{}
	}

	// Consumimos los dos ':'
	i += 2

	tokenFinded := token.Token{
		Word:     "::",
		Category: constants.ACCESO_ESTATICO, // Debes tener esta constante en tu paquete constants
		Index:    fmt.Sprintf("%d-%d", startPos, i),
	}

	return true, i, tokenFinded
}
