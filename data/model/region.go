package model

import (
	"encoding/json"
	"fmt"
)

type Region struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func (r *Region) Scan(src interface{}) error {
	switch v := src.(type) {
	case []uint8:
		err := json.Unmarshal([]byte(v), &r)
		return err
	}

	return fmt.Errorf("failed to scan region")
}

func (Region) IsEntity() {}
