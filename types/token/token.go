package token

// estructura para almacenar la informacion de  un token
type Token struct {
	// guarda la palabra escrita
	Word string `json:"word"`
	// guarda la categoria a la que pertenece
	Category string `json:"category"`
	// guarda el indice de la palabra
	Index string `json:"index"`
}
