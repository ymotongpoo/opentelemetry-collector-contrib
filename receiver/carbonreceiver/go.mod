module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opencensus.io v0.23.0
	go.opentelemetry.io/collector v0.9.1-0.20200903224024-3eb3b664a832
	go.uber.org/zap v1.16.0
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
	google.golang.org/protobuf v1.25.0
)
