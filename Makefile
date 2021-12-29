app.start:
	go run main.go

docker.start:
	docker-compose -f build/docker-compose.yml up -d

docker.ps:
	docker-compose -f build/docker-compose.yml ps

docker.db.start:
	docker-compose -f build/docker-compose.yml up -d db

docker.db.stop:
	docker-compose -f build/docker-compose.yml rm -fs db

docker.db.sql:
	docker-compose -f build/docker-compose.yml run db psql -h db -U alura -d alura_db
