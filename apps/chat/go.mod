module github.com/urodstvo/messenger-server/apps/chat

go 1.22.4

replace (
	github.com/urodstvo/messenger-server/libs/app => ../../libs/app
	github.com/urodstvo/messenger-server/libs/config => ../../libs/config
	github.com/urodstvo/messenger-server/libs/grpc => ../../libs/grpc
	github.com/urodstvo/messenger-server/libs/jwt => ../../libs/jwt
	github.com/urodstvo/messenger-server/libs/logger => ../../libs/logger
)

require (
	github.com/olahol/melody v1.2.1
	github.com/urodstvo/messenger-server/libs/app v0.0.0-00010101000000-000000000000
	github.com/urodstvo/messenger-server/libs/config v0.0.0-00010101000000-000000000000
	github.com/urodstvo/messenger-server/libs/grpc v0.0.0-20240615113456-7cfbe55c0ba9
	github.com/urodstvo/messenger-server/libs/jwt v0.0.0-00010101000000-000000000000
	github.com/urodstvo/messenger-server/libs/logger v0.0.0-00010101000000-000000000000
	github.com/urodstvo/messenger-server/libs/models v0.0.0-20240615113456-7cfbe55c0ba9
	go.uber.org/fx v1.22.1
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
	gorm.io/gorm v1.25.10
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/redis/go-redis/v9 v9.5.3 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	gorm.io/driver/postgres v1.5.9 // indirect
)
