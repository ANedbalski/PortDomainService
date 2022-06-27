package memory

import (
	"ports/domain"
	"sync"
)

type Port struct {
	ports map[string]*domain.Port
	sync.Mutex
}

func NewPort() *Port {
	return &Port{
		ports: make(map[string]*domain.Port),
	}
}

func (p *Port) GetById(id string) (*domain.Port, error) {
	if port, ok := p.ports[id]; ok {
		return port, nil
	}
	return nil, domain.ErrPortNotFound
}

func (p *Port) Save(port *domain.Port) error {
	if _, ok := p.ports[port.Code]; ok {
		p.Lock()
		p.ports[port.Code] = port
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
	p.ports[port.Code] = port
	p.Unlock()
	return nil
}
