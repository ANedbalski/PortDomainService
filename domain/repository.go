package domain

import "errors"

var (
	ErrPortNotFound = errors.New("port with this id wasn't found")
)

type PortRepository interface {
	GetById(id uint64) (*Port, error)
	Save(*Port) error
	Add(port *Port) error
	UpdateOrCreate(*Port) error
}
