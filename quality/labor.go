package quality

import (
	"github.com/richardwilkes/gcs/v5/model/fxp"
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

func init() {
	LaborQualities = make([]LaborQuality, len(Qualities))
	for i, k := range Qualities {
		LaborQualities[i] = LaborQuality(k)
	}
}
