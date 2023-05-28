package connections

type Status int

const (
	Unknown Status = iota
	Inactive
	Active
)

type Connection struct {
	Name   string
	Status Status
}
