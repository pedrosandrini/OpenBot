createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrateup:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(openbot_mysql:3306)/chatservice" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(openbot_mysql:3306)/chatservice" -verbose  drop

grpc:
	protoc --go_out=. --go-grpc_out=. proto/chat.proto --experimental_allow_proto3_optional

.PHONY: migrate createmigration migratedown grpc