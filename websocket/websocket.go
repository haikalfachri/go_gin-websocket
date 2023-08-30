package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer func() {
		conn.Close()

		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
	}()

	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
	}
}

func StartWebSocketServer() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading connection:", err)
			return
		}
		defer func() {
			conn.Close()

			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()
		}()

		clientsMu.Lock()
		clients[conn] = true
		clientsMu.Unlock()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}
		}
	})

	http.ListenAndServe(":8081", nil)
}

func BroadcastMessage(message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error sending message to client:", err)
		}
	}
}

func SendWebSocketUpdate(message string) {
	go BroadcastMessage(message)
}
