package domain

type PortRepository interface {
	GetById(id uint64) (*Port, error)
	Save(*Port) error
	Add(port *Port) error
	UpdateOrCreate(*Port) error
}
