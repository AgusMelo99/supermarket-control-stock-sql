package JSONs

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
)

//esto seguramente se puede hacer con funciones y debe quedar mas bonito de ver y leer, pero no quiero hacer cagada
func Unmalvaviscador(){
	//Unmarshall clientes.json
	jsonFile, err := ioutil.ReadFile("JSONs/clientes.json")
    if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
	var clientes [] Cliente
	err = json.Unmarshal(jsonFile, &clientes)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println("clientes.json leido correctamente \n")
	fmt.Println("Contenido de clientes.json: \n")
	fmt.Printf("%v\n", clientes)
	
	//Unmarshall direcciones.json
	jsonFile, err = ioutil.ReadFile("JSONs/direcciones.json")
	if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
    var direcciones [] Direccion_entrega
    err = json.Unmarshal(jsonFile, &direcciones)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println("direcciones.json leido correctamente \n")
	fmt.Println("Contenido de direcciones.json: \n")
	fmt.Printf("%v\n", direcciones)
    
    //Unmarshall productos.json
    jsonFile, err = ioutil.ReadFile("JSONs/productos.json")
	if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
    var productos [] Producto
    err = json.Unmarshal(jsonFile, &productos)
	if err != nil { 
		fmt.Println("error:", err)
	}
	
	fmt.Println("productos.json leido correctamente \n")
	fmt.Println("Contenido de productos.json: \n")
	fmt.Printf("%v\n", productos)
	
	//Unmarshall tarifas_entrega.json
	jsonFile, err = ioutil.ReadFile("JSONs/tarifas_entrega.json")
	if err != nil {
        log.Fatalf("Error al leer el archivo JSON: %s", err)
    }
    var tarifas_entrega [] Tarifa_entrega
    err = json.Unmarshal(jsonFile, &tarifas_entrega)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println("tarifas_entrega.json leido correctamente \n")
	fmt.Println("Contenido de tarifas_entrega.json: \n")
	fmt.Printf("%v\n", tarifas_entrega)
}
