module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/go-redis/redis/v7 v7.4.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.9.1-0.20200903224024-3eb3b664a832
	go.uber.org/zap v1.18.1
	google.golang.org/protobuf v1.25.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common
