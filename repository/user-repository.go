package repository

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(entity.User)
	Login(int) entity.User
}

type userDB struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userDB{
		connection: db,
	}
}

func (db *userDB) Register(newuser entity.User) {
	db.connection.Create(&newuser)
}
func (db *userDB) Login(id int) entity.User {
	var usuario entity.User
	db.connection.Find(&usuario, id)
	return usuario
}
