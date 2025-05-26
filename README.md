# Analizador Léxico para Rust

## Descripción
Este es un analizador léxico escrito en Go como proyecto final de la asignatura Teoría de Lenguajes Formales. Este proyecto está diseñado específicamente para analizar código fuente del lenguaje de programación Rust, descomponiéndolo en tokens léxicos y categorizándolos según su tipo.

## Propósito
El objetivo principal es procesar código fuente de Rust y descomponerlo en tokens léxicos, identificando cada elemento según su categoría gramatical. Esto sirve como base para futuros análisis sintácticos y semánticos del código.

## Herramientas Utilizadas
- **Lenguaje**: Go (Golang)
- **Dependencias**:
  - `github.com/gorilla/websocket` - Para la comunicación en tiempo real detectando lo que el usuario ingresa para analizar

## Cómo Ejecutarlo

### Requisitos Previos
- Go 1.23 o superior instalado
- Navegador web moderno

### Instalación
```bash
go mod tidy
```

### Ejecución
```bash
go run main.go
```

### Acceso
- Abrir el navegador en `http://localhost:8081`

## Estructura del Proyecto

### Estructura de Carpetas

```bash
analizador_lexico_go/
├── main.go # Punto de entrada y servidor web
├── types/ # Definiciones de tipos y tokenizador
│ ├── token/ # Estructuras de tokens
│ └── tokenizer.go # Lógica principal del analizador
├── afd/ # Autómatas Finitos Deterministas
├── constants/ # Constantes y definiciones
└── views/ # Interfaz de usuario
  └─── index.html # Página web principal`
```

## Análisis Caracter por Caracter

El proceso de análisis léxico se realiza de la siguiente manera:

1. **Entrada**: El texto se convierte en un array de runes (caracteres Unicode)
2. **Procesamiento**:
   - Se ignora espacios en blanco, tabulaciones y saltos de línea
   - Se analiza cada carácter secuencialmente
   - Se utilizan AFDs específicos para cada tipo de token
3. **Salida**: Lista de tokens categorizados

## Autómatas Finitos Deterministas (AFD)

El proyecto implementa los siguientes AFDs:

### 1. Números (`number.go`)
- Detecta números naturales y reales
- Maneja punto decimal y notación científica
- Ejemplos: `42`, `3.14`, `1e-10`

### 2. Operadores (`operator.go`)
- Aritméticos: `+`, `-`, `*`, `/`, `%`
- Comparación: `==`, `!=`, `>`, `<`, `>=`, `<=`
- Lógicos: `&&`, `||`, `!`
- Asignación: `=`, `+=`, `-=`, `*=`, `/=`, `%=`

### 3. Identificadores (`identifier.go`)
- Comienza con letra o guion bajo
- Puede contener letras, números y guiones bajos
- No puede ser una palabra reservada
- Ejemplos: `variable`, `_private`, `user123`

### 4. Comentarios
- **Comentarios de Línea** (`line_comentary.go`):
  - Comienza con `//`
  - Continúa hasta el final de la línea
  - Ejemplo: `// Este es un comentario`

- **Comentarios de Bloque** (`blockComment.go`):
  - Comienza con `/*`
  - Termina con `*/`
  - Puede contener múltiples líneas
  - Ejemplo: `/* Este es un comentario de bloque */`

### 5. Delimitadores
- **Paréntesis** (`parenthesis.go`):
  - Detecta `(` y `)`
  - Usados para agrupación y llamadas de función

- **Llaves** (`brace.go`):
  - Detecta `{` y `}`
  - Usadas para bloques de código y estructuras de datos

- **Punto y Coma** (`semicolon.go`):
  - Detecta `;`
  - Terminador de sentencias en Rust

- **Coma** (`comma.go`):
  - Detecta `,`
  - Separador de elementos en listas y argumentos

### 6. Cadenas de Texto (`chain.go`)
- Detecta cadenas de caracteres
- Comienza y termina con comillas dobles `"`
- Puede contener caracteres escapados
- Ejemplo: `"Hola, mundo!"`

### 7. Tipos Primitivos (`primitive.go`)
- Detecta tipos de datos básicos de Rust
- Incluye:
  - Enteros: `i32`, `u64`, `isize`, `usize`
  - Flotantes: `f32`, `f64`
  - Booleanos: `bool`
  - Caracteres: `char`
  - Cadenas: `str`

### 8. Asignación de Tipo (`var_type_assignation.go`)
- Detecta declaraciones de variables con tipo
- Patrón: `let identificador: tipo`
- Ejemplos:
  - `let x: i32 = 5;`
  - `let nombre: String = "Rust";`

#### 9 Tipos de Retorno (`return_type.go`)
- Detecta declaraciones de tipos de retorno en funciones
- Patrón: `-> tipo`
- Ejemplos:
  - `fn suma() -> i32`
  - `fn procesar() -> Result<T, E>`

#### 10 Acceso Estático (`static_access.go`)
- Detecta acceso a miembros estáticos
- Patrón: `::`
- Ejemplos:
  - `String::new()`
  - `Vec::<i32>::new()`

#### 11 Genéricos (`generics.go`)
- Detecta declaraciones y uso de tipos genéricos
- Patrón: `<T>`, `<T: Trait>`, `<T, U>`
- Ejemplos:
  - `struct Punto<T>`
  - `fn procesar<T: Display>()`
  - `Vec<i32>`

#### 12 Paso por Referencia (`pass_for_reference.go`)
- Detecta operadores de referencia y préstamo
- Patrón: `&`, `&mut`
- Ejemplos:
  - `&variable`
  - `&mut contador`

### 13. Palabras Reservadas
Implementadas en `constants/constants.go`, incluye:
- Control de flujo: `if`, `else`, `while`, `for`, `loop`, `match`
- Declaraciones: `let`, `const`, `fn`, `struct`, `enum`
- Modificadores: `pub`, `mut`, `static`
- Tipos: `type`, `trait`, `impl`
- Y otras palabras clave específicas de Rust

## Interfaz de AFD
Cada AFD implementa la siguiente interfaz:
```go
func Detect[TokenType](chars []rune, startIndex int) (bool, int, token.Token)
```

Donde:
- `bool`: Indica si se encontró un token válido
- `int`: Nueva posición en el texto
- `token.Token`: Token encontrado con su categoría

### Orden de Precedencia
Los AFDs se evalúan en el siguiente orden para evitar ambigüedades:
1. Asignación de tipo
2. Tipos de retorno
3. Acceso estático
4. Genéricos
5. Paso por referencia
6. Cadenas de texto
7. Comentarios (línea y bloque)
8. Paréntesis y llaves
9. Punto y coma y coma
10. Operadores
11. Números
12. Tipos primitivos
13. Identificadores