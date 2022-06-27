package valueobject

type GPS struct {
	Lat  float64
	Long float64
}

func NewGPS(lat, long float64) GPS {
	return GPS{
		Lat:  lat,
		Long: long,
	}
}
