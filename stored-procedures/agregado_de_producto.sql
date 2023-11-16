create or replace function agregado_de_producto (pedido_id int, producto_id int, cantidad_producto int) returns boolean as $$

declare
	
	estado_pedido char(10);
	stock_producto int;
	precio_unitario_producto decimal(12,2);

begin
	RAISE NOTICE 'El nivel de aislamiento actual es: %', current_setting('transaction_isolation', true);
	-- verificar si existe el pedido
	if not exists (select 1 from pedido p where p.id_pedido = pedido_id) then
		insert into error(id_error, id_pedido, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?id de pedido no valido'); 
		
		return false;
	end if;

	-- verificar que el pedido se encuentre en estado 'ingresado'
	select estado into estado_pedido from pedido p where p.id_pedido = pedido_id;
	
	if estado_pedido != 'ingresado' then
		insert into error(id_error, id_pedido, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?pedido cerrado'); 
		
		return false;
	end if;
	
	-- verificar que el producto exista
	if not exists (select 1 from producto p where p.id_producto = producto_id) then
		insert into error(id_error, id_pedido, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?id de producto no valido'); 
		
		return false;
	end if;

	-- verificar si hay stock del producto
	select stock_disponible into stock_producto from producto p where p.id_producto = producto_id;
		
	if cantidad_producto > stock_producto then
		insert into error(id_error, id_pedido, f_error, motivo)
			values (coalesce((select max(id_error) from error), 0) + 1, pedido_id, current_date, '?stock no disponible para el producto' || ' ' || producto_id);  
		
		return false;
	end if;
		
	-- si llego hasta aca es porque se puede agregar el producto
		
	-- obtener el precio unitario del producto
	select precio_unitario into precio_unitario_producto from producto p where p.id_producto = producto_id;
	
	-- paso 1: descontar el stock disponible y agregarlo al stock reservado
	update producto set stock_disponible = (stock_disponible - cantidad_producto),
		stock_reservado = (stock_reservado + cantidad_producto) where id_producto = producto_id;
		
		
	-- paso 2: actualizar el monto total del pedido
	update pedido set monto_total = (coalesce(monto_total, 0) + (precio_unitario_producto * cantidad_producto)) where id_pedido = pedido_id;
	-- verificar si el producto ya esta en el pedido
	if not exists (select 1 from pedido_detalle p where p.id_producto = producto_id) then 
		-- paso 3: insertar el producto en pedido_detalle si no existe
		insert into pedido_detalle values (pedido_id, producto_id, cantidad_producto, precio_unitario_producto);
	else
		-- si existe, actualizar la cantidad
		update pedido_detalle set cantidad = (cantidad + cantidad_producto) where id_pedido = pedido_id and id_producto = producto_id;
	end if;
	return true;
end;

$$ language plpgsql;
