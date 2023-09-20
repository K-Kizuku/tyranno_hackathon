package repository

type IClientRepository interface {
	// ServeWs(hub *model.Hub, w http.ResponseWriter, r *http.Request)
	// ReadPump()
	// WritePump()
}

type clientRepository struct {
}

func NewClientRepository() IClientRepository {
	return &clientRepository{}
}
