package model

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	ID          string          `json:"id"`
	Slug        string          `json:"slug"`
	Name        string          `json:"name"`
	Cost        int             `json:"cost"`
	FlingPower  int             `json:"fling_power"`
	FlingEffect string          `json:"fling_effect"`
	Effect      string          `json:"effect"`
	Sprite      string          `json:"sprite"`
	Category    ItemCategory    `json:"category"`
	Attributes  []ItemAttribute `json:"attributes"`
}

func (l *Item) Scan(src interface{}) error {
	switch v := src.(type) {
	case []uint8:
		err := json.Unmarshal([]byte(v), &l)
		return err
	}

	return fmt.Errorf("failed to scan item")
}

func (Item) IsEntity() {}
