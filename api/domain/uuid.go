package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// UUID is an Object to handle uuid format checking
type UUID struct {
	value string
}

// NewUUID is an Constructor for UUID Object generating UUID from Google UUID
func NewUUID() *UUID {
	return &UUID{
		value: uuid.New().String(),
	}
}

// NewUUIDFromString is an Constructor for UUID Object from String
func NewUUIDFromString(uuid string) (*UUID, error) {

	uuidObj := &UUID{
		value: uuid,
	}

	if !uuidObj.validateUUID() {
		return nil, fmt.Errorf("UUID Format Invalid")
	}

	return uuidObj, nil

}

// GetValue is a Getter Function for Value
func (obj *UUID) GetValue() string {
	return obj.value
}

func (obj *UUID) validateUUID() bool {
	_, err := uuid.Parse(obj.value)

	if err != nil {
		return false
	}

	return true

}

func (obj *UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.value)
}

func (obj *UUID) UnmarshalJSON(data []byte) error {

	var err error
	err = json.Unmarshal(data, &obj.value)

	if err != nil {
		return err
	}

	if !obj.validateUUID() {
		return fmt.Errorf("UUID Format Invalid")
	}

	return nil

}

func (obj *UUID) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	var byteValue []byte
	var ok bool

	byteValue, ok = value.([]byte)

	if !ok {
		return fmt.Errorf("Column is not a string")
	}

	obj.value = string(byteValue)

	if !obj.validateUUID() {
		return fmt.Errorf("UUID Format Invalid")
	}

	return nil

}

func (obj UUID) Value() (driver.Value, error) {
	if len(obj.value) == 0 {
		return nil, nil
	}
	return string(obj.GetValue()), nil
}
