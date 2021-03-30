package model

import (
	"testing"
)

type TestDataDamageClass struct {
	input DamageClass
	exp   interface{}
}

type TestDataUnmarshal struct {
	input    DamageClass
	hasError bool
}

type TestDataMarshal struct {
	input DamageClass
	exp   string
}

func TestDamageClass_IsValid(t *testing.T) {
	damageClasses := []TestDataDamageClass{
		{Physical, true},
		{Special, true},
		{Status, true},
		{"Something else", false},
	}

	for _, item := range damageClasses {
		if got := item.input.IsValid(); item.exp != got {
			t.Errorf("expected '%t' but got '%t' instead", item.exp, got)
		}
	}
}

func TestDamageClass_String(t *testing.T) {
	damageClasses := []TestDataDamageClass{
		{Physical, "physical"},
		{Special, "special"},
		{Status, "status"},
	}

	for _, item := range damageClasses {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}

func TestDamageClass_UnmarshalGQL(t *testing.T) {
	damageClasses := []TestDataUnmarshal{
		{Physical, false},
		{Special, false},
		{Status, false},
		{"Something else", true},
	}

	for _, item := range damageClasses {
		if got := (&item.input).UnmarshalGQL(item.input.String()); item.hasError && got == nil {
			t.Error("expected an error but didn't get one")
		}
	}
}

func TestDamageClass_MarshalGQL(t *testing.T) {
	damageClasses := []TestDataMarshal{
		{Physical, "physical"},
		{Special, "special"},
		{Status, "status"},
		{"Something else", "Something else"},
	}

	for _, item := range damageClasses {
		if got := item.input.String(); item.exp != got {
			t.Errorf("expected '%s' but got '%s' instead", item.exp, got)
		}
	}
}
