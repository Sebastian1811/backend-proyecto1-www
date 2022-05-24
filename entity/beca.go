package entity

import "time"

type Beca struct {
	ID                     uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Categoria              string `json:"categoria" binding:"required,gte=1,lte=40" gorm:"type:varchar(40)"`
	Nombre                 string `json:"nombre" binding:"required,gte=1,lte=30" gorm:"type:varchar(30)"`
	PorcentajeFinanciacion int    `json:"porcentajeF" binding:"required,gte=1,lte=100" gorm:"type:int"`
	Pais                   string `json:"pais" binding:"required,gte=1,lte=40" gorm:"type:varchar(40)"`
	//Requisitos
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
