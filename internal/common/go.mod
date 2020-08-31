module github.com/open-telemetry/opentelemetry-collector-contrib/internal/common

go 1.14

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/containerd/containerd v1.3.6 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20200514230353-811a247d06e8+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/gotestyourself/gotestyourself v1.4.0 // indirect
	github.com/imdario/mergo v0.3.10 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/collector v0.8.1-0.20200815205113-8e5c6065eb0e
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gotest.tools v1.4.0 // indirect
	k8s.io/client-go v0.19.0
)

// Yet another hack that we need until kubernetes client moves to the new github.com/googleapis/gnostic
replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
