package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
	"unicode"
)

// DetectIdentifier verifica si el carácter actual inicia un identificador o palabra reservada.
// Retorna true si detecta uno, el nuevo índice y el token correspondiente.
func DetectIdentifier(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)

	if !isIdentifierStart(chars[i]) {
		return false, i, token.Token{}
	}

	startPos := i
	i++

	// AFD para identificadores: letra seguida de letras, dígitos o guiones bajos
	for i < textLen && isIdentifierPart(chars[i]) {
		i++
	}

	word := string(chars[startPos:i])
	category := constants.IDENTIFICADOR

	// Verificar si es palabra reservada
	if isReservedWord(word) {
		category = constants.PALABRA_RESERVADA
	} else if len(word) > 10 {
		category = constants.DESCONOCIDO
	}

	tokenFinded := token.Token{
		Word:     word,
		Category: category,
		Index:    fmt.Sprintf("%d-%d", startPos, i),
	}

	return true, i, tokenFinded
}

// isIdentifierStart verifica si un carácter puede ser el inicio de un identificador
func isIdentifierStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

// isIdentifierPart verifica si un carácter puede ser parte de un identificador
func isIdentifierPart(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// Verificación de palabras reservadas implementada con búsqueda lineal
func isReservedWord(token string) bool {
	for _, word := range constants.ReservedWords {
		if token == word {
			return true
		}
	}
	return false
}
