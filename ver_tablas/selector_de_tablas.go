package ver_tablas

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
	structs	"trabajo-practico-base-de-datos/datajson"
	"strconv"
)

// 
func VerClientes(db *sql.DB){
    //hace la sentencia en sql
    rows, err := db.Query(`select * from cliente`)
    if err != nil {
		log.Fatal(err)
		}
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

func VerProductos(db *sql.DB){
    //hace la sentencia en sql
    rows, err := db.Query(`select * from producto`)
    if err != nil {
		log.Fatal(err)
		}
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

func VerPedidos(db *sql.DB){
    //hace la sentencia en sql
    rows, err := db.Query(`select * from pedido`)
    if err != nil {
		log.Fatal(err)
		}
    defer rows.Close()
    
    //utilizamos el struct p para guardar los datos
    var p structs.Pedido
    fmt.Printf("A continuacion se mostrara el contenido de la tabla pedido: \n")
    for rows.Next() {
		if err := rows.Scan(&p.Id_pedido, &p.F_pedido, &p.Fecha_entrega, &p.Hora_entrega_desde, &p.Hora_entrega_hasta, &p.Id_usuarie, &p.Id_direccion_entrega, &p.Monto_total, &p.Costo_envio, &p.Estado); err != nil {
			log.Fatal(err)
		}
		var ret string;
		ret = "---------------------------------------------------\n"
		ret = ret + "Pedido ID: " + strconv.Itoa(p.Id_pedido) + "\n"
		ret = ret + "Fecha de pedido: " + p.F_pedido + "\n"
		if p.Fecha_entrega.Valid {
			ret = ret + "Fecha de entrega: " + p.Fecha_entrega.String + "\n"	
		}
		if p.Hora_entrega_desde.Valid { //al llamar al metodo .Valid verifico si el contenido de 
										//p.Hora_entrega_desde es null (false) o tiene un valor (true)
			ret = ret + "Hora entrega desde: " + p.Hora_entrega_desde.String + "\n"	
		}
		if p.Hora_entrega_hasta.Valid {
			ret = ret + "Hora entrega hasta: " + p.Hora_entrega_hasta.String + "\n"	
		}
		ret = ret + "Usuarie ID: " + strconv.Itoa(p.Id_usuarie) + "\n"
		ret = ret + "Direccion de entrega ID: " + strconv.Itoa(p.Id_direccion_entrega) + "\n"
		if p.Monto_total.Valid {
			ret = ret + "Monto total: " + strconv.FormatFloat(p.Monto_total.Float64, 'f', -1, 64) + "\n"	
		}
		ret = ret + "Costo de envio: " + strconv.FormatFloat(p.Costo_envio, 'f', -1, 64) + "\n"
		ret = ret + "Estado: " + p.Estado + "\n"
		
		fmt.Printf(ret)
	}
}

func VerPedidosDetalle(db *sql.DB){
    //hace la sentencia en sql
    rows, err := db.Query(`select * from pedido_detalle`)
    if err != nil {
		log.Fatal(err)
		}
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

func VerReposicion(db *sql.DB){
    //hace la sentencia en sql
    rows, err := db.Query(`select * from reposicion`)
    if err != nil {
		log.Fatal(err)
		}
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

func VerEmails(db *sql.DB){
    //hace la sentencia en sql
    rows, err := db.Query(`select * from envio_email`)
    if err != nil {
		log.Fatal(err)
		}
    defer rows.Close()
    
    //utilizamos el struct e para guardar los datos
    var e structs.Envio_email
    fmt.Printf("A continuacion se mostrara el contenido de la tabla envio_email: \n")
    for rows.Next() {
		if err := rows.Scan(&e.Id_email, &e.F_generacion, &e.Email_cliente, &e.Asunto, &e.Cuerpo, &e.F_envio, &e.Estado); err != nil {
			log.Fatal(err)
		}
		var ret string
		ret = "---------------------------------------------------\n"
		ret = ret + "Email ID: " + strconv.Itoa(e.Id_email) + "\n"
		ret = ret + "Fecha de generacion: " + e.F_generacion + "\n"
		ret = ret + "Email del cliente: " + e.Email_cliente + "\n"
		ret = ret + "Asunto: " + e.Asunto + "\n"
		if e.Cuerpo.Valid {//al llamar al metodo .Valid verifico si el contenido de 
							//p.Cuerpo es null (false) o tiene un valor (true)
			ret = ret + "Cuerpo: " + e.Cuerpo.String + "\n"	
		}
		if e.F_envio.Valid {
			ret = ret + "Fecha de envio: " + e.F_envio.String + "\n"	
		}
		ret = ret + "Estado: " + e.Estado + "\n"
		
		fmt.Printf(ret)
	}
}

func VerErrores(db *sql.DB){
    //hace la sentencia en sql
    rows, err := db.Query(`select * from error`)
    if err != nil {
		log.Fatal(err)
		}
    defer rows.Close()
    
    //utilizamos el struct e para guardar los datos
    var e structs.Error
    fmt.Printf("A continuacion se mostrara el contenido de la tabla error: \n")
    for rows.Next() {
		if err := rows.Scan(&e.Id_error, &e.Id_pedido, &e.F_pedido, &e.Id_usuarie, &e.Id_direccion_entrega, &e.Direccion, &e.Localidad, &e.Codigo_postal, &e.Id_producto, &e.Cantidad, &e.Operacion, &e.F_error, &e.Motivo); err != nil {
			log.Fatal(err)
		}
		var ret string
		ret = "---------------------------------------------------\n"
		ret = ret + "Error ID: " + strconv.Itoa(e.Id_error) + "\n"

		if e.Id_pedido.Valid {//al llamar al metodo .Valid verifico si el contenido de 
							//e.Id_pedido es null (false) o tiene un valor (true)
							//y asi con los demas valores
			ret = ret + "Pedido ID: " + strconv.FormatInt(e.Id_pedido.Int64, 10) + "\n"
		}
		if e.F_pedido.Valid{
			ret = ret + "Fecha de pedido: " + e.F_pedido.String + "\n"
		}
		if e.Id_usuarie.Valid{
			ret = ret + "Usuarie ID: " + strconv.FormatInt(e.Id_usuarie.Int64, 10) + "\n"
		}
		if e.Id_direccion_entrega.Valid {
			ret = ret + "Direccion de entrega ID: " + strconv.FormatInt(e.Id_direccion_entrega.Int64, 10) + "\n"
		}
		if e.Direccion.Valid {
			ret = ret + "Direccion: " + e.F_pedido.String + "\n"
		}
		if e.Localidad.Valid {
			ret = ret + "Localidad: " + e.F_pedido.String + "\n"
		}
		if e.Codigo_postal.Valid {
			ret = ret + "Codigo postal: " + e.Codigo_postal.String + "\n"
		}
		if e.Id_producto.Valid {
			ret = ret + "Producto ID: " + strconv.FormatInt(e.Id_producto.Int64, 10) + "\n"
		}
		if e.Cantidad.Valid {
			ret = ret + "Cantidad: " + strconv.FormatInt(e.Cantidad.Int64, 10) + "\n"
		}
		if e.Operacion.Valid {
			ret = ret + "Operacion: " + e.Operacion.String + "\n"
		}
		if e.F_error.Valid {
			ret = ret + "Fecha del error: " + e.F_error.String + "\n"
		}
		if e.Motivo.Valid {
			ret = ret + "Motivo: " + e.Motivo.String + "\n"
		}
		fmt.Printf(ret)
	}
}
