package controllers

import (
	"net/http"
	"strconv"
	"time"

	"docker/api/models"
	"docker/api/utils"

	"github.com/gorilla/mux"
)

func InsertNivelAcesso(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	nivel := &models.NivelAcesso{}
	utils.ParseBody(r, nivel)
	nivel.DataCriacao = time.Now()

	var campos = []string{"descricao", "status", "data_criacao"}
	rows := models.Insert("nivel_acesso", campos, nivel)

	if rows > 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso inserido com sucesso!", nil)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao inserir o nível de acesso!", nil)
	}
}

func UpdateNivelAcesso(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	var alteraNivel = &models.NivelAcesso{}
	utils.ParseBody(r, alteraNivel)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id > 0 {
		nivel := models.SelectNivelAcessoById(id)

		if nivel.ID > 0 {
			if alteraNivel.Descricao != "" {
				nivel.Descricao = alteraNivel.Descricao
			}
			if alteraNivel.Status != nivel.Status {
				nivel.Status = alteraNivel.Status
			}
			nivel.DataAlteracao = time.Now()

			var campos = []string{"descricao", "status", "data_alteracao"}
			clausula := map[string]interface{}{"id": id}
			rows := models.Update("nivel_acesso", campos, clausula, nivel)
			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Nível de Acesso alterado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao alterar o nível de acesso!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o nível de acesso!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do nível de acesso!", nil)
	}
}

func DeleteNivelAcesso(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id > 0 {
		verificaNivel := models.SelectNivelAcessoById(id)

		if verificaNivel.ID > 0 {

			clausula := map[string]interface{}{"id": id}
			rows := models.Delete("nivel_acesso", clausula, verificaNivel)

			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso deletado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 200, "Error", "Ocorreu um erro ao deletar o nível de acesso!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o nível de acesso!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do nível de acesso!", nil)
	}
}

func SelectAllNivelAcesso(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)
	nivel := models.SelectAllNivelAcesso()
	utils.SetJsonReturn(w, true, 200, "Success", "Níveis de acesso encontrados!", nivel)
}

func SelectNivelAcessoById(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	nivel := models.SelectNivelAcessoById(id)

	if nivel.ID != 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Nível de acesso encontrado!", nivel)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi possível encontrar o nível de acesso!", nil)
	}
}
