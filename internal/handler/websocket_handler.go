package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WebsocketServer struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

type Client struct {
	Username  string
	Socket    *websocket.Conn
	Broadcast chan *RequestMessage
}

var server = &WebsocketServer{
	clients:    make(map[string]*Client),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type RequestMessage struct {
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

const workerPool = 10

func HandleWebsocket(c *gin.Context) {
	// nâng cấp request http lên websocket
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(fmt.Errorf("websocket upgrader error: %v", err))
		return
	}
	defer conn.Close()
	/////
	go server.HandleRegistration()
	client := &Client{
		Username: getUsernameFromToken(c),
		Socket:   conn,
	}
	server.register <- client
	go client.receiveMessage()
	for {
		//Nhận các message được đến endpoint /ws
		var mess RequestMessage
		err := conn.ReadJSON(&mess)
		if err != nil {
			log.Println("error reading message:", err)
			break
		}

		server.sendMessageToUser(*client, mess)

	}
	server.unregister <- client
	fmt.Println("client disconnect")
}

func (w *WebsocketServer) HandleRegistration() {
	for {
		select {
		case client := <-w.register:
			w.mu.Lock()
			if _, ok := w.clients[client.Username]; !ok {
				log.Println("register client:", client.Username)
				w.clients[client.Username] = client
				client.Broadcast = make(chan *RequestMessage, 100)
			} else {
				log.Printf("%s already registered", client.Username)
			}
			w.mu.Unlock()
		case client := <-w.unregister:
			w.mu.Lock()
			if _, ok := w.clients[client.Username]; ok {
				log.Println("unregister client:", client.Username)
				delete(w.clients, client.Username)
				close(client.Broadcast)
			}
			w.mu.Unlock()
		}
	}
}

func OnlineClient(c *gin.Context) {
	var onlineUser []string
	server.mu.RLock()
	defer server.mu.RUnlock()
	for key, _ := range server.clients {
		onlineUser = append(onlineUser, key)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": onlineUser})
}

func (w *WebsocketServer) sendMessageToUser(client Client, message RequestMessage) {
	w.mu.Lock()
	defer w.mu.Unlock()
	// Kiểm tra người nhận có phải chính bản thân người gửi không
	if client.Username == message.Receiver {
		log.Println("Cannot send message to yourself")
		return
	}
	//Lấy connection user người gửi tin, kiểm tra tồn tai của client có đang online không
	sender, ok := w.clients[client.Username]
	if !ok {
		log.Println("client not found")
		return
	}
	//Lấy connection user người nhận tin, kiểm tra tồn tại của receiver có đang online không
	receiver, ok := server.clients[message.Receiver]
	if !ok {
		log.Println("receiver not found")
		return
	}
	receiver.Broadcast <- &message
	sender.Broadcast <- &message
	// Gửi tin nhắn cho receiver và hiển thị tin nhắn đã gửi đến client gửi tin nhắn
	//err := receiver.Socket.WriteMessage(websocket.TextMessage, []byte(message.Message))
	//if err != nil {
	//	log.Println("error writing message:", err)
	//	return
	//}
	//err = sender.Socket.WriteMessage(websocket.TextMessage, []byte(message.Message))
	//if err != nil {
	//	log.Println("error writing message:", err)
	//	return
	//}
}

func (c *Client) receiveMessage() {
	messageChannel := make(chan RequestMessage, 100)
	for i := 0; i < workerPool; i++ {
		go func() {
			for message := range messageChannel {
				err := c.Socket.WriteMessage(websocket.TextMessage, []byte(message.Message))
				if err != nil {
					log.Printf("Error sending message to %s: %v\n", c.Username, err)
					break
				}
			}
		}()
	}
	for message := range c.Broadcast {
		select {
		case messageChannel <- *message:
		case <-c.Broadcast:
			log.Println("broadcast channel closed")
			return
		default:
			log.Println("Message channel full, skipping message.")
		}
	}
}
