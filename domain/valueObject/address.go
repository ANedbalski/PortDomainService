package valueObject

import (
	"ports/domain/entity"
	"time"
)

type Address struct {
	Country  entity.Country
	City     string
	Province string
	Timezone time.Location
	GPS      GPS
}
