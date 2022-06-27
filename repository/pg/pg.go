package pg

import "ports/domain"

type Port struct {
}

func NewPort() (*Port, error) {
	return &Port{}, nil
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
