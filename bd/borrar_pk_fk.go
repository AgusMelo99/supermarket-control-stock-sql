package bd

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
)

func BorrarFK(db *sql.DB){
	_, err := db.Exec(`alter table direccion_entrega drop constraint direccion_entrega_id_usuarie_fk;
					alter table pedido_detalle drop constraint pedido_detalle_id_pedido_fk;
					alter table pedido_detalle drop constraint pedido_detalle_id_producto_fk;
					alter table reposicion drop constraint reposicion_id_producto_fk`)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("FK's borradas exitosamente \n")
}

func BorrarPK(db *sql.DB){
    _, err := db.Exec(`alter table cliente drop constraint cliente_pk;
					alter table direccion_entrega drop constraint direccion_entrega_pk;
					alter table tarifa_entrega drop constraint tarifa_entrega_pk;
					alter table producto drop constraint producto_pk;
					alter table pedido drop constraint pedido_pk;
					alter table pedido_detalle drop constraint pedido_detalle_pk;
					alter table reposicion drop constraint reposicion_pk;
					alter table error drop constraint error_pk;
					alter table envio_email drop constraint envio_email_pk;`)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("PK's borradas exitosamente \n")
}
