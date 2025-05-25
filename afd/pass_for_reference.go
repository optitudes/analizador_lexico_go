package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

func DetectPassForReference(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)
	startPos := i

	if i >= textLen || chars[i] != '&' {
		return false, i, token.Token{}
	}
	i++ // consumimos el '&'

	isMutable := false

	// Verificar si sigue "mut"
	if i+2 < textLen && string(chars[i:i+3]) == "mut" {
		isMutable = true
		i += 3
	}

	if i >= textLen || !isIdentifierStart(chars[i]) {
		// No hay identificador después de & o &mut
		return false, startPos, token.Token{}
	}

	idStart := i
	i++

	for i < textLen && isIdentifierPart(chars[i]) {
		i++
	}

	identifier := string(chars[idStart:i])
	fullWord := string(chars[startPos:i]) // incluye & o &mut

	category := constants.REFERENCIA
	if isMutable {
		category = constants.REFERENCIA_MUT
	}
	if isReservedWord(identifier) {
		category = constants.REFERENCIA_PALABRA_RESERVADA
	}

	tokenFinded := token.Token{
		Word:     fullWord,
		Category: category,
		Index:    fmt.Sprintf("%d-%d", startPos, i),
	}

	return true, i, tokenFinded
}
