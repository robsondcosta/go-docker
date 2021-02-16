package controllers

import (
	"net/http"
	"strconv"
	"time"

	"docker/api/models"
	"docker/api/utils"

	"github.com/gorilla/mux"
)

func InsertMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	menu := &models.Menu{}
	utils.ParseBody(r, menu)
	menu.DataCriacao = time.Now()

	var campos = []string{"descricao", "url", "icone", "status", "menu_pai", "administrador_login", "data_criacao"}
	rows := models.Insert("menu", campos, menu)

	if rows > 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Menu inserido com sucesso!", nil)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao inserir o menu!", nil)
	}
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	var alteraMenu = &models.Menu{}
	utils.ParseBody(r, alteraMenu)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id > 0 {
		menu := models.SelectMenuById(id)

		if menu.ID > 0 {
			if alteraMenu.Descricao != "" {
				menu.Descricao = alteraMenu.Descricao
			}
			if alteraMenu.URL != "" {
				menu.URL = alteraMenu.URL
			}
			if alteraMenu.Icone != "" {
				menu.Icone = alteraMenu.Icone
			}
			if alteraMenu.Status != menu.Status {
				menu.Status = alteraMenu.Status
			}
			if alteraMenu.MenuPai != menu.MenuPai {
				menu.MenuPai = alteraMenu.MenuPai
			}
			menu.DataAlteracao = time.Now()

			campos := []string{"descricao", "url", "icone", "status", "menu_pai", "data_alteracao"}
			clausula := map[string]interface{}{"id": id}
			rows := models.Update("ensino", campos, clausula, menu)

			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Menu alterado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao alterar o menu!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o menu!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do menu!", nil)
	}
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id > 0 {
		verificaMenu := models.SelectMenuById(id)

		if verificaMenu.ID > 0 {

			clausula := map[string]interface{}{"id": id}
			rows := models.Delete("menu", clausula, verificaMenu)

			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Menu deletado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 200, "Error", "Ocorreu um erro ao deletar o menu!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o menu!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do menu!", nil)
	}
}

func SelectAllMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)
	menu := models.SelectAllMenu()
	utils.SetJsonReturn(w, true, 200, "Success", "Menus encontrados!", menu)
}

func SelectMenuById(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	menu := models.SelectMenuById(id)

	if menu.ID != 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Menu encontrado!", menu)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi possível encontrar o menu!", nil)
	}
}
