package service

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/repository"
)

type BecaService interface {
	Save(entity.Beca)
	GetAll() []entity.Beca
	Update(entity.Beca)
	Delete(int)
	GetById(int) entity.Beca
	GetByCategoria(string) []entity.Beca
}

type becaService struct {
	becaRepository repository.BecaRepository
}

func New(repo repository.BecaRepository) BecaService {
	return &becaService{
		becaRepository: repo,
	}
}

func (service *becaService) Save(beca entity.Beca) {
	service.becaRepository.Save(beca)

}

func (service *becaService) Update(beca entity.Beca) {
	service.becaRepository.Update(beca)
}

func (service *becaService) Delete(id int) {
	service.becaRepository.Delete(id)
}

func (service *becaService) GetAll() []entity.Beca {
	return service.becaRepository.GetAll()
}

func (service *becaService) GetById(id int) entity.Beca {
	return service.becaRepository.GetById(id)
}
func (service *becaService) GetByCategoria(categoria string) []entity.Beca {
	return service.becaRepository.GetByCategoria(categoria)
}
