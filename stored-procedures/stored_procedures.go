package stored_procedures

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
	"io/ioutil"
)

func SubirStoredProceduresyTriggers(){ 
	subirStoredProceduresyTriggers("stored-procedures/creacion_del_pedido.sql")
	
	subirStoredProceduresyTriggers("stored-procedures/agregado_de_producto.sql")
	
	subirStoredProceduresyTriggers("stored-procedures/enviar_mail.sql")
    
	subirStoredProceduresyTriggers("stored-procedures/enviar_mail_trigger.sql")
	
	subirStoredProceduresyTriggers("stored-procedures/entrega_pedido.sql")
	
	subirStoredProceduresyTriggers("stored-procedures/cierre_del_pedido.sql")
	
	subirStoredProceduresyTriggers("stored-procedures/cancelacion_de_pedido.sql")
}

func subirStoredProceduresyTriggers(localizacion string) {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //lee el contenido del archivo sql y lo guarda en el tipo de datos []byte
    contenidoSql, err := ioutil.ReadFile(localizacion)
    if err != nil {
        fmt.Printf("Error al leer el archivo %v, \n %v", localizacion, err)
        log.Fatal()
    }
    
	//Convierte el contenido del archivo .sql en un string para despues ser ejecutado
	contenidoSqlEnString := string(contenidoSql)
    _, err = db.Exec(contenidoSqlEnString)
    if err != nil {
        fmt.Printf("Error al ejecutar el sp/trigger %v, \n %v", localizacion, err)
        log.Fatal()
    }
}
