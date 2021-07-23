package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Timestamp struct {
	value time.Time
}

func NewTimestamp(timestamp time.Time) (*Timestamp, error) {

	newObj := &Timestamp{
		value: timestamp,
	}

	return newObj, nil

}

func NewNowTimestamp() (*Timestamp, error) {

	var err error
	var loc *time.Location

	loc, err = time.LoadLocation("Asia/Jakarta")

	if err != nil {
		return nil, err
	}

	newObj := &Timestamp{
		value: time.Now().In(loc),
	}

	return newObj, nil

}

func (obj *Timestamp) GetValue() time.Time {
	return obj.value
}

func (obj *Timestamp) SetValue(timestamp time.Time) {
	obj.value = timestamp
}

func (obj *Timestamp) GetString() string {
	return obj.value.Format("2006-01-02 15:04:05")
}

func (obj *Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.GetString())
}

func (obj *Timestamp) UnmarshalJSON(data []byte) error {

	var err error
	var value string

	err = json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	obj.value, err = time.Parse("2006-01-02 15:04:05", value)

	return err

}

func (obj *Timestamp) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	var val time.Time
	var ok bool

	val, ok = value.(time.Time)

	if !ok {
		return fmt.Errorf("Column is not a timestamp")
	}

	obj.value = val

	return nil

}

func (obj Timestamp) Value() (driver.Value, error) {

	var emptyTimestamp time.Time = time.Time{}

	if obj.value == emptyTimestamp {
		return nil, nil

	}
	return obj.GetValue(), nil
}
