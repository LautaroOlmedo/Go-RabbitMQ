DOCKER && RABBITMQ COMMANDS

*sudo docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management ---> // START CONTAINER

*sudo docker exec rabbitmq rabbitmqctl ---> // RABBITMQ MANAGEMENT 

*sudo docker exec rabbitmq rabbitmqctl add_user lautaro secret ---> // add a new RABBITMQ user

*sudo docker exec rabbitmq rabbitmqctl set_user_tags

*sudo docker exec rabbitmq rabbitmqctl set_user_tags lautaro administrator ---> // assigns the new user the administrator role

*sudo docker exec rabbitmq rabbitmqctl delete_user guest ---> // delete a default RABBITMQ user

*sudo docker exec rabbitmq rabbitmqctl add_vhost customers ---> // make a new VIRTUAL HOST

*sudo docker exec rabbitmq rabbitmqctl set_permissions -p customers lautaro ".*" ".*" ".*" ---> // set permissions to a user on a virtual host




*sudo docker exec rabbitmq rabbitmqadmin

*sudo docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u lautaro -p secret durable=true