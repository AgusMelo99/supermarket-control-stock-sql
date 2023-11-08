create or replace function creacion_del_pedido (id_cliente int, id_direccion int) returns boolean as $$

declare

	tarifa int;
	cp_corto char(4);
	
begin
	-- verificar si existe el usuario
	if not exists (select 1 from cliente where id_usuarie = id_cliente) then
		--coalesce hace que si la tabla error este vacia devuelva como valor predeterminado 
		--el numero que se pone al final de la sentencia, en este caso 0
		insert into error(id_error, id_usuarie, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, id_cliente, current_date, '?id de usuarie no valido'); 
		raise 'id de usuario invalido.';
		return false;
	end if;
	
	-- verificar si la direccion de entrega existe
	if not exists (select 1 from direccion_entrega where id_direccion_entrega = id_direccion) then
		--coalesce hace que si la tabla error este vacia devuelva como valor predeterminado 
		--el numero que se pone al final de la sentencia, en este caso 0
		insert into error(id_error, id_direccion_entrega, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, id_direccion, current_date, '?id de direccion no valido'); 
		raise 'id de direccion no valido.';
		return false;
	end if;
	
	-- obtener el codigo postal corto
	select substring(codigo_postal, 2, 4) into cp_corto from direccion_entrega where id_direccion_entrega = id_direccion;
	
	-- verificar si el cp_corto tiene tarifa
	select costo into tarifa from tarifa_entrega where codigo_postal_corto = cp_corto;
	
	-- verificar si la tarifa es valida
	if tarifa is null then
		--coalesce hace que si la tabla error este vacia devuelva como valor predeterminado 
		--el numero que se pone al final de la sentencia, en este caso 0
		insert into error(id_error, id_direccion_entrega, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, id_direccion, current_date, '?direccion de entrega fuera del area de atencion'); 
		raise 'direccion de entrega fuera del area de atencion';
		return false;
	end if;
	
	-- insertar pedido si todos los datos son correctos
	insert into pedido (id_pedido, f_pedido, id_usuario, id_direccion_entrega, costo_envio, estado) 
		values (coalesce((select max(id_pedido) from pedido), 0) + 1, current_timestamp, id_usuario, id_direccion, tarifa, 'ingresado');
	return true;
end;

$$ language plpgsql;
