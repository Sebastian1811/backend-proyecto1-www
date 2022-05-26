package entity

type Requisito struct {
	ID           int    `json:"id" gorm:"primary_key;auto_increment"`
	Descripicion string `json:"descripcion" binding:"required,gte=1,lte=250" gorm:"type:varchar(250)"`
	BecaID       int    `json:"becaid" gorm:"type:int" `
}
