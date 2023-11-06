package main

import (
//	"database/sql"
	"fmt"
	_"github.com/lib/pq"
//	"log"
	database 			"trabajo-practico-base-de-datos/BD"
	pk_fk	 			"trabajo-practico-base-de-datos/PKyFK"
	tables	 			"trabajo-practico-base-de-datos/TABLES"
	datos	 			"trabajo-practico-base-de-datos/JSONs"
	stored_procedures	"trabajo-practico-base-de-datos/StoredProcedures"
)

func printMenu() {
	fmt.Println("1. Crear base de datos \n")
	fmt.Println("2. Crear Tablas \n")
	fmt.Println("3. Aniadir PK's y FK's \n")
	fmt.Println("4. Eliminar PK's y FK's \n")
	fmt.Println("5. Agregar SP's y Triggers \n")
	fmt.Println("6. Unmarshalear JSON's y subir los datos \n")
	fmt.Println("0. Cerrar sesion\n")
	}

func continuar (opcion int){
	if opcion == 1 {
		fmt.Println("Elegiste la opcion 1, crear base de datos \n")
		database.CreateDatabase()
			} else if opcion == 2 {
				fmt.Println("Elegiste la opcion 2, se crearan las tablas correspondientes \n")
				tables.Create_tables()
			} else if opcion == 3 {
				fmt.Println("Elegiste la opcion 3, se crearan las PK's y FK's correspondientes \n")
				pk_fk.CreatePK()
				pk_fk.CreateFK()
			} else if opcion == 4 {
				fmt.Println("Elegiste la opcion 4, se eliminaran las PK's y FK's correspondientes \n")
				pk_fk.BorrarFK()
				pk_fk.BorrarPK()
			} else if opcion == 5 {
				fmt.Println("solo hecho creacion del pedido y agregado de producto")
				stored_procedures.Creacion_del_pedido()
				stored_procedures.Agregado_de_producto()
			} else if opcion == 6 {
				fmt.Println("Elegiste la opcion 6, se unmarshalearan los JSON's y se subiran los datos a las tablas \n")
				datos.Unmalvaviscador()
			}
	}

func main() {
	fmt.Println("Bienvenido a la base de datos de ChangoOnline!\n")
	fmt.Println("Elija entre las siguientes opciones para continuar!\n")
	var user int
	for {
		printMenu()
		_, err := fmt.Scanf("%d", &user)
		if err != nil {
			fmt.Println("Error: ", err)
			break
			}
			
		if user == 0 {
			fmt.Println("Chau!")
			break
			}
		
		continuar(user)
	}	
}
