package main

import (
//	"database/sql"
	"fmt"
	_"github.com/lib/pq"
//	"log"
	database "trabajo-practico-base-de-datos/BD"
)

func printMenu() {
	fmt.Println("Bienvenido a la base de datos de ChangoOnline!\n")
	fmt.Println("Elija entre las siguientes opciones para continuar!\n")
	fmt.Println("1. Crear base de datos \n")
	fmt.Println("2. Crear Tablas \n")
	fmt.Println("3. Aniadir PK's y FK's \n")
	fmt.Println("4. Eliminar PK's y FK's \n")
	fmt.Println("5. Agregar SP's y Triggers \n")
	fmt.Println("0. Cerrar sesion\n")
	}

func continuar (opcion int){
	if opcion == 1 {
		fmt.Println("Elegiste la opcion 1, crear base de datos \n")
		database.CreateDatabase()
		}
	}

func main() {
	var user int
	fmt.Println("hola")
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
