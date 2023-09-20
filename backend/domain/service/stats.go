package service

import (
	"tyranno/backend/domain/model"
	"tyranno/backend/domain/repository"
)

type Stats struct {
	repo repository.Stats
}

func (s *Stats) CalcPos(input model.Input) {

}

func (s *Stats) CalcScore(input model.Input) {

}

func (s *Stats) DeleteMino(input model.Input) {

}
