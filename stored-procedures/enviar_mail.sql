create or replace function enviar_mail() returns trigger as $$

declare
	this_pedido	pedido%rowtype;
	email_del_cliente text;
	nombre_del_cliente text;
	cuerpo text;
	direccion_de_entrega varchar(60);
	detalle_del_producto text;
	cantidad int;
	precio_unitario decimal(12,2);
	nombre_del_producto text;
	
begin

	select new into this_pedido;
	select email into email_del_cliente from cliente where id_usuarie = this_pedido.id_usuarie;
	select nombre into nombre_del_cliente from cliente where id_usuarie = this_pedido.id_usuarie;
	select direccion into direccion_de_entrega from direccion_entrega d where d.id_usuarie = this_pedido.isuarie;
	select p.nombre into nombre_del_producto from producto p where p.id_producto = (select d.id_producto from pedido_detalle d where d.id_pedido = this_pedido.id_pedido);
	select d.cantidad into cantidad from pedido_detalle d where d.id_pedido = this_pedido.id_pedido;
	select d.precio_unitario into precio_unitario from pedido_detalle d where d.id_pedido = this_pedido.id_pedido;
	
	if this_pedido.estado = 'completado' then
		-- los valores son de prueba
		cuerpo:= 'Te enviamos este mail para confirmarte que el pedido con el ID numero:' || ' ' || this_pedido.id || ' ' ||
				'ha sido aceptado, a continuacion los detalles del mismo: \n' || ' ' ||
				'ID: ' || this_pedido.id || '\n' ||
				'Fecha del pedido: ' || this_pedido.f_pedido || '\n' ||
				'Direccion de entrega: ' || direccion_de_entrega || '\n' ||
				'Nombre del producto: ' || nombre_del_producto || '\n' ||
				'Cantidad: ' || cantidad || '\n' ||
				'Precio por unidad: ' || precio_unitario || '\n' ||
				'Precio total: ' || this_pedido.monto_total || '\n';
				
		insert into envio_email values (1, current_timestamp, email_del_cliente, 'Pedido aceptado', cuerpo, current_timestamp, 'enviado');
		
	elsif estado_pedido = 'cancelado' then
		insert into envio_email values (1, current_timestamp, email_del_cliente, 'Pedido cancelado', 'cuerpo', current_timestamp, 'cancelado');
			
	elsif estado_pedido = 'entregado' then
		insert into envio_email values (1, current_timestamp, email_del_cliente, 'Pedido entregado', 'cuerpo', current_timestamp, 'entregado');
			-- el segundo current_timestamp deberia ser la fecha del dia siguiente.
			-- falta validar que la fecha actual current_timestamp sea igual al registro de la fecha siguiente para updatear el estado del envio a "enviado"
		insert into envio_email values (1, current_timestamp, email_del_cliente, 'encuesta', 'cuerpo', (current_timestamp + interval '1 day'), 'pendiente');
	end if;

	return new;
end;
$$ language plpgsql;
