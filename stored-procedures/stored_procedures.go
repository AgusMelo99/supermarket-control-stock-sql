package stored_procedures

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
	"io/ioutil"
)

func SubirSPyTriggers(){
	subirSPyTriggers("stored-procedures/creacion_del_pedido.sql")

	subirSPyTriggers("stored-procedures/agregado_de_producto.sql")
	
	subirSPyTriggers("stored-procedures/enviar_mail.sql")
    
	subirSPyTriggers("stored-procedures/enviar_mail_trigger.sql")
}

func subirSPyTriggers(localizacion string) {
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //lee el contenido del archivo sql y lo guarda en el tipo de datos []byte
    contenidoSql, err := ioutil.ReadFile(localizacion)
    if err != nil {
        log.Fatal(err)
    }
    
	//Convierte el contenido del archivo .sql en un string para despues ser ejecutado
	contenidoSqlEnString := string(contenidoSql)
    _, err = db.Exec(contenidoSqlEnString)
    if err != nil {
        log.Fatal(err)
    }
}

func CreacionDelPedido() {
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //lee el contenido del archivo sql y lo guarda en el tipo de datos []byte
    contenidoSql, err := ioutil.ReadFile("stored-procedures/creacion_del_pedido.sql")
    if err != nil {
		log.Fatal(err)
	}
	
	//Convierte el contenido del archivo .sql en un string para despues ser ejecutado
	contenidoSqlEnString := string(contenidoSql)
    _, err = db.Exec(contenidoSqlEnString)
    if err != nil {
		log.Fatal(err)
	}
   
    // iniciar la transaccion serializable
    tx, err := db.Begin()
    if err != nil {
		log.Fatal(err)
	}
    
    // variable para almacenar el resultado de la funcion
    var result bool
    
    id_cliente := 1; //VALOR NO VALIDO (PRUEBA), HAY QUE LEER DE LOS SLICES
    id_direccion := 2; //VALOR NO VALIDO (PRUEBA), HAY QUE LEER DE LOS SLICES
    // llamar al stored procedure que crea el pedido
    // falta agregar aca como tomar el id del cliente y de la direccion
    
    /* agrego for update al final del select para bloquear todas filas involucradas en la funcion, evitando que otra transaccion se ejecute 
     * al mismo tiempo. Y cambio el db por tx para que utilice la transaccion*/
    query := fmt.Sprintf(`select creacion_del_pedido(%v, %v) for update;`, id_cliente, id_direccion)
    err = tx.QueryRow(query).Scan(&result)
    
    if err != nil {
		// En caso de error deshacer la transaccion
		tx.Rollback()
		log.Fatal(err)
		fmt.Println("No se pudo crear el pedido.")
	}
	
	if result {
		// Confirmar la transaccion
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Pedido creado con exito.")
	}
}

func AgregadoDeProducto() {
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //lee el contenido del archivo sql y lo guarda en el tipo de datos []byte
    contenidoSql, err := ioutil.ReadFile("stored-procedures/agregado_de_producto.sql")
    if err != nil {
		log.Fatal(err)
	}
	
	//Convierte el contenido del archivo .sql en un string para despues ser ejecutado
	contenidoSqlEnString := string(contenidoSql)
    _, err = db.Exec(contenidoSqlEnString)
    if err != nil {
		log.Fatal(err)
	}
    
    // iniciar la transaccion serializable
    tx, err := db.Begin()
    if err != nil {
		log.Fatal(err)
	}
    
    // variable para almacenar el resultado de la funcion
    var result bool
    
    pedido_id := 1 //VALOR NO VALIDO (PRUEBA), HAY QUE LEER DE LOS SLICES
    producto_id := 2 //VALOR NO VALIDO (PRUEBA), HAY QUE LEER DE LOS SLICES
    cantidad_producto_id := 3; //VALOR NO VALIDO (PRUEBA), HAY QUE LEER DE LOS SLICES
    
    // llamar al stored procedure que agrega el producto
    // falta agregar aca como tomar el id del pedido, producto y cantidad
    
    // agrego otro for update y tambien cambio db por tx

    query := fmt.Sprintf(`select agregado_de_producto(%v, %v, %v) for update`, pedido_id, producto_id, cantidad_producto_id)
    err = tx.QueryRow(query).Scan(&result)
    
    if err != nil {
		// En caso de error deshacer la transaccion
		tx.Rollback()
		log.Fatal(err)
		fmt.Println("No se pudo crear el pedido.")
	}
	
	if result {
		// Confirmar la transaccion
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Pedido creado con exito.")
	}
}

func EnvioDeMail(){
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //lee el contenido del archivo sql y lo guarda en el tipo de datos []byte
    contenidoSql, err := ioutil.ReadFile("stored-procedures/enviar_mail_trigger.sql")
    if err != nil {
		log.Fatal(err)
	}
	
	//Convierte el contenido del archivo .sql en un string para despues ser ejecutado
	contenidoSqlEnString := string(contenidoSql)
    _, err = db.Exec(contenidoSqlEnString)
    if err != nil {
		log.Fatal(err)
	}
}
