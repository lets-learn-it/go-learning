# RabbitMQ

## Setup

### Run rabbitmq as container

```sh
podman run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management-alpine
```

**Note**: Default user is `guest` and password is `guest`.

### Configure Rabbitmq

```sh
# add user
podman exec rabbitmq rabbitmqctl add_user pskp secret

# make new use admin
podman exec rabbitmq rabbitmqctl set_user_tags pskp administrator

# delete guest user
podman exec rabbitmq rabbitmqctl delete_user guest

# add virtual host
podman exec rabbitmq rabbitmqctl add_vhost customers

# give new user permissions for vhost customers
podman exec rabbitmq rabbitmqctl set_permissions -p customers pskp ".*" ".*" ".*"

# create exchange
podman exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u pskp -p secret durable=true

# give user read/write permissions to topic
podman exec rabbitmq rabbitmqctl set_topic_permissions -p customers pskp customer_events "^customers.*" "^customers.*"
```

## References

[[00] https://www.youtube.com/watch?v=1yC_bw0tWhQ&t=2760s](https://www.youtube.com/watch?v=1yC_bw0tWhQ&t=2760s)

[[01] https://www.rabbitmq.com/tutorials/amqp-concepts.html](https://www.rabbitmq.com/tutorials/amqp-concepts.html)

