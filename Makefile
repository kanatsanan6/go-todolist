migrateup:
	migrate -path migration -database "postgresql://postgres:@localhost:5432/go_todo_list?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://postgres:@localhost:5432/go_todo_list?sslmode=disable" -verbose down

.PHONY: migrateup migratedown
