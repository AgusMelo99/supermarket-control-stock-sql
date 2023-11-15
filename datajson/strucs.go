package datajson

import (
	"database/sql"
)
/* sql.NullString, sql.NullINt64 y sql.NullFloat64 son tipos de datos que a diferencia de int, string y float64 pueden tomar Null como valor
 * por eso las columnas que en algun momento de la ejecucion pueden ser null toman esos tipos de valores
 * estos tienen una funcion .Valid (retorna true o false) la cual verifica si el valor contenido es Null o tiene un valor insertado
*/
type Cliente struct {
	Id_usuarie	int
	Nombre		string
	Apellido	string
	Dni			int
	Fecha_nacimiento string
	Telefono	string
	Email		string
	}
	
type Direccion_entrega struct {
	Id_usuarie	int
	Id_direccion_entrega int
	Direccion	string
	Localidad	string
	Codigo_postal string
	} 
	
type Tarifa_entrega struct {
	Codigo_postal_corto string
	Costo		float64
	}
	
type Producto struct {
	Id_producto	int
	Nombre		string
	Precio_unitario float64
	Stock_disponible	int
	Stock_reservado	int
	Punto_reposicion	int
	Stock_maximo	int
	}
	
type Pedido struct {
	Id_pedido	int
	F_pedido	string
	Fecha_entrega	sql.NullString 
	Hora_entrega_desde	sql.NullString
	Hora_entrega_hasta	sql.NullString
	Id_usuarie	int
	Id_direccion_entrega int
	Monto_total	sql.NullFloat64
	Costo_envio	float64
	Estado		string
	}
	
type Pedido_detalle struct{
	Id_pedido	int
	Id_producto	int
	Cantidad	int
	Precio_unitario float64
	}
	
type Reposicion struct{
	Id_producto	int
	Fecha_solicitud string
	Cantidad_a_reponer int
	Fecha_reposicion	sql.NullString
	Estado	string
	}

type Envio_email struct{
	Id_email int
	F_generacion	string
	Email_cliente	string
	Asunto	string
	Cuerpo	sql.NullString
	F_envio	sql.NullString
	Estado	string
	}
	
type Error struct{
	Id_error int
	Id_pedido	sql.NullInt64
	F_pedido	sql.NullString
	Id_usuarie	sql.NullInt64
	Id_direccion_entrega sql.NullInt64
	Direccion	sql.NullString
	Localidad	sql.NullString
	Codigo_postal	sql.NullString
	Id_producto sql.NullInt64
	Cantidad sql.NullInt64
	Operacion sql.NullString
	F_error sql.NullString
	Motivo	sql.NullString
	}
	
type Entrada_trx_pedido struct {
	Id_orden int
	Operacion string
	Id_usuarie int
	Id_direccion_entrega int
	Id_pedido int
	Id_producto int
	Cantidad int
	Fecha_hora_entrega string
	}


	
	
	
	
