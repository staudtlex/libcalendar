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

// The following Go code is translated from the Lisp code discussed in:
// - Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical
//   Calculations", Software - Practice and Experience, 20 (9), 899-928.
//   https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.17.4274
// - Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993.
//   "Calendrical Calculations, II: Three Historical Calendars",
//   Software - Practice & Experience, 23 (4), 383-404.
//   https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.9215

package libcalendar

import "math"

// Holidays

// Secular holidays

// IndependenceDay returns the absolute (fixed) date of the US
// Independence Day.
func IndependenceDay(year float64) (date float64) {
	return AbsoluteFromGregorian(GregorianDate{year, july, 4})
}

// NthKDay computes the absolute (fixed) date of the nth kth day in a given
// month in a given Gregorian year.
func NthKDay(n float64, k float64, month float64, year float64) (date float64) {
	if n > 0 {
		return KDayOnOrBefore(
			AbsoluteFromGregorian(GregorianDate{year, month, 7}), k) +
			(7 * (n - 1))
	} else {
		return KDayOnOrBefore(
			AbsoluteFromGregorian(
				GregorianDate{year, month, LastDayOfGregorianMonth(month, year)}), k) +
			(7 * (n + 1))
	}
}

// LaborDay returns the absolute (fixed) date of US Labor Day in a given
// Gregorian year.
func LaborDay(year float64) (date float64) {
	return NthKDay(1, 1, september, year)
}

// MemorialDay returns the absolute (fixed) date of US Memorial Day in a
// given Gregorian year.
func MemorialDay(year float64) (date float64) {
	return NthKDay(-1, 1, may, year)
}

// DaylightSavingsStart returns the absolute (fixed) date of the start of US
// daylight savings time.
func DaylightSavingsStart(year float64) (date float64) {
	//return NthKDay(1, 0, april, year) // before 2007
	return NthKDay(2, 0, march, year) // since 2007
}

// DaylightSavingsEnd returns the absolute (fixed) date of the end of US
// daylight savings time.
func DaylightSavingsEnd(year float64) (date float64) {
	//return NthKDay(-1, 0, october, year) // before 2007
	return NthKDay(1, 0, november, year) // since 2007
}

// Christian holidays

// Christmas returns the absolute (fixed) date of Gregorian Christmas in a
// given Gregorian year.
func Christmas(year float64) (date float64) {
	return AbsoluteFromGregorian(GregorianDate{year, december, 25})
}

// Advent returns the absolute (fixed) date of Advent in a given Gregorian year.
func Advent(year float64) (date float64) {
	return KDayOnOrBefore(
		AbsoluteFromGregorian(GregorianDate{year, december, 3}), 0)
}

// Epiphany returns the absolute (fixed) date of Epiphany in a given
// Gregorian year.
func Epiphany(year float64) (date float64) {
	return Christmas(year) + 12
}

// EasternOrthodoxChristmas returns the absolute (fixed) date of Eastern
// Orthodox Christmas in a given Gregorian year.
func EasternOrthodoxChristmas(year float64) (dates []float64) {
	jan_1 := AbsoluteFromGregorian(GregorianDate{year, january, 1})
	dec_31 := AbsoluteFromGregorian(GregorianDate{year, december, 31})
	y := JulianFromAbsolute(jan_1).Year
	c_1 := AbsoluteFromJulian(JulianDate{y, december, 25})
	c_2 := AbsoluteFromJulian(JulianDate{y + 1, december, 25})
	dates = make([]float64, 0, 2)
	if jan_1 <= c_1 && c_1 <= dec_31 {
		dates = append(dates, c_1)
	}
	if jan_1 <= c_2 && c_2 <= dec_31 {
		dates = append(dates, c_2)
	}
	return dates
}

// NicaeanRuleEaster computes the absolute (fixed) date of Easter in a given
// Julian year.
func NicaeanRuleEaster(year float64) (date float64) {
	shiftedEpact := mod(14+(11*mod(year, 19)), 30)
	paschalMoon := AbsoluteFromJulian(JulianDate{year, april, 19}) - shiftedEpact
	return KDayOnOrBefore(paschalMoon+7, 0)
}

// Easter computes the absolute (fixed) date of Easter in a given Gregorian year.
func Easter(year float64) (date float64) {
	century := math.Floor(year/100) + 1
	shiftedEpact := mod(14+
		(11*mod(year, 19))-
		math.Floor((3*century)/4)+
		math.Floor((5+(8*century))/25)+
		30*century, 30)
	adjustedEpact := 0.0
	if shiftedEpact == 0 || (shiftedEpact == 1 && 10 < mod(year, 19)) {
		adjustedEpact = shiftedEpact + 1
	} else {
		adjustedEpact = shiftedEpact
	}
	paschalMoon := AbsoluteFromGregorian(GregorianDate{year, april, 19}) - adjustedEpact
	return KDayOnOrBefore(paschalMoon+7, 0)
}

