package handler

import (
	"database/sql"
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
	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")
	dataNascimento := r.FormValue("datanascimento")

	c.ID = int64(id)
	c.Nome = nome
	c.DataNascimento = dataNascimento

	clientes, err := c.GetClientes(db)

	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, clientes, 0, 0)

}

//DeleteCliente ...
func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	var c model.Cliente
	var d db.DB
	var t util.App

	err := d.Connection()

	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Erro ao tentar abrir conexão", "Teste")
	}
	db := d.DB
	defer db.Close()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	c.ID = int64(id)
	if err := c.DeleteCliente(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, c, 0, 0)
	//log.Println(string(""))

}

//GetCliente ...
func GetCliente(w http.ResponseWriter, r *http.Request) {
	var c model.Cliente
	var d db.DB
	var t util.App

	err := d.Connection()

	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Erro ao tentar abrir conexão", "Teste")
	}
	db := d.DB
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Erro ao tentar abrir conexão", "Teste")
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	c.ID = int64(id)
	if err := c.GetCliente(db); err != nil {
		if err == sql.ErrNoRows {
			log.Println("[handler/GetCliente] - Não existe cliente com este id", err.Error())
			t.ResponseWithError(w, http.StatusNotFound, "Não existe cliente para este id", "id")
		} else {
			log.Println("[handler/GetCliente] - Erro ao tentar consultar cliente", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
			return
		}

	}
	t.ResponseWithJSON(w, http.StatusOK, c, 0, 0)
	//log.Println(string(""))
}
