module github.com/Bubotka/Microservices/proxy

go 1.19

require github.com/Bubotka/Microservices/geo v0.0.0

require (
	github.com/go-chi/chi/v5 v5.0.11 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/ptflp/godecoder v0.0.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
)

replace github.com/Bubotka/Microservices/geo v0.0.0 => ../geo
