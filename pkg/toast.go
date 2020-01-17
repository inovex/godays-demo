package pkg

import "time"

type Toast struct {
	Name    string       `json:"name"`
	Weekday time.Weekday `json:"weekday"`
}
