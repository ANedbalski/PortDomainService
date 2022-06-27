package services

import (
	"io"
	"ports/domain"
	"ports/repository/pg"
)

type PortImportConfiguration func(os *PortImport) error

type PortImport struct {
	ports domain.PortRepository
}

//NewPortImport takes a variable amount of PortImportConfiguration functions and returns a new PortImport
// Each PortImportConfiguration will be called in the order they are passed in
func NewPortImport(cfgs ...PortImportConfiguration) (*PortImport, error) {
	pi := &PortImport{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(pi)
		if err != nil {
			return nil, err
		}
	}

	return pi, nil
}

func WithInMemoryPortImportRepository(r domain.PortRepository) PortImportConfiguration {
	return func(pi *PortImport) error {
		pi.ports = r
		return nil
	}
}

func WithPGPortImportRepository(dsn string) PortImportConfiguration {
	return func(pi *PortImport) error {
		// call to create PG connection from DSN
		p, err := pg.New()
		if err != nil {
			return err
		}
		pi.ports = p
		return nil
	}
}

func (pi *PortImport) Import(in io.Reader) error {
	return nil
}
