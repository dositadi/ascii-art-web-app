migrate-create:
	migrate create -ext sql -dir database/migrations -seq $name

migrate-up:
	migrate -path database/migrations -database "mysql://username:password@tcp(host)/db_name" up 

migrate-down: 
	migrate -path database/migrations -database "mysql://username:password@tcp(host)/db_name" down 1

seed-into-users:
	mysql -u username -pPASSWORD blog_db < database/seed/insert_users.sql