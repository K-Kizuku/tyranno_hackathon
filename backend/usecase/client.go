package usecase

import (
	"net/http"
	"tyranno/backend/domain/model"
	"tyranno/backend/domain/service"
)

type IClientUsecase interface {
	ServeWs(hub *model.Hub, w http.ResponseWriter, r *http.Request)
}

type clientUsecase struct {
	service service.IClientService
}

func NewClientUsecase(ss service.IClientService) IClientUsecase {
	return &clientUsecase{
		service: ss,
	}
}

func (c *clientUsecase) ServeWs(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	c.service.ServeWs(hub, w, r)
}
