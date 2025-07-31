package router

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

//---

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//---

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Client says: %s", msg)

		// Write message back (echo)
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

//---

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Server running"))
}

//---

func Router() {
	http.HandleFunc("/api/v1/conn", HandleConn)
	http.HandleFunc("/api/v1/conn/verification", HandleConnVerification)
	http.HandleFunc("/api/v1/ws", handleWS)
	http.HandleFunc("/", handleHomePage)
}
