package task

import (
	"context"
	"go.uber.org/zap"
	"ports/domain/services"
)

type ChanDispatcher struct {
	portImportService *services.PortImport
	taskChan          chan *NewPortsListFile
	log               *zap.SugaredLogger
}

func NewDispatcher(piService *services.PortImport, log *zap.SugaredLogger) *ChanDispatcher {
	return &ChanDispatcher{
		portImportService: piService,
		taskChan:          make(chan *NewPortsListFile),
		log:               log,
	}
}

func (d *ChanDispatcher) Sub() {

}

func (d *ChanDispatcher) Pub(t *NewPortsListFile) {
	d.taskChan <- t
}

func (d *ChanDispatcher) Run(ctx context.Context) {
	go func() {
		d.run(ctx)
	}()
}

func (d *ChanDispatcher) run(ctx context.Context) {
	for {
		select {
		case t := <-d.taskChan:
			err := d.portImportService.Import(t.src)
			if err != nil {
				d.log.Errorw("error in importing data ", "err", err)
			}
		case <-ctx.Done():
			close(d.taskChan)
			return
		}
	}
}
