package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	versionHeader = "Accept"
	appName       = "BitmediaGo"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps Dependencies) (router *mux.Router) {

	router = mux.NewRouter()

	// No version requirement for /ping
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	router.HandleFunc("/users", getAllUsersHandle(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/users", createUserHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)

	return
}
