MIGRATE=docker exec -it cake-store migrate -path=migration -database "mysql://root:root@tcp(cake-store-db)/cake-store" -verbose

migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down
drop:
		$(MIGRATE) drop


.PHONY: migrate-up migrate-down drop
