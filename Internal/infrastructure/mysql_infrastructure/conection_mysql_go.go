package mysql_infrastructure

import (
	"database/sql"
	"fmt"
)

//username:password@tcp(host:port)/database?charset=utf8
const url = "root:secret@tcp(localhost:9000)/goweb_db"

//Guarda la conexion
var db *sql.DB

//Realizar lac conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection
}

//Cerrar la Conexion
func Close() {
	db.Close()
}

//Verificar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
