package router

import (
	"github.com/gorilla/mux"
	"goStudyProject/handler/sd"
	"net/http"
	"goStudyProject/handler/user"
	"goStudyProject/handler/zk"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Basic check
	router.HandleFunc("/check/health", sd.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/check/disk", sd.DiskCheck).Methods(http.MethodGet)
	router.HandleFunc("/check/cpu", sd.CPUCheck).Methods(http.MethodGet)
	router.HandleFunc("/check/ram", sd.RAMCheck).Methods(http.MethodGet)

	// User
	router.HandleFunc("/user/get/{id}", user.Get).Methods(http.MethodGet)
	router.HandleFunc("/user/post", user.Post).Methods(http.MethodPost)
	router.HandleFunc("/user/put", user.Put).Methods(http.MethodPut)
	router.HandleFunc("/user/delete/{id}", user.Delete).Methods(http.MethodDelete)

	// ZK
	router.HandleFunc("/zk/services", zk.GetServiceNodes).Methods(http.MethodGet)
	router.HandleFunc("/zk/service/{node}/data", zk.GetServiceNodeData).Methods(http.MethodGet)

	return router
}