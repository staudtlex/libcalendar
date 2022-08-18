// Copyright (C) 2022  Alexander Staudt
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package libcalendar_test

import (
	"fmt"

	lc "staudtlex.de/libcalendar"
)

// Convert from a Gregorian date to the corresponding date of the Julian, ISO,
// Islamic, Hebrew, French Revolutionary, Mayan (long count, haab, tzolkin),
// and Old Hindu (solar, lunar) calendars.
func Example() {
	// Convert dates
	// 1. set Gregorian date
	gregorianDate := lc.GregorianDate{Year: 2022, Month: 06, Day: 15}

	// 2. convert Gregorian date to absolute (fixed) date
	absoluteDate := lc.AbsoluteFromGregorian(gregorianDate)

	// 3. convert absolute dates into corresponding calendar dates
	fmt.Println("Converting from a Gregorian date to other calendar dates")
	fmt.Println("Gregorian:\t\t", gregorianDate)
	fmt.Println("Absolute date:\t\t", absoluteDate)
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

	// Utility functions
	// 1. Marshal a Date to JSON:
	gregorian_json := gregorianDate.Date().Json()
	fmt.Println(gregorian_json)

	// 2. Unmarshal JSON into a Date
	json_1 := `{
		"components": [2022,6,15],
		"calendar": "gregorian",
		"componentNames" :[],
		"monthNames": []
	}`
	fmt.Println(json_1)
	unmarshal_json_1 := lc.JsonToDate(json_1)
	fmt.Println(unmarshal_json_1)

	// 3. Convert a Date to an absolute date, and compute a Mayan Long count
	// date from that absolute date
	mayanLongCount :=
		lc.FromAbsolute(
			lc.AbsoluteFromDate(unmarshal_json_1), "mayanLongCount")
	fmt.Println("Mayan long count: ", mayanLongCount)
}
