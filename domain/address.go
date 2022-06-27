package domain

import "time"

type Address struct {
	Country  Country
	City     string
	Province string
	Timezone time.Location
	GPS      GPS
}
