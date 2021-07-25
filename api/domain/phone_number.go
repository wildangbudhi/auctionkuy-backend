package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"auctionkuy.wildangbudhi.com/utils"
)

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(phoneNumber string) (*PhoneNumber, error) {

	phoneNumberObj := &PhoneNumber{
		value: phoneNumber,
	}

	if !phoneNumberObj.validateFormat() {
		return phoneNumberObj, fmt.Errorf("Phone number format invalid")
	}

	return phoneNumberObj, nil

}

func (obj *PhoneNumber) GetValue() string {
	return obj.value
}

func (obj *PhoneNumber) validateFormat() bool {
	re := regexp.MustCompile(`\+[0-9]+`)
	return re.MatchString(obj.value)
}

func (obj *PhoneNumber) SanitizePhoneNumber(countryMap *utils.CountryDict, userLocale string) error {

	reg, err := regexp.Compile("[^0-9]+")

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Failed to process phone number")
	}

	var phone string

	phone = reg.ReplaceAllString(obj.value, "")

	var isPhoneFormatValid bool
	isPhoneFormatValid = countryMap.PhoneNumberPrefixChecker.ValidateStringPrefix(phone)

	if isPhoneFormatValid {
		phone = "+" + phone
		obj.value = phone
		return nil
	}

	if phone[0] == '0' {
		phone = phone[1:]
	}

	var localePrefix string
	var isLocaleExist bool

	localePrefix, isLocaleExist = countryMap.PhoneNumberMaps[userLocale]

	if !isLocaleExist {
		return fmt.Errorf("User locale invalid")
	}

	phone = "+" + localePrefix + phone
	obj.value = phone

	return nil

}

func (obj *PhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.value)
}

func (obj *PhoneNumber) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &obj.value)
}

func (obj *PhoneNumber) Scan(value interface{}) error {

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

	if !obj.validateFormat() {
		return fmt.Errorf("Phone number format invalid")
	}

	return nil

}

func (obj PhoneNumber) Value() (driver.Value, error) {
	if len(obj.value) == 0 {
		return nil, nil
	}
	return string(obj.GetValue()), nil
}
