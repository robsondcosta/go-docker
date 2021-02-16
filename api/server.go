package api

import (
	"log"
	"net/http"

	"docker/api/config/migration"
	"docker/api/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Run() {
	r := mux.NewRouter()
	routes.AutenticacaoRoutes(r)
	routes.AdministradorRoutes(r)
	routes.EnsinoRoutes(r)
	routes.NivelAcessoRoutes(r)
	routes.MenuRoutes(r)
	routes.NivelMenuRoutes(r)
	http.Handle("/", r)
	migration.AutoMigration()
	log.Fatal(http.ListenAndServe(":3000", r))
}
