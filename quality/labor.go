package quality

import (
	"strings"

	"github.com/richardwilkes/construction/fxp"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/txt"
)

type LaborQuality Quality

const (
	Unskilled = LaborQuality(Cheap)
	Skilled   = LaborQuality(Normal)
	Masterful = LaborQuality(Fine)
)

var LaborQualities []LaborQuality

func (q LaborQuality) EnsureValid() LaborQuality {
	if q <= LaborQualities[len(LaborQualities)-1] {
		return q
	}
	return 0
}

func (q LaborQuality) String() string {
	return qualities[Quality(q.EnsureValid())].LaborName
}

func (q LaborQuality) Key() string {
	return strings.ReplaceAll(txt.ToSnakeCase(q.String()), " ", "")
}

func (q LaborQuality) CFAdjustment() fxp.Int {
	return Quality(q.EnsureValid()).CFAdjustment()
}

func (q LaborQuality) HTAdjustment() fxp.Int {
	return Quality(q.EnsureValid()).HTAdjustment()
}

func (q LaborQuality) DRMultiplier() fxp.Int {
	return Quality(q.EnsureValid()).DRMultiplier()
}

func (q LaborQuality) HPMultiplier() fxp.Int {
	return Quality(q.EnsureValid()).HPMultiplier()
}

func (q LaborQuality) MarshalText() ([]byte, error) {
	return []byte(q.Key()), nil
}

func (q *LaborQuality) UnmarshalText(text []byte) error {
	s := string(text)
	for _, k := range LaborQualities {
		if strings.EqualFold(s, k.Key()) {
			*q = k
			return nil
		}
	}
	return errs.Newf("invalid Labor Quality: %q", s)
}

func init() {
	LaborQualities = make([]LaborQuality, len(Qualities))
	for i, k := range Qualities {
		LaborQualities[i] = LaborQuality(k)
	}
}
