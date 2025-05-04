package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
	"unicode"
)

// DetectPrimitive intenta detectar tipos primitivos de Rust como i32, u8, f64, etc.
// Retorna true si lo detecta, el nuevo valor de `i` y el Token generado.
func DetectPrimitive(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)

	// Solo entra si empieza con una letra válida para tipos primitivos (como 'i', 'u', 'f', etc.)
	if unicode.IsLetter(chars[i]) {
		startPos := i
		i++

		// Avanzar sobre dígitos (como el 32 de i32, u64, etc.)
		for i < textLen && unicode.IsDigit(chars[i]) {
			i++
		}

		word := string(chars[startPos:i])

		// Lista de tipos primitivos válidos de Rust
		primitives := map[string]bool{
			"i8": true, "i16": true, "i32": true, "i64": true, "i128": true,
			"u8": true, "u16": true, "u32": true, "u64": true, "u128": true,
			"f32": true, "f64": true,
			"bool": true, "char": true, "str": true,
		}

		if primitives[word] {
			tokenFinded := token.Token{
				Word:     word,
				Category: constants.PRIMITIVO,
				Index:    fmt.Sprintf("%d-%d", startPos, i),
			}
			return true, i, tokenFinded
		}
	}

	return false, i, token.Token{}
}
