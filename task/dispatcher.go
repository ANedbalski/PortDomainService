package task

import (
	"context"
	"ports/domain/services"
)

type Dispatcher struct {
	portImportService *services.PortImport
}

func NewDispatcher(piService *services.PortImport) *Dispatcher {
	return &Dispatcher{
		portImportService: piService,
	}
}

func (d *Dispatcher) Sub() {

}

func (d *Dispatcher) Run(ctx context.Context) {
	go func() {
		d.run(ctx)
	}()
}

func (d *Dispatcher) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
