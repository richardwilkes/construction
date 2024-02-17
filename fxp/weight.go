/*
 * Copyright ©1998-2023 by Richard A. Wilkes. All rights reserved.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, version 2.0. If a copy of the MPL was not distributed with
 * this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * This Source Code Form is "Incompatible With Secondary Licenses", as
 * defined by the Mozilla Public License, version 2.0.
 */

package fxp

import (
	"strings"
)

// Weight contains a fixed-point value in pounds.
type Weight Int

// WeightFromString creates a new Weight. May have any of the known Weight suffixes or no notation at all, in which case
// defaultUnits is used.
func WeightFromString(text string, defaultUnits WeightUnit) (Weight, error) {
	text = strings.TrimLeft(strings.TrimSpace(text), "+")
	for _, unit := range WeightUnits {
		if strings.HasSuffix(text, unit.Key()) {
			value, err := FromString(strings.TrimSpace(strings.TrimSuffix(text, unit.Key())))
			if err != nil {
				return 0, err
			}
			return Weight(unit.ToPounds(value)), nil
		}
	}
	// No matches, so let's use our passed-in default units
	value, err := FromString(strings.TrimSpace(text))
	if err != nil {
		return 0, err
	}
	return Weight(defaultUnits.ToPounds(value)), nil
}

func (w Weight) String() string {
	return Pound.Format(w)
}

func (w Weight) MarshalText() ([]byte, error) {
	return []byte(w.String()), nil
}

func (w *Weight) UnmarshalText(in []byte) error {
	var err error
	*w, err = WeightFromString(string(in), Pound)
	return err
}
