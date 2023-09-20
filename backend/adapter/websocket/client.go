package websocket

import (
	"net/http"
	"tyranno/backend/domain/model"
	"tyranno/backend/usecase"
)

type clientHandler struct {
	usecase usecase.IClientUsecase
}

func NewClientHandler(su usecase.IClientUsecase) *clientHandler {
	return &clientHandler{
		usecase: su,
	}
}

func (c *clientHandler) WSHandler(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	c.usecase.ServeWs(hub, w, r)
}
