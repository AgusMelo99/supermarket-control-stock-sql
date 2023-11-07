create or replace function enviar_mail() returns trigger as $$

declare

	estado_pedido char(10);
	
begin

	select estado into estado_pedido from pedido;
	
	if estado_pedido = 'ingresado' then
		-- los valores son de prueba
		insert into envio_mail values (1, current_timestamp, 'mail_cliente', 'Pedido aceptado', 'cuerpo', current_timestamp, 'enviado');
		
		else if estado_pedido = 'cancelado' then
			insert into envio_mail values (1, current_timestamp, 'mail_cliente', 'Pedido cancelado', 'cuerpo', current_timestamp, 'cancelado');
			
		else if estado_pedido = 'entregado' then
			insert into envio_mail values (1, current_timestamp, 'mail_cliente', 'Pedido entregado', 'cuerpo', current_timestamp, 'entregado');
			
			-- el segundo current_timestamp deberia ser la fecha del dia siguiente.
			insert into envio_mail values (1, current_timestamp, 'mail_cliente', 'encuesta', 'cuerpo', current_timestamp, 'pendiente');

	return new;
end;
$$ language plpgsql;
