package main

import (
	"analizadorLexico/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("Servidor iniciado en http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	// valida si hubo un error al momento de recibir el contenido del socket
	if err != nil {
		fmt.Println("Error al actualizar a WebSocket:", err)
		return
	}
	// setea lun defer para que se asegure el cierre de la conexion
	defer conn.Close()

	// inicia un for infinito para estar siempre analizando y leyendo los mensajes del socket
	for {
		// se obtiene un mensaje del socket o un error
		_, msg, err := conn.ReadMessage()
		// en caso de que haya error se imprime y se rompe el ciclo infinito
		if err != nil {
			fmt.Println("Error leyendo mensaje:", err)
			break
		}

		// se inicializa el tokenizador con el texto que llega del socket
		tokenizer := types.NewTokenizer(string(msg))

		// si el tokenizer no se creo' correctamente imprimer el error y pasa a la siguiente iteracion del ciclo
		if tokenizer == nil {
			fmt.Printf("Tokenizer inválido para mensaje: %q\n", msg)
			continue
		}
		// obtiene los tokens del texto que llego' al socket
		tokens := tokenizer.GetTokens()
		// devuelve los tokens obtenidos para que le front los pueda renderizar en la tabla
		respJSON, err := json.Marshal(tokens)
		if err != nil {
			fmt.Println("Error al convertir a JSON:", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, respJSON); err != nil {
			fmt.Println("Error enviando mensaje:", err)
			break
		}
	}
}
