package model

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Region Region `json:"region"`
}

func (l *Location) Scan(src interface{}) error {
	switch v := src.(type) {
	case []uint8:
		err := json.Unmarshal([]byte(v), &l)
		return err
	}

	return fmt.Errorf("failed to scan location")
}

func (Location) IsEntity() {}
