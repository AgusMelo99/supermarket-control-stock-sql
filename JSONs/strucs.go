package JSONs

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

	
	
	
	
