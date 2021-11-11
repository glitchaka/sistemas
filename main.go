package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "admin"
	Contrasenia := "1234"
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion

}

var plantillas = template.Must(template.ParseGlob("plantilla/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	log.Println("El servidor esta corriendo en el puerto 8090")
	http.ListenAndServe(":8090", nil)
}
func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionDB()
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre, correo) VALUES('JuliANO', 'tuanno@gemail.kam')")
	if err != nil {
		panic(err.Error())
	}
	insertarRegistros.Exec()
	//fmt.Fprintf(w, "Hola mundo")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	plantillas.ExecuteTemplate(w, "crear", nil)
}
