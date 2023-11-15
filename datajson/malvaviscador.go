package datajson // datajson

import (
    "encoding/json"
    "database/sql"
    "fmt"
    "log"
    "io/ioutil"
)

var clientes [] Cliente
var productos [] Producto
var direcciones [] Direccion_entrega
var tarifas_entrega [] Tarifa_entrega
var entrada_trx_pedidos [] Entrada_trx_pedido

func InsertarClientes(db *sql.DB){
	//Inserta los registros en la tabla cliente
	if clientes != nil{
		for _, c := range clientes {
				
				insert := fmt.Sprintf(`insert into cliente values (%v, '%v', '%v', %v, '%v', '%v', '%v');` , c.Id_usuarie, c.Nombre, c.Apellido, c.Dni, c.Fecha_nacimiento, c.Telefono, c.Email)
			    _, err := db.Exec(insert)
					
				if err != nil {
					log.Fatal("Error al insertar cliente \n", err)
				}
		}
		fmt.Println("Clientes insertados correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON Clientes \n")
	}	
}

func InsertarProductos(db *sql.DB){
	if productos != nil{
		for _, p := range productos {
				insert := fmt.Sprintf(`insert into producto values (%v,'%v', %v , %v , %v ,%v, %v );` , p.Id_producto, p.Nombre, p.Precio_unitario, p.Stock_disponible, p.Stock_reservado, p.Punto_reposicion,p.Stock_maximo)
			    _, err := db.Exec(insert)
				
				if err != nil {
					log.Fatal("Error al insertar producto \n", err)
				}
		}
		fmt.Println("Productos insertados correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON Productos \n")
	}	
}

func InsertarTarifaEntrega(db *sql.DB){
	if tarifas_entrega != nil{
		for _, t := range tarifas_entrega {
				insert := fmt.Sprintf(`insert into tarifa_entrega values ('%v',%v );` , t.Codigo_postal_corto, t.Costo) 
			    _, err := db.Exec(insert)
				
				if err != nil {
					log.Fatal(err)
				}
		}
		fmt.Println("Tarifas de entrega insertadas correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON tarifa_entrega \n")
	}	
}

func InsertarDireccionesEntrega(db *sql.DB){
	if direcciones != nil{
		for _, d := range direcciones {
				
				insert := fmt.Sprintf(`insert into direccion_entrega values (%v, %v, '%v', '%v', '%v');` , d.Id_usuarie, d.Id_direccion_entrega, d.Direccion, d.Localidad, d.Codigo_postal)
			    _, err := db.Exec(insert)
				
				if err != nil {
					log.Fatal("Error al insertar direccion \n", err)
				}
		}
		fmt.Println("Direcciones insertadas correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON direcciones_entrega \n")
	}		
}

func InsertarOperaciones(db *sql.DB){
	if entrada_trx_pedidos != nil{
		var resultado bool
		for _, d := range entrada_trx_pedidos {
				fmt.Printf("RESULTADO ORDEN ID %v \n", d.Id_orden)
				if d.Operacion == "creación" {
					insert := fmt.Sprintf(`select * from creacion_del_pedido(%v, %v)`, d.Id_usuarie, d.Id_direccion_entrega)
					err := db.QueryRow(insert).Scan(&resultado)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(resultado)
				} else if d.Operacion == "producto" {
					insert := fmt.Sprintf(`begin; 
											set transaction isolation level serializable;
											select * from agregado_de_producto(%v, %v, %v); 
											commit;`, d.Id_pedido, d.Id_producto, d.Cantidad)
					err := db.QueryRow(insert).Scan(&resultado)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(resultado)
				} else if d.Operacion == "cierre" {
					insert := fmt.Sprintf(`begin; 
											set transaction isolation level serializable;
											select * from cierre_del_pedido(%v, '%v'); 
											commit;`, d.Id_pedido, d.Fecha_hora_entrega)
					err := db.QueryRow(insert).Scan(&resultado)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(resultado)
				} else if d.Operacion == "entrega" {
					insert := fmt.Sprintf(`begin; 
											set transaction isolation level serializable;
											select * from entrega_pedido(%v) 
											commit;`, d.Id_pedido)
					err := db.QueryRow(insert).Scan(&resultado)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(resultado)
				} else if d.Operacion == "cancelación" {
					insert := fmt.Sprintf(`begin; 
											set transaction isolation level serializable;
											select * from cancelacion_de_pedido(%v) 
											commit;`, d.Id_pedido)
					err := db.QueryRow(insert).Scan(&resultado)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(resultado)
				}
		}
		fmt.Println("Direcciones insertadas correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON trx_pedidos \n")
	}		
}

func GenerarSolicitudesDeReposicion(db *sql.DB){
	var ret bool
	insert := fmt.Sprintf(`select * from generar_solicitud_reposicion()`)
	err := db.QueryRow(insert).Scan(&ret)
	if err != nil {
		log.Fatal(err)
	}
	if ret {
		fmt.Println("Se han generado solicitudes de reposicion \n")
	} else {
		fmt.Println("No se han generado solicitudes de reposicion debido a que ya fueron generadas el dia de hoy \n")
		}
}


	
//esto seguramente se puede hacer con funciones y debe quedar mas bonito de ver y leer, pero no quiero hacer cagada
func Unmalvaviscador(){
	
	//Unmarshall clientes.json
	jsonFile, err := ioutil.ReadFile("datajson/clientes.json")
    if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
	
	err = json.Unmarshal(jsonFile, &clientes)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println("clientes.json leido correctamente \n")
	
	//Unmarshall direcciones.json
	jsonFile, err = ioutil.ReadFile("datajson/direcciones.json")
	if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
    
    err = json.Unmarshal(jsonFile, &direcciones)
	if err != nil {
		fmt.Println("error:", err) 
	}
	
	fmt.Println("direcciones.json leido correctamente \n")
    
    //Unmarshall productos.json
    jsonFile, err = ioutil.ReadFile("datajson/productos.json")
	if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
    
    err = json.Unmarshal(jsonFile, &productos)
	if err != nil { 
		fmt.Println("error:", err)
	}
	
	fmt.Println("productos.json leido correctamente \n")
	
	//Unmarshall tarifas_entrega.json
	jsonFile, err = ioutil.ReadFile("datajson/tarifas_entrega.json")
	if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
    
    err = json.Unmarshal(jsonFile, &tarifas_entrega)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println("tarifas_entrega.json leido correctamente \n")
	//fmt.Printf("%v\n", tarifas_entrega)
	
	jsonFile, err = ioutil.ReadFile("datajson/trx_pedido.json")
	if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
    
	err = json.Unmarshal(jsonFile, &entrada_trx_pedidos)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println("trx_pedidos.json leido correctamente \n")
}
