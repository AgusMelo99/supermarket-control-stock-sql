package stored_procedures

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
)

func CreacionDelPedido() {
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // iniciar la transaccion serializable
    tx, err := db.Begin()
    if err != nil {
		log.Fatal(err)
	}
    
    // variable para almacenar el resultado de la funcion
    var result bool
    
    // llamar al stored procedure que crea el pedido
    // falta agregar aca como tomar el id del cliente y de la direccion
    
    /* agrego for update al final del select para bloquear todas filas involucradas en la funcion, evitando que otra transaccion se ejecute 
     * al mismo tiempo. Y cambio el db por tx para que utilice la transaccion*/
    err = tx.QueryRow("select creacion_del_pedido(%v, %v), id_cliente, id_direccion for update").Scan(&result)
    
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
    
    // iniciar la transaccion serializable
    tx, err := db.Begin()
    if err != nil {
		log.Fatal(err)
	}
    
    // variable para almacenar el resultado de la funcion
    var result bool
    
    // llamar al stored procedure que agrega el producto
    // falta agregar aca como tomar el id del pedido, producto y cantidad
    
    // agrego otro for update y tambien cambio db por tx
    err = tx.QueryRow("select agregado_de_producto(%v, %v, %v), id_pedido, id_producto, cantidad_producto for update").Scan(&result)
    
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
