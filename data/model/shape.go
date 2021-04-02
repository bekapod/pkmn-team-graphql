package model

import (
	"fmt"
	"io"
	"strconv"
)

type Shape string

const (
	Ball      Shape = "ball"
	Squiggle  Shape = "squiggle"
	Fish      Shape = "fish"
	Arms      Shape = "arms"
	Blob      Shape = "blob"
	Upright   Shape = "upright"
	Legs      Shape = "legs"
	Quadruped Shape = "quadruped"
	Wings     Shape = "wings"
	Tentacles Shape = "tentacles"
	Heads     Shape = "heads"
	Humanoid  Shape = "humanoid"
	BugWings  Shape = "bug-wings"
	Armor     Shape = "armor"
)

func (s Shape) IsValid() bool {
	switch s {
	case Ball, Squiggle, Fish, Arms, Blob, Upright, Legs, Quadruped, Wings, Tentacles, Heads, Humanoid, BugWings, Armor:
		return true
	}
	return false
}

func (s Shape) String() string {
	return string(s)
}

func (s *Shape) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*s = Shape(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid Shape", str)
	}
	return nil
}

func (s Shape) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(s.String()))
}

func (Shape) IsEntity() {}
