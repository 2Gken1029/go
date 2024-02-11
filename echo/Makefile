init:
	@echo "初期設定を開始します"
	docker compose build

start:
	@echo "コンテナを立ち上げます"
	docker compose up

down:
	@echo "コンテナを落とします"
	docker compose down
	
go:
	@echo "コンテナの中に入ります"
	docker compose exec go bash

include ./mysql/.env
migrate-up:
	docker-compose exec go migrate -path=db/migrations -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:3306)/$(MYSQL_DATABASE)??multiStatements=true"  up 1
