package memory

import (
	"ports/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPort_GetById(t *testing.T) {
	testCases := []struct {
		name   string
		id     uint64
		exp    *domain.Port
		expErr error
	}{
		{
			name:   "port doesn't exists in the DB",
			id:     5,
			exp:    nil,
			expErr: ErrPortNotFound,
		},
		{
			name:   "port exists in the DB",
			id:     1,
			exp:    &domain.Port{ID: 1, Name: "A1"},
			expErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			port, err := db.GetById(tt.id)

			assert.Equal(t, tt.exp, port)
			assert.Equal(t, tt.expErr, err)
		})
	}
}

func TestPort_Add(t *testing.T) {
	testCases := []struct {
		name   string
		port   *domain.Port
		expErr error
	}{
		{
			name:   "validate no error",
			port:   &domain.Port{ID: 10, Name: "A10"},
			expErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			err := db.Add(tt.port)

			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.port, db.ports[tt.port.ID])
		})
	}
}

func TestPort_Save(t *testing.T) {
	testCases := []struct {
		name   string
		port   *domain.Port
		expErr error
	}{
		{
			name:   "Port exists in DB and record updated",
			port:   &domain.Port{ID: 1, Name: "A1111"},
			expErr: nil,
		},
		{
			name:   "Port don't exists in DB return error",
			port:   &domain.Port{ID: 10, Name: "A10"},
			expErr: ErrPortNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			err := db.Save(tt.port)

			assert.Equal(t, tt.expErr, err)
			if err == nil {
				assert.Equal(t, db.ports[tt.port.ID], tt.port)
			}
		})
	}
}

func TestPort_UpdateOrCreate(t *testing.T) {
	testCases := []struct {
		name string
		port *domain.Port
	}{
		{
			name: "update existing port",
			port: &domain.Port{ID: 1, Name: "A1111"},
		},
		{
			name: "add new port",
			port: &domain.Port{ID: 10, Name: "A10"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			err := db.UpdateOrCreate(tt.port)

			if !assert.Nil(t, err) {
				return
			}
			assert.Equal(t, db.ports[tt.port.ID], tt.port)
		})
	}

}

func initDBWithFixtures() *Port {
	db := &Port{
		ports: map[uint64]*domain.Port{
			1: &domain.Port{ID: 1, Name: "A1"},
			2: &domain.Port{ID: 2, Name: "A2"},
			3: &domain.Port{ID: 3, Name: "A3"},
		},
	}
	return db
}
