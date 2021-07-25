package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Image struct {
	value  string
	prefix string
}

func NewImage(imagePath string, prefix *string) (*Image, error) {

	newObj := &Image{
		value: imagePath,
	}

	if prefix != nil {
		newObj.prefix = *prefix
	}

	return newObj, nil

}

func (obj *Image) GetValue() string {
	if obj.prefix == "" {
		return fmt.Sprintf("%s", obj.value)
	}

	return fmt.Sprintf("%s/%s", obj.prefix, obj.value)
}

func (obj *Image) SetPrefix(prefix string) {
	obj.prefix = prefix
}

func (obj *Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.GetValue())
}

func (obj *Image) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &obj.value)
}

func (obj *Image) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	var byteValue []byte
	var ok bool

	byteValue, ok = value.([]byte)

	if !ok {
		return fmt.Errorf("Column is not a string")
	}

	obj.prefix = ""
	obj.value = string(byteValue)

	return nil

}

func (obj Image) Value() (driver.Value, error) {
	if len(obj.value) == 0 {
		return nil, nil
	}
	return string(obj.GetValue()), nil
}
