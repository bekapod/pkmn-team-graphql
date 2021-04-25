package model

import (
	"fmt"
	"io"
	"strconv"
)

type Gender string

const (
	Female  Gender = "female"
	Male    Gender = "male"
	Unknown Gender = "unknown"
)

func (g Gender) IsValid() bool {
	switch g {
	case Female, Male, Unknown:
		return true
	}
	return false
}

func (g Gender) String() string {
	return string(g)
}

func (g *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*g = Gender(str)
	if !g.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (g Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(g.String()))
}

func (g *Gender) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		*g = Gender(v)
		if !g.IsValid() {
			return fmt.Errorf("%s is not a valid Gender", src)
		}
		return nil
	case nil:
		return nil
	}

	return fmt.Errorf("failed to scan gender")
}

func (Gender) IsEntity() {}
