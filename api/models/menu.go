package models

import (
	"time"

	"docker/api/config"
)

func (Menu) TableName() string {
	return "menu"
}

type Menu struct {
	ID                 int    `json:"id"`
	Descricao          string `gorm:"not null" json:"descricao"`
	URL                string `json:"url"`
	Icone              string `json:"icone"`
	Status             bool   `gorm:"default:1" json:"status"`
	MenuPai            int    `gorm:"foreignKey: menu_id;not null" json:"menu_pai"`
	AdministradorLogin string `gorm:"foreignKey: administrador_login;size:25;not null" json:"administrador_login"`
	DataCriacao        time.Time
	DataAlteracao      time.Time
}

func SelectAllMenu() []Menu {
	db = config.Connect()
	defer db.Close()

	var array []Menu
	db.Table("menu").Find(&array)
	return array
}

func SelectMenuById(id int) *Menu {
	db = config.Connect()
	defer db.Close()

	var item Menu
	db.Table("menu").Where("id = ?", id).Find(&item)
	return &item
}

func SelectMenuFilhos(id int) []Menu {
	db = config.Connect()
	defer db.Close()

	var array []Menu
	db.Table("menu").Where("menu_pai = ?", id).Find(&array)
	return array
}
