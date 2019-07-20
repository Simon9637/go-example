package user

import (
	"net/http"
	"go-example/model"
	"go-example/util"
	"github.com/gorilla/mux"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := model.User{
		Id:   id,
		Name: "zhang",
		Age:  20,
	}
	util.ResponseJSON(w, http.StatusOK, user)
}

func Post(w http.ResponseWriter, r *http.Request) {
	payload := model.User{}
	util.RequestJSON(r, &payload)
	util.ResponseJSON(w, http.StatusOK, payload)
}

func Put(w http.ResponseWriter, r *http.Request) {
	payload := model.User{}
	util.RequestJSON(r, &payload)
	util.ResponseJSON(w, http.StatusOK, payload)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := model.User{
		Id:   id,
		Name: "zhang",
		Age:  20,
	}
	util.ResponseJSON(w, http.StatusOK, user)
}
