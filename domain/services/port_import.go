package services

import (
	"go.uber.org/zap"
	"io"
	"ports/domain"
	"ports/domain/services/importer"
	"ports/repository/pg"
)

type PortImportConfiguration func(os *PortImport) error

type Importer interface {
	Watch() <-chan importer.Entry
	Start(src io.Reader)
}

type PortImport struct {
	ports    domain.PortRepository
	importer Importer
	log      *zap.SugaredLogger
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
		p, err := pg.NewPort()
		if err != nil {
			return err
		}
		pi.ports = p
		return nil
	}
}

func WithStreamingPortImporter(i Importer) PortImportConfiguration {
	return func(pi *PortImport) error {
		pi.importer = i
		return nil
	}
}

func WithSugaredLogger(l *zap.SugaredLogger) PortImportConfiguration {
	return func(pi *PortImport) error {
		pi.log = l
		return nil
	}
}

func (pi *PortImport) Import(in io.Reader) error {
	go func() {
		for data := range pi.importer.Watch() {
			if data.Err != nil {
				pi.log.Errorw("error parsing json", "err", data.Err)
			}
			err := pi.ports.UpdateOrCreate(data.Port.ToDomain())
			if err != nil {
				pi.log.Errorw("error storing to DB", "err", err)
			}
		}
	}()
	pi.importer.Start(in)
	return nil
}
