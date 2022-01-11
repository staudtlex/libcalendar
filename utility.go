// Copyright (C) 2021  Alexander Staudt
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

package libcalendar

import "fmt"

// FromAbsolute converts a given absolute (fixed) date to the date
// representation specified in `calendar`.
//
// The list of supported calendar strings is:
//  - gregorian
//  - iso
//  - julian
//  - islamic
//  - hebrew
//  - mayanLongCount
//  - mayanHaab
//  - mayanTzolkin
//  - french
//  - oldHinduSolar
//  - oldHinduLunar
//
// For more information about these calendars, see:
//
// Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical Calculations", Software - Practice and Experience, 20 (9), 899-928. https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.17.4274.
//
// Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993. "Calendrical Calculations, II: Three Historical Calendars", Software - Practice & Experience, 23 (4), 383-404. https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.9215.
func FromAbsolute(date float64, calendar string) string {
	switch calendar {
	case "gregorian":
		return fmt.Sprint(GregorianFromAbsolute(date))
	case "julian":
		return fmt.Sprint(JulianFromAbsolute(date))
	case "iso":
		return fmt.Sprint(IsoFromAbsolute(date))
	case "islamic":
		return fmt.Sprint(IslamicFromAbsolute(date))
	case "hebrew":
		return fmt.Sprint(HebrewFromAbsolute(date))
	case "mayanLongCount":
		return fmt.Sprint(MayanLongCountFromAbsolute(date))
	case "mayanHaab":
		return fmt.Sprint(MayanHaabFromAbsolute(date))
	case "mayanTzolkin":
		return fmt.Sprint(MayanTzolkinFromAbsolute(date))
	case "french":
		return fmt.Sprint(FrenchFromAbsolute(date))
	case "oldHinduSolar":
		return fmt.Sprint(OldHinduSolarFromAbsolute(date))
	case "oldHinduLunar":
		return fmt.Sprint(OldHinduLunarFromAbsolute(date))
	default:
		return ""
	}
}
