package services

import (
	"ports/domain"
	"ports/repository/pg"
)

type PortConfiguration func(os *Port) error

type Port struct {
	ports domain.PortRepository
}

//NewPort takes a variable amount of PortImport functions and returns a new Port
// Each PortImport will be called in the order they are passed in
func NewPort(cfgs ...PortConfiguration) (*Port, error) {
	p := &Port{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(p)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func WithInMemoryPortRepository(r domain.PortRepository) PortConfiguration {
	return func(pi *Port) error {
		pi.ports = r
		return nil
	}
}

func WithPGPortRepository(dsn string) PortConfiguration {
	return func(pi *Port) error {
		// call to create PG connection from DSN
		p, err := pg.New()
		if err != nil {
			return err
		}
		pi.ports = p
		return nil
	}
}
