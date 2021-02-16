package models

import (
	"time"

	"docker/api/config"
)

func (NivelAcesso) TableName() string {
	return "nivel_acesso"
}

type NivelAcesso struct {
	ID            int    `json:"id"`
	Descricao     string `gorm:"not null" json:"descricao"`
	Status        bool   `gorm:"default:1" json:"status"`
	DataCriacao   time.Time
	DataAlteracao time.Time
}

func SelectAllNivelAcesso() []NivelAcesso {
	db = config.Connect()
	defer db.Close()

	var array []NivelAcesso
	db.Table("nivel_acesso").Find(&array)
	return array
}

func SelectNivelAcessoById(id int) *NivelAcesso {
	db = config.Connect()
	defer db.Close()

	var item NivelAcesso
	db.Table("nivel_acesso").Where("id = ?", id).Find(&item)
	return &item
}
