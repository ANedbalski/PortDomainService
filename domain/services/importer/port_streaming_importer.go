package importer

import (
	"encoding/json"
	"fmt"
	"io"
	"ports/domain"
	"ports/domain/entity"
	"ports/domain/valueobject"
	"time"
)

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

// Entry represents each stream. If the stream fails, an error will be present.
type Entry struct {
	Err  error
	Port Port
}

// Stream helps transmit each streams withing a channel.
type Stream struct {
	stream chan Entry
}

func (p *Port) ToDomain() *domain.Port {
	// trying to find timezone. For simplicity at error skip this and use UTC
	timezone, err := time.LoadLocation(p.Timezone)
	if err != nil {
		timezone = time.UTC
	}
	return &domain.Port{
		Code: p.Code,
		Name: p.Name,
		Address: valueobject.NewAddress(entity.NewCountry(p.Country),
			p.City,
			p.Province,
			timezone,
			valueobject.NewGPS(p.Coordinates[0], p.Coordinates[1])),
	}
}

// NewJSONStream returns a new `Stream` type.
func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

// Watch watches JSON streams. Each stream entry will either have an error or a
// User object. Client code does not need to explicitly exit after catching an
// error as the `Start` method will close the channel automatically.
func (s Stream) Watch() <-chan Entry {
	return s.stream
}

// Start starts streaming JSON file line by line. If an error occurs, the channel
// will be closed.
func (s Stream) Start(src io.Reader) {
	// Stop streaming channel as soon as nothing left to read in the file.
	defer close(s.stream)

	decoder := json.NewDecoder(src)

	// Read opening delimiter. `[` or `{`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Err: fmt.Errorf("decode opening delimiter: %w", err)}
		return
	}

	// Read file content as long as there is something.
	i := 1
	for decoder.More() {
		var port Port
		if err := decoder.Decode(&port); err != nil {
			s.stream <- Entry{Err: fmt.Errorf("decode line %d: %w", i, err)}
			return
		}
		s.stream <- Entry{Port: port}

		i++
	}

	// Read closing delimiter. `]` or `}`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Err: fmt.Errorf("decode closing delimiter: %w", err)}
		return
	}
}
