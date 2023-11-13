create or replace function cierre_del_pedido (pedido_id int, fecha_entrega timestamp) returns boolean as $$

declare
	cantidad_productos int;
	hora_entrega time;
begin

	-- valida que el id del pedido exista en la tabla.
	if not exists (select 1 from pedido p where p.id_pedido = pedido_id) then
		-- se carga en la tabla el error
		insert into error(id_error, id_pedido, f_error, motivo)
		values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?id de pedido no valido.'); 
		
		raise 'id de pedido no valido.';
		return false;
	end if;
	
	-- valida que el pedido tenga al menos un producto agregado.
	select cantidad into cantidad_productos from pedido_detalle pd where pd.id_pedido = pedido_id;
	
	if cantidad_productos < 1 then
	insert into error(id_error, id_pedido, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?pedido vacio'); 
		raise 'pedido vacio';
		return false;
	end if;
	
	-- valida que la fecha de entrega del pedido sea mayor a la fecha actual.
	if fecha_entrega > current_timestamp then 
	insert into error(id_error, id_pedido, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?fecha de entrega no válida'); 
		raise 'fecha de entrega no válida';
		return false;
	end if;
	
	-- despues de validar todo:
	-- actualiza la fecha de entrega del pedido desde y hasta que hora puede llegar.
	-- casteo la fecha de entrega a time para sacar la hora
	select cast(fecha_entrega as time) into hora_entrega;

	update pedido set hora_entrega_desde = hora_entrega where id_pedido = pedido_id;
	update pedido set hora_entrega_hasta = (hora_entrega + interval '2 hour') where id_pedido = pedido_id;	
	
	--actualiza el estado del pedido
	update pedido set estado = 'completado' where id_pedido = pedido_id;	
	
	return true;

end;

$$ language plpgsql;
