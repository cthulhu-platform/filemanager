module github.com/cthulhu-platform/filemanager

go 1.25.4

require (
	github.com/cthulhu-platform/common v0.0.0
	github.com/wagslane/go-rabbitmq v0.15.0
)

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
)

replace github.com/cthulhu-platform/common => ../../common
