package todocontroller

import (
	"encoding/json"
	"github.com/manjeshpv/qgotify/server/api/todo/dao"
	todo "github.com/manjeshpv/qgotify/server/api/todo/model"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"fmt"
)

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("getall")
	ts, err := tododao.All()
	fmt.Println("ts",ts)
	w.Header().Set("Content-Type", "application/json")


	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tsm, err := json.Marshal(ts)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write(tsm)
}

func NewTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	t := todo.Todo{}

	tf, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(tf, &t)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	nt, err := tododao.NewTodo(t)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ntm, err := json.Marshal(nt)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write(ntm)
}

func RemoveTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")

	err := tododao.DeleteTodo(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
