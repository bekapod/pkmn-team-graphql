package model

import (
	"fmt"
	"io"
	"strconv"
)

type Color string

const (
	Black  Color = "black"
	Blue   Color = "blue"
	Brown  Color = "brown"
	Gray   Color = "gray"
	Green  Color = "green"
	Pink   Color = "pink"
	Purple Color = "purple"
	Red    Color = "red"
	White  Color = "white"
	Yellow Color = "yellow"
)

func (c Color) IsValid() bool {
	switch c {
	case Black, Blue, Brown, Gray, Green, Pink, Purple, Red, White, Yellow:
		return true
	}
	return false
}

func (c Color) String() string {
	return string(c)
}

func (c *Color) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*c = Color(str)
	if !c.IsValid() {
		return fmt.Errorf("%s is not a valid Color", str)
	}
	return nil
}

func (c Color) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(c.String()))
}

func (Color) IsEntity() {}
