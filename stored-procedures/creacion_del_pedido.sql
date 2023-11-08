create or replace function creacion_del_pedido (id_cliente int, id_direccion int) returns boolean as $$

declare

	tarifa int;
	cp_corto char(4);
	
begin
	-- verificar si existe el usuario
	if not exists (select 1 from cliente where id_usuarie = id_cliente) then
		raise 'id de usuario invalido.';
		return false;
	end if;
	
	-- verificar si la direccion de entrega existe
	if not exists (select 1 from direccion_entrega where id_direccion_entrega = id_direccion) then
		raise 'id de direccion no valido.';
		return false;
	end if;
	
	-- obtener el codigo postal corto
	select substring(codigo_postal, 3, 4) into cp_corto from direccion_entrega where id_direccion_entrega = id_direccion;
	
	-- verificar si el cp_corto tiene tarifa
	select costo into tarifa from tarifa_entrega where codigo_postal_corto = cp_corto;
	
	-- verificar si la tarifa es valida
	if tarifa is null then
		raise 'direccion de entrega fuera del area de atencion';
		return false;
	end if;
	
	-- insertar pedido si todos los datos son correctos
	insert into pedido (id_usuario, id_direccion_entrega, fecha_y_hora_pedido, costo_envio, estado) values (id_cliente, id_direccion, current_date, tarifa, 'ingresado');
	return true;
end;

$$ language plpgsql;
