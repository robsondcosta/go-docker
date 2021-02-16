package controllers

import (
	"fmt"
	"net/http"
	"time"

	"docker/api/models"
	"docker/api/utils"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func InsertAdministrador(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	Administrador := &models.Administrador{}
	utils.ParseBody(r, Administrador)

	hash, err := bcrypt.GenerateFromPassword([]byte(Administrador.Senha), 5)

	if err != nil {
		utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro interno", nil)
		return
	}

	Administrador.Senha = string(hash)
	Administrador.DataCriacao = time.Now()

	var campos = []string{"login", "senha", "nome", "email", "status", "nivel_acesso_id", "data_criacao"}
	rows := models.Insert("administrador", campos, Administrador)

	if rows > 0 {
		utils.SetJsonReturn(w, true, 200, "Success", "Administrador inserido com sucesso!", nil)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao inserir o administrador!", nil)
	}
}

func UpdateAdministrador(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	var alteraAdm = &models.Administrador{}
	utils.ParseBody(r, alteraAdm)
	vars := mux.Vars(r)
	login := vars["login"]

	if login != "" {
		administrador := models.SelectAdministradorByLogin(login)

		if administrador.Login != "" {
			if alteraAdm.Senha != "" {
				administrador.Senha = alteraAdm.Senha
			}
			if alteraAdm.Nome != "" {
				administrador.Nome = alteraAdm.Nome
			}
			if alteraAdm.Email != "" {
				administrador.Email = alteraAdm.Email
			}
			if alteraAdm.Status != administrador.Status {
				administrador.Status = alteraAdm.Status
			}
			if alteraAdm.NivelAcessoID != administrador.NivelAcessoID {
				administrador.NivelAcessoID = alteraAdm.NivelAcessoID
			}
			administrador.DataAlteracao = time.Now()

			var campos = []string{"senha", "nome", "email", "status", "nivel_acesso_id", "data_alteracao"}
			clausula := map[string]interface{}{"login": login}
			rows := models.Update("administrador", campos, clausula, administrador)

			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Administrador alterado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao alterar o administrador!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o administrador!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o login do administrador!", nil)
	}
}

func DeleteAdministrador(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	login := vars["login"]

	if login != "" {
		verificaAdm := models.SelectAdministradorByLogin(login)

		if verificaAdm.Login != "" {

			clausula := map[string]interface{}{"login": login}
			rows := models.Delete("administrador", clausula, verificaAdm)

			if rows > 0 {
				utils.SetJsonReturn(w, true, 200, "Success", "Administrador deletado com sucesso!", nil)
			} else {
				utils.SetJsonReturn(w, false, 400, "Error", "Ocorreu um erro ao deletar o administrador!", nil)
			}
		} else {
			utils.SetJsonReturn(w, false, 400, "Error", "Não foi possivel localizar o administrador!", nil)
		}
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi informado o login do administrador!", nil)
	}
}

func SelectAllAdministrador(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	adms := models.SelectAll("administrador")

	utils.SetJsonReturn(w, true, 200, "Success", "Administradores encontrados!", adms)
}

func SelectAdministradorByLogin(w http.ResponseWriter, r *http.Request) {
	utils.SetJsonHeader(w)

	vars := mux.Vars(r)
	login := vars["login"]
	fmt.Println(login)

	adm := &models.Administrador{}
	campos := []string{"login", "nome", "email"}
	clausula := map[string]interface{}{"login": login}
	models.Select("administrador", campos, clausula, adm)

	if adm.Login != "" {
		utils.SetJsonReturn(w, true, 200, "Success", "Administrador encontrado!", adm)
	} else {
		utils.SetJsonReturn(w, false, 400, "Error", "Não foi possível encontrar o administrador!", nil)
	}
}
