package repository

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"gorm.io/gorm"
)

type RequisitoRepository interface {
	Save(entity.Requisito)
	Update(entity.Requisito)
	Delete(int)
	GetAll(int) []entity.Requisito
	GetById(int, int) entity.Requisito
}

type requisitoDB struct {
	connection *gorm.DB
}

func NewRequesitoRepository(db *gorm.DB) RequisitoRepository {
	return &requisitoDB{
		connection: db,
	}
}

func (db *requisitoDB) Update(requisito entity.Requisito) {
	db.connection.Save(&requisito)

}
func (db *requisitoDB) Save(requisito entity.Requisito) {
	db.connection.Create(&requisito)
}

func (db *requisitoDB) Delete(id int) {
	db.connection.Where("beca_id = ?", id).Delete(&entity.Requisito{})
	//Sdb.connection.Delete(&entity.Requisito{}, id)
}

func (db *requisitoDB) GetById(idbeca int, idreq int) entity.Requisito {
	var requisito entity.Requisito
	db.connection.Where("beca_id = ? AND id", idbeca, idreq).First(&requisito)
	return requisito
}

func (db *requisitoDB) GetAll(idBeca int) []entity.Requisito {
	var requesitos []entity.Requisito
	db.connection.Where("beca_id = ?", idBeca).Find(&requesitos)
	return requesitos
}
