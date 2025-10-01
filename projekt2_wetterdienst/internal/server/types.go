package server

import (
	"time"
)

type TempMessage struct {
	Temp int       `json:"temp"`
	Time time.Time `json:"time"`
}
