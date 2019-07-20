package zk

import (
	"net/http"
	"go-example/util"
	"go-example/zk"
	"log"
	"github.com/gorilla/mux"
)

func GetServiceNodes(w http.ResponseWriter, r *http.Request) {
	data, err := zk.ZKClient.GetServiceNodeList()
	if err != nil {
		log.Printf("Get zk services node err :%s", err)
	}
	util.ResponseJSON(w, http.StatusOK, data)
}

func GetServiceNodeData(w http.ResponseWriter, r *http.Request) {
	serviceNodeName := mux.Vars(r)["node"]

	data, err := zk.ZKClient.GetServiceNodeData(serviceNodeName)
	if err != nil {
		log.Printf("Get zk services node data err :%s", err)
	}

	util.ResponseJSON(w, http.StatusOK, data)
}
