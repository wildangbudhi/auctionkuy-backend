package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
)

// Email is an Object to handle email format checking
type Email struct {
	value string
}

// NewEmail is an Constructor for Email Object
func NewEmail(email string) (*Email, error) {

	emailObj := &Email{
		value: email,
	}

	if !emailObj.validateEmail() {
		return nil, fmt.Errorf("Email Format Invalid")
	}

	return emailObj, nil
}

// GetValue is a Getter Function for Value
func (obj *Email) GetValue() string {
	return obj.value
}

func (obj *Email) validateEmail() bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(obj.value)
}

func (obj *Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.value)
}

func (obj *Email) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &obj.value)
}

func (obj *Email) Scan(value interface{}) error {

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

	if !obj.validateEmail() {
		return fmt.Errorf("Email Format Invalid")
	}

	return nil

}

func (obj Email) Value() (driver.Value, error) {
	if len(obj.value) == 0 {
		return nil, nil
	}
	return string(obj.GetValue()), nil
}
