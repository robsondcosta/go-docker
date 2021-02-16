package models

import (
	"docker/api/config"
)

func (NivelAcessoMenu) TableName() string {
	return "nivel_acesso_menu"
}

type NivelAcessoMenu struct {
	NivelAcessoID int `gorm:"foreignKey: nivel_acesso_id;not null" json:"nivel_acesso_id"`
	MenuID        int `gorm:"foreignKey: menu_id;not null" json:"menu_id"`
}

//gorm:"foreign_key;not null"

func SelectAllNivelMenu() []NivelAcessoMenu {
	db = config.Connect()
	defer db.Close()

	var array []NivelAcessoMenu
	db.Table("nivel_acesso_menu").Find(&array)
	return array
}

func SelectNivelMenuByMenuId(id int) []NivelAcessoMenu {
	db = config.Connect()
	defer db.Close()

	var array []NivelAcessoMenu
	db.Table("nivel_acesso_menu").Where("menu_id = ?", id).Find(&array)
	return array
}

func SelectNivelMenuByNivelId(id int) []NivelAcessoMenu {
	db = config.Connect()
	defer db.Close()

	var array []NivelAcessoMenu
	db.Table("nivel_acesso_menu").Where("nivel_acesso_id = ?", id).Find(&array)
	return array
}

func SelectNivelMenu(nivelId, menuId int) *NivelAcessoMenu {
	db = config.Connect()
	defer db.Close()

	var item NivelAcessoMenu
	db.Table("nivel_acesso_menu").Where("nivel_acesso_id = ? and menu_id = ?", nivelId, menuId).Find(&item)
	return &item
}
