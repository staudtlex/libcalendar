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

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

// FromAbsoluteToString converts a given absolute (fixed) date to the date
// representation specified in `calendar`.
//
// The list of supported calendar strings is:
//  - "gregorian"
//  - "iso"
//  - "julian"
//  - "islamic"
//  - "hebrew"
//  - "mayanLongCount"
//  - "mayanHaab"
//  - "mayanTzolkin"
//  - "french"
//  - "oldHinduSolar"
//  - "oldHinduLunar"
//
// For more information about these calendars, see:
//
// Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical Calculations", Software - Practice and Experience, 20 (9), 899-928. https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.17.4274.
//
// Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993. "Calendrical Calculations, II: Three Historical Calendars", Software - Practice & Experience, 23 (4), 383-404. https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.9215.
func FromAbsolute(absoluteDate float64, calendar string) string {
	switch calendar {
	case "gregorian":
		return fmt.Sprint(GregorianFromAbsolute(absoluteDate))
	case "julian":
		return fmt.Sprint(JulianFromAbsolute(absoluteDate))
	case "iso":
		return fmt.Sprint(IsoFromAbsolute(absoluteDate))
	case "islamic":
		return fmt.Sprint(IslamicFromAbsolute(absoluteDate))
	case "hebrew":
		return fmt.Sprint(HebrewFromAbsolute(absoluteDate))
	case "mayanLongCount":
		return fmt.Sprint(MayanLongCountFromAbsolute(absoluteDate))
	case "mayanHaab":
		return fmt.Sprint(MayanHaabFromAbsolute(absoluteDate))
	case "mayanTzolkin":
		return fmt.Sprint(MayanTzolkinFromAbsolute(absoluteDate))
	case "french":
		return fmt.Sprint(FrenchFromAbsolute(absoluteDate))
	case "oldHinduSolar":
		return fmt.Sprint(OldHinduSolarFromAbsolute(absoluteDate))
	case "oldHinduLunar":
		return fmt.Sprint(OldHinduLunarFromAbsolute(absoluteDate))
	default:
		return ""
	}
}

// Utilities for a more generic approach to converting dates

// Date represents a generic container for dates, holding information about
// arbitrary dates. Possible calendar names are:
type Date struct {
	Calendar       string    `json:"calendar"`       // calendar name, e.g. "gregorian"
	Components     []float64 `json:"components"`     // e.g. []float64{2022, 05, 28}
	ComponentNames []string  `json:"componentNames"` // e.g. []string{"year", "month", "day"}
	MonthNames     []string  `json:"monthNames"`     // e.g. []string{"January", ..., "December"}
}

// Json serializes the receiver into a JSON-formatted string.
func (d Date) Json() string {
	jsond, _ := json.Marshal(d)
	return string(jsond)
}

// String creates a string representation of its receiver.
//
// This function reuses the String() method of GregorianDate{} etc. by
// internally converting the given Date struct into a GregorianDate{},
// IsoDate{} or other appropriate struct.
func (d Date) String() string {
	switch d.Calendar {
	case "gregorian":
		return fmt.Sprint(gregorianFromDate(d))
	case "iso":
		return fmt.Sprint(isoFromDate(d))
	case "julian":
		return fmt.Sprint(gregorianFromDate(d))
	case "islamic":
		return fmt.Sprint(islamicFromDate(d))
	case "hebrew":
		return fmt.Sprint(hebrewFromDate(d))
	case "mayanLongCount":
		return fmt.Sprint(mayanLongCountFromDate(d))
	case "mayanHaab":
		return fmt.Sprint(mayanHaabFromDate(d))
	case "mayanTzolkin":
		return fmt.Sprint(mayanTzolkinFromDate(d))
	case "french":
		return fmt.Sprint(frenchFromDate(d))
	case "oldHinduSolar":
		return fmt.Sprint(oldHinduSolarFromDate(d))
	case "oldHinduLunar":
		return fmt.Sprint(oldHinduLunarFromDate(d))
	default:
		return ""
	}
}

// values returns the values of map[float64]string x as a slice of type string, sorted
// by the keys.
func values(x map[float64]string) []string {
	values := make([]string, 0, len(x))
	for _, key := range keys(x) {
		values = append(values, x[key])
	}

	return values
}

// keys returns the keys of map[float64]string x as a sorted slice of type float64.
func keys(x map[float64]string) []float64 {
	keys := make([]float64, 0, len(x))
	for key := range x {
		keys = append(keys, key)
	}
	sort.Float64s(keys)
	return keys
}

// lastValidDate checks if the day provided in the parameters is valid
// for a given year and month. If the day is valid, returns that day,
// otherwise returns the last valid day of the given year and month.
func LastValidDayOfMonth(year, month, day float64, calendar string) (validDay float64) {
	lastValidDay := 0.0
	switch calendar {
	case "gregorian":
		lastValidDay = LastDayOfGregorianMonth(month, year)
	case "julian":
		lastValidDay = LastDayOfJulianMonth(month, year)
	case "islamic":
		lastValidDay = LastDayOfIslamicMonth(month, year)
	case "hebrew":
		lastValidDay = LastDayOfHebrewMonth(month, year)
	case "french":
		lastValidDay = FrenchLastDayOfMonth(month, year)
	default:
		lastValidDay = day
	}
	if day < lastValidDay {
		return day
	} else {
		return lastValidDay
	}
}

