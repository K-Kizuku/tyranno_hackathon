package service

import (
	"tyranno/backend/domain/model"
	"tyranno/backend/domain/repository"
)

type HubService struct {
	hub  model.Hub
	repo repository.IHubRepository
}

func NewHubModel() *model.Hub {
	return &model.Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *model.Client),
		Unregister: make(chan *model.Client),
		Clients:    make(map[*model.Client]bool),
	}
}

func (h *HubService) Run() {
	for {
		select {
		case client := <-h.hub.Register:
			h.hub.Clients[client] = true
		case client := <-h.hub.Unregister:
			if _, ok := h.hub.Clients[client]; ok {
				delete(h.hub.Clients, client)
				close(client.Send)
			}
		case message := <-h.hub.Broadcast:
			for client := range h.hub.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.hub.Clients, client)
				}
			}
		}
	}
}
