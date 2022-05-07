package repository

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"gorm.io/driver/postgres"
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
	dsn := "postgres://xjwhscksgtbqtr:e44cca7b7ba6e6843ece37c3f2681747b4de38ea29d7ca76955c5a81926114b1@ec2-3-217-113-25.compute-1.amazonaws.com:5432/d6lcnv0p7n75e8"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Beca{}, &entity.User{})

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