// Date() creates a Date from its receiver.
func (d GregorianDate) Date() Date {
	return Date{
		Calendar: "gregorian",
		Components: []float64{
			d.Year,
			d.Month,
			d.Day,
		},
		ComponentNames: []string{
			"year", "month", "day",
		},
		MonthNames: values(gregorianMonths),
	}
}

// Date() creates a Date from its receiver.
func (d JulianDate) Date() Date {
	return Date{
		Calendar: "julian",
		Components: []float64{
			d.Year,
			d.Month,
			d.Day,
		},
		ComponentNames: []string{
			"year", "month", "day",
		},
		MonthNames: values(gregorianMonths),
	}
}

// Date() creates a Date from its receiver.
func (d IsoDate) Date() Date {
	return Date{
		Calendar: "iso",
		Components: []float64{
			d.Year,
			d.Week,
			d.Day,
		},
		ComponentNames: []string{
			"year", "week", "day",
		},
		MonthNames: []string{},
	}
}

// Date() creates a Date from its receiver.
func (d IslamicDate) Date() Date {
	return Date{
		Calendar: "islamic",
		Components: []float64{
			d.Year,
			d.Month,
			d.Day,
		},
		ComponentNames: []string{
			"year", "month", "day",
		},
		MonthNames: values(islamicMonths),
	}
}

// Date() creates a Date from its receiver.
func (d HebrewDate) Date() Date {
	return Date{
		Calendar: "hebrew",
		Components: []float64{
			d.Year,
			d.Month,
			d.Day,
		},
		ComponentNames: []string{
			"year", "month", "day",
		},
		MonthNames: HebrewMonthNames(d),
	}
}

// HebrewMonthNames returns the month names for a given Hebrew year, taking
// into account the Hebrew leap year rule.
func HebrewMonthNames(d HebrewDate) []string {
	if !HebrewLeapYear(d.Year) {
		return values(hebrewMonths)
	} else {
		months := values(hebrewMonths)
		months[11] = "Adar I"
		return append(months, "Adar II")
	}
}

// Date() creates a Date from its receiver.
func (d MayanLongCount) Date() Date {
	return Date{
		Calendar: "mayanLongCount",
		Components: []float64{
			d.Baktun,
			d.Katun,
			d.Tun,
			d.Uinal,
			d.Kin,
		},
		ComponentNames: []string{
			"baktun", "katun", "tun", "uinal", "kin",
		},
		MonthNames: []string{},
	}
}

// Date() creates a Date from its receiver.
func (d MayanHaabDate) Date() Date {
	return Date{
		Calendar: "mayanHaab",
		Components: []float64{
			d.Day,
			d.Month,
		},
		ComponentNames: []string{
			"day", "month",
		},
		MonthNames: values(mayanHaabMonths),
	}
}

// Date() creates a Date from its receiver.
func (d MayanTzolkinDate) Date() Date {
	return Date{
		Calendar: "mayanTzolkin",
		Components: []float64{
			d.Number,
			d.Name,
		},
		ComponentNames: []string{
			"number", "name",
		},
		MonthNames: values(mayanTzolkinNames),
	}
}

// Date() creates a Date from its receiver.
func (d FrenchDate) Date() Date {
	return Date{
		Calendar: "french",
		Components: []float64{
			d.Year,
			d.Month,
			d.Day,
		},
		ComponentNames: []string{
			"year", "month", "day",
		},
		MonthNames: values(frenchMonths),
	}
}

// Date() creates a Date from its receiver.
func (d OldHinduSolarDate) Date() Date {
	return Date{
		Calendar: "oldHinduSolar",
		Components: []float64{
			d.Year,
			d.Month,
			d.Day,
		},
		ComponentNames: []string{
			"year", "month", "day",
		},
		MonthNames: values(hinduSolarMonths),
	}
}

// Date() creates a Date from its receiver.
func (d OldHinduLunarDate) Date() Date {
	return Date{
		Calendar: "oldHinduLunar",
		Components: []float64{
			d.Year,
			d.Month,
			d.Day,
		},
		ComponentNames: []string{
			"year", "month", "day",
		},
		MonthNames: values(hinduLunarMonths),
	}
}

func gregorianFromDate(d Date) GregorianDate {
	return GregorianDate{
		Year:  d.Components[0],
		Month: d.Components[1],
		Day:   d.Components[2],
	}
}

func julianFromDate(d Date) JulianDate {
	return JulianDate{
		Year:  d.Components[0],
		Month: d.Components[1],
		Day:   d.Components[2],
	}
}

