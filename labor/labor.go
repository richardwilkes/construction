package labor

import (
	"cmp"
	"slices"

	"github.com/richardwilkes/construction/quality"
	"github.com/richardwilkes/gcs/v5/model/fxp"
	"github.com/richardwilkes/toolbox/collection/dict"
)

type Type byte

type labor struct {
	Name        string
	Quality     quality.LaborQuality
	MonthlyWage fxp.Int
}

const (
	Architect Type = iota
	BuildingLaborer
	Carpenter
	Mason
	MasterCarpenter
	MasterMason
	MasterShipwright
	MasterSmith
	Miner
	Shipwright
	Smith
)

var Types []Type

func (t Type) EnsureValid() Type {
	if t <= Types[len(Types)-1] {
		return t
	}
	return 0
}

func (t Type) String() string {
	return types[t.EnsureValid()].Name
}

func (t Type) Quality() quality.LaborQuality {
	return types[t.EnsureValid()].Quality
}

func (t Type) MonthlyWage() fxp.Int {
	return types[t.EnsureValid()].MonthlyWage
}

func init() {
	keys := dict.Keys(types)
	slices.SortFunc(keys, cmp.Compare[Type])
	Types = make([]Type, len(keys))
	for i, k := range keys {
		Types[i] = k
	}
}

var types = map[Type]labor{
	Architect: {
		Name:        "Architect",
		Quality:     quality.Masterful,
		MonthlyWage: fxp.From(4000),
	},
	BuildingLaborer: {
		Name:        "Building Laborer",
		Quality:     quality.Unskilled,
		MonthlyWage: fxp.From(400),
	},
	Carpenter: {
		Name:        "Carpenter",
		Quality:     quality.Skilled,
		MonthlyWage: fxp.From(790),
	},
	MasterCarpenter: {
		Name:        "Master Carpenter",
		Quality:     quality.Masterful,
		MonthlyWage: fxp.From(1580),
	},
	Mason: {
		Name:        "Mason",
		Quality:     quality.Skilled,
		MonthlyWage: fxp.From(900),
	},
	MasterMason: {
		Name:        "Master Mason",
		Quality:     quality.Masterful,
		MonthlyWage: fxp.From(1800),
	},
	Miner: {
		Name:        "Miner",
		Quality:     quality.Skilled,
		MonthlyWage: fxp.From(420),
	},
	Shipwright: {
		Name:        "Shipwright",
		Quality:     quality.Skilled,
		MonthlyWage: fxp.From(850),
	},
	MasterShipwright: {
		Name:        "Master Shipwright",
		Quality:     quality.Masterful,
		MonthlyWage: fxp.From(1700),
	},
	Smith: {
		Name:        "Smith",
		Quality:     quality.Skilled,
		MonthlyWage: fxp.From(900),
	},
	MasterSmith: {
		Name:        "Master Smith",
		Quality:     quality.Masterful,
		MonthlyWage: fxp.From(1800),
	},
}
