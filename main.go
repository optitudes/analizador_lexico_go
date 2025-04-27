package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
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
	data, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		http.Error(w, "No se pudo cargar el HTML", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error al actualizar a WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error leyendo mensaje:", err)
			break
		}

		upper := strings.ToUpper(string(msg))
		if err = conn.WriteMessage(websocket.TextMessage, []byte(upper)); err != nil {
			fmt.Println("Error enviando mensaje:", err)
			break
		}
	}
}
