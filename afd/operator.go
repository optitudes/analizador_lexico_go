package afd

import (
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
	"fmt"
)

// DetectOperator detecta operadores simples o compuestos.
func DetectOperator(chars []rune, i int) (bool, int, token.Token) {
	textLen := len(chars)

	if !isOperatorChar(chars[i]) {
		return false, i, token.Token{}
	}

	startPos := i
	i++
	category := ""

	// AFD para operadores compuestos
	if i < textLen && isOperatorContinuation(chars[startPos], chars[i]) {
		i++
		// Verificar operadores de tres caracteres (<<=, >>=)
		if i < textLen && ((chars[startPos] == '<' && chars[startPos+1] == '<') ||
			(chars[startPos] == '>' && chars[startPos+1] == '>')) &&
			chars[i] == '=' {
			i++
		}
	}

	word := string(chars[startPos:i])

	switch {
	case isArithmeticOperator(word):
		category = constants.OP_ARITMETICO
	case isComparisonOperator(word):
		category = constants.OP_COMPARACION
	case isLogicalOperator(word):
		category = constants.OP_LOGICO
	case isAssignmentOperator(word):
		category = constants.OP_ASIGNACION
	case isIncrementDecrementOperator(word):
		category = constants.OP_INC_DEC
	default:
		category = constants.DESCONOCIDO
	}

	tokenFinded := token.Token{
		Word:     word,
		Category: category,
		Index:    fmt.Sprintf("%d-%d", startPos, i),
	}

	return true, i, tokenFinded
}

// Verificación de operadores de incremento/decremento
func isIncrementDecrementOperator(token string) bool {
	for _, op := range constants.IncrementDecrementOperators {
		if token == op {
			return true
		}
	}
	return false
}

// Verificación de operadores de asignación
func isAssignmentOperator(token string) bool {
	for _, op := range constants.AssignmentOperators {
		if token == op {
			return true
		}
	}
	return false
}

// Verificación de operadores aritméticos
func isArithmeticOperator(token string) bool {
	for _, op := range constants.ArithmeticOperators {
		if token == op {
			return true
		}
	}
	return false
}

// Verificación de operadores de comparación
func isComparisonOperator(token string) bool {
	for _, op := range constants.ComparisonOperators {
		if token == op {
			return true
		}
	}
	return false
}

// Verificación de operadores lógicos
func isLogicalOperator(token string) bool {
	for _, op := range constants.LogicalOperators {
		if token == op {
			return true
		}
	}
	return false
}

// isOperatorChar verifica si un carácter puede ser parte de un operador
func isOperatorChar(r rune) bool {
	ops := "+-*/%=!<>&|^"
	for _, op := range ops {
		if r == op {
			return true
		}
	}
	return false
}

// isOperatorContinuation verifica si el segundo carácter puede formar un operador compuesto con el primero
func isOperatorContinuation(first, second rune) bool {
	if first == '=' && second == '=' {
		return true
	}
	if first == '!' && second == '=' {
		return true
	}
	if first == '<' && (second == '=' || second == '<') {
		return true
	}
	if first == '>' && (second == '=' || second == '>') {
		return true
	}
	if first == '&' && (second == '&' || second == '=') {
		return true
	}
	if first == '|' && (second == '|' || second == '=') {
		return true
	}
	if (first == '+' || first == '-' || first == '*' || first == '/' ||
		first == '%' || first == '^') && second == '=' {
		return true
	}
	return false
}
