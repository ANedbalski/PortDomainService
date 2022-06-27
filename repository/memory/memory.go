package memory

import (
	"ports/domain"
	"sync"
)

type Port struct {
	ports map[uint64]*domain.Port
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *Port {
	return &Port{
		ports: make(map[uint64]*domain.Port),
	}
}

func (p *Port) GetById(id uint64) (*domain.Port, error) {
	if port, ok := p.ports[id]; ok {
		return port, nil
	}
	return nil, domain.ErrPortNotFound
}

func (p *Port) Save(port *domain.Port) error {
	if _, ok := p.ports[port.ID]; ok {
		p.Lock()
		p.ports[port.ID] = port
		p.Unlock()
		return nil
	}
	return domain.ErrPortNotFound
}

func (p *Port) UpdateOrCreate(port *domain.Port) error {
	err := p.Save(port)
	if err == domain.ErrPortNotFound {
		return p.Add(port)
	}
	return err
}

func (p *Port) Add(port *domain.Port) error {
	p.Lock()
	p.ports[port.ID] = port
	p.Unlock()
	return nil
}
