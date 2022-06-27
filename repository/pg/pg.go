package pg

import "ports/domain"

type Port struct {
}

func New() *Port {
	return &Port{}
}

func (p *Port) GetById(id uint64) (*domain.Port, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Port) Save(port *domain.Port) error {
	//TODO implement me
	panic("implement me")
}

func (p *Port) Add(port *domain.Port) error {
	//TODO implement me
	panic("implement me")
}

func (p *Port) UpdateOrCreate(port *domain.Port) error {
	//TODO implement me
	panic("implement me")
}
