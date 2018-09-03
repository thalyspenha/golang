package model

import (
	"database/sql"
	"strings"
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
func (c *Cliente) GetCliente(db *sql.DB) error {
	err := db.QueryRow(`select id, nome, data_nascimento
				from clientes
				 where id = ? `, c.ID).Scan(&c.ID, &c.Nome, &c.DataNascimento)
	if err != nil {
		return err
	}
	return err

}

//GetClientes função get clientes
func (c *Cliente) GetClientes(db *sql.DB) ([]Cliente, error) {
	var values []interface{}
	var where []string

	where = append(where, "id != ? ")
	values = append(values, 0)

	if c.ID != 0 {
		where = append(where, "id = ? ")
		values = append(values, c.ID)
	}

	if c.DataNascimento != "" {
		where = append(where, "data_nascimento = ? ")
		values = append(values, c.DataNascimento)
	}

	// if c.Nome != "" {
	// 	where = append(where, "nome = ? ")
	// 	values = append(values, c.Nome)
	// }

	if c.Nome != "" {
		where = append(where, "nome like ?")
		//values = append(values, fmt.Sprintf("%%%", c.Nome, "%"))
	}

	rows, err := db.Query(`select id, nome, data_nascimento
					  from clientes
					   where `+strings.Join(where, " and "), values...)

	if err != nil {
		return nil, err
	}

	clientes := []Cliente{}
	defer rows.Close()

	for rows.Next() {
		var cl Cliente
		if err := rows.Scan(&cl.ID, &cl.Nome, &cl.DataNascimento); err != nil {
			return nil, err
		}
		clientes = append(clientes, cl)
	}
	return clientes, nil
}

//DeleteCliente ...
func (c *Cliente) DeleteCliente(db *sql.DB) error {
	smt, err := db.Prepare(`delete from clientes where id = ? `)
	if err != nil {
		return err
	}
	_, err = smt.Exec(c.ID)
	return err
}
