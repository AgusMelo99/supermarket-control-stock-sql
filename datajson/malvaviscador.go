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

func InsertarClientes(){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
	
	//Inserta los registros en la tabla cliente
	if clientes != nil{
		for _, c := range clientes {
				
				insert := fmt.Sprintf(`insert into cliente values (%v, '%v', '%v', %v, '%v', '%v', '%v');` , c.Id_usuarie, c.Nombre, c.Apellido, c.Dni, c.Fecha_nacimiento, c.Telefono, c.Email)
			    _, err = db.Exec(insert)
					
				if err != nil {
					log.Fatal("Error al insertar cliente \n", err)
				}
		}
		fmt.Println("Clientes insertados correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON Clientes \n")
	}	
}

func InsertarProductos(){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
	if productos != nil{
		for _, p := range productos {
				insert := fmt.Sprintf(`insert into producto values (%v,'%v', %v , %v , %v ,%v, %v );` , p.Id_producto, p.Nombre, p.Precio_unitario, p.Stock_disponible, p.Stock_reservado, p.Punto_reposicion,p.Stock_maximo)
			    _, err = db.Exec(insert)
				
				if err != nil {
					log.Fatal("Error al insertar producto \n", err)
				}
		}
		fmt.Println("Productos insertados correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON Productos \n")
	}	
}

func InsertarTarifaEntrega(){
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()	
	if tarifas_entrega != nil{
		for _, t := range tarifas_entrega {
				insert := fmt.Sprintf(`insert into tarifa_entrega values ('%v',%v );` , t.Codigo_postal_corto, t.Costo) 
			    _, err = db.Exec(insert)
				
				if err != nil {
					log.Fatal(err)
				}
		}
		fmt.Println("Tarifas de entrega insertadas correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON tarifa_entrega \n")
	}	
}

func InsertarDireccionesEntrega(){
	//Inserta los registros en la tabla direccion_entrega
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
	if direcciones != nil{
		for _, d := range direcciones {
				
				insert := fmt.Sprintf(`insert into direccion_entrega values (%v, %v, '%v', '%v', '%v');` , d.Id_usuarie, d.Id_direccion_entrega, d.Direccion, d.Localidad, d.Codigo_postal)
			    _, err = db.Exec(insert)
				
				if err != nil {
					log.Fatal("Error al insertar direccion \n", err)
				}
		}
		fmt.Println("Direcciones insertadas correctamente \n")
	}else{
		fmt.Println("Todavia no leiste el JSON direcciones_entrega \n")
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
	fmt.Println("Contenido de clientes.json: \n")
	//fmt.Printf("%v\n", clientes)
	
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
	fmt.Println("Contenido de direcciones.json: \n")
	//fmt.Printf("%v\n", direcciones)
    
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
	fmt.Println("Contenido de productos.json: \n")
	//fmt.Printf("%v\n", productos)
	
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
	fmt.Println("Contenido de tarifas_entrega.json: \n")
	//fmt.Printf("%v\n", tarifas_entrega)
}
