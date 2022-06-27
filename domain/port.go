package domain

import "ports/domain/valueobject"

type Port struct {
	// Port code is used as ID of port
	Code    string
	Name    string
	Address valueobject.Address
}
