package model

import (
	"testing"

	"github.com/go-test/deep"
)

func TestLocation_Scan(t *testing.T) {
	exp := Location{
		ID:   "1823b8bc-2a3a-473d-9baa-277be67427bc",
		Name: "Some location",
		Slug: "some-location",
		Region: Region{
			ID:   "1823b8bc-2a3a-473d-9baa-277be67427bc",
			Name: "Galar",
			Slug: "galar",
		},
	}
	got := Location{}
	err := (&got).Scan([]uint8{123, 10, 9, 9, 34, 105, 100, 34, 58, 32, 32, 32, 34, 49, 56, 50, 51, 98, 56, 98, 99, 45, 50, 97, 51, 97, 45, 52, 55, 51, 100, 45, 57, 98, 97, 97, 45, 50, 55, 55, 98, 101, 54, 55, 52, 50, 55, 98, 99, 34, 44, 10, 9, 9, 34, 110, 97, 109, 101, 34, 58, 32, 34, 83, 111, 109, 101, 32, 108, 111, 99, 97, 116, 105, 111, 110, 34, 44, 10, 9, 9, 34, 115, 108, 117, 103, 34, 58, 32, 34, 115, 111, 109, 101, 45, 108, 111, 99, 97, 116, 105, 111, 110, 34, 44, 10, 9, 9, 34, 114, 101, 103, 105, 111, 110, 34, 58, 32, 123, 10, 9, 9, 9, 34, 105, 100, 34, 58, 32, 32, 32, 34, 49, 56, 50, 51, 98, 56, 98, 99, 45, 50, 97, 51, 97, 45, 52, 55, 51, 100, 45, 57, 98, 97, 97, 45, 50, 55, 55, 98, 101, 54, 55, 52, 50, 55, 98, 99, 34, 44, 10, 9, 9, 9, 34, 110, 97, 109, 101, 34, 58, 32, 34, 71, 97, 108, 97, 114, 34, 44, 10, 9, 9, 9, 34, 115, 108, 117, 103, 34, 58, 32, 34, 103, 97, 108, 97, 114, 34, 10, 9, 9, 125, 10, 9, 125})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestLocation_Scan_Error(t *testing.T) {
	got := Location{}
	err := (&got).Scan([]uint8{123, 34, 115, 108, 111, 116, 34, 58, 32, 49, 44, 32, 34, 116, 121, 112, 101, 34, 58, 32, 123, 34, 105, 100, 34, 58, 32, 34, 48, 55, 98, 57, 101, 98, 125})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

func TestLocation_Scan_LocationError(t *testing.T) {
	got := Location{}
	err := (&got).Scan(5)

	if err == nil {
		t.Error("expected an error but got nil")
	}
}
