package main

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

type Alarm struct {
	alarmType string
}

func (m *Mail) onRecv() {
	m.listener.OnFire()
}
