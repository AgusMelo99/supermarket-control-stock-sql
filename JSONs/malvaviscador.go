package JSONs

import (
    "encoding/json"
    "database/sql"
    "fmt"
    "log"
    "io/ioutil"
)
	
//esto seguramente se puede hacer con funciones y debe quedar mas bonito de ver y leer, pero no quiero hacer cagada
func Unmalvaviscador(){
	
	//Unmarshall clientes.json
	var clientes []Cliente
	jsonFile, err := ioutil.ReadFile("JSONs/clientes.json")
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
	
	db,err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
        if err != nil {
        log.Fatal(err)
    }
	
	//Inserta los registros en la tabla cliente
	if clientes != nil{
		for _, c := range clientes {
				
				insert := fmt.Sprintf("insert into cliente values (%v, '%v', '%v', %v, '%v', '%v', '%v');" , c.Id_usuarie, c.Nombre, c.Apellido, c.Dni, c.Fecha_nacimiento, c.Telefono, c.Email)
			    _, err = db.Exec(insert)
				
				if err != nil {
					log.Fatal(err)
				}
		}
	}
	
	//Unmarshall direcciones.json
	var direcciones [] Direccion_entrega
	jsonFile, err = ioutil.ReadFile("JSONs/direcciones.json")
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
	
	//Inserta los registros en la tabla direccion_entrega
	if direcciones != nil{
		for _, d := range direcciones {
				
				insert := fmt.Sprintf("insert into direccion_entrega values (%v, %v, '%v', '%v', '%v');" , d.Id_usuarie, d.Id_direccion_entrega, d.Direccion, d.Localidad, d.Codigo_postal)
			    _, err = db.Exec(insert)
				
				if err != nil {
					log.Fatal(err)
				}
		}
	}
    
    //Unmarshall productos.json
    var productos [] Producto
    jsonFile, err = ioutil.ReadFile("JSONs/productos.json")
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
	
	//Inserta los registros en la tabla producto
	if productos != nil{
		for _, p := range productos {
				insert := fmt.Sprintf("insert into producto values (%v,'%v', %v , %v , %v ,%v, %v );" , p.Id_producto, p.Nombre, p.Precio_unitario, p.Stock_disponible, p.Stock_reservado, p.Punto_reposicion,p.Stock_maximo)
			    _, err = db.Exec(insert)
				
				if err != nil {
					log.Fatal(err)
				}
		}
	}
	
	//Unmarshall tarifas_entrega.json
	var tarifas_entrega [] Tarifa_entrega
	
	jsonFile, err = ioutil.ReadFile("JSONs/tarifas_entrega.json")
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
	
	//Inserta los registros en la tabla tarifa_entrega
	if tarifas_entrega != nil{
		for _, t := range tarifas_entrega {
				insert := fmt.Sprintf("insert into tarifa_entrega values ('%v',%v );" , t.Codigo_postal_corto, t.Costo) 
			    _, err = db.Exec(insert)
				
				if err != nil {
					log.Fatal(err)
				}
		}
	}
}
