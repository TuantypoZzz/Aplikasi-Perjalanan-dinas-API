package enum

import (
	"errors"
)

type Gender struct {
	value string
}

const (
	MALE   string = "Laki-Laki"
	FEMALE string = "Perempuan"
)

func NewGender(val string) (Gender, error) {
	gender := Gender{value: val}
	if val != "" {
		if !gender.isAllowedGender() {
			return Gender{}, errors.New("unsupport gender ")
		}
	}
	return gender, nil
}

func (gender Gender) isAllowedGender() bool {
	allowedGenders := []string{MALE, FEMALE}
	for _, allowedGender := range allowedGenders {

		if allowedGender == gender.value {
			return true
		}
	}
	return false
}

func (gender Gender) String() string {
	g := ""
	switch gender.value {
	case MALE:
		g = "Laki-Laki"
	case FEMALE:
		g = "Perempuan"
	}
	return g
}

// func (gender *Gender) Scan(value interface{}) error {
// 	*gender = Gender(value.([]byte))
// 	return nil
// }

// func (gender Gender) Value() (driver.Value, error) {
// 	return string(gender), nil
// }
