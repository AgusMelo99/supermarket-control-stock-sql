package BD // los nombres de paquetes van en minúsculas

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
)

func CreateDatabase() {
    db,err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    _, err = db.Exec(`drop database if exists nissero_albarracin_melo_banchero_db1`)
	if err != nil {
		log.Fatal(err)
		}

    _, err = db.Exec(`create database nissero_albarracin_melo_banchero_db1`)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Base de datos creada exitosamente! nombre: nissero_albarracin_melo_banchero_db1\n")
}
