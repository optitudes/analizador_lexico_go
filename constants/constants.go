package constants

// Categorías de tokens
const (
	NUMERO_NATURAL               = "NÚMERO NATURAL"
	NUMERO_REAL                  = "NÚMERO REAL"
	IDENTIFICADOR                = "IDENTIFICADOR"
	PALABRA_RESERVADA            = "PALABRA RESERVADA"
	OP_ARITMETICO                = "OPERADOR ARITMÉTICO"
	OP_COMPARACION               = "OPERADOR DE COMPARACIÓN"
	OP_LOGICO                    = "OPERADOR LÓGICO"
	OP_ASIGNACION                = "OPERADOR DE ASIGNACIÓN"
	OP_INC_DEC                   = "OPERADOR DE INCREMENTO/DECREMENTO"
	PARENTESIS_ABRE              = "PARÉNTESIS DE APERTURA"
	PARENTESIS_CIERRA            = "PARÉNTESIS DE CIERRE"
	LLAVE_ABRE                   = "LLAVE DE APERTURA"
	LLAVE_CIERRA                 = "LLAVE DE CIERRE"
	TERMINAL                     = "TERMINAL"
	SEPARADOR                    = "SEPARADOR"
	CADENA                       = "CADENA DE CARACTERES"
	COMENTARIO_LINEA             = "COMENTARIO DE LÍNEA"
	COMENTARIO_BLOQUE            = "COMENTARIO DE BLOQUE"
	DESCONOCIDO                  = "DESCONOCIDO"
	PRIMITIVO                    = "PRIMITIVO"
	ASIGNACION_TIPO_VARIABLE     = "ASIGNACION TIPO VARIABLE"
	MACRO                        = "MACRO"
	REFERENCIA                   = "REFERENCIA"
	REFERENCIA_MUT               = "REFERENCIA_MUT"
	REFERENCIA_PALABRA_RESERVADA = "REFERENCIA_PALABRA_RESERVADA"
	ACCESO_ESTATICO              = "ACCESO_ESTATICO"
	GENERICO                     = "GENERICO"
	TIPO_RETORNO                 = "TIPO_RETORNO"
)

var (
	// Separators and symbols
	Separators = []rune{' ', '\t', '\n', '(', ')', '{', '}', ';', ','}

	// Rust reserved keywords (strict)
	ReservedWords = []string{
		"as", "break", "const", "continue", "crate", "else", "enum",
		"extern", "false", "fn", "for", "if", "impl", "in", "let",
		"loop", "match", "mod", "move", "mut", "pub", "ref", "return",
		"self", "Self", "static", "struct", "super", "trait", "true",
		"type", "unsafe", "use", "where", "while", "dyn", "await",
		"async", "abstract", "become", "box", "do", "final", "macro",
		"override", "priv", "typeof", "unsized", "virtual", "yield", "try",
	}

	// Arithmetic operators
	ArithmeticOperators = []string{"+", "-", "*", "/", "%"}

	// Comparison operators
	ComparisonOperators = []string{"==", "!=", ">", "<", ">=", "<="}

	// Logical operators
	LogicalOperators = []string{"&&", "||", "!"}

	// Assignment operators
	AssignmentOperators = []string{"=", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "<<=", ">>="}

	// Increment/decrement: Rust does NOT have ++ or --
	// We'll keep this empty to be accurate
	IncrementDecrementOperators = []string{}

	// Parentheses
	Parentheses = []rune{'(', ')'}

	// Braces
	Braces = []rune{'{', '}'}

	// Statement terminator
	StatementTerminator = ';'

	// Separator (comma)
	Comma = ','

	// String delimiter
	StringDelimiter = '"'

	// Comment patterns
	LineCommentStart  = "//"
	BlockCommentStart = "/*"
	BlockCommentEnd   = "*/"
)
