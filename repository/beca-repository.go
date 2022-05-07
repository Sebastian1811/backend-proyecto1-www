package repository

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BecaRepository interface {
	Save(entity.Beca)
	Update(entity.Beca)
	Delete(int)
	GetAll() []entity.Beca
	GetById(int) entity.Beca
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewBecaRepository() BecaRepository {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Beca{})

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	sqldb, _ := db.connection.DB()
	sqldb.Close()
}

func (db *database) Save(beca entity.Beca) {
	db.connection.Create(&beca)
}
func (db *database) Update(beca entity.Beca) {
	db.connection.Save(&beca)
}
func (db *database) Delete(id int) {
	db.connection.Delete(&entity.Beca{}, id)
}
func (db *database) GetAll() []entity.Beca {
	var becas []entity.Beca
	db.connection.Find(&becas)
	return becas
}
func (db *database) GetById(id int) entity.Beca {
	var beca entity.Beca
	db.connection.First(&beca, id)
	return beca
}
