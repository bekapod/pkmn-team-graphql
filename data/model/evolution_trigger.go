package model

import (
	"fmt"
	"io"
	"strconv"
)

type EvolutionTrigger string

const (
	LevelUp EvolutionTrigger = "level-up"
	Other   EvolutionTrigger = "other"
	Shed    EvolutionTrigger = "shed"
	Trade   EvolutionTrigger = "trade"
	UseItem EvolutionTrigger = "use-item"
)

func (e EvolutionTrigger) IsValid() bool {
	switch e {
	case LevelUp, Other, Shed, Trade, UseItem:
		return true
	}
	return false
}

func (e EvolutionTrigger) String() string {
	return string(e)
}

func (e *EvolutionTrigger) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EvolutionTrigger(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EvolutionTrigger", str)
	}
	return nil
}

func (e EvolutionTrigger) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *EvolutionTrigger) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		*e = EvolutionTrigger(v)
		if !e.IsValid() {
			return fmt.Errorf("%s is not a valid EvolutionTrigger", src)
		}
		return nil
	case nil:
		return nil
	}

	return fmt.Errorf("failed to scan evolution trigger")
}

func (EvolutionTrigger) IsEntity() {}
