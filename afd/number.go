package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
	"unicode"
)

// DetectNumber detecta números naturales y reales.
// Retorna true si detecta un número, el nuevo índice y el token generado.
func DetectNumber(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)

	if !unicode.IsDigit(chars[i]) {
		return false, i, token.Token{}
	}

	startPos := i
	i++
	category := constants.NUMERO_NATURAL
	hasDecimalPoint := false

	// AFD para números
	for i < textLen && (unicode.IsDigit(chars[i]) || (!hasDecimalPoint && chars[i] == '.')) {
		if chars[i] == '.' {
			hasDecimalPoint = true
			category = constants.NUMERO_REAL
		}
		i++
	}

	// Validar formato de número real
	if hasDecimalPoint {
		// Debe haber al menos un dígito antes y después del punto
		dotIndex := -1
		for j := startPos; j < i; j++ {
			if chars[j] == '.' {
				dotIndex = j
				break
			}
		}
		if dotIndex == startPos || dotIndex == i-1 || dotIndex == -1 {
			category = constants.DESCONOCIDO
		}
	}

	tokenFinded := token.Token{
		Word:     string(chars[startPos:i]),
		Category: category,
		Index:    fmt.Sprintf("%d-%d", startPos, i),
	}

	return true, i, tokenFinded
}
