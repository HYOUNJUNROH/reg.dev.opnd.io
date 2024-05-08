package pqcustom

import (
	"database/sql/driver"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(src interface{}) error {
	return pq.Array(a).Scan(src)
}
func (a UUIDArray) Value() (driver.Value, error) {
	return pq.Array(a).Value()
}

func (a *UUIDArray) StringArray() []string {
	var s []string
	for _, v := range *a {
		s = append(s, v.String())
	}
	return s
}

func StringArrayToUUIDArray(s []string) (UUIDArray, error) {
	var a UUIDArray
	for _, v := range s {
		uuidVal, err := uuid.FromString(v)
		if err != nil {
			return nil, err
		}
		a = append(a, uuidVal)
	}
	return a, nil
}
