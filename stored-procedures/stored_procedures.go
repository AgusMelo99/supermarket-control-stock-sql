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
    
    // variable para almacenar el resultado de la funcion
    var result bool
    
    // llamar al stored procedure que crea el pedido
    // falta agregar aca como tomar el id del cliente y de la direccion
    err = db.QueryRow("select creacion_del_pedido(%v, %v), id_cliente, id_direccion").Scan(%result)
    
    if err != nil {
		fmt.Println("Error", err)
	}
	
	if result {
		fmt.Println("Pedido creado con exito.")
	} else {
		fmt.Println("Creacion del pedido fallida")
	}
}

func AgregadoDeProducto() {
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // variable para almacenar el resultado de la funcion
    var result bool
    
    // llamar al stored procedure que agrega el producto
    // falta agregar aca como tomar el id del pedido, producto y cantidad
    err = db.QueryRow("select agregado_de_producto(%v, %v, %v), id_pedido, id_producto, cantidad_producto").Scan(%result)
    
    if err != nil {
		fmt.Println("Error", err)
	}
	
	if result {
		fmt.Println("Producto agregado con exito.")
	} else {
		fmt.Println("El producto no pudo ser agregado.")
	}
}
