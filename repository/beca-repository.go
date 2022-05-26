package repository

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"gorm.io/gorm"
)

type BecaRepository interface {
	Save(entity.Beca)
	Update(entity.Beca)
	Delete(int)
	GetAll() []entity.Beca
	GetById(int) entity.Beca
	GetByCategoria(string) []entity.Beca
}

type becaDB struct {
	connection *gorm.DB
}

func NewBecaRepository(db *gorm.DB) BecaRepository {
	return &becaDB{
		connection: db,
	}
}

func (db *becaDB) Save(beca entity.Beca) {
	db.connection.Create(&beca)
}
func (db *becaDB) Update(beca entity.Beca) {
	db.connection.Save(&beca)
}
func (db *becaDB) Delete(id int) {
	db.connection.Delete(&entity.Beca{}, id)
}
func (db *becaDB) GetAll() []entity.Beca {
	var becas []entity.Beca
	db.connection.Find(&becas)
	return becas
}
func (db *becaDB) GetById(id int) entity.Beca {
	var beca entity.Beca
	db.connection.First(&beca, id)
	return beca
}
func (db *becaDB) GetByCategoria(categoria string) []entity.Beca {
	var becas []entity.Beca
	db.connection.Where("categoria = ?", categoria).Find(&becas)
	return becas
}
