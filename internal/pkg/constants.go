package pkg

import (
	"github.com/cthulhu-platform/common/pkg/env"
)

const (
	STORAGE_PATH = "./storage"
)

var (
	// AMQP config
	AMQP_USER  = env.GetEnv("AMQP_USER", "guest")
	AMQP_PASS  = env.GetEnv("AMQP_PASS", "guest")
	AMQP_HOST  = env.GetEnv("AMQP_HOST", "localhost")
	AMQP_PORT  = env.GetEnv("AMQP_PORT", "5672")
	AMQP_VHOST = env.GetEnv("AMQP_VHOST", "/")
)
