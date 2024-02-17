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
	"github.com/richardwilkes/toolbox/xmath"
	"github.com/richardwilkes/toolbox/xmath/fixed"
	"github.com/richardwilkes/toolbox/xmath/fixed/f64"
)

// Common values that can be reused.
var (
	Fifth             = FromStringForced("0.2")
	Quarter           = FromStringForced("0.25")
	ThreeTenths       = FromStringForced("0.3")
	Half              = FromStringForced("0.5")
	One               = From(1)
	Two               = From(2)
	Three             = From(3)
	Four              = From(4)
	Five              = From(5)
	Six               = From(6)
	Seven             = From(7)
	Eight             = From(8)
	Nine              = From(9)
	Twelve            = From(12)
	Thirteen          = From(13)
	Sixteen           = From(16)
	ThirtySix         = From(36)
	Hundred           = From(100)
	FiveHundred       = From(500)
	TwoThousand       = From(2000)
	ThirtySixThousand = From(36000)
	MileInInches      = From(63360)
)

// DP is an alias for the fixed-point decimal places configuration we are using.
type DP = fixed.D4

// Int is an alias for the fixed-point type we are using.
type Int = f64.Int[DP]

// From creates an Int from a numeric value.
func From[T xmath.Numeric](value T) Int {
	return f64.From[DP](value)
}

// FromString creates an Int from a string.
func FromString(value string) (Int, error) {
	return f64.FromString[DP](value)
}

// FromStringForced creates an Int from a string, ignoring any conversion inaccuracies.
func FromStringForced(value string) Int {
	return f64.FromStringForced[DP](value)
}

// As returns the equivalent value in the destination type.
func As[T xmath.Numeric](value Int) T {
	return f64.As[DP, T](value)
}
