create or replace function generar_solicitud_reposicion() returns boolean as $$

declare

	registro record;
	cantidad_reposicion int;

begin 
	
	-- valida que no exista una solicitud de reposicion ya cargada con la fecha actual
	if exists (select 1 from reposicion where estado = 'pendiente' and fecha_solicitud = current_date) then
		return false;
	end if;
	
	-- recorre cada producto verificando si está bajo de stock
	for registro in select * from producto loop
		if (registro.stock_disponile + registro.stock_reservado) <= registro.punto_reposicion then
			-- calcula la cantidad que se necesita reponer
			cantidad_reposicion = registro.stock_maximo - (registro.stock_disponile + registro.stock_reservado);
			
			-- se inserta un nuevo registro de solicitud en la tabla reposicion
			insert into reposicion (registro.id_producto, current_date, cantidad_reposicion, null, 'pendiente');
		end if;
	end loop;
	
	return true;
	
end;

$$ language plpgsql;
