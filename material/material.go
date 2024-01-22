package material

import (
	"cmp"
	"slices"

	"github.com/richardwilkes/gcs/v5/model/fxp"
	"github.com/richardwilkes/toolbox/collection/dict"
)

type Type byte

type material struct {
	Name          string
	CostPerInch   func(totalThickness fxp.Length) fxp.Int
	WeightPerInch fxp.Int
	DRPerInch     fxp.Int
}

const (
	Ashlar Type = iota
	Brick
	Concrete
	HardEarth
	Rubble
	Thatch
	Wood
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

func (t Type) CostPerInch(totalThickness fxp.Length) fxp.Int {
	return types[t.EnsureValid()].CostPerInch(totalThickness)
}

func (t Type) WeightPerInch() fxp.Int {
	return types[t.EnsureValid()].WeightPerInch
}

func (t Type) DRPerInch() fxp.Int {
	return types[t.EnsureValid()].DRPerInch
}

func init() {
	keys := dict.Keys(types)
	slices.SortFunc(keys, cmp.Compare[Type])
	Types = make([]Type, len(keys))
	for i, k := range keys {
		Types[i] = k
	}
}

func woodCost(totalThickness fxp.Length) fxp.Int {
	inches := fxp.Int(totalThickness)
	switch {
	case inches <= fxp.One:
		return fxp.FromStringForced("7.75")
	case inches <= fxp.Two:
		return fxp.FromStringForced("4.14")
	case inches <= fxp.Three:
		return fxp.FromStringForced("3.1")
	case inches <= fxp.Four:
		return fxp.FromStringForced("2.47")
	case inches <= fxp.Five:
		return fxp.FromStringForced("1.99")
	case inches <= fxp.Six:
		return fxp.FromStringForced("1.66")
	case inches <= fxp.Seven:
		return fxp.FromStringForced("1.43")
	default:
		return fxp.FromStringForced("1.26")
	}
}

var types = map[Type]material{
	Ashlar: {
		Name:          "Ashlar",
		CostPerInch:   func(_ fxp.Length) fxp.Int { return fxp.FromStringForced("11.72") },
		WeightPerInch: fxp.FromStringForced("15.5"),
		DRPerInch:     fxp.Thirteen,
	},
	Brick: {
		Name:          "Brick",
		CostPerInch:   func(_ fxp.Length) fxp.Int { return fxp.FromStringForced("3.34") },
		WeightPerInch: fxp.FromStringForced("7.7"),
		DRPerInch:     fxp.Eight,
	},
	Concrete: {
		Name:          "Concrete",
		CostPerInch:   func(_ fxp.Length) fxp.Int { return fxp.FromStringForced("9.98") },
		WeightPerInch: fxp.FromStringForced("15.5"),
		DRPerInch:     fxp.Nine,
	},
	HardEarth: {
		Name:          "Hard Earth",
		CostPerInch:   func(_ fxp.Length) fxp.Int { return fxp.FromStringForced("0.85") },
		WeightPerInch: fxp.FromStringForced("3.75"),
		DRPerInch:     fxp.One,
	},
	Rubble: {
		Name:          "Rubble",
		CostPerInch:   func(_ fxp.Length) fxp.Int { return fxp.FromStringForced("3.82") },
		WeightPerInch: fxp.From(14),
		DRPerInch:     fxp.Twelve,
	},
	Thatch: {
		Name:          "Thatch",
		CostPerInch:   func(_ fxp.Length) fxp.Int { return fxp.FromStringForced("0.74") },
		WeightPerInch: fxp.FromStringForced("1.3"),
		DRPerInch:     fxp.Half,
	},
	Wood: {
		Name:          "Wood",
		CostPerInch:   woodCost,
		WeightPerInch: fxp.FromStringForced("2.67"),
		DRPerInch:     fxp.One,
	},
}
