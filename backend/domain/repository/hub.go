package repository

type IHubRepository interface {
}

type hubRepository struct{}

func NewHubRepository() IHubRepository {
	return &hubRepository{}
}
