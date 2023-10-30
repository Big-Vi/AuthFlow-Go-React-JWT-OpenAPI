# authflow
Building event ticketing system.  

migrate create -ext sql -dir migrations -seq create_users_table  
migrate -path migrations/ -database "postgres://postgres:root@localhost:5433/authflow-store-db?sslmode=disable" -verbose up  
migrate -path migrations/ -database "postgres://postgres:root@localhost:5433/authflow-store-db?sslmode=disable" -verbose down  
migrate -path migrations/ -database "postgres://postgres:root@localhost:5433/authflow-store-db?sslmode=disable" force 1  

docker exec -it authflow-store_redis redis-cli  
