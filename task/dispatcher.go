package task

import (
	"context"
)

type Dispatcher struct {
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
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
