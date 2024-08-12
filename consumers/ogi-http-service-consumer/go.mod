module ogi-http-service-consumer

go 1.22.6

require (
	github.com/OpenChaos/ogi v0.0.0-20240812160549-cb906bb70707
	github.com/gol-gol/golenv v0.0.0-20230302172901-210791b57f21
	github.com/oklog/ulid/v2 v2.1.0
	github.com/prometheus/client_golang v1.19.1
	github.com/pseidemann/finish v1.2.0
)

replace github.com/OpenChaos/ogi v0.0.0-20240812160549-cb906bb70707 => ../../../ogi/

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/gol-gol/golconv v0.0.0-20230302172921-7c155f24578e // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.24.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
