create or replace function entrega_pedido(pedido_id int) returns boolean as $$

declare

	estado_pedido pedido.estado%type;
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
	
	--consulta el estado actual del pedido.
	select estado_pedido into estado_actual from pedido p where p.id_pedido = pedido_id;
	
	-- valida si el pedido ya se encuentra en estado completado.
	if estado_pedido != 'completado' then
		-- se carga en la tabla el error
		insert into error(id_error, id_pedido, f_error, motivo)
		values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?pedido sin cerrar o ya entregado'); 
		
		raise '?pedido sin cerrar o ya entregado';
		return false;
	end if;
	
	-- POST VALIDACIONES
	-- actualiza el estado del pedido como entregado.
	update pedido set estado = 'entregado' where id_pedido = pedido_id;	
	
	-- por cada fila que esté relacionada con el id_pedido que fue entregado, se guarda el id_producto y cantidad en la variable registro(record).
	for registro in select id_producto, cantidad from pedido_detalle where id_pedido = pedido_id loop
	
	-- actualiza la tabla producto restando al stock_reservado la cantidad del pedido entregado.
		update producto set stock_reservado = (stock_reservado - registro.cantidad) where id_producto = registro.id_producto;
	end loop;
	
	return true;
end;

$$ language plpgsql;
