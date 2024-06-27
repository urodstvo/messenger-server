module github.com/urodstvo/messenger-server/apps/api

go 1.22.4

replace (
	github.com/urodstvo/messenger-server/libs/app => ../../libs/app
	github.com/urodstvo/messenger-server/libs/config => ../../libs/config
	github.com/urodstvo/messenger-server/libs/grpc => ../../libs/grpc
	github.com/urodstvo/messenger-server/libs/logger => ../../libs/logger
)

require (
	github.com/urodstvo/messenger-server/libs/app v0.0.0-00010101000000-000000000000
	github.com/urodstvo/messenger-server/libs/config v0.0.0-00010101000000-000000000000
	github.com/urodstvo/messenger-server/libs/grpc v0.0.0-20240615113456-7cfbe55c0ba9
	github.com/urodstvo/messenger-server/libs/logger v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.22.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gorm.io/driver/postgres v1.5.9 // indirect
	gorm.io/gorm v1.25.10 // indirect
)
