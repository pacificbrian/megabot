module github.com/pacificbrian/megabot

go 1.18

require (
	gobot.io/x/gobot/v2 v2.1.1
	google.golang.org/protobuf v1.28.1
)

replace gobot.io/x/gobot/v2 => ../gobot

require (
	github.com/creack/goselect v0.1.2 // indirect
	github.com/eclipse/paho.mqtt.golang v1.4.2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.bug.st/serial v1.5.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
