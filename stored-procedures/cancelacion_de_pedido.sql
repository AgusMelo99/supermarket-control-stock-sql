create or replace function cancelacion_de_pedido (pedido_id int) returns boolean as $$ 

declare
	
	registro record;
	
begin
	-- valida que el id del pedido exista en la tabla.
	if not exists (select 1 from pedido p where p.id_pedido = pedido_id) then
		-- se carga en la tabla el error
		insert into error(id_error, id_pedido, f_error, motivo)
		values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?id de pedido no valido.'); 
		
		raise 'id de pedido no valido.';
		return false;
	end if;
	
	--valida que el estado del pedido no sea ingresado o completado
	if not exists (select 1 from pedido p where p.id_pedido = pedido_id and p.estado = 'ingresado' or p.estado = 'completado') then
		-- se carga en la tabla el error
		insert into error(id_error, id_pedido, f_error, motivo)
		values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?pedido ya entregado o completado.'); 
		
		raise 'pedido ya entregado o completado.';
		return false;
	end if;
	
	--terminan validacion
	
	--cambia el estado del pedido a cancelado
	update pedido set estado = 'cancelado' where id_pedido = pedido_id;	
	
	-- por cada fila que esté relacionada con el id_pedido que fue cancelado, se guarda el id_producto y cantidad en la variable registro(record).
	for registro in select id_producto, cantidad from pedido_detalle where id_pedido = pedido_id loop
	
	-- actualiza la tabla producto restando al stock_reservado la cantidad del pedido cancelado.
		update producto set stock_reservado = (stock_reservado - registro.cantidad) where id_producto = registro.id_producto;
	-- actualiza la tabla producto sumando al stock disponible la cantidad del pedido cancelado
		update producto set stock_disponible = (stock_disponible + registro.cantidad) where id_producto = registro.id_producto;
	end loop;

	return true;
end;

$$ language plpgsql;
