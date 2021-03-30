package model

import (
	"fmt"
	"io"
	"strconv"
)

type DamageClass string

const (
	Physical DamageClass = "physical"
	Special  DamageClass = "special"
	Status   DamageClass = "status"
)

func (dc DamageClass) IsValid() bool {
	switch dc {
	case Physical, Special, Status:
		return true
	}
	return false
}

func (dc DamageClass) String() string {
	return string(dc)
}

func (dc *DamageClass) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*dc = DamageClass(str)
	if !dc.IsValid() {
		return fmt.Errorf("%s is not a valid DamageClass", str)
	}
	return nil
}

func (dc DamageClass) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(dc.String()))
}

func (DamageClass) IsEntity() {}
