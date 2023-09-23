DOCKER && RABBITMQ COMMANDS

*sudo docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management ---> // Inicializa el contenedor. En el puerto 15672 se encuentra la interfaz de rabbitMQ (usuario por default user=guest, password=guest)

*sudo docker exec rabbitmq rabbitmqctl ---> // uno de los dos comando principal para manejar rabbitMQ 

*sudo docker exec rabbitmq rabbitmqctl add_user lautaro secret ---> // agregar un nuevo usuario a RabbitMQ (en este caso user=lautaro, password=secret)


*sudo docker exec rabbitmq rabbitmqctl set_user_tags lautaro administrator ---> // Asignar usuario como administrador 

*sudo docker exec rabbitmq rabbitmqctl delete_user guest ---> // emilinar usuario que viene por defecto

*sudo docker exec rabbitmq rabbitmqctl add_vhost customers ---> // crea un VIRTUAL HOST, en este caso "customers"

*sudo docker exec rabbitmq rabbitmqctl set_permissions -p customers lautaro ".*" ".*" ".*" ---> // le da permisos al usuario para poder ESCRIBIR, LEER Y CONFIGURAR virtual host


*sudo docker exec rabbitmq rabbitmqadmin // uno de los dos comando principal para manejar rabbitMQ 

*sudo docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u lautaro -p secret durable=true // ---> Crea un EXCHANGE dentro de un virtual host
