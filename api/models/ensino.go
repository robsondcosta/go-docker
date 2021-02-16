package models

import (
	"time"

	"docker/api/config"
)

func (Ensino) TableName() string {
	return "ensino"
}

type Ensino struct {
	ID            int64  `json:"id"`
	Descricao     string `gorm:"not null" json:"descricao"`
	DataCriacao   time.Time
	DataAlteracao time.Time
}

func SelectAllEnsino() []Ensino {
	db = config.Connect()
	defer db.Close()

	var array []Ensino
	db.Table("ensino").Find(&array)
	return array
}

func SelectEnsinoById(id int) *Ensino {
	db = config.Connect()
	defer db.Close()

	var item Ensino
	db.Table("ensino").Where("id = ?", id).Find(&item)
	return &item
}
