package model

import (
	"fmt"
	"io"
	"strconv"
)

type Habitat string

const (
	Cave         Habitat = "cave"
	Forest       Habitat = "forest"
	Grassland    Habitat = "grassland"
	Mountain     Habitat = "mountain"
	Rare         Habitat = "rare"
	RoughTerrain Habitat = "rough-terrain"
	Sea          Habitat = "sea"
	Urban        Habitat = "urban"
	WatersEdge   Habitat = "waters-edge"
)

func (h Habitat) IsValid() bool {
	switch h {
	case Cave, Forest, Grassland, Mountain, Rare, RoughTerrain, Sea, Urban, WatersEdge:
		return true
	}
	return false
}

func (h Habitat) String() string {
	return string(h)
}

func (h *Habitat) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*h = Habitat(str)
	if !h.IsValid() {
		return fmt.Errorf("%s is not a valid Habitat", str)
	}
	return nil
}

func (h Habitat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(h.String()))
}

func (Habitat) IsEntity() {}
