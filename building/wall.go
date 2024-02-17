package building

import (
	"fmt"
	"math"

	"github.com/richardwilkes/construction/fxp"
	"github.com/richardwilkes/construction/labor"
	"github.com/richardwilkes/construction/material"
	"github.com/richardwilkes/construction/quality"
)

type Wall struct {
	Name            string             `yaml:"name"`
	Length          fxp.Length         `yaml:"length"`
	Height          fxp.Length         `yaml:"height"`
	Thickness       fxp.Length         `yaml:"thickness"`
	Material        material.Type      `yaml:"material"`
	MaterialQuality quality.Quality    `yaml:"material_quality"`
	Labor           map[labor.Type]int `yaml:"labor"`
	Cleanup         bool               `yaml:"cleanup"`
}

func (w *Wall) Cost() fxp.Int {
	cost := toFeet(w.Length).Mul(toFeet(w.Height)).Mul(fxp.Int(w.Thickness)).Mul(w.Material.CostPerInch(w.Thickness)).Mul(w.MaterialQuality.CFAdjustment() + w.OverallLaborQuality().CFAdjustment() + fxp.One)
	if w.Cleanup {
		cost = cost.Mul(fxp.From(0.05))
	}
	return cost.Ceil()

}

func (w *Wall) Weight() fxp.Weight {
	return fxp.Weight(toFeet(w.Length).Mul(toFeet(w.Height)).Mul(fxp.Int(w.Thickness)).Mul(w.Material.WeightPerInch()).Ceil())
}

func (w *Wall) DR() int {
	base := fxp.Int(w.Thickness).Mul(w.Material.DRPerInch())
	return fxp.As[int](base.Mul(w.OverallLaborQuality().DRMultiplier()).Mul(w.MaterialQuality.DRMultiplier()).Round())
}

func (w *Wall) HT() int {
	base := fxp.Int(w.Thickness).Mul(w.Material.DRPerInch())
	return fxp.As[int](base + w.OverallLaborQuality().HTAdjustment() + w.MaterialQuality.HTAdjustment())
}

func (w *Wall) HP() int {
	return int(math.Round(math.Cbrt(fxp.As[float64](w.Material.WeightPerInch().Mul(fxp.Int(w.Thickness)).Div(fxp.Hundred))) * 80 * fxp.As[float64](w.MaterialQuality.HPMultiplier()) * fxp.As[float64](w.OverallLaborQuality().HPMultiplier())))
}

func (w *Wall) DaysToBuild() int {
	cost := w.Cost()
	var wages fxp.Int
	for l, qty := range w.Labor {
		if qty > 0 {
			wages += l.MonthlyWage().Mul(fxp.From(qty))
		}
	}
	return fxp.As[int](cost.Div(wages.Div(fxp.From(28))).Ceil())
}

func (w *Wall) OverallLaborQuality() quality.LaborQuality {
	total := 0
	count := 0
	for l, qty := range w.Labor {
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

func (w *Wall) String() string {
	if w.Cleanup {
		return fmt.Sprintf("%s (%v long, %v high, %v thick %s Debris; $%s, %s, %d days with %s labor)",
			w.Name,
			w.Length,
			w.Height,
			w.Thickness,
			w.Material,
			w.Cost().Comma(),
			w.Weight(),
			w.DaysToBuild(),
			w.OverallLaborQuality(),
		)
	}
	return fmt.Sprintf("%s (%v long, %v high, %v thick %s Wall; $%s, %s, DR %d, HP %d, HT %d, %d days with %s labor and %s materials)",
		w.Name,
		w.Length,
		w.Height,
		w.Thickness,
		w.Material,
		w.Cost().Comma(),
		w.Weight(),
		w.DR(),
		w.HP(),
		w.HT(),
		w.DaysToBuild(),
		w.OverallLaborQuality(),
		w.MaterialQuality,
	)
}

func toFeet(length fxp.Length) fxp.Int {
	return fxp.Int(length).Div(fxp.Twelve)
}
