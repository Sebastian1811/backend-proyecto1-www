package controller

import "github.com/Sebastian1811/backend-proyecto1-www/entity"

type RequisitoController interface {
	Save(entity.Requisito)
	Update(entity.Requisito)
	Delete(int)
	GetAll(int)
	GetById(int, int)
}
