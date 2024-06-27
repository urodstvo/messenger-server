module github.com/urodstvo/messenger-server/libs/app

go 1.22.4

replace (
	github.com/urodstvo/messenger-server/libs/config => ../../libs/config
	github.com/urodstvo/messenger-server/libs/logger => ../../libs/logger
	github.com/urodstvo/messenger-server/libs/jwt => ../../libs/jwt
)

require (
	github.com/redis/go-redis/v9 v9.5.3
	github.com/urodstvo/messenger-server/libs/config v0.0.0-00010101000000-000000000000
	github.com/urodstvo/messenger-server/libs/logger v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.22.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.10
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
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
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
