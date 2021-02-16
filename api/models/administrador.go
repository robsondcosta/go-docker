package models

import (
	"time"

	"docker/api/config"
)

func (Administrador) TableName() string {
	return "administrador"
}

type Administrador struct {
	Login         string    `gorm:"primary_key;size:25;not null" json:"login"`
	Senha         string    `gorm:"not null" json:"senha"`
	Nome          string    `gorm:"not null" json:"nome"`
	Email         string    `json:"email"`
	Status        bool      `json:"status"`
	NivelAcessoID int       `gorm:"foreignKey: nivel_acesso_id;not null" json:"nivel_acesso_id"`
	DataCriacao   time.Time `json:"data_criacao"`
	DataAlteracao time.Time `json:"data_alteracao"`
}

func SelectAllAdministrador() []Administrador {
	db = config.Connect()
	defer db.Close()

	var arrayAdm []Administrador
	db.Table("administrador").Find(&arrayAdm)
	return arrayAdm
}

func SelectAdministradorByLogin(login string) *Administrador {
	db = config.Connect()
	defer db.Close()

	var adm Administrador
	db.Table("administrador").Where("login = ?", login).Find(&adm)
	return &adm
}
