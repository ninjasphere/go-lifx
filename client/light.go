package client

import (
	"github.com/bjeanes/go-lifx/protocol/payloads"
)

type Light struct {
	client
}

func (l Light) TurnOff() {
	l.client.SendMessage(payloads.DeviceSetPower{Level: 0})
}

func (l Light) TurnOn() {
	l.client.SendMessage(payloads.DeviceSetPower{Level: 1})
}
