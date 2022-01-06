// Copyright 2021 The libcalcal author. All rights reserved.
// Use of this source code is governed by the license that
// can be found in the file LICENSE.md.

package main

import (
	"fmt"

	lc "staudtlex.de/libcalendar"
)

// convert from a Gregorian date to the corresponding date of the Julian, ISO,
// Islamic, Hebrew, French Revolutionary, Mayan (long count, haab, tzolkin),
// and Old Hindu (solar, lunar) calendars.
func main() {
	// 1. set Gregorian date
	gregorianDate := lc.GregorianDate{2022, 01, 06}
	// 2. convert Gregorian date to absolute (fixed) date
	absoluteDate := lc.AbsoluteFromGregorian(gregorianDate)
	// 3. convert absolute dates into corresponding calendar dates
	fmt.Println("Converting from a Gregorian date to other calendar dates")
	fmt.Println("Gregorian:\t\t", gregorianDate)
	fmt.Println("Julian:\t\t\t", lc.JulianFromAbsolute(absoluteDate))
	fmt.Println("ISO:\t\t\t", lc.IsoFromAbsolute(absoluteDate))
	fmt.Println("Islamic:\t\t", lc.IslamicFromAbsolute(absoluteDate))
	fmt.Println("Hebrew:\t\t\t", lc.HebrewFromAbsolute(absoluteDate))
	fmt.Println("French Revolutionary:\t", lc.FrenchFromAbsolute(absoluteDate))
	fmt.Println("Mayan Long Count:\t", lc.MayanLongCountFromAbsolute(absoluteDate))
	fmt.Println("Mayan Haab:\t\t", lc.MayanHaabFromAbsolute(absoluteDate))
	fmt.Println("Mayan Tzolkin:\t\t", lc.MayanTzolkinFromAbsolute(absoluteDate))
	fmt.Println("Hindu Solar:\t\t", lc.OldHinduSolarFromAbsolute(absoluteDate))
	fmt.Println("Hindu Lunar:\t\t", lc.OldHinduLunarFromAbsolute(absoluteDate))
}
