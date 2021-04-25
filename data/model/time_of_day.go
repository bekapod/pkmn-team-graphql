package model

import (
	"fmt"
	"io"
	"strconv"
)

type TimeOfDay string

const (
	Any   TimeOfDay = "any"
	Day   TimeOfDay = "day"
	Night TimeOfDay = "night"
)

func (t TimeOfDay) IsValid() bool {
	switch t {
	case Any, Day, Night:
		return true
	}
	return false
}

func (t TimeOfDay) String() string {
	return string(t)
}

func (t *TimeOfDay) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = TimeOfDay(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid TimeOfDay", str)
	}
	return nil
}

func (t TimeOfDay) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}

func (t *TimeOfDay) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		*t = TimeOfDay(v)
		if !t.IsValid() {
			return fmt.Errorf("%s is not a valid TimeOfDay", src)
		}
		return nil
	case nil:
		return nil
	}

	return fmt.Errorf("failed to scan time of day")
}

func (TimeOfDay) IsEntity() {}
