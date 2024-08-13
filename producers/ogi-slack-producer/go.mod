module ogi-slack-producer

go 1.22.6

require (
	github.com/OpenChaos/ogi v0.0.0-20240812160549-cb906bb70707
	github.com/gol-gol/golenv v0.0.0-20230302172901-210791b57f21
	github.com/slack-go/slack v0.13.1
)

require (
	github.com/gol-gol/golconv v0.0.0-20230302172921-7c155f24578e // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.24.0 // indirect
)

replace github.com/OpenChaos/ogi v0.0.0-20240812160549-cb906bb70707 => ../../../ogi/
