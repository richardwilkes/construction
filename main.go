package main

import (
	"fmt"
	"os"

	"github.com/richardwilkes/construction/building"
	"github.com/richardwilkes/construction/fxp"
	"github.com/richardwilkes/construction/labor"
	"github.com/richardwilkes/construction/material"
	"github.com/richardwilkes/construction/quality"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/fatal"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Walls     []building.Wall
	Buildings []building.Building
}

func main() {
	cl := cmdline.New(true)
	var file, dump string
	cl.NewGeneralOption(&file).SetName("config").SetSingle('c').SetArg("file").SetUsage("The file to load a configuration from")
	cl.NewGeneralOption(&dump).SetName("dump").SetSingle('d').SetArg("file").SetUsage("The file to write the loaded configuration to")
	cl.Parse(os.Args[1:])
	var config Config
	if file == "" {
		fillDefaults(&config)
	} else {
		data, err := os.ReadFile(file)
		fatal.IfErr(err)
		fatal.IfErr(yaml.Unmarshal(data, &config))
	}
	var totalCost fxp.Int
	var totalDays int
	for _, w := range config.Walls {
		totalCost += w.Cost()
		totalDays += w.DaysToBuild()
		fmt.Println(w.String())
	}
	for _, b := range config.Buildings {
		totalCost += b.Cost()
		totalDays += b.DaysToBuild()
		fmt.Println(b.String())
	}
	fmt.Println("========================================")
	fmt.Printf("Total Cost: $%s\n", totalCost.Comma())
	fmt.Println("Total Days to Build:", totalDays)
	if dump != "" {
		data, err := yaml.Marshal(&config)
		fatal.IfErr(err)
		fatal.IfErr(os.WriteFile(dump, data, 0644))
	}
}

func fillDefaults(config *Config) {
	materialQuality := quality.Normal
	workers := map[labor.Type]int{
		labor.Architect:       1,
		labor.BuildingLaborer: 20,
		labor.MasterCarpenter: 3,
		labor.Carpenter:       30,
	}
	config.Walls = append(config.Walls, building.Wall{
		Name:            "Walls",
		Length:          fxp.LengthFromInteger(140*3+60*2, fxp.Feet),
		Height:          fxp.LengthFromInteger(10, fxp.Feet),
		Thickness:       fxp.LengthFromInteger(1, fxp.Feet),
		Material:        material.HardEarth,
		MaterialQuality: materialQuality,
		Labor:           workers,
	})
	for _, name := range []string{"Northeast", "Northwest", "Southeast", "Southwest"} {
		config.Buildings = append(config.Buildings, building.Building{
			Name:            name + " Corner Tower",
			Length:          fxp.LengthFromInteger(20, fxp.Feet),
			Width:           fxp.LengthFromInteger(20, fxp.Feet),
			Height:          fxp.LengthFromInteger(20, fxp.Feet),
			WallThickness:   fxp.LengthFromInteger(1, fxp.Feet),
			PartitionFactor: fxp.Quarter,
			Material:        material.Wood,
			MaterialQuality: materialQuality,
			Labor:           workers,
		})
	}
	config.Buildings = append(config.Buildings,
		building.Building{
			Name:            "Gatehouse",
			Length:          fxp.LengthFromInteger(20, fxp.Feet),
			Width:           fxp.LengthFromInteger(20, fxp.Feet),
			Height:          fxp.LengthFromInteger(20, fxp.Feet),
			WallThickness:   fxp.LengthFromInteger(1, fxp.Feet),
			PartitionFactor: fxp.ThreeTenths,
			Material:        material.Wood,
			MaterialQuality: materialQuality,
			Labor:           workers,
		},
		building.Building{
			Name:            "Stables",
			Length:          fxp.LengthFromInteger(60, fxp.Feet),
			Width:           fxp.LengthFromInteger(20, fxp.Feet),
			Height:          fxp.LengthFromInteger(15, fxp.Feet),
			WallThickness:   fxp.LengthFromInteger(6, fxp.Inch),
			PartitionFactor: fxp.ThreeTenths,
			Material:        material.Wood,
			MaterialQuality: materialQuality,
			Labor:           workers,
		},
		building.Building{
			Name:            "Barracks",
			Length:          fxp.LengthFromInteger(60, fxp.Feet),
			Width:           fxp.LengthFromInteger(25, fxp.Feet),
			Height:          fxp.LengthFromInteger(10, fxp.Feet),
			WallThickness:   fxp.LengthFromInteger(6, fxp.Inch),
			PartitionFactor: fxp.Quarter,
			Material:        material.Wood,
			MaterialQuality: materialQuality,
			Labor:           workers,
		},
		building.Building{
			Name:            "Officer's Quarters",
			Length:          fxp.LengthFromInteger(50, fxp.Feet),
			Width:           fxp.LengthFromInteger(30, fxp.Feet),
			Height:          fxp.LengthFromInteger(10, fxp.Feet),
			WallThickness:   fxp.LengthFromInteger(6, fxp.Inch),
			PartitionFactor: fxp.Half,
			Material:        material.Wood,
			MaterialQuality: materialQuality,
			Labor:           workers,
		},
	)
}
