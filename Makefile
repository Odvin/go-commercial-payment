paymentService:
	DATA_SOURCE_URL=root:verysecretpass@tcp\(localhost:3306\)/payment APPLICATION_PORT=3001 ENV=development go run cmd/main.go

.PHONY: paymentService