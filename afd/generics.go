package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
	"strings"
	"unicode"
)

func DetectGenerics(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)
	startPos := i

	if i >= textLen || chars[i] != '<' {
		return false, i, token.Token{}
	}

	var sb strings.Builder
	sb.WriteRune(chars[i]) // agregamos el '<'
	i++

	for i < textLen {
		c := chars[i]

		if c == '>' {
			// Cerramos el primer nivel
			sb.WriteRune(c)
			i++ // consumimos el '>'
			break
		}

		// Agregamos cualquier caracter válido dentro sin aumentar nesting
		// Por si quieres validar caracteres permitidos, ajusta aquí:
		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == ',' || c == ':' || c == ' ' || c == '_' || c == '<' || c == '>' {
			sb.WriteRune(c)
		} else {
			// Caracter inválido dentro del genérico
			return false, startPos, token.Token{}
		}
		i++
	}

	word := sb.String()

	// Si no cerramos con '>', entonces no es válido
	if len(word) < 2 || word[len(word)-1] != '>' {
		return false, startPos, token.Token{}
	}

	tok := token.Token{
		Word:     word,
		Category: constants.GENERICO,
		Index:    fmt.Sprintf("%d-%d", startPos, i),
	}

	fmt.Println(tok)
	return true, i, tok
}
