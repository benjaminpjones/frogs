package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"
)

// Websocket upgrader to upgrade HTTP requests to WebSocket connections
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

// Handle WebSocket connections
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    // Read and send messages (this is a simple echo example)
    for {
        // Read message from WebSocket
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        // Echo the message back to the client
        if err := conn.WriteMessage(msgType, msg); err != nil {
            log.Println(err)
            return
        }
    }
}

// Create and start a game (REST API)
func createGame(w http.ResponseWriter, r *http.Request) {
    // TODO: implement logic to start the game
    w.Write([]byte("Game started!"))
}

func main() {
    // Initialize router
    r := mux.NewRouter()

    // REST APIs
    r.HandleFunc("/api/create", createGame).Methods("POST")

    // WebSocket endpoint for real-time updates
    r.HandleFunc("/ws", handleWebSocket)

    // Start the server
    http.Handle("/", r)
    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