func isoFromDate(d Date) IsoDate {
	return IsoDate{
		Year: d.Components[0],
		Week: d.Components[1],
		Day:  d.Components[2],
	}
}

func islamicFromDate(d Date) IslamicDate {
	return IslamicDate{
		Year:  d.Components[0],
		Month: d.Components[1],
		Day:   d.Components[2],
	}
}

func hebrewFromDate(d Date) HebrewDate {
	return HebrewDate{
		Year:  d.Components[0],
		Month: d.Components[1],
		Day:   d.Components[2],
	}
}

func mayanLongCountFromDate(d Date) MayanLongCount {
	return MayanLongCount{
		Baktun: d.Components[0],
		Katun:  d.Components[1],
		Tun:    d.Components[2],
		Uinal:  d.Components[3],
		Kin:    d.Components[4],
	}
}

func mayanHaabFromDate(d Date) MayanHaabDate {
	return MayanHaabDate{
		Day:   d.Components[0],
		Month: d.Components[1],
	}
}

func mayanTzolkinFromDate(d Date) MayanTzolkinDate {
	return MayanTzolkinDate{
		Number: d.Components[0],
		Name:   d.Components[1],
	}
}

func frenchFromDate(d Date) FrenchDate {
	return FrenchDate{
		Year:  d.Components[0],
		Month: d.Components[1],
		Day:   d.Components[2],
	}
}

func oldHinduSolarFromDate(d Date) OldHinduSolarDate {
	return OldHinduSolarDate{
		Year:  d.Components[0],
		Month: d.Components[1],
		Day:   d.Components[2],
	}
}

func oldHinduLunarFromDate(d Date) OldHinduLunarDate {
	return OldHinduLunarDate{
		Year:  d.Components[0],
		Month: d.Components[1],
		Day:   d.Components[2],
	}
}

// AbsoluteFromDate returns the absolute (fixed) date from a given calendar
// date. Note that no checks are performed as to whether the given date is
// valid, i.e. in the date range for which the calendar was defined. For
// instance, conversion of dates preceding the 19.07.622 (1 Muharram 1 A.H.)
// into the Islamic calendar or dates preceding the 22.09.1792 (1 Vendémiare
// an 1) may yield non-sensical results.
func AbsoluteFromDate(d Date) (absoluteDate float64) {
	switch d.Calendar {
	case "gregorian":
		return AbsoluteFromGregorian(gregorianFromDate(d))
	case "iso":
		return AbsoluteFromIso(isoFromDate(d))
	case "julian":
		return AbsoluteFromJulian(julianFromDate(d))
	case "islamic":
		return AbsoluteFromIslamic(islamicFromDate(d))
	case "hebrew":
		return AbsoluteFromHebrew(hebrewFromDate(d))
	case "mayanLongCount":
		return AbsoluteFromMayanLongCount(mayanLongCountFromDate(d))
	case "french":
		return AbsoluteFromFrench(frenchFromDate(d))
	case "oldHinduSolar":
		return AbsoluteFromOldHinduSolar(oldHinduSolarFromDate(d))
	case "oldHinduLunar":
		return AbsoluteFromOldHinduLunar(oldHinduLunarFromDate(d))
	default:
		return math.NaN()
	}
}

// JsonToDate unmarshals a JSON-string into a Date struct.
func JsonToDate(s string) Date {
	c := Date{}
	json.Unmarshal([]byte(s), &c)
	return (c)
}

// DateFromAbsolute converts a given absolute (fixed) date into the date
// representation specified in `calendar`
func DateFromAbsolute(absoluteDate float64, calendar string) Date {
	switch calendar {
	case "gregorian":
		return GregorianFromAbsolute(absoluteDate).Date()
	case "iso":
		return IsoFromAbsolute(absoluteDate).Date()
	case "julian":
		return JulianFromAbsolute(absoluteDate).Date()
	case "islamic":
		return IslamicFromAbsolute(absoluteDate).Date()
	case "hebrew":
		return HebrewFromAbsolute(absoluteDate).Date()
	case "mayanLongCount":
		return MayanLongCountFromAbsolute(absoluteDate).Date()
	case "mayanHaab":
		return MayanHaabFromAbsolute(absoluteDate).Date()
	case "mayanTzolkin":
		return MayanTzolkinFromAbsolute(absoluteDate).Date()
	case "french":
		return FrenchFromAbsolute(absoluteDate).Date()
	case "oldHinduSolar":
		return OldHinduSolarFromAbsolute(absoluteDate).Date()
	case "oldHinduLunar":
		return OldHinduLunarFromAbsolute(absoluteDate).Date()
	default:
		return Date{}
	}
}

// JsonDateFromAbsolute converts a given absolute (fixed) date into the date
// representation specified in `calendar`. It returns a Date marshalled into
// a JSON-string.
func JsonDateFromAbsolute(rd float64, calendar string) string {
	return DateFromAbsolute(rd, calendar).Json()
}
