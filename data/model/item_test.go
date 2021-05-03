package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"testing"

	"github.com/go-test/deep"
)

func TestNewItemFromDb_WithNulls(t *testing.T) {
	item := db.ItemModel{
		InnerItem: db.InnerItem{
			ID:         "123",
			Slug:       "some-item",
			Name:       "Some Item",
			Category:   db.ItemCategoryCOLLECTIBLES,
			Attributes: []db.ItemAttribute{},
		},
	}
	exp := Item{
		ID:          item.ID,
		Slug:        item.Slug,
		Name:        item.Name,
		Cost:        nil,
		FlingPower:  nil,
		FlingEffect: nil,
		Effect:      nil,
		Sprite:      nil,
		Category:    ItemCategoryCollectibles,
		Attributes:  []ItemAttribute{},
	}

	got := NewItemFromDb(item)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewItemFromDb_WithFullData(t *testing.T) {
	cost := 3000
	flingPower := 80
	flingEffect := "Some effect when flung"
	effect := "Some item effect"
	sprite := "item.png"
	item := db.ItemModel{
		InnerItem: db.InnerItem{
			ID:          "123",
			Slug:        "some-item",
			Name:        "Some Item",
			Cost:        &cost,
			FlingPower:  &flingPower,
			FlingEffect: &flingEffect,
			Effect:      &effect,
			Sprite:      &sprite,
			Category:    db.ItemCategoryCOLLECTIBLES,
			Attributes:  []db.ItemAttribute{db.ItemAttributeCONSUMABLE, db.ItemAttributeCOUNTABLE},
		},
	}
	exp := Item{
		ID:          item.ID,
		Slug:        item.Slug,
		Name:        item.Name,
		Cost:        &cost,
		FlingPower:  &flingPower,
		FlingEffect: &flingEffect,
		Effect:      &effect,
		Sprite:      &sprite,
		Category:    ItemCategoryCollectibles,
		Attributes:  []ItemAttribute{ItemAttributeConsumable, ItemAttributeCountable},
	}

	got := NewItemFromDb(item)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}
