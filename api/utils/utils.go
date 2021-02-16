package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Padrao struct {
	ID            int64 `json:"id"`
	DataCriacao   time.Time
	DataAlteracao time.Time
}

type Description struct {
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
}

type Response struct {
	Status  bool        `json:"status"`
	Codigo  int64       `json:"codigo"`
	Message Description `json:"message"`
	Dados   interface{} `json:"dados"`
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SetJsonReturn(w http.ResponseWriter, status bool, codigo int64, titulo, descricao string, dados interface{}) {
	desc := Description{Titulo: titulo, Descricao: descricao}
	resp := Response{Status: status, Codigo: codigo, Message: desc, Dados: dados}
	json.NewEncoder(w).Encode(resp)
}
