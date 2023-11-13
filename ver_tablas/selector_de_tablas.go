package ver_tablas

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
	structs	"trabajo-practico-base-de-datos/datajson"
)

// 
func VerClientes(){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //hace la sentencia en sql
    rows, err := db.Query(`select * from cliente`)
    defer rows.Close()
    
    //utilizamos el struct c para guardar los datos
    var c structs.Cliente
    fmt.Printf("A continuacion se mostrara el contenido de la tabla cliente: \n")
	for rows.Next() { //por cada fila...
		if err := rows.Scan(&c.Id_usuarie, &c.Nombre, &c.Apellido, &c.Dni, &c.Fecha_nacimiento, &c.Telefono, &c.Email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Usuario ID: %v, Nombre: %v, Apellido: %v, DNI:%v, Fecha de nacimiento:%v, Telefono: %v, Email: %v \n",
					c.Id_usuarie, c.Nombre, c.Apellido, c.Dni, c.Fecha_nacimiento, c.Telefono, c.Email)
		}
}

func VerProductos(){
		db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //hace la sentencia en sql
    rows, err := db.Query(`select * from producto`)
    defer rows.Close()
    
    //utilizamos el struct p para guardar los datos
    var p structs.Producto
    fmt.Printf("A continuacion se mostrara el contenido de la tabla producto: \n")
    for rows.Next() {
		if err := rows.Scan(&p.Id_producto, &p.Nombre, &p.Precio_unitario, &p.Stock_disponible, &p.Stock_reservado, &p.Punto_reposicion, &p.Stock_maximo); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Producto ID: %v, Nombre: %v, Precio unitario: %v, Stock disponible:%v, Stock reservado:%v, Punto de reposicion: %v, Stock maximo: %v \n",
					p.Id_producto, p.Nombre, p.Precio_unitario, p.Stock_disponible, p.Stock_reservado, p.Punto_reposicion, p.Stock_maximo)
	}
}

func VerPedidos(){
		db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //hace la sentencia en sql
    rows, err := db.Query(`select * from pedido`)
    defer rows.Close()
    
    //utilizamos el struct p para guardar los datos
    var p structs.Pedido
    fmt.Printf("A continuacion se mostrara el contenido de la tabla pedido: \n")
    for rows.Next() {
		if err := rows.Scan(&p.Id_pedido, &p.F_pedido, &p.Fecha_entrega, &p.Hora_entrega_desde, &p.Hora_entrega_hasta, &p.Id_usuarie, &p.Id_direccion_entrega, &p.Monto_total, &p.Costo_envio, &p.Estado); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Pedido ID: %v, Fecha del pedido: %v, Fecha de entrega: %v, Desde:%v, Hasta:%v, Id de usuarie: %v, Id direccion de entrega: %v, Monto total: %v, Costo de envio: %v, Estado de envio %v\n",
					p.Id_pedido, p.F_pedido, p.Fecha_entrega, p.Hora_entrega_desde, p.Hora_entrega_hasta, p.Id_usuarie, p.Id_direccion_entrega, p.Monto_total, p.Costo_envio, p.Estado)
	}
}

func VerPedidosDetalle(){
		db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //hace la sentencia en sql
    rows, err := db.Query(`select * from pedido_detalle`)
    defer rows.Close()
    
    //utilizamos el struct p para guardar los datos
    var p structs.Pedido_detalle
    fmt.Printf("A continuacion se mostrara el contenido de la tabla pedido_detalle: \n")
    for rows.Next() {
		if err := rows.Scan(&p.Id_pedido, &p.Id_producto, &p.Cantidad, &p.Precio_unitario); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Pedido ID: %v, Producto ID: %v, Cantidad: %v, Precio unitario: %v\n",
					p.Id_pedido, p.Id_producto, p.Cantidad, p.Precio_unitario)
	}
}

func VerReposicion(){
		db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //hace la sentencia en sql
    rows, err := db.Query(`select * from reposicion`)
    defer rows.Close()
    
    //utilizamos el struct r para guardar los datos
    var r structs.Reposicion
    fmt.Printf("A continuacion se mostrara el contenido de la tabla reposicion: \n")
    for rows.Next() {
		if err := rows.Scan(&r.Id_producto, &r.Fecha_solicitud, &r.Cantidad_a_reponer, &r.Fecha_reposicion, &r.Estado); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Producto ID: %v, Fecha de solicitud: %v, Cantidad a reponer: %v, Fecha de reposicion: %v, Estado %v\n",
					r.Id_producto, r.Fecha_solicitud, r.Cantidad_a_reponer, r.Fecha_reposicion, r.Estado)
	}
}

func VerEmails(){
		db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //hace la sentencia en sql
    rows, err := db.Query(`select * from envio_email`)
    defer rows.Close()
    
    //utilizamos el struct e para guardar los datos
    var e structs.Envio_email
    fmt.Printf("A continuacion se mostrara el contenido de la tabla envio_email: \n")
    for rows.Next() {
		if err := rows.Scan(&e.Id_email, &e.F_generacion, &e.Asunto, &e.Cuerpo, &e.F_envio, &e.Estado); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Email ID: %v, Fecha de generacion: %v, Asunto: %v, Cuerpo: %v, Fecha de envio: %v, Estado: %v\n",
					e.Id_email, e.F_generacion, e.Asunto, e.Cuerpo, e.F_envio, e.Estado)
	}
}

func VerErrores(){
		db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    //hace la sentencia en sql
    rows, err := db.Query(`select * from error`)
    defer rows.Close()
    
    //utilizamos el struct e para guardar los datos
    var e structs.Error
    fmt.Printf("A continuacion se mostrara el contenido de la tabla error: \n")
    for rows.Next() {
		if err := rows.Scan(&e.Id_error, &e.Id_pedido, &e.F_pedido, &e.Id_usuarie, &e.Id_direccion_entrega, &e.Direccion, &e.Localidad, &e.Codigo_postal, &e.Id_producto, &e.Cantidad, &e.Operacion, &e.Operacion, &e.F_error, &e.Motivo); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Error ID: %v, Pedido ID: %v, Fecha de pedido: %v, Usuarie ID: %v, Direccion de entrega ID: %v, Direccion: %v, Localidad : %v, Codigo postal: %v, Producto ID: %v, Cantidad: %v\n",
					e.Id_error, e.Id_pedido, e.F_pedido, e.Id_usuarie, e.Id_direccion_entrega, e.Direccion, e.Localidad, e.Codigo_postal, e.Id_producto, e.Cantidad, e.Operacion, e.Operacion, e.F_error, e.Motivo)
	}
}
