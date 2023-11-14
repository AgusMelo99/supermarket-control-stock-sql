package datajson

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
	Fecha_entrega	string
	Hora_entrega_desde	string
	Hora_entrega_hasta	string
	Id_usuarie	int
	Id_direccion_entrega int
	Monto_total	float64
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
	Fecha_reposicion	string
	Estado	string
	}

type Envio_email struct{
	Id_email int
	F_generacion	string
	Email_cliente	string
	Asunto	string
	Cuerpo	string
	F_envio	string
	Estado	string
	}
	
type Error struct{
	Id_error int
	Id_pedido	int
	F_pedido	string
	Id_usuarie	int
	Id_direccion_entrega int
	Direccion	string
	Localidad	string
	Codigo_postal	string
	Id_producto int
	Cantidad int
	Operacion string
	F_error string
	Motivo	string
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


	
	
	
	
