package routes

import (
	"docker/api/controllers"
	auth "docker/api/controllers/auth"

	"github.com/gorilla/mux"
)

var AutenticacaoRoutes = func(router *mux.Router) {
	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.HandleFunc("/primeiroacesso", controllers.InsertAdministrador).Methods("POST")
}

var AdministradorRoutes = func(router *mux.Router) {
	router.HandleFunc("/administrador", auth.ValidateMiddleware(controllers.InsertAdministrador)).Methods("POST")
	router.HandleFunc("/administrador", auth.ValidateMiddleware(controllers.SelectAllAdministrador)).Methods("GET")
	router.HandleFunc("/administrador/{login}", auth.ValidateMiddleware(controllers.SelectAdministradorByLogin)).Methods("GET")
	router.HandleFunc("/administrador/{login}", auth.ValidateMiddleware(controllers.UpdateAdministrador)).Methods("PUT")
	router.HandleFunc("/administrador/{login}", auth.ValidateMiddleware(controllers.DeleteAdministrador)).Methods("DELETE")
}

var EnsinoRoutes = func(router *mux.Router) {
	router.HandleFunc("/ensino", auth.ValidateMiddleware(controllers.InsertEnsino)).Methods("POST")
	router.HandleFunc("/ensino", auth.ValidateMiddleware(controllers.SelectAllEnsino)).Methods("GET")
	router.HandleFunc("/ensino/{id}", auth.ValidateMiddleware(controllers.SelectEnsinoById)).Methods("GET")
	router.HandleFunc("/ensino/{id}", auth.ValidateMiddleware(controllers.UpdateEnsino)).Methods("PUT")
	router.HandleFunc("/ensino/{id}", auth.ValidateMiddleware(controllers.DeleteEnsino)).Methods("DELETE")
}

var NivelAcessoRoutes = func(router *mux.Router) {
	router.HandleFunc("/nivelacesso", auth.ValidateMiddleware(controllers.InsertNivelAcesso)).Methods("POST")
	router.HandleFunc("/nivelacesso", auth.ValidateMiddleware(controllers.SelectAllNivelAcesso)).Methods("GET")
	router.HandleFunc("/nivelacesso/{id}", auth.ValidateMiddleware(controllers.SelectNivelAcessoById)).Methods("GET")
	router.HandleFunc("/nivelacesso/{id}", auth.ValidateMiddleware(controllers.UpdateNivelAcesso)).Methods("PUT")
	router.HandleFunc("/nivelacesso/{id}", auth.ValidateMiddleware(controllers.DeleteNivelAcesso)).Methods("DELETE")
}

var MenuRoutes = func(router *mux.Router) {
	router.HandleFunc("/menu", auth.ValidateMiddleware(controllers.InsertMenu)).Methods("POST")
	router.HandleFunc("/menu", auth.ValidateMiddleware(controllers.SelectAllMenu)).Methods("GET")
	router.HandleFunc("/menu/{id}", auth.ValidateMiddleware(controllers.SelectMenuById)).Methods("GET")
	router.HandleFunc("/menu/{id}", auth.ValidateMiddleware(controllers.UpdateMenu)).Methods("PUT")
	router.HandleFunc("/menu/{id}", auth.ValidateMiddleware(controllers.DeleteMenu)).Methods("DELETE")
}

var NivelMenuRoutes = func(router *mux.Router) {
	router.HandleFunc("/nivelmenu", auth.ValidateMiddleware(controllers.InsertNivelMenu)).Methods("POST")
	router.HandleFunc("/nivelmenu", auth.ValidateMiddleware(controllers.SelectAllNivelMenu)).Methods("GET")
	router.HandleFunc("/nivelmenu/{nivelId}/{menuId}", auth.ValidateMiddleware(controllers.SelectNivelMenu)).Methods("GET")
	router.HandleFunc("/nivelmenu/{nivelId}/{menuId}", auth.ValidateMiddleware(controllers.UpdateNivelMenu)).Methods("PUT")
	router.HandleFunc("/nivelmenu/{nivelId}/{menuId}", auth.ValidateMiddleware(controllers.DeleteNivelMenu)).Methods("DELETE")
	router.HandleFunc("/nivelmenu/{nivelId}", auth.ValidateMiddleware(controllers.SelectNivelMenuByNivelId)).Methods("GET")
	router.HandleFunc("/nivelmenu/{menuId}", auth.ValidateMiddleware(controllers.SelectNivelMenuByMenuId)).Methods("GET")
}
