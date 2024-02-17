package building

import (
	"fmt"
	"math"

	"github.com/richardwilkes/construction/fxp"
	"github.com/richardwilkes/construction/labor"
	"github.com/richardwilkes/construction/material"
	"github.com/richardwilkes/construction/quality"
)

type Building struct {
	Name            string             `yaml:"name"`
	Length          fxp.Length         `yaml:"length"`
	Width           fxp.Length         `yaml:"width"`
	Height          fxp.Length         `yaml:"height"`
	WallThickness   fxp.Length         `yaml:"wall_thickness"`
	PartitionFactor fxp.Int            `yaml:"partition_factor"`
	Material        material.Type      `yaml:"material"`
	MaterialQuality quality.Quality    `yaml:"material_quality"`
	Labor           map[labor.Type]int `yaml:"labor"`
}

func (b *Building) volume() fxp.Int {
	return toFeet(b.Length).Mul(toFeet(b.Width)).Mul(toFeet(b.Height))
}

func (b *Building) Cost() fxp.Int {
	return b.volume().Mul(b.PartitionFactor).Mul(fxp.Int(b.WallThickness)).Mul(b.Material.CostPerInch(b.WallThickness)).Mul(b.MaterialQuality.CFAdjustment() + b.OverallLaborQuality().CFAdjustment() + fxp.One).Ceil()
}

func (b *Building) Weight() fxp.Weight {
	return fxp.Weight(b.volume().Mul(b.PartitionFactor).Mul(fxp.Int(b.WallThickness)).Mul(b.Material.WeightPerInch()).Ceil())
}

func (b *Building) DR() int {
	base := fxp.Int(b.WallThickness).Mul(b.Material.DRPerInch())
	return fxp.As[int](base.Mul(b.OverallLaborQuality().DRMultiplier()).Mul(b.MaterialQuality.DRMultiplier()).Round())
}

func (b *Building) HT() int {
	base := fxp.Int(b.WallThickness).Mul(b.Material.DRPerInch())
	return fxp.As[int](base + b.OverallLaborQuality().HTAdjustment() + b.MaterialQuality.HTAdjustment())
}

func (b *Building) HP() int {
	return int(math.Round(math.Cbrt(fxp.As[float64](b.Material.WeightPerInch().Mul(fxp.Int(b.WallThickness)).Div(fxp.Hundred))) * 80 * fxp.As[float64](b.MaterialQuality.HPMultiplier()) * fxp.As[float64](b.OverallLaborQuality().HPMultiplier())))
}

func (b *Building) DaysToBuild() int {
	cost := b.Cost()
	var wages fxp.Int
	for l, qty := range b.Labor {
		if qty > 0 {
			wages += l.MonthlyWage().Mul(fxp.From(qty))
		}
	}
	return fxp.As[int](cost.Div(wages.Div(fxp.From(28))).Ceil())
}

func (b *Building) OverallLaborQuality() quality.LaborQuality {
	total := 0
	count := 0
	for l, qty := range b.Labor {
		if qty > 0 {
			count += qty
			total += int(l.Quality()) * qty
		}
	}
	if count <= 0 {
		return quality.Unskilled
	}
	q := float64(total) / float64(count)
	if q <= 0.5 {
		return quality.Unskilled
	}
	if q <= 1.5 {
		return quality.Skilled
	}
	return quality.Masterful
}

func (b *Building) String() string {
	return fmt.Sprintf("%s (%v long, %v wide, %v high, %v thick %s Wall; $%s, %s, DR %d, HP %d, HT %d, %d days with %s labor and %s materials)",
		b.Name,
		b.Length,
		b.Width,
		b.Height,
		b.WallThickness,
		b.Material,
		b.Cost().Comma(),
		b.Weight(),
		b.DR(),
		b.HP(),
		b.HT(),
		b.DaysToBuild(),
		b.OverallLaborQuality(),
		b.MaterialQuality,
	)
}
