package migration

import (
	db "docker/api/config"

	"docker/api/models"
)

func AutoMigration() bool {
	db := db.Connect()
	defer db.Close()

	db.AutoMigrate(models.NivelAcesso{})
	db.AutoMigrate(models.Administrador{})
	db.AutoMigrate(models.Menu{})
	db.AutoMigrate(models.NivelAcessoMenu{})
	db.AutoMigrate(models.Ensino{})

	//Definindo as chaves estrangeiras (foreign key)
	//Administrador
	db.Model(&models.Administrador{}).AddForeignKey("nivel_acesso_id", "nivel_acesso(id)", "RESTRICT", "RESTRICT")
	//Menu
	db.Model(&models.Menu{}).AddForeignKey("menu_pai", "menu(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Menu{}).AddForeignKey("administrador_login", "administrador(login)", "RESTRICT", "RESTRICT")
	//Nivel Acesso Menu
	db.Model(&models.NivelAcessoMenu{}).AddForeignKey("nivel_acesso_id", "nivel_acesso(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.NivelAcessoMenu{}).AddForeignKey("menu_id", "menu(id)", "RESTRICT", "RESTRICT")

	return true
}
