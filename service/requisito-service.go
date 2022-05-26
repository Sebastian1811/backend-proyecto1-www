package service

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/repository"
)

type RequisitoService interface {
	Save(entity.Requisito)
	Update(entity.Requisito)
	Delete(int)
	GetAll(int) []entity.Requisito
	GetById(int, int) entity.Requisito
}

type requisitoService struct {
	requisitoRepository repository.RequisitoRepository
}

func NewRequisitoService(repo repository.RequisitoRepository) RequisitoService {
	return &requisitoService{
		requisitoRepository: repo,
	}
}

func (service *requisitoService) Update(requisito entity.Requisito) {
	service.requisitoRepository.Update(requisito)

}
func (service *requisitoService) Save(requisito entity.Requisito) {
	service.requisitoRepository.Save(requisito)
}

func (service *requisitoService) Delete(id int) {
	service.requisitoRepository.Delete(id)
}

func (service *requisitoService) GetById(idbeca int, idreq int) entity.Requisito {
	requisito := service.requisitoRepository.GetById(idbeca, idreq)
	return requisito
}

func (service *requisitoService) GetAll(idBeca int) []entity.Requisito {
	requesitos := service.requisitoRepository.GetAll(idBeca)
	return requesitos
}
