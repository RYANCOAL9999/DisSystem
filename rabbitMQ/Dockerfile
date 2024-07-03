# Use the official RabbitMQ image from Docker Hub
FROM rabbitmq:3.9.9-management-alpine

# Optionally, copy custom configuration files
COPY rabbitmq.conf /etc/rabbitmq/rabbitmq.conf

# Expose ports for AMQP (5672) and RabbitMQ management UI (15672)
# These ports are already exposed in the base image, so this step is optional
EXPOSE 5672 15672

# Set environment variables if needed (e.g., RabbitMQ default user and password)
ENV RABBITMQ_DEFAULT_USER=admin
ENV RABBITMQ_DEFAULT_PASS=password

# Example of how to enable plugins (e.g., rabbitmq_management)
RUN rabbitmq-plugins enable rabbitmq_management

# Start RabbitMQ server on container startup
CMD ["rabbitmq-server"]