module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/golang/protobuf v1.4.2
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opencensus.io v0.22.4
	go.opentelemetry.io/collector v0.8.1-0.20200815205113-8e5c6065eb0e
	go.uber.org/zap v1.16.0
	google.golang.org/grpc/examples v0.0.0-20200728194956-1c32b02682df // indirect
)
