package main

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
	database 			"trabajo-practico-base-de-datos/bd"
	datos	 			"trabajo-practico-base-de-datos/datajson"
	sp					"trabajo-practico-base-de-datos/stored-procedures"
	ver					"trabajo-practico-base-de-datos/ver_tablas"
)

var ConexionDataBase *sql.DB

func printMenu() {
	fmt.Println("1. Crear base de datos \n")
	fmt.Println("2. Crear Tablas \n")
	fmt.Println("3. Agregar PK's y FK's \n")
	fmt.Println("4. Eliminar PK's y FK's \n")
	fmt.Println("5. Agregar Stored Procedures y Triggers \n")
	fmt.Println("6. Unmarshalear JSON's y cargar los datos \n")
	fmt.Println("7. Ver clientes \n")
	fmt.Println("8. Ver productos \n")
	fmt.Println("9. Ver pedidos \n")
	fmt.Println("10. Ver detalles de los pedidos \n")
	fmt.Println("11. Ver reposición \n")
	fmt.Println("12. Ver envio de mails \n")
	fmt.Println("13. Ver errores \n")
	fmt.Println("14. Cargar pedidos de prueba\n")
	fmt.Println("0. Cerrar sesion\n")
	}

func continuar (opcion int){
	if opcion == 1 {
		if ConexionDataBase != nil {
			fmt.Println("Ya creaste la base de datos previamente!\n")
		} else {
			fmt.Println("Elegiste la opcion 1, crear base de datos \n")
			database.CreateDatabase()
		}
			} else if opcion == 2 {
				fmt.Println("Elegiste la opcion 2, se crearan las tablas correspondientes \n")
				database.CreateTables(ConexionDataBase)
			} else if opcion == 3 {
				fmt.Println("Elegiste la opcion 3, se crearan las PK's y FK's correspondientes \n")
				database.CreatePK(ConexionDataBase)
				database.CreateFK(ConexionDataBase)
			} else if opcion == 4 {
				fmt.Println("Elegiste la opcion 4, se eliminaran las PK's y FK's correspondientes \n")
				database.BorrarFK(ConexionDataBase)
				database.BorrarPK(ConexionDataBase)
			} else if opcion == 5 {
				fmt.Println("Elegiste la opcion 5, creacion del pedido, agregado de producto y envio de mails")
				sp.SubirStoredProceduresyTriggers(ConexionDataBase)
			} else if opcion == 6 {
				fmt.Println("Elegiste la opcion 6, se unmarshalearan los JSON's y se subiran los datos a las tablas \n")
				datos.Unmalvaviscador()
				datos.InsertarClientes(ConexionDataBase)
				datos.InsertarDireccionesEntrega(ConexionDataBase)
				datos.InsertarProductos(ConexionDataBase)
				datos.InsertarTarifaEntrega(ConexionDataBase)
			} else if opcion == 7 {
				fmt.Println("Elegiste la opcion 7, a continuacion se mostrara el contenido de la tabla cliente\n")
				ver.VerClientes(ConexionDataBase)
			} else if opcion == 8 {
				fmt.Println("Elegiste la opcion 8, a continuacion se mostrara el contenido de la tabla producto\n")
				ver.VerProductos(ConexionDataBase)
			} else if opcion == 9 {
				fmt.Println("Elegiste la opcion 9, a continuacion se mostrara el contenido de la tabla pedido\n")
				ver.VerPedidos(ConexionDataBase)
			} else if opcion == 10 {
				fmt.Println("Elegiste la opcion 10, a continuacion se mostrara el contenido de la tabla pedido_detalle\n")
				ver.VerPedidosDetalle(ConexionDataBase)
			} else if opcion == 11 {
				fmt.Println("Elegiste la opcion 11, a continuacion se mostrara el contenido de la tabla reposicion\n")
				ver.VerReposicion(ConexionDataBase)
			} else if opcion == 12 {
				fmt.Println("Elegiste la opcion 12, a continuacion se mostrara el contenido de la tabla envio_email\n")
				ver.VerEmails(ConexionDataBase)
			} else if opcion == 13 {
				fmt.Println("Elegiste la opcion 13, a continuacion se mostrara el contenido de la tabla error\n")
				ver.VerErrores(ConexionDataBase)
			} else if opcion == 14 {
				fmt.Println("Elegiste la opcion 14, a continuacion se insertaran las operaciones sobre la base de datos\n")
				datos.InsertarOperaciones(ConexionDataBase)
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
		if user == 1 {
			var err error
			ConexionDataBase,err = sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
				if err != nil {
				log.Fatal(err)
			}
			defer ConexionDataBase.Close()
		}
	}	
}
