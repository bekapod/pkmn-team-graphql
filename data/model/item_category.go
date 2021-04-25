package model

import (
	"fmt"
	"io"
	"strconv"
)

type ItemCategory string

const (
	AllMachines     ItemCategory = "all-machines"
	AllMail         ItemCategory = "all-mail"
	ApricornBalls   ItemCategory = "apricorn-balls"
	ApricornBox     ItemCategory = "apricorn-box"
	BadHeldItems    ItemCategory = "bad-held-items"
	BakingOnly      ItemCategory = "baking-only"
	Choice          ItemCategory = "choice"
	Collectibles    ItemCategory = "collectibles"
	DataCards       ItemCategory = "data-cards"
	DexCompletion   ItemCategory = "dex-completion"
	EffortDrop      ItemCategory = "effort-drop"
	EffortTraining  ItemCategory = "effort-training"
	EventItems      ItemCategory = "event-items"
	Evolution       ItemCategory = "evolution"
	Flutes          ItemCategory = "flutes"
	Gameplay        ItemCategory = "gameplay"
	Healing         ItemCategory = "healing"
	HeldItems       ItemCategory = "held-items"
	InAPinch        ItemCategory = "in-a-pinch"
	Jewels          ItemCategory = "jewels"
	Loot            ItemCategory = "loot"
	Medicine        ItemCategory = "medicine"
	MegaStones      ItemCategory = "mega-stones"
	Memories        ItemCategory = "memories"
	MiracleShooter  ItemCategory = "miracle-shooter"
	Mulch           ItemCategory = "mulch"
	OtherCategory   ItemCategory = "other"
	PickyHealing    ItemCategory = "picky-healing"
	Plates          ItemCategory = "plates"
	PlotAdvancement ItemCategory = "plot-advancement"
	PPRecovery      ItemCategory = "pp-recovery"
	Revival         ItemCategory = "revival"
	Scarves         ItemCategory = "scarves"
	SpecialBalls    ItemCategory = "special-balls"
	SpeciesSpecific ItemCategory = "species-specific"
	Spelunking      ItemCategory = "spelunking"
	StandardBalls   ItemCategory = "standard-balls"
	StatBoosts      ItemCategory = "stat-boosts"
	StatusCures     ItemCategory = "status-cures"
	Training        ItemCategory = "training"
	TypeEnhancement ItemCategory = "type-enhancement"
	TypeProtection  ItemCategory = "type-protection"
	Unused          ItemCategory = "unused"
	Vitamins        ItemCategory = "vitamins"
	ZCrystals       ItemCategory = "z-crystals"
)

func (c ItemCategory) IsValid() bool {
	switch c {
	case AllMachines, AllMail, ApricornBalls, ApricornBox, BadHeldItems, BakingOnly, Choice, Collectibles, DataCards, DexCompletion, EffortDrop, EffortTraining, EventItems, Evolution, Flutes, Gameplay, Healing, HeldItems, InAPinch, Jewels, Loot, Medicine, MegaStones, Memories, MiracleShooter, Mulch, OtherCategory, PickyHealing, Plates, PlotAdvancement, PPRecovery, Revival, Scarves, SpecialBalls, SpeciesSpecific, Spelunking, StandardBalls, StatBoosts, StatusCures, Training, TypeEnhancement, TypeProtection, Unused, Vitamins, ZCrystals:
		return true
	}
	return false
}

func (c ItemCategory) String() string {
	return string(c)
}

func (c *ItemCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*c = ItemCategory(str)
	if !c.IsValid() {
		return fmt.Errorf("%s is not a valid ItemCategory", str)
	}
	return nil
}

func (c ItemCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(c.String()))
}

func (c *ItemCategory) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		*c = ItemCategory(v)
		if !c.IsValid() {
			return fmt.Errorf("%s is not a valid ItemCategory", src)
		}
		return nil
	case nil:
		return nil
	}

	return fmt.Errorf("failed to scan item category")
}

func (ItemCategory) IsEntity() {}
