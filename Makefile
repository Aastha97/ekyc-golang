.PHONY: setup
setup:
	docker-compose up -d

.PHONY: setup-down
setup-down:
	docker-compose down --volumes

.PHONY: connect
connect:
	docker exec -it ekyc-db psql -U postgres -d postgres
