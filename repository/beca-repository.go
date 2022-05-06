package repository

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BecaRepository interface {
	Save(beca entity.Beca)
	Update(beca entity.Beca)
	Delete(beca entity.Beca)
	GetAll() []entity.Beca
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
func (db *database) Delete(beca entity.Beca) {
	db.connection.Delete(&beca)
}
func (db *database) GetAll() []entity.Beca {
	var becas []entity.Beca
	db.connection.Find(&becas)
	return becas
}
