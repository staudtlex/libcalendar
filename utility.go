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
