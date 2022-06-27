package domain

import "ports/domain/valueObject"

type Port struct {
	// Port code is used as ID of port
	Code    string
	Name    string
	Address valueObject.Address
}
