package main

import (
	"fmt"
	"log"
	bolt "go.etcd.io/bbolt"
	funciones "trabajo-practico-base-de-datos/bolt-db"
	"strconv"
	"encoding/json"
)

type Cliente struct {
	Id_usuarie int
	Nombre string
	Apellido string
	Dni int
	Fecha_nacimiento string
	Telefono string
	Email string
}

type Direccion_entrega struct {
	Id_usuarie int
	Id_direccion_entrega int
	Direccion string
	Localidad string
	Codigo_postal string
}

type Producto struct {
	Id_producto int
	Nombre string
	Precio_unitario int
	Stock_disponible int
	Stock_reservado int
	Punto_reposicion int
	Stock_maximo int
}

type Pedido struct {
	Id_pedido int
	F_pedido string
	Fecha_entrega string
	Hora_entrega_desde string
	Hora_entrega_hasta string
	Id_usuarie int
	Id_direccion_entrega int
	Monto_total int
	Costo_envio int
	Estado string
}

func main() {
	db, err := bolt.Open("nissero_albarracin_melo_banchero_db1", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer.db.Close()
	
	// aca se crean los datos
	cliente_uno := Cliente{1, "Juan", "Perez", 4082390123, "08-08-2000", "1120975982", "juan_perez@gmail.com"}
	cliente_dos := Cliente{2, "Claudia", "Gonzalez", 41672309, "02-01-2002", "117416082", "claudia_gonzalez@gmail.com"}
	
	data_c_uno, err := json.Marshal(cliente_uno)
	if err != nil {
		log.Fatal(err)
	}
	
	data_c_dos, err := json.Marshal(cliente_dos)
	if err != nil {
		log.Fatal(err)
	}
	
	direccion_cliente_uno := Direccion_entrega{1, 1, "Calle falsa 123", "Sprinfield", "2394"}
	direccion_cliente_dos := Direccion_entrega{2, 2, "Avenida Siempre Viva 742", "Sprinfield", "2394"}
	
	data_dc_uno, err := json.Marshal(direccion_cliente_uno)
	if err != nil {
		log.Fatal(err)
	}
	
	data_dc_dos, err := json.Marshal(direccion_cliente_dos)
	if err != nil {
		log.Fatal(err)
	}
	
	producto_uno := Producto{1, "Pan lactal", 450, 20, 10, 5, 30}
	producto_dos := Producto{2, "Oreo", 600, 50, 20, 15, 70}
	producto_tres := Producto{3, "Durazno en lata", 200, 10, 5, 2, 15}
	producto_cuatro := Producto{4, "Coca-Cola", 1200, 25, 5, 8, 30}
	producto_cinco := Producto{5, "Agua Villavicencio", 350, 15, 5, 4, 20}
	producto_seis := Producto{6, "Nutella", 900, 10, 2, 3, 12}
	
	data_pr_uno, err := json.Marshal(producto_uno)
	if err != nil {
		log.Fatal(err)
	}
	
	data_pr_dos, err := json.Marshal(producto_dos)
	if err != nil {
		log.Fatal(err)
	}
	
	data_pr_tres, err := json.Marshal(producto_tres)
	if err != nil {
		log.Fatal(err)
	}
	
	data_pr_cuatro, err := json.Marshal(producto_cuatro)
	if err != nil {
		log.Fatal(err)
	}
	
	data_pr_cinco, err := json.Marshal(producto_cinco)
	if err != nil {
		log.Fatal(err)
	}
	
	data_pr_seis, err := json.Marshal(producto_seis)
	if err != nil {
		log.Fatal(err)
	}
	
	pedido_uno := Pedido{1, "14-11-2023", "15-11-2023", "12:00", "14:00", 1, 1, 2000, 200, "estado"}
	pedido_dos := Pedido{2, "14-11-2023", "15-11-2023", "15:00", "17:00", 2, 2, 1700, 150, "estado"}
	
	data_pe_uno, err := json.Marshal(pedido_uno)
	if err != nil {
		log.Fatal(err)
	}
	
	data_pe_dos, err := json.Marshal(pedido_dos)
	if err != nil {
		log.Fatal(err)
	}
	
	//agregar clientes al respectivo bucket
	funciones.CrearBucket(db, "cliente", []byte(strconv.Itoa(cliente_uno.Id_usuarie)), data_c_uno)
	funciones.CrearBucket(db, "cliente", []byte(strconv.Itoa(cliente_dos.Id_usuarie)), data_c_dos)
	
	//agregar direcciones de entrega al respectivo bucket
	funciones.CrearBucket(db, "direccion_entrega", []byte(strconv.Itoa(direccion_cliente_uno.Id_direccion_entrega)), data_dc_uno)
	funciones.CrearBucket(db, "direccion_entrega", []byte(strconv.Itoa(direccion_cliente_dos.Id_direccion_entrega)), data_dc_dos)
	
	//agregar productos al respectivo bucket
	funciones.CrearBucket(db, "producto", []byte(strconv.Itoa(producto_uno.Id_producto)), data_pr_uno)
	funciones.CrearBucket(db, "producto", []byte(strconv.Itoa(producto_dos.Id_producto)), data_pr_dos)
	funciones.CrearBucket(db, "producto", []byte(strconv.Itoa(producto_tres.Id_producto)), data_pr_tres)
	funciones.CrearBucket(db, "producto", []byte(strconv.Itoa(producto_cuatro.Id_producto)), data_pr_cuatro)
	funciones.CrearBucket(db, "producto", []byte(strconv.Itoa(producto_cinco.Id_producto)), data_pr_cinco)
	funciones.CrearBucket(db, "producto", []byte(strconv.Itoa(producto_seis.Id_producto)), data_pr_seis)
	
	//agregar pedidos al respectivo bucket
	funciones.CrearBucket(db, "pedido", []byte(strconv.Itoa(pedido_uno.Id_pedido)), data_pe_uno)
	funciones.CrearBucket(db, "pedido", []byte(strconv.Itoa(pedido_dos.Id_pedido)), data_pe_dos)
	
	//leer los buckets pedido
	p_uno, err := funcoines.LeerBucket(db, "pedido", []byte(strconv.Itoa(pedido_uno.Id_usuarie)))
	p_dos, err := funcoines.LeerBucket(db, "pedido", []byte(strconv.Itoa(pedido_dos.Id_usuarie)))
}
