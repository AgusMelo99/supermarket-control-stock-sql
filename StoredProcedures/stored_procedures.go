package storedProcedures

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
)

func Creacion_del_pedido() {
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // llamar al stored procedure que crea el pedido
    var result bool
    
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

func Agregado_de_producto() {
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // llamar al stored procedure que agrega el producto
    var result bool
    
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
