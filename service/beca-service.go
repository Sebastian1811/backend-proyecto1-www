package service

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/repository"
)

type BecaService interface {
	Save(entity.Beca) entity.Beca
	GetAll() []entity.Beca
	Update(beca entity.Beca)
	Delete(beca entity.Beca)
}

type becaService struct {
	becaRepository repository.BecaRepository
}

func New(repo repository.BecaRepository) BecaService {
	return &becaService{
		becaRepository: repo,
	}
}

func (service *becaService) Save(beca entity.Beca) entity.Beca {
	service.becaRepository.Save(beca)
	return beca
}

func (service *becaService) Update(beca entity.Beca) {
	service.becaRepository.Update(beca)
}

func (service *becaService) Delete(beca entity.Beca) {
	service.becaRepository.Delete(beca)
}

func (service *becaService) GetAll() []entity.Beca {
	return service.becaRepository.GetAll()
}
