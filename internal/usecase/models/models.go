package models

type Car struct {
	Id          string
	Brand       string
	Model       string
	Type        string
	Speed       uint32
	Seats       uint32
	Color       string
	Description string
}

type Update struct {
	Color       string
	Description string
}
