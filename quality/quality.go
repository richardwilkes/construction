package quality

import (
	"cmp"
	"slices"
	"strings"

	"github.com/richardwilkes/construction/fxp"
	"github.com/richardwilkes/toolbox/collection/dict"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/txt"
)

type Quality byte

type quality struct {
	MaterialName string
	LaborName    string
	CFAdjustment fxp.Int
	HTAdjustment fxp.Int
	DRMultiplier fxp.Int
	HPMultiplier fxp.Int
}

const (
	Cheap Quality = iota
	Normal
	Fine
)

var Qualities []Quality

func (q Quality) EnsureValid() Quality {
	if q <= Qualities[len(Qualities)-1] {
		return q
	}
	return 0
}

func (q Quality) String() string {
	return qualities[q.EnsureValid()].MaterialName
}

func (q Quality) Key() string {
	return strings.ReplaceAll(txt.ToSnakeCase(q.String()), " ", "")
}

func (q Quality) CFAdjustment() fxp.Int {
	return qualities[q.EnsureValid()].CFAdjustment
}

func (q Quality) HTAdjustment() fxp.Int {
	return qualities[q.EnsureValid()].HTAdjustment
}

func (q Quality) DRMultiplier() fxp.Int {
	return qualities[q.EnsureValid()].DRMultiplier
}

func (q Quality) HPMultiplier() fxp.Int {
	return qualities[q.EnsureValid()].HPMultiplier
}

func (q Quality) MarshalText() ([]byte, error) {
	return []byte(q.Key()), nil
}

func (q *Quality) UnmarshalText(text []byte) error {
	s := string(text)
	for _, k := range Qualities {
		if strings.EqualFold(s, k.Key()) {
			*q = k
			return nil
		}
	}
	return errs.Newf("invalid Quality: %q", s)
}

func init() {
	keys := dict.Keys(qualities)
	slices.SortFunc(keys, cmp.Compare[Quality])
	Qualities = make([]Quality, len(keys))
	LaborQualities = make([]LaborQuality, len(Qualities))
	for i, k := range keys {
		Qualities[i] = k
		LaborQualities[i] = LaborQuality(k)
	}
}

var qualities = map[Quality]quality{
	Fine: {
		MaterialName: "Fine",
		LaborName:    "Masterful",
		CFAdjustment: fxp.Quarter,
		HTAdjustment: fxp.One,
		DRMultiplier: fxp.One,
		HPMultiplier: fxp.One,
	},
	Normal: {
		MaterialName: "Normal",
		LaborName:    "Skilled",
		DRMultiplier: fxp.One,
		HPMultiplier: fxp.One,
	},
	Cheap: {
		MaterialName: "Cheap",
		LaborName:    "Unskilled",
		CFAdjustment: -fxp.Fifth,
		HTAdjustment: -fxp.One,
		DRMultiplier: fxp.FromStringForced("0.95"),
		HPMultiplier: fxp.FromStringForced("0.95"),
	},
}
