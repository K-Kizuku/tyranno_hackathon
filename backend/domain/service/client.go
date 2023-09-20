package service

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"tyranno/backend/domain/model"
	"tyranno/backend/domain/repository"

	"github.com/gorilla/websocket"
)

type IClientService interface {
	ServeWs(hub *model.Hub, w http.ResponseWriter, r *http.Request)
}

type clientService struct {
	Client model.Client
	repo   repository.IClientRepository
}

func NewClientService(cr repository.IClientRepository) IClientService {
	return &clientService{
		repo: cr,
	}
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *clientService) ReadPump() {
	defer func() {
		c.Client.Hub.Unregister <- &c.Client
		c.Client.Conn.Close()
	}()
	c.Client.Conn.SetReadLimit(maxMessageSize)
	c.Client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Client.Conn.SetPongHandler(func(string) error { c.Client.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	log.Println("tets")
	for {
		_, message, err := c.Client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Client.Hub.Broadcast <- message
	}
}

func (c *clientService) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Client.Send:
			c.Client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Client.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *clientService) ServeWs(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client_ := &model.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client := &clientService{Client: *client_, repo: nil}
	log.Printf("%v", client)
	client.Client.Hub.Register <- &client.Client
	go client.WritePump()
	go client.ReadPump()
}
