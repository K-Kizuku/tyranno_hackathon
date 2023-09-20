package websocket

import (
	"net/http"
	"tyranno/backend/domain/model"
	"tyranno/backend/domain/repository"
	"tyranno/backend/domain/service"
	"tyranno/backend/usecase"
)

func InitWS(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	clientRepository := repository.NewClientRepository()
	clientService := service.NewClientService(clientRepository)
	clientUsecase := usecase.NewClientUsecase(clientService)

	clientHandler := NewClientHandler(clientUsecase)
	clientHandler.WSHandler(hub, w, r)
}
