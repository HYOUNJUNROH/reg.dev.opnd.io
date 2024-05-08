package service

import (
	"database/sql/driver"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func GetOffset(size, page int) int {
	if page == -1 || size == -1 {
		return -1
	}
	return size * (page - 1)
}

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(src interface{}) error {
	return pq.Array(a).Scan(src)
}
func (a UUIDArray) Value() (driver.Value, error) {
	return pq.Array(a).Value()
}
