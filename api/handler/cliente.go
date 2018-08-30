package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang/db"
	"github.com/golang/model"
	"github.com/golang/util"
	"github.com/gorilla/mux"
)

//GravaCliente funçao grava cliente
func GravaCliente(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var c model.Cliente
	var d db.DB

	err := d.Connection()

	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Erro ao acessar bd", "")
		return
	}

	db := d.DB

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&c)

	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}

	defer r.Body.Close()
	err = c.GravaCliente(db)

	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, c, 0, 0)

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }

	log.Println(c)

}

//UpdateCliente ...
func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	var a model.Cliente
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateCliente] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	a.ID = int64(id)
	if err := a.UpdateCliente(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, a, 0, 0)

}

//GetClientes ...
func GetClientes(w http.ResponseWriter, r *http.Request) {
	var c model.Cliente
	var t util.App
	var d db.DB

	err := d.Connection()
	db := d.DB
	defer db.Close()

	clientes := c.GetClientes(db)

	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, clientes, 0, 0)

}
