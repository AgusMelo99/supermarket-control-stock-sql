create trigger enviar_mail_trigger
after insert or update on pedido
for each row
execute procedure enviar_mail();
