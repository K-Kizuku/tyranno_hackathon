package repository

import "tyranno/backend/domain/model"

type Stats interface {
	CalcPos(input model.Input)
	CalcScore(input model.Input)
	DeleteMino(input model.Input)
}
