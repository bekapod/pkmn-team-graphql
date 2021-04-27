package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestRegion_Scan(t *testing.T) {
	exp := Region{
		ID:   "1823b8bc-2a3a-473d-9baa-277be67427bc",
		Name: "Galar",
		Slug: "galar",
	}
	got := Region{}
	err := (&got).Scan([]uint8{123, 34, 105, 100, 34, 58, 32, 34, 49, 102, 48, 57, 53, 56, 97, 48, 45, 52, 56, 99, 97, 45, 52, 49, 54, 48, 45, 57, 102, 49, 56, 45, 55, 101, 53, 102, 48, 54, 100, 57, 54, 100, 50, 55, 34, 44, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 70, 97, 105, 114, 121, 34, 44, 32, 34, 115, 108, 117, 103, 34, 58, 32, 34, 102, 97, 105, 114, 121, 34, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestRegion_Scan_Error(t *testing.T) {
	got := Region{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestRegion_Scan_RegionError(t *testing.T) {
	got := Region{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
