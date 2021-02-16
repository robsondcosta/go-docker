package controllers

import (
	"net/http"
	"strconv"

	"docker/api/models"
	"docker/api/utils"

	"github.com/gorilla/mux"
)

func InsertNivelMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	nivel := &models.NivelAcessoMenu{}
	utils.ParseBody(r, nivel)
	campos := []string{"nivel_acesso_id", "menu_id"}
	rows := models.Insert("nivel_acesso", campos, nivel)

	if rows > 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso do menu inserido com sucesso!", nil)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao inserir o nível de acesso do menu!", nil)
	}
}

func UpdateNivelMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	var alteraNivel = &models.NivelAcessoMenu{}
	utils.ParseBody(r, alteraNivel)
	vars := mux.Vars(r)
	idNivel, _ := strconv.Atoi(vars["nivelId"])
	idMenu, _ := strconv.Atoi(vars["menuId"])

	if idNivel > 0 && idMenu > 0 {
		nivel := models.SelectNivelMenu(idNivel, idMenu)

		if nivel.NivelAcessoID > 0 && nivel.MenuID > 0 {
			if nivel.NivelAcessoID != alteraNivel.NivelAcessoID {
				nivel.NivelAcessoID = alteraNivel.NivelAcessoID
			}
			if nivel.MenuID != alteraNivel.MenuID {
				nivel.MenuID = alteraNivel.MenuID
			}

			campos := []string{"nivel_acesso_id", "menu_id"}
			clausula := map[string]interface{}{"nivel_acesso_id": idNivel, "menu_id": idMenu}
			rows := models.Update("nivel_acesso_menu", campos, clausula, nivel)
			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso do menu alterado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao alterar o nível de acesso do menu!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o nível de acesso do menu!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do nível de acesso do menu!", nil)
	}
}

func DeleteNivelMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	idNivel, _ := strconv.Atoi(vars["nivelId"])
	idMenu, _ := strconv.Atoi(vars["menuId"])

	if idNivel > 0 && idMenu > 0 {
		verificaNivel := models.SelectNivelMenu(idNivel, idMenu)

		if verificaNivel.NivelAcessoID > 0 && verificaNivel.MenuID > 0 {

			clausula := map[string]interface{}{"nivel_acesso_id": idNivel, "menu_id": idMenu}
			rows := models.Delete("nivel_acesso", clausula, verificaNivel)

			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso do menu deletado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 200, "Error", "Ocorreu um erro ao deletar o nível de acesso do menu!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o nível de acesso do menu!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do nível de acesso do menu!", nil)
	}
}

func SelectAllNivelMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)
	nivel := models.SelectAllNivelMenu()
	utils.SetJsonReturn(w, true, 200, "Success", "Níveis de acesso do menu encontrados!", nivel)
}

func SelectNivelMenu(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	idNivel, _ := strconv.Atoi(vars["nivelId"])
	idMenu, _ := strconv.Atoi(vars["menuId"])
	nivel := models.SelectNivelMenu(idNivel, idMenu)

	if nivel.NivelAcessoID != 0 && nivel.MenuID != 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso do menu encontrado!", nivel)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi possível encontrar o nível de acesso do menu!", nil)
	}
}

func SelectNivelMenuByNivelId(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["nivelId"])
	nivel := models.SelectNivelMenuByMenuId(id)

	if nivel != nil {
		utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso do menu encontrado!", nivel)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi possível encontrar o nível de acesso do menu!", nil)
	}
}

func SelectNivelMenuByMenuId(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["menuId"])
	nivel := models.SelectNivelMenuByMenuId(id)

	if nivel != nil {
		utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso do menu encontrado!", nivel)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi possível encontrar o nível de acesso do menu!", nil)
	}
}
