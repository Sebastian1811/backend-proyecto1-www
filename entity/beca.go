package entity

// Añadir requerimientos, quizas como otra entidad porque son muchos.

type Beca struct {
	Categoria              string `json:"categoria"`
	Nombre                 string `json:"nombre"`
	PorcentajeFinanciacion int    `json:"porcentajeF"`
	Pais                   string `json:"pais"`
	//Requisitos
}
