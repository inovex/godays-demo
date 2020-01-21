package pkg

import "time"

type Toast struct {
	Name    string       `json:"name"`
	Weekday time.Weekday `json:"weekday"`
}

func GetToasts() []Toast {
	return toasts
}

var toasts = []Toast{
	{
		Name:    "Hawaii",
		Weekday: time.Monday,
	},
	{
		Name:    "Peperoni",
		Weekday: time.Tuesday,
	},
	{
		Name:    "Cheese",
		Weekday: time.Wednesday,
	},
	{
		Name:    "Ham",
		Weekday: time.Thursday,
	},
	{
		Name:    "Caprese",
		Weekday: time.Friday,
	},
	{
		Name:    "Avocado",
		Weekday: time.Saturday,
	},
	{
		Name:    "Honey",
		Weekday: time.Sunday,
	},
}
