package stored_procedures

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
	"io/ioutil"
)

func SubirStoredProceduresyTriggers(db *sql.DB){ 
	subirStoredProceduresyTriggers("stored-procedures/creacion_del_pedido.sql", db)
	
	subirStoredProceduresyTriggers("stored-procedures/agregado_de_producto.sql", db)
	
	subirStoredProceduresyTriggers("stored-procedures/cierre_del_pedido.sql", db)
	
	subirStoredProceduresyTriggers("stored-procedures/cancelacion_de_pedido.sql", db)
	
	subirStoredProceduresyTriggers("stored-procedures/entrega_pedido.sql", db)
	
	subirStoredProceduresyTriggers("stored-procedures/generar_solicitud_reposicion.sql", db)
	
	subirStoredProceduresyTriggers("stored-procedures/enviar_mail.sql", db)
    
	subirStoredProceduresyTriggers("stored-procedures/enviar_mail_trigger.sql", db)
}

func subirStoredProceduresyTriggers(localizacion string, db *sql.DB) {
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
