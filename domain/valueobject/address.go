package valueobject

import (
	"ports/domain/entity"
	"time"
)

type Address struct {
	Country  entity.Country
	City     string
	Province string
	Timezone *time.Location
	GPS      GPS
}

func NewAddress(c entity.Country, city, province string, timezone *time.Location, gps GPS) Address {
	return Address{
		Country:  c,
		City:     city,
		Province: province,
		Timezone: timezone,
		GPS:      gps,
	}
}
