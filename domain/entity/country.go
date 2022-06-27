package entity

//Represents Country Entity. Supposed to be a table in the DB
//With own Repository and moved to aggregators
//In the current implementation it is skipped
type Country struct {
	ID           uint16
	Name         string
	Abbreviation string
	Code         uint16
}

func NewCountry(name string) Country {
	return Country{Name: name}
}
