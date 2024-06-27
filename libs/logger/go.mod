module github.com/urodstvo/messenger-server/libs/logger

go 1.22.4

replace github.com/urodstvo/messenger-server/libs/config => ../../libs/config

require github.com/urodstvo/messenger-server/libs/config v0.0.0-00010101000000-000000000000

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
)
