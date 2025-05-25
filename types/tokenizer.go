package types

import (
	"analizadorLexico/afd"
	"analizadorLexico/constants"
	"analizadorLexico/types/token"
)

type Tokenizer struct {
	text string
}

func NewTokenizer(text string) *Tokenizer {
	// en caso de que el texto sea vacio retorna una instancia nula
	if text == "" {
		return nil
	}

	// en
	lastChar := rune(text[len(text)-1])

	if !IsSeparator(lastChar) {
		return nil
	}

	return &Tokenizer{text: text}
}

func (t *Tokenizer) GetTokens() []token.Token {
	return TokenizeAndCategorize(t.text)
}

func TokenizeAndCategorize(text string) []token.Token {
	tokenList := []token.Token{}
	chars := []rune(text)
	textLen := len(chars)
	tokenIndex := 0
	i := 0

	for i < textLen {
		// Ignorar espacios en blanco
		if chars[i] == ' ' || chars[i] == '\t' || chars[i] == '\n' {
			i++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectStaticAccess(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectVarTypeAssignation(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectPassForReference(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectChain(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectLineComment(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectBlockComment(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectOpenParenthesis(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectCloseParenthesis(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectOpenBrace(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectCloseBrace(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectSemicolon(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectComma(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectGenerics(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectOperator(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectNumber(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}
		if ok, newI, tokenFinded := afd.DetectPrimitive(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}

		if ok, newI, tokenFinded := afd.DetectIdentifier(chars, i); ok {
			tokenList = append(tokenList, tokenFinded)
			i = newI
			tokenIndex++
			continue
		}

		// Si llegamos aquí, es un carácter desconocido
		tokenList = append(tokenList, token.Token{
			Word:     string(chars[i]),
			Category: constants.DESCONOCIDO,
			Index:    string(tokenIndex),
		})
		tokenIndex++
		i++
	}

	return tokenList
}

// IsSeparator verifica si un caracter es un separador
func IsSeparator(r rune) bool {
	for _, sep := range constants.Separators {
		if r == sep {
			return true
		}
	}
	return false
}
