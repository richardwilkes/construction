package main

import (
	"fmt"

	"github.com/richardwilkes/construction/building"
	"github.com/richardwilkes/construction/labor"
	"github.com/richardwilkes/construction/material"
	"github.com/richardwilkes/construction/quality"
	"github.com/richardwilkes/gcs/v5/model/fxp"
)

func main() {
	materialQuality := quality.Normal
	workers := map[labor.Type]int{
		labor.Architect:       1,
		labor.BuildingLaborer: 20,
		labor.Carpenter:       30,
	}
	w := building.Wall{
		Name:            "Walls",
		Length:          fxp.LengthFromInteger(140*3+60*2, fxp.Feet),
		Height:          fxp.LengthFromInteger(10, fxp.Feet),
		Thickness:       fxp.LengthFromInteger(1, fxp.Feet),
		Material:        material.Wood,
		MaterialQuality: materialQuality,
		Labor:           workers,
	}
	totalCost := w.Cost()
	totalDays := w.DaysToBuild()
	fmt.Println(w.String())

	for _, name := range []string{"Northeast", "Northwest", "Southeast", "Southwest"} {
		b := building.Building{
			Name:            name + " Corner Tower",
			Length:          fxp.LengthFromInteger(20, fxp.Feet),
			Width:           fxp.LengthFromInteger(20, fxp.Feet),
			Height:          fxp.LengthFromInteger(20, fxp.Feet),
			WallThickness:   fxp.LengthFromInteger(1, fxp.Feet),
			PartitionFactor: fxp.Quarter,
			Material:        material.Wood,
			MaterialQuality: materialQuality,
			Labor:           workers,
		}
		totalCost += b.Cost()
		totalDays += b.DaysToBuild()
		fmt.Println(b.String())
	}

	b := building.Building{
		Name:            "Gatehouse",
		Length:          fxp.LengthFromInteger(20, fxp.Feet),
		Width:           fxp.LengthFromInteger(20, fxp.Feet),
		Height:          fxp.LengthFromInteger(20, fxp.Feet),
		WallThickness:   fxp.LengthFromInteger(1, fxp.Feet),
		PartitionFactor: fxp.ThreeTenths,
		Material:        material.Wood,
		MaterialQuality: materialQuality,
		Labor:           workers,
	}
	totalCost += b.Cost()
	totalDays += b.DaysToBuild()
	fmt.Println(b.String())

	b = building.Building{
		Name:            "Stables",
		Length:          fxp.LengthFromInteger(60, fxp.Feet),
		Width:           fxp.LengthFromInteger(20, fxp.Feet),
		Height:          fxp.LengthFromInteger(15, fxp.Feet),
		WallThickness:   fxp.LengthFromInteger(6, fxp.Inch),
		PartitionFactor: fxp.ThreeTenths,
		Material:        material.Wood,
		MaterialQuality: materialQuality,
		Labor:           workers,
	}
	totalCost += b.Cost()
	totalDays += b.DaysToBuild()
	fmt.Println(b.String())

	b = building.Building{
		Name:            "Barracks",
		Length:          fxp.LengthFromInteger(60, fxp.Feet),
		Width:           fxp.LengthFromInteger(25, fxp.Feet),
		Height:          fxp.LengthFromInteger(10, fxp.Feet),
		WallThickness:   fxp.LengthFromInteger(6, fxp.Inch),
		PartitionFactor: fxp.Quarter,
		Material:        material.Wood,
		MaterialQuality: materialQuality,
		Labor:           workers,
	}
	totalCost += b.Cost()
	totalDays += b.DaysToBuild()
	fmt.Println(b.String())

	b = building.Building{
		Name:            "Officer's Quarters",
		Length:          fxp.LengthFromInteger(50, fxp.Feet),
		Width:           fxp.LengthFromInteger(30, fxp.Feet),
		Height:          fxp.LengthFromInteger(10, fxp.Feet),
		WallThickness:   fxp.LengthFromInteger(6, fxp.Inch),
		PartitionFactor: fxp.Half,
		Material:        material.Wood,
		MaterialQuality: materialQuality,
		Labor:           workers,
	}
	totalCost += b.Cost()
	totalDays += b.DaysToBuild()
	fmt.Println(b.String())

	fmt.Println("========================================")
	fmt.Printf("Total Cost: $%s\n", totalCost.Comma())
	fmt.Println("Total Days to Build:", totalDays)
}
