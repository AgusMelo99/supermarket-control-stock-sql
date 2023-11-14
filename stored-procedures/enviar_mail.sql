create or replace function enviar_mail() returns trigger as $$

declare
	this_pedido	pedido%rowtype;
	email_del_cliente text;
	nombre_del_cliente text;
	cuerpo text;
	direccion_de_entrega varchar(60);
	detalles text;
	
begin

	this_pedido := new;
	select email into email_del_cliente from cliente where id_usuarie = this_pedido.id_usuarie;
	select nombre into nombre_del_cliente from cliente where id_usuarie = this_pedido.id_usuarie;
	select d.direccion into direccion_de_entrega from direccion_entrega d where d.id_usuarie = this_pedido.id_usuarie and d.id_direccion_entrega = this_pedido.id_direccion_entrega;
	select string_agg('ID del producto: ' || id_producto || ', cantidad: ' || cantidad || ', precio unitario: ' || precio_unitario, chr(10)) into detalles from pedido_detalle where id_pedido = this_pedido.id_pedido;
	
	if this_pedido.estado = 'completado' then
		-- los valores son de prueba
		cuerpo := 'Hola ' || nombre_del_cliente || ', te enviamos este mail para confirmarte que el pedido con el ID numero: ' || this_pedido.id_pedido || ' ha sido aceptado, a continuacion te dejamos los detalles del mismo' || chr(10)
				|| 'Fecha de pedido: ' || this_pedido.f_pedido || chr(10)
				|| 'Fecha de entrega: ' || this_pedido.fecha_entrega || chr(10)
				|| 'Recibiras tu envio entre las: ' || this_pedido.hora_entrega_desde|| chr(10)
				|| 'Hasta las: ' || this_pedido.hora_entrega_hasta|| chr(10)
				|| 'Costo de envio: ' || this_pedido.costo_envio|| chr(10)
				|| 'Monto total: ' || this_pedido.monto_total|| chr(10)
				|| 'Direccion de entrega: ' || direccion_de_entrega|| chr(10)
				|| 'Detalles del pedido: ' || chr(10) || detalles || chr(10);
				
		insert into envio_email values (COALESCE((select max(id_email) + 1 from envio_email), 0), current_timestamp, email_del_cliente, 'Pedido aceptado', cuerpo, current_timestamp, 'enviado');
		
	elsif this_pedido.estado = 'cancelado' then
		cuerpo := 'Hola ' || nombre_del_cliente || ', te enviamos este mail para confirmarte que el pedido con el ID numero: ' 
				|| this_pedido.id_pedido || ' ha sido cancelado, a continuacion te dejamos los detalles del mismo' || chr(10)
				|| 'Fecha de pedido: ' || this_pedido.f_pedido || chr(10);
					
		insert into envio_email values (COALESCE((select max(id_email) + 1 from envio_email), 0), current_timestamp, email_del_cliente, 'Pedido cancelado', cuerpo, current_timestamp, 'cancelado');
			
	elsif this_pedido.estado = 'entregado' then
		cuerpo := 'Hola ' || nombre_del_cliente || ', te enviamos este mail para confirmarte que el pedido con el ID numero: ' || this_pedido.id_pedido || ' ha sido enviado, a continuacion te dejamos los detalles del mismo' || chr(10)
				|| 'Fecha de pedido: ' || this_pedido.f_pedido || chr(10)
				|| 'Fecha de entrega: ' || this_pedido.fecha_entrega || chr(10)
				|| 'Recibiras tu envio entre las: ' || this_pedido.hora_entrega_desde|| chr(10)
				|| 'Hasta las: ' || this_pedido.hora_entrega_hasta|| chr(10)
				|| 'Costo de envio: ' || this_pedido.costo_envio|| chr(10)
				|| 'Monto total: ' || this_pedido.monto_total|| chr(10)
				|| 'Direccion de entrega: ' || direccion_de_entrega|| chr(10)
				|| 'Detalles del pedido: ' || chr(10) || detalles || chr(10);
				
		insert into envio_email values (COALESCE((select max(id_email) + 1 from envio_email), 0), current_timestamp, email_del_cliente, 'Pedido entregado', cuerpo, current_timestamp, 'entregado');
			-- el segundo current_timestamp deberia ser la fecha del dia siguiente.
			-- falta validar que la fecha actual current_timestamp sea igual al registro de la fecha siguiente para updatear el estado del envio a "enviado"
			
		cuerpo:= 'Te enviamos esta encuesta del pedido:' || ' ' || this_pedido.id_pedido || ' ' ||
				'para saber el nivel de satisfaccion durante el periodo de compra y envio' || chr(10);
				
		insert into envio_email values (COALESCE((select max(id_email) + 1 from envio_email), 0), current_timestamp, email_del_cliente, 'encuesta', cuerpo, null, 'pendiente');
	end if;

	return new;
end;
$$ language plpgsql;