// Pentecost returns the absolute (fixed) date of Pentecost in a given
// Gregorian year.
func Pentecost(year float64) (date float64) {
	return Easter(year) + 49
}

// Islamic holidays

// IslamicDatesInGregorianYear returns a slice of absolute dates of a given
// Islamic date (month, day) that occur in a given Gregorian year.
func IslamicDatesInGregorianYear(month float64, day float64, year float64) (dates []float64) {
	jan_1 := AbsoluteFromGregorian(GregorianDate{year, january, 1})
	dec_31 := AbsoluteFromGregorian(GregorianDate{year, december, 31})
	y := IslamicFromAbsolute(jan_1).Year
	date_1 := AbsoluteFromIslamic(IslamicDate{y, month, day})
	date_2 := AbsoluteFromIslamic(IslamicDate{y + 1, month, day})
	date_3 := AbsoluteFromIslamic(IslamicDate{y + 2, month, day})
	dates = make([]float64, 0, 3)
	if jan_1 <= date_1 && date_1 <= dec_31 {
		dates = append(dates, date_1)
	}
	if jan_1 <= date_2 && date_2 <= dec_31 {
		dates = append(dates, date_2)
	}
	if jan_1 <= date_3 && date_3 <= dec_31 {
		dates = append(dates, date_3)
	}
	return dates
}

// MuladAlNabi computes slice of absolute (fixed) dates of Mulad al Nabi that
// occur in a given Gregorian year.
func MuladAlNabi(year float64) (dates []float64) {
	return IslamicDatesInGregorianYear(rabi_i, 12, year)
}

// Jewish holidays

// YomKippur returns the absolute (fixed) date of Yom Kippur in a given
// Gregorian year.
func YomKippur(year float64) (date float64) {
	return AbsoluteFromHebrew(HebrewDate{year + 3761, tishri, 10})
}

// Passover returns the abolute (fixed) date of Passover in a given Gregorian year.
func Passover(year float64) (date float64) {
	return AbsoluteFromHebrew(HebrewDate{year + 3760, nisan, 15})
}

// Purim returns the absolute (fixed) date of Purim in a given Gregorian year.
func Purim(year float64) (date float64) {
	return AbsoluteFromHebrew(
		HebrewDate{year + 3760, LastMonthOfHebrewYear(year + 3760), 14})
}

// TaAnitEsther returns the absolute (fixed) date of TaAnitEsther in a given
// Gregorian year.
func TaAnitEsther(year float64) (date float64) {
	purimDate := Purim(year)
	if mod(purimDate, 7) == 0 {
		return purimDate - 3
	} else {
		return purimDate - 1
	}
}

// TishaBAv returns the absolute (fixed) date of Tisha B'Av in a given
// Gregorian year.
func TishaBAv(year float64) (date float64) {
	ninthOfAv := AbsoluteFromHebrew(HebrewDate{year + 3760, av, 9})
	if mod(ninthOfAv, 7) == 6 {
		return ninthOfAv + 1
	} else {
		return ninthOfAv
	}
}

// HebrewBirthday determines the absolute (fixed) date of the anniversary of a
// given Hebrew birth date in a given Hebrew year.
func HebrewBirthday(birthdate HebrewDate, year float64) (date float64) {
	birthYear := birthdate.Year
	birthMonth := birthdate.Month
	birthDay := birthdate.Day
	if birthMonth == LastMonthOfHebrewYear(birthYear) {
		return AbsoluteFromHebrew(HebrewDate{year, LastMonthOfHebrewYear(year), birthDay})
	} else {
		return AbsoluteFromHebrew(HebrewDate{year, birthMonth, birthDay})
	}
}

// Yahrzeit determines the absolute (fixed) date of the anniversary of a given
// Hebrew death-date in a given Hebrew year
func Yahrzeit(deathDate HebrewDate, year float64) (date float64) {
	deathYear := deathDate.Year
	deathMonth := deathDate.Month
	deathDay := deathDate.Day
	switch {
	case deathMonth == 8 && deathDay == 30 && !LongHeshvan(deathYear+1):
		return AbsoluteFromHebrew(HebrewDate{year, kislev, 1})
	case deathMonth == 9 && deathDay == 30 && ShortKislev(deathYear+1):
		return AbsoluteFromHebrew(HebrewDate{year, teveth, 1})
	case deathMonth == 13:
		return AbsoluteFromHebrew(HebrewDate{year, LastMonthOfHebrewYear(year), deathDay})
	case deathDay == 30 && deathMonth == 12 && !HebrewLeapYear(deathYear):
		return AbsoluteFromHebrew(HebrewDate{year, shevat, 30})
	default:
		return AbsoluteFromHebrew(HebrewDate{year, deathMonth, deathDay})
	}
}
