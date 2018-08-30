package model

import (
	"database/sql"
)

//Cliente estrutura
type Cliente struct {
	ID             int64  `json:"id"`
	Nome           string `json:"nome"`
	DataNascimento string `json:"dataNascimento"`
}

//GravaCliente função grava cliente
func (c *Cliente) GravaCliente(db *sql.DB) error {
	smt, err := db.Prepare(`insert into clientes 
							 (nome,data_nascimento)
							 values(?,?)`)
	if err != nil {
		return err
	}
	res, err := smt.Exec(c.Nome, c.DataNascimento)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	c.ID = id
	return nil
	//log.Println(c)

}

//UpdateCliente função update cliente
func (c *Cliente) UpdateCliente(db *sql.DB) error {
	smt, err := db.Prepare(`update clientes 
						  set nome        = ?,
						  data_nascimento = ? 
						  where  id =      ?`)
	if err != nil {
		return err
	}
	_, err = smt.Exec(c.Nome, c.DataNascimento, c.ID)
	return err
}

//GetCliente função get cliente
func GetCliente() {

}

//GetClientes função get clientes
func (c *Cliente) GetClientes(db *sql.DB) []Cliente {

	//smt, err := db.Prepare(`Select * from clientes `)
	// if err != nil {
	// 	return nil, err
	// }

	smt, err := db.Query(`Select * from clientes`)
	//defer db.Close()

	//clientes, err := smt.Exec()
	clientes := []Cliente{}
	if err != nil {
		return nil
	}

	for smt.Next() {
		var c Cliente
		var nome sql.NullString
		var dataNascimento sql.NullString
		var id int64

		c.Nome = nome.String
		c.DataNascimento = dataNascimento.String
		c.ID = id
		clientes = append(clientes, c)
	}

	return clientes
}
