package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/demo_go" // Cambia los credenciales según sea necesario
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("No se pudo establecer una conexión estable:", err)
	}

	log.Println("Conexión a la base de datos exitosa!")
}
