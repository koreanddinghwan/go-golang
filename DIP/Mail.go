package main

type Mail struct {
	alarm Alarm
}

type Alarm struct {
	alarmType string
}

func (m *Mail) onRecv() {
	m.alarm.alarmType = "mail"
}
