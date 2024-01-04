start-user-management-service:
	go run cmd/user-management-service/main.go
create-user:
	go run cmd/client/main.go user create --email=$(EMAIL) --password=$(PASSWORD)