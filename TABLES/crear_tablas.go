package create_table

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
)

func Create_tables(){ // fixme
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    _,err = db.Exec(`create table cliente(id_usuarie int, nombre text, apellido text, dni int, fecha_nacimento date, telefono char(12), email text);`)
    if err != nil {
        log.Fatal(err)
    }
    
    
    _,err = db.Exec(`create table direccion_entrega(id_usuarie int, id_direccion_entrega int, direccion varchar(60), localidad varchar(30), codigo_postal char(8));`)
    if err != nil {
        log.Fatal(err)
    }
    
    _,err = db.Exec(`create table tarifa_entrega(codigo_postal_corto char(4), costo decimal(12,2));`)
    if err != nil {
        log.Fatal(err)
    }
    
    _,err = db.Exec(`create table producto(id_producto int, nombre text, precio_unitario decimal(12,2), stock_disponible int, stock_reservado int, punto_reposicion int, stock_maximo int);`)
    if err != nil {
        log.Fatal(err)
    }
    
    _,err = db.Exec(`create table pedido(id_pedido int, f_pedido timestamp, fecha_entrega date, hora_entrega_desde time, hora_entrega_hasta time, id_usuarie int, id_direccion_entrega int, monto_total decimal(12,2), costo_envio decimal(12,2), estado char(10) );`)
    if err != nil {
        log.Fatal(err)
    }
    
    _,err = db.Exec(`create table pedido_detalle(id_pedido int, id_producto int, cantidad int, precio_unitario decimal(12,2));`)
    if err != nil {
        log.Fatal(err)
    }
    
    _,err = db.Exec(`create table reposicion(id_producto int, fecha_solicitud date, cantidad_a_reponer int, fecha_reposicion date, estado char(12));`)
    if err != nil {
        log.Fatal(err)
    }
    
    _,err = db.Exec(`create table error(id_error int, id_pedido int, f_pedido timestamp, id_usuarie int, id_direccion_entrega int, direccion varchar(60), localidad varchar(30), codigo_postal char(8), id_producto int, cantidad int, operacion char(12), f_error timestamp,motivo varchar(64));`)
    if err != nil {
        log.Fatal(err)
    }
    
    _,err = db.Exec(`create table envio_email(id_email int, f_generacion timestamp, email_cliente text, asunto text, cuerpo text, f_envio timestamp, estado char(10));`)
    if err != nil {
        log.Fatal(err)
    }
    _,err = db.Exec(`create table entrada_trx_pedido(id_orden int, operacion char(12), id_usuarie int, id_direccion_entrega int, id_pedido int, id_producto int, cantidad int, fecha_hora_entrega timestamp);`)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Las siguientes tablas fueron creadas con exito: \n-cliente \n-direccion_entrega \n-tarifa_entrega \n-producto \n-pedido_detalle \n-reposicion \n-error \n-envio_mail \n-entrada_trx_pedido")
}
