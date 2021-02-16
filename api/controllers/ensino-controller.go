package controllers

import (
	"net/http"
	"strconv"
	"time"

	"docker/api/models"
	"docker/api/utils"

	"github.com/gorilla/mux"
)

func InsertEnsino(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	Ensino := &models.Ensino{}
	utils.ParseBody(r, Ensino)
	Ensino.DataCriacao = time.Now()

	var campos = []string{"descricao", "data_criacao"}
	rows := models.Insert("ensino", campos, Ensino)

	if rows > 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Ensino inserido com sucesso!", nil)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao inserir o ensino!", nil)
	}
}

func UpdateEnsino(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	var alteraEns = &models.Ensino{}
	utils.ParseBody(r, alteraEns)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id > 0 {
		ens := models.SelectEnsinoById(id)

		if ens.ID > 0 {
			if alteraEns.Descricao != "" {
				ens.Descricao = alteraEns.Descricao
			}
			ens.DataAlteracao = time.Now()

			campos := []string{"descricao", "data_alteracao"}
			clausula := map[string]interface{}{"id": id}
			rows := models.Update("ensino", campos, clausula, ens)
			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Ensino alterado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao alterar o ensino!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o ensino!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do ensino!", nil)
	}
}

func DeleteEnsino(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id > 0 {
		verificaEns := models.SelectEnsinoById(id)

		if verificaEns.ID > 0 {

			clausula := map[string]interface{}{"id": id}
			rows := models.Delete("ensino", clausula, verificaEns)

			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Ensino deletado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 200, "Error", "Ocorreu um erro ao deletar o ensino!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o ensino!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o id do ensino!", nil)
	}
}

func SelectEnsinoById(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ensino := &models.Ensino{}
	campos := []string{"id", "descricao"}
	clausula := map[string]interface{}{"id": id}
	models.Select("ensino", campos, clausula, ensino)

	if ensino.ID != 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Ensino encontrado!", ensino)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi possível encontrar o ensino!", nil)
	}
}

func SelectAllEnsino(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	// ensino := &models.Ensino{}
	array := models.SelectAll("ensino")

	utils.SetJsonReturn(w, true, 200, "Success", "Ensinos encontrados!", array)
}
