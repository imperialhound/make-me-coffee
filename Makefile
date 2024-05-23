CLUSTER  := coffee-rabbit

.PHONY: start-rabbit stop-rabbit

start-rabbit:
	@echo "Starting RabbitMQ Cluster"
	docker run --rm -it -d --network host --name ${CLUSTER} rabbitmq:3-management 

stop-rabbit:
	@echo "Stopping RabbitMQ Cluster"
	docker stop ${CLUSTER} 
