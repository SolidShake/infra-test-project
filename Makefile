.PHONY: build up down restart rebuild logs clean help app-restart filebeat-restart


build:
	@echo "Building all containers..."
	docker compose build

# Start all containers
up:
	@echo "Starting all containers..."
	docker compose up -d

# Stop all containers
down:
	@echo "Stopping all containers..."
	docker compose down

# Restart all containers
restart:
	@echo "Restarting all containers..."
	docker compose restart

# Rebuild and restart all containers
rebuild:
	@echo "Rebuilding and restarting all containers..."
	docker compose down
	docker compose build
	docker compose up -d

# Rebuild and restart only the app container
app-restart:
	@echo "Rebuilding and restarting app container..."
	docker compose build app
	docker compose stop app
	docker compose rm -f app
	docker compose up -d app

# Restart only the filebeat container
filebeat-restart:
	@echo "Restarting filebeat container..."
	docker compose restart filebeat

# Show logs from all containers
logs:
	@echo "Showing logs from all containers..."
	docker compose logs -f

# Show logs from the app container
logs-app:
	@echo "Showing logs from app container..."
	docker compose logs -f app

# Show logs from Elasticsearch, Logstash, and Kibana
logs-elk:
	@echo "Showing logs from ELK stack..."
	docker compose logs -f elasticsearch logstash kibana

# Show logs from Filebeat
logs-filebeat:
	@echo "Showing logs from Filebeat..."
	docker compose logs -f filebeat

# Remove all containers, volumes, and networks
clean:
	@echo "Removing all containers, volumes, and networks..."
	docker compose down -v
	docker compose rm -f 