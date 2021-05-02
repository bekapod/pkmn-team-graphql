package model

import "bekapod/pkmn-team-graphql/data/db"

func NewItemFromDb(dbItem db.ItemModel) Item {
	i := Item{
		ID:       dbItem.ID,
		Slug:     dbItem.Slug,
		Name:     dbItem.Name,
		Category: ItemCategory(dbItem.Category),
	}

	if cost, ok := dbItem.Cost(); ok {
		i.Cost = &cost
	} else {
		i.Cost = nil
	}

	if flingPower, ok := dbItem.FlingPower(); ok {
		i.FlingPower = &flingPower
	} else {
		i.FlingPower = nil
	}

	if flingEffect, ok := dbItem.FlingEffect(); ok {
		i.FlingEffect = &flingEffect
	} else {
		i.FlingEffect = nil
	}

	if value, ok := dbItem.Effect(); ok {
		i.Effect = &value
	} else {
		i.Effect = nil
	}

	if sprite, ok := dbItem.Sprite(); ok {
		i.Sprite = &sprite
	} else {
		i.Sprite = nil
	}

	attributes := make([]ItemAttribute, 0)

	for _, a := range dbItem.Attributes {
		attribute := ItemAttribute(a)
		attributes = append(attributes, attribute)
	}

	i.Attributes = attributes

	return i
}
