module github.com/austinrgray/signalcodex

go 1.22.2

require (
	github.com/nats-io/nats.go v1.38.0
	go.uber.org/zap v1.27.0
	google.golang.org/protobuf v1.36.2
)

require (
	github.com/google/uuid v1.6.0
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/nats-io/nkeys v0.4.9 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
)

retract (
	v0.1.0
	v0.1.1
	v0.0.1
)
