package handler

import (
	"awesomeProject/internal/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WebsocketServer struct {
	clients    map[string]*websocket.Conn
	register   chan Client
	unregister chan Client
	mu         sync.Mutex
}

type Client struct {
	Username string
	Socket   *websocket.Conn
}

var server = &WebsocketServer{
	clients:    make(map[string]*websocket.Conn),
	register:   make(chan Client),
	unregister: make(chan Client),
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessageRequest struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func HandleWebsocket(c *gin.Context) {
	// nâng cấp request http lên websocket
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(fmt.Errorf("websocket upgrader error: %v", err))
		return
	}
	defer conn.Close()
	/////
	client := &Client{
		Username: getUsernameFromToken(c),
		Socket:   conn,
	}
	server.mu.Lock()
	server.clients[client.Username] = client.Socket
	server.mu.Unlock()
	for k, v := range server.clients {
		log.Println(k, v)
	}
	for {
		//Nhận các message được đến endpoint /ws
		var mess MessageRequest
		err := conn.ReadJSON(&mess)
		if err != nil {
			log.Println("error reading message:", err)
			break
		}
		if server.clients[mess.Receiver] != nil {
			log.Println(mess.Receiver)
			err := server.clients[mess.Receiver].WriteJSON(mess)
			if err != nil {
				delete(server.clients, client.Username)
				break
			}
		}
	}
	fmt.Println("client disconnect")
}

func getUsernameFromToken(c *gin.Context) string {
	username, err := helpers.GetUsername(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return username
}
