package model

import (
	"encoding/json"
	"errors"
)

type DamageClass string

const (
	Special  DamageClass = "special"
	Physical DamageClass = "physical"
	Status   DamageClass = "status"
)

func (dc *DamageClass) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	damageClass := DamageClass(s)
	switch damageClass {
	case Special, Physical, Status:
		*dc = damageClass
		return nil
	}
	return errors.New("invalid damage class")
}

func (DamageClass) IsEntity() {}
