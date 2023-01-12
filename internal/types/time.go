package types

import (
	"database/sql/driver"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return t.Time.MarshalJSON()
}

func (t *Time) UnmarshalJSON(data []byte) error {
	return t.Time.UnmarshalJSON(data)
}

func (t *Time) Value() (driver.Value, error) {
	return t.UnixMilli(), nil
}

func (t *Time) Scan(src interface{}) error {
	t.Time = time.UnixMilli(src.(int64))
	return nil
}

func NewTime(t time.Time) Time {
	return Time{t}
}
