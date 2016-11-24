package model

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"servidor/domain"
)

var db *sql.DB = AbrirDB()

func AbrirDB() *sql.DB{
	var db, err = sql.Open("mysql", "root:root@/alarmaSys")
	if(err != nil){
		panic(err.Error())
	}
	return db
}

func AddCliente(c domain.Cliente) {
	_,err := db.Exec("INSERT INTO Cliente (nombre, apellido,categoria,fechaNacimiento) VALUES (?, ?,?,?)",
					c.Nombre,c.Apellido,c.Categoria,c.FechaNacimiento)
	if err != nil{
		panic(err.Error())
	}
}

func GetClientes() []domain.Cliente{
	rows, err := db.Query("select * from Cliente")

	if err != nil{
		panic(err.Error())
	}

	clientes := []domain.Cliente{}

	for rows.Next() {
		var cliente domain.Cliente
		err = rows.Scan(&cliente.Id, &cliente.Nombre,&cliente.Apellido,&cliente.Categoria,&cliente.FechaNacimiento)
		//fmt.Println(cliente)
		clientes = append(clientes,cliente)
		if err != nil {
			panic(err.Error())
		}
	}
	return clientes
}

func GetCliente(id string) domain.Cliente{
	rows, err := db.Query("select * from Cliente where id = ?",id)

	if err != nil{
		panic(err.Error())
	}

	cliente := domain.Cliente{}

	for rows.Next() {
		err = rows.Scan(&cliente.Id, &cliente.Nombre,&cliente.Apellido,&cliente.Categoria,&cliente.FechaNacimiento)
		if err != nil {
			panic(err.Error())
		}
	}
	return cliente
}

func MostrarClientes(){
	rows, err := db.Query("select * from Cliente")

	if err != nil{
		panic(err.Error())
	}

	clientes := []domain.Cliente{}
	columns, err := rows.Columns()

	if err != nil{
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns)) // CREA UN SLICE PARA LOS VALORES

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	var cliente domain.Cliente
	for rows.Next() {	//mientras que haya siguiente
		// get RawBytes from data
		err = rows.Scan(&cliente.Id, &cliente.Nombre,&cliente.Apellido,&cliente.Categoria,&cliente.FechaNacimiento)
		clientes = append(clientes,cliente)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	fmt.Println("Clientes: \n")
	for _, clienteAux := range clientes {

		fmt.Println(clienteAux, "\n")
	}
}