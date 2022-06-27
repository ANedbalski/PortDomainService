package memory

import (
	"ports/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPort_GetById(t *testing.T) {
	testCases := []struct {
		name   string
		code   string
		exp    *domain.Port
		expErr error
	}{
		{
			name:   "port doesn't exists in the DB",
			code:   "5",
			exp:    nil,
			expErr: domain.ErrPortNotFound,
		},
		{
			name:   "port exists in the DB",
			code:   "1",
			exp:    &domain.Port{Code: "1", Name: "A1"},
			expErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			port, err := db.GetById(tt.code)

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
			port:   &domain.Port{Code: "10", Name: "A10"},
			expErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			err := db.Add(tt.port)

			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.port, db.ports[tt.port.Code])
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
			port:   &domain.Port{Code: "1", Name: "A1111"},
			expErr: nil,
		},
		{
			name:   "Port don't exists in DB return error",
			port:   &domain.Port{Code: "10", Name: "A10"},
			expErr: domain.ErrPortNotFound,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			err := db.Save(tt.port)

			assert.Equal(t, tt.expErr, err)
			if err == nil {
				assert.Equal(t, db.ports[tt.port.Code], tt.port)
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
			port: &domain.Port{Code: "1", Name: "A1111"},
		},
		{
			name: "add new port",
			port: &domain.Port{Code: "10", Name: "A10"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := initDBWithFixtures()

			err := db.UpdateOrCreate(tt.port)

			if !assert.Nil(t, err) {
				return
			}
			assert.Equal(t, db.ports[tt.port.Code], tt.port)
		})
	}

}

func initDBWithFixtures() *Port {
	db := &Port{
		ports: map[string]*domain.Port{
			"1": {Code: "1", Name: "A1"},
			"2": {Code: "2", Name: "A2"},
			"3": {Code: "3", Name: "A3"},
		},
	}
	return db
}
