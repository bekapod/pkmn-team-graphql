package model

import (
	"fmt"
	"io"
	"strconv"
)

type ItemAttribute string

const (
	Consumable      ItemAttribute = "consumable"
	Countable       ItemAttribute = "countable"
	Holdable        ItemAttribute = "holdable"
	HoldableActive  ItemAttribute = "holdable-active"
	HoldablePassive ItemAttribute = "holdable-passive"
	Underground     ItemAttribute = "underground"
	UsableInBattle  ItemAttribute = "usable-in-battle"
	UsableOverworld ItemAttribute = "usable-overworld"
)

func (a ItemAttribute) IsValid() bool {
	switch a {
	case Consumable, Countable, Holdable, HoldableActive, HoldablePassive, Underground, UsableInBattle, UsableOverworld:
		return true
	}
	return false
}

func (a ItemAttribute) String() string {
	return string(a)
}

func (a *ItemAttribute) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*a = ItemAttribute(str)
	if !a.IsValid() {
		return fmt.Errorf("%s is not a valid ItemAttribute", str)
	}
	return nil
}

func (a ItemAttribute) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(a.String()))
}

func (a *ItemAttribute) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		*a = ItemAttribute(v)
		if !a.IsValid() {
			return fmt.Errorf("%s is not a valid ItemAttribute", src)
		}
		return nil
	case nil:
		return nil
	}

	return fmt.Errorf("failed to scan item attribute")
}

func (ItemAttribute) IsEntity() {}
