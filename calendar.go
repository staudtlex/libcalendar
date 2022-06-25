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

// Package libcalendar implements functions to compute and convert dates
// from various calendars. These are the
// Gregorian, ISO, Julian, Islamic, Hebrew, Mayan (long count, haab, tzolkin),
// French Revolutionary, and Old Hindu (solar, lunar) calendars.
//
// The calendrical algorithms are a translation of the Lisp code discussed in:
//
// Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical Calculations",
// Software - Practice and Experience, 20 (9), 899-928.
// https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.17.4274
//
// Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993. "Calendrical
// Calculations, II: Three Historical Calendars", Software - Practice &
// Experience, 23 (4), 383-404.
// https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.9215
package libcalendar // import "staudtlex.de/libcalendar"

import (
	"math"
	"math/big"
)

// The Gregorian Calendar

// Gregorian and Julian months
const (
	january   = 1
	february  = 2
	march     = 3
	april     = 4
	may       = 5
	june      = 6
	july      = 7
	august    = 8
	september = 9
	october   = 10
	november  = 11
	december  = 12
)

// Gregorian date
type GregorianDate struct {
	Year  float64
	Month float64
	Day   float64
}

// LastDayOfGregorianMonth returns the last day (number of days) of a given
// Gregorian month.
func LastDayOfGregorianMonth(month float64, year float64) (day float64) {
	daysInMonths := []float64{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (month == 2) &&
		(mod(year, 4) == 0) &&
		!member(mod(year, 400), []float64{100, 200, 300}) {
		return 29
	} else {
		return daysInMonths[int(month)-1]
	}
}

// AbsoluteFromGregorian computes the absolute (fixed) date from a
// Gregorian date.
func AbsoluteFromGregorian(d GregorianDate) (absoluteDate float64) {
	year := d.Year
	month := d.Month
	day := d.Day
	f := func(m float64) float64 { return LastDayOfGregorianMonth(m, year) }
	p := func(m float64) bool { return m < month }
	return day +
		sum(f, 1, p) +
		(365 * (year - 1)) +
		math.Floor((year-1)/4) +
		(-math.Floor((year - 1) / 100)) +
		(math.Floor((year - 1) / 400))
}

// GregorianFromAbsolute computes the Gregorian date corresponding to a
// given absolute date.
func GregorianFromAbsolute(absoluteDate float64) GregorianDate {
	// compute year
	d_0 := absoluteDate - 1
	n_400 := math.Floor(d_0 / 146097)
	d_1 := mod(d_0, 146097)
	n_100 := math.Floor(d_1 / 36524)
	d_2 := mod(d_1, 36524)
	n_4 := math.Floor(d_2 / 1461)
	d_3 := mod(d_2, 1461)
	n_1 := math.Floor(d_3 / 365)
	//d_4 := mod(d_3, 365)
	year := 400*n_400 + 100*n_100 + 4*n_4 + n_1
	if n_100 == 4 || n_1 == 4 {
		// do nothing
	} else {
		year = year + 1
	}
	// compute month
	f := func(float64) float64 { return 1 }
	p := func(m float64) bool {
		d := GregorianDate{year, m, LastDayOfGregorianMonth(m, year)}
		return absoluteDate > AbsoluteFromGregorian(d)
	}
	month := sum(f, 1, p) + 1
	// compute day
	day := absoluteDate - (AbsoluteFromGregorian(GregorianDate{year, month, 1}) - 1)
	// return date
	return GregorianDate{year, month, day}
}

// The ISO Calendar

// ISO date
type IsoDate struct {
	Year float64
	Week float64
	Day  float64
}

// KDayOnOrBefore computes the absolute date of a given week day in the
// seven-day interval ending on date.
func KDayOnOrBefore(absoluteDate float64, k float64) float64 {
	return absoluteDate - mod(absoluteDate-k, 7)
}

// AbsoluteFromIso computes the absolute (fixed) date from an ISO date.
func AbsoluteFromIso(d IsoDate) (absoluteDate float64) {
	year := d.Year
	week := d.Week
	day := d.Day
	return KDayOnOrBefore(
		AbsoluteFromGregorian(GregorianDate{year, january, 4}), 1) +
		(7 * (week - 1)) +
		(day - 1)
}

// IsoFromAbsolute computes the IsoDate corresponding to a given absolute
// (fixed) date.
func IsoFromAbsolute(absoluteDate float64) IsoDate {
	approx := GregorianFromAbsolute(absoluteDate - 3).Year
	year := 0.0
	if absoluteDate >= AbsoluteFromIso(IsoDate{approx + 1, 1, 1}) {
		year = approx + 1
	} else {
		year = approx
	}
	week := 1 + math.Floor(
		(absoluteDate-AbsoluteFromIso(IsoDate{year, 1, 1}))/
			7)
	day := 0.0
	if mod(absoluteDate, 7) == 0 {
		day = 7
	} else {
		day = mod(absoluteDate, 7)
	}
	return IsoDate{year, week, day}
}

// The Julian Calendar

// LastDayOfJulianMonth returns the last day (number of days) of a given
// Julian month.
func LastDayOfJulianMonth(month float64, year float64) (day float64) {
	daysInMonths := []float64{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (month == 2) && (mod(year, 4) == 0) {
		return 29
	} else {
		return daysInMonths[int(month)-1]
	}
}

// Julian date
type JulianDate GregorianDate

// AbsoluteFromJulian computes the absolute (fixed) date corresponding to a
// given Julian date.
func AbsoluteFromJulian(d JulianDate) (absoluteDate float64) {
	year := d.Year
	month := d.Month
	day := d.Day
	f := func(m float64) float64 { return LastDayOfJulianMonth(m, year) }
	p := func(m float64) bool { return m < month }
	return day +
		sum(f, 1, p) +
		(365 * (year - 1)) +
		math.Floor((year-1)/4) -
		2
}

// JulianFromAbsolute computes the Julian date corresponding to a
// given absolute date.
func JulianFromAbsolute(absoluteDate float64) JulianDate {
	approx := math.Floor((absoluteDate + 2) / 366)
	f := func(float64) float64 { return 1 }
	p := func(y float64) bool {
		return absoluteDate >= AbsoluteFromJulian(JulianDate{y + 1, january, 1})
	}
	year := approx + sum(f, approx, p)
	month := 1 + sum(f, 1, func(m float64) bool {
		return absoluteDate > AbsoluteFromJulian(
			JulianDate{year, m, LastDayOfJulianMonth(m, year)})
	})
	day := absoluteDate - (AbsoluteFromJulian(JulianDate{year, month, 1}) - 1)
	return JulianDate{year, month, day}
}

// The Islamic Calendar

// Islamic months
const (
	muharram    = 1
	safar       = 2
	rabi_i      = 3
	rabi_ii     = 4
	jumada_i    = 5
	jumada_ii   = 6
	rajab       = 7
	shaBan      = 8
	ramadan     = 9
	shawwal     = 10
	dhuAlQada   = 11
	dhuAlHijjah = 12
)

// Islamic date
type IslamicDate struct {
	Year  float64
	Month float64
	Day   float64
}

// IslamicLeapYear returns true if a given Islamic year is leap, and
// false otherwise.
func IslamicLeapYear(year float64) bool {
	return mod(14+(11*year), 30) < 11
}

// LastDayOfIslamicMonth determines the last day of an Islamic month.
func LastDayOfIslamicMonth(month float64, year float64) (day float64) {
	if mod(month, 2) != 0 || (month == 12 && IslamicLeapYear(year)) {
		return 30
	} else {
		return 29
	}
}

// AbsoluteFromIslamic computes the absolute date corresponding to a given
// Islamic date.
func AbsoluteFromIslamic(d IslamicDate) (absoluteDate float64) {
	year := d.Year
	month := d.Month
	day := d.Day
	return day +
		(29 * (month - 1)) +
		math.Floor(month/2) +
		((year - 1) * 354) +
		math.Floor((3+(11*year))/30) +
		227014
}

// IslamicFromAbsolute computes the Islamic date corresponding to a given
// absolute date.
func IslamicFromAbsolute(absoluteDate float64) IslamicDate {
	if absoluteDate <= 227014 {
		return IslamicDate{0, 0, 0}
	}
	approx := math.Floor((absoluteDate - 227014) / 355)
	f := func(float64) float64 { return 1 }
	p := func(y float64) bool {
		return absoluteDate >= AbsoluteFromIslamic(IslamicDate{y + 1, muharram, 1})
	}
	year := approx + sum(f, approx, p)
	month := 1 + sum(f, 1, func(m float64) bool {
		return absoluteDate > AbsoluteFromIslamic(
			IslamicDate{year, m, LastDayOfIslamicMonth(m, year)})
	})
	day := absoluteDate - (AbsoluteFromIslamic(IslamicDate{year, month, 1}) - 1)
	return IslamicDate{year, month, day}
}

// The Hebrew Calendar

// Hebrew months
const (
	nisan   = 1
	iyyar   = 2
	sivan   = 3
	tammuz  = 4
	av      = 5
	elul    = 6
	tishri  = 7
	heshvan = 8
	kislev  = 9
	teveth  = 10
	shevat  = 11
	adar    = 12 // only in leap years, otherwise adar is skipped
	adar_ii = 13
)

// Hebrew date
type HebrewDate struct {
	Year  float64
	Month float64
	Day   float64
}

// HebrewLeapYear returns true if year is a Hebrew leap year.
func HebrewLeapYear(year float64) bool {
	return mod(1+(7*year), 19) < 7
}

// LastMonthOfHebrewYear returns the last month of a given Hebrew year.
func LastMonthOfHebrewYear(year float64) (month float64) {
	if HebrewLeapYear(year) {
		return 13
	} else {
		return 12
	}
}

// LastDayOfHebrewMonth returns the day (number of days) of a given
// Hebrew month.
func LastDayOfHebrewMonth(month float64, year float64) (day float64) {
	if member(month, []float64{2, 4, 6, 10, 13}) ||
		(month == 12 && !HebrewLeapYear(year)) ||
		(month == 8 && !LongHeshvan(year)) ||
		(month == 9 && ShortKislev(year)) {
		return 29
	} else {
		return 30
	}
}

// HebrewCalendarElapsedDays computes the number of days elapsed from the
// Sunday prior to the start of the Hebrew calendar to the mean conjunction of
// Tishri of a given Hebrew year.
func HebrewCalendarElapsedDays(year float64) (days float64) {
	// months in complete cycle so far
	monthsElapsed := (235 * math.Floor((year-1)/19)) +
		(12 * mod(year-1, 19)) +
		math.Floor((7*mod(year-1, 19)+1)/19)
	partsElapsed := 204 + 793*mod(monthsElapsed, 1080)
	hoursElapsed := 5 +
		12*monthsElapsed +
		793*math.Floor(monthsElapsed/1080) +
		math.Floor(partsElapsed/1080)
	// conjunction day
	day := 1 + 29*monthsElapsed + math.Floor(hoursElapsed/24)
	// conjunction parts
	parts := 1080*mod(hoursElapsed, 24) + mod(partsElapsed, 1080)
	alternativeDay := 0.0
	if parts >= 19440 || // if new moon is at or after midday
		// or is on a Tuesday at 9 hours, 224 parts later of a common year
		(mod(day, 7) == 2 && parts >= 9924 && !HebrewLeapYear(year)) ||
		// or is on a Monday at 15 hours, 589 parts later at the end of a leap year
		(mod(day, 7) == 1 && parts >= 16789 && HebrewLeapYear(year-1)) {
		// postpone Rosh HaShanah one day
		alternativeDay = day + 1
	} else {
		alternativeDay = day
	}
	// Ih Rosh HaShanah would occur on Sunday, Wednesday, or Friday
	if member(mod(alternativeDay, 7), []float64{0, 3, 5}) {
		// postpone it one (more day)
		return alternativeDay + 1
	} else {
		return alternativeDay
	}
}

// DaysInHebrewYear computes the number of days in a given Hebrew year.
func DaysInHebrewYear(year float64) (days float64) {
	return HebrewCalendarElapsedDays(year+1) - HebrewCalendarElapsedDays(year)
}

// LongHeshvan returns true if Heshvan is long in a given Hebrew year.
func LongHeshvan(year float64) bool {
	return mod(DaysInHebrewYear(year), 10) == 5
}

// ShortKislev returns true if Kislev is short in a given Hebrew year.
func ShortKislev(year float64) bool {
	return mod(DaysInHebrewYear(year), 10) == 3
}

// AbsoluteFromHebrew computes the absolute (fixed) date from a given
// Hebrew date.
func AbsoluteFromHebrew(d HebrewDate) (absoluteDate float64) {
	year := d.Year
	month := d.Month
	day := d.Day
	f := func(m float64) float64 { return LastDayOfHebrewMonth(m, year) }
	if month < 7 {
		return day +
			sum(f, 7, func(m float64) bool { return m <= LastMonthOfHebrewYear(year) }) +
			sum(f, 1, func(m float64) bool { return m < month }) +
			HebrewCalendarElapsedDays(year) +
			-1373429
	} else {
		return day +
			sum(f, 7, func(m float64) bool { return m < month }) +
			HebrewCalendarElapsedDays(year) +
			-1373429
	}
}

// HebrewFromAbsolute computes the Hebrew date corresponding to a given
// absolute (fixed) date
func HebrewFromAbsolute(absoluteDate float64) HebrewDate {
	approx := math.Floor((absoluteDate + 1373429) / 366)
	f := func(float64) float64 { return 1 }
	year := approx + sum(f, approx, func(y float64) bool {
		return absoluteDate >= AbsoluteFromHebrew(HebrewDate{y + 1, 7, 1})
	})
	start := 0.0
	if (absoluteDate < AbsoluteFromHebrew(HebrewDate{year, 1, 1})) {
		start = 7
	} else {
		start = 1
	}
	month := start + sum(f, start, func(m float64) bool {
		return absoluteDate > AbsoluteFromHebrew(HebrewDate{year, m, LastDayOfHebrewMonth(m, year)})
	})
	day := absoluteDate - (AbsoluteFromHebrew(HebrewDate{year, month, 1}) - 1)
	return HebrewDate{year, month, day}
}

// The Mayan Calendars

// Mayan haab months
const (
	pop    = 1
	uo     = 2
	zip    = 3
	zotz   = 4
	tzec   = 5
	xul    = 6
	yaxkin = 7
	mol    = 8
	chen   = 9
	yax    = 10
	zac    = 11
	ceh    = 12
	mac    = 13
	kankin = 14
	muan   = 15
	pax    = 16
	kayab  = 17
	cumku  = 18
)

// Mayan tzolkin names
const (
	imix    = 1
	ik      = 2
	akbal   = 3
	kan     = 4
	chiccan = 5
	cimi    = 6
	manik   = 7
	lamat   = 8
	muluc   = 9
	oc      = 10
	chuen   = 11
	eb      = 12
	ben     = 13
	ix      = 14
	men     = 15
	cib     = 16
	caban   = 17
	etznab  = 18
	cauac   = 19
	ahau    = 20
)

// Mayan long count date type
type MayanLongCount struct {
	Baktun float64
	Katun  float64
	Tun    float64
	Uinal  float64
	Kin    float64
}

// Mayan Haab date type
type MayanHaabDate struct {
	Day   float64
	Month float64
}

// Mayan Tzolkin date type
type MayanTzolkinDate struct {
	Number float64
	Name   float64
}

// Number of days of the Mayan calendar epoch before absolute day 0, according
// to the Goodman-Martinez-Thompson correlation (see Reingold/Dershowitz 2018).
const MayanDaysBeforeAbsoluteZero float64 = 1137142

// Number of days of the Mayan calendar epoch before absolute day 0, according
// to the Spinden correlation.
//const MayanDaysBeforeAbsoluteZero float64 = 1232041

// AbsoluteFromMayanLongCount returns the absolute (fixed) date of a given
// Mayan long count.
func AbsoluteFromMayanLongCount(d MayanLongCount) (date float64) {
	return d.Baktun*144000 +
		d.Katun*7200 +
		d.Tun*360 +
		d.Uinal*20 +
		d.Kin -
		MayanDaysBeforeAbsoluteZero
}

// MayanLongCountFromAbsolute computes the Mayan long count corresponding to
// the given absolute date.
func MayanLongCountFromAbsolute(absoluteDate float64) MayanLongCount {
	longCount := absoluteDate + MayanDaysBeforeAbsoluteZero
	baktun := math.Floor(longCount / 144000)
	dayOfBaktun := mod(longCount, 144000)
	katun := math.Floor(dayOfBaktun / 7200)
	dayOfKatun := mod(dayOfBaktun, 7200)
	tun := math.Floor(dayOfKatun / 360)
	dayOfTun := mod(dayOfKatun, 360)
	uinal := math.Floor(dayOfTun / 20)
	kin := mod(dayOfTun, 20)
	return MayanLongCount{baktun, katun, tun, uinal, kin}
}

// MayanHaabAtEpoch denotes the haab date at long count 0.0.0.0.0.
var MayanHaabAtEpoch MayanHaabDate = MayanHaabDate{8, cumku}

// MayanHaabFromAbsolute returns the Mayan haab date corresponding to a given
// absolute (fixed) date.
func MayanHaabFromAbsolute(absoluteDate float64) MayanHaabDate {
	longCount := absoluteDate + MayanDaysBeforeAbsoluteZero
	dayOfHaab := mod(longCount+MayanHaabAtEpoch.Day+(20*(MayanHaabAtEpoch.Month-1)), 365)
	day := mod(dayOfHaab, 20)
	month := math.Floor(dayOfHaab/20) + 1
	return MayanHaabDate{day, month}
}

// MayanHaabDifference computes the number of days between two haab dates.
func MayanHaabDifference(d1, d2 MayanHaabDate) (days float64) {
	return mod(20*(d2.Day-d1.Day)+(d2.Month-d1.Month), 365)
}

// MayanHaabOnOrBefore returns the absolute (fixed) date of a Mayan haab date
// on or before a given absolute date.
func MayanHaabOnOrBefore(haab MayanHaabDate, d float64) (date float64) {
	return d - mod(date-MayanHaabDifference(MayanHaabFromAbsolute(0), haab), 365)
}

// MayanTzolkinAtEpoch denotes tha tzolkin date at long count 0.0.0.0.0.
var MayanTzolkinAtEpoch MayanTzolkinDate = MayanTzolkinDate{4, ahau}

// MayanTzolkinFromAbsolute returns a Mayan tzolkin date corresponding to
// a given absolute (fixed) date.
func MayanTzolkinFromAbsolute(absoluteDate float64) MayanTzolkinDate {
	longCount := absoluteDate + MayanDaysBeforeAbsoluteZero
	number := amod(longCount+MayanTzolkinAtEpoch.Number, 13)
	name := amod(longCount+MayanTzolkinAtEpoch.Name, 20)
	return MayanTzolkinDate{number, name}
}

// MayanTzolkinDifference returns the number of days between two given Mayan
// tzolkin dates.
func MayanTzolkinDifference(d1, d2 MayanTzolkinDate) (days float64) {
	numberDifference := d2.Number - d1.Number
	nameDifference := d2.Name - d1.Name
	return mod(numberDifference+(13*mod(3*(numberDifference-nameDifference), 20)), 260)
}

// MayanHaabTzolkinOnOrBefore returns the absolute date of the latest date on
// or before a given haab date and a given tzolkin date. Returns NaN when no
// such combination is found.
func MayanHaabTzolkinOnOrBefore(haab MayanHaabDate, tzolkin MayanTzolkinDate, d float64) (absoluteDate float64) {
	haabDifference := MayanHaabDifference(MayanHaabFromAbsolute(0), haab)
	tzolkinDifference := MayanTzolkinDifference(MayanTzolkinFromAbsolute(0), tzolkin)
	difference := tzolkinDifference - haabDifference
	if mod(difference, 5) == 0 {
		return d - mod(d-(haabDifference+(365*difference)), 18980)
	} else {
		return math.NaN()
	}
}

// The French Revolutionary Calendar

// French Revolutionary calendar months
const (
	vendémiaire = 1
	brumaire    = 2
	frimaire    = 3
	nivôse      = 4
	pluviôse    = 5
	ventôse     = 6
	germinal    = 7
	floréal     = 8
	prairial    = 9
	messidor    = 10
	thermidor   = 11
	fructidor   = 12
)

// French date type
type FrenchDate struct {
	Year  float64
	Month float64
	Day   float64
}

// FrenchLastDayOfMonth returns the last day of a given French Revolutionary
// month in a given French Revolutionary year
func FrenchLastDayOfMonth(month, year float64) (day float64) {
	if month < 13 {
		return 30
	} else {
		if FrenchLeapYear(year) {
			return 6
		} else {
			return 5
		}
	}
}

// FrenchLeapYear returns true if a given year is a leap year, and
// false otherwise
func FrenchLeapYear(year float64) bool {
	return member(year, []float64{3, 7, 11}) ||
		member(year, []float64{15, 20}) ||
		(year > 20 &&
			mod(year, 4) == 0 &&
			!member(mod(year, 400), []float64{100, 200, 300}) &&
			mod(year, 4000) != 0)
}

// AbsoluteFromFrench returns the absolute (fixed) date from a given French
// Revolutionary date.
func AbsoluteFromFrench(d FrenchDate) (absoluteDate float64) {
	year := d.Year
	month := d.Month
	day := d.Day
	if year < 20 {
		return 654414 + (365 * (year - 1)) +
			math.Floor(year/4) +
			(30 * (month - 1)) + day
	} else {
		return 654414 + (365 * (year - 1)) +
			(math.Floor((year-1)/4) -
				math.Floor((year-1)/100) +
				math.Floor((year-1)/400) -
				math.Floor((year-1)/4000)) +
			(30 * (month - 1)) + day
	}
}

// FrenchFromAbsolute returns the French Revolutionary date corresponding to a
// given absolute (fixed) date.
func FrenchFromAbsolute(absoluteDate float64) FrenchDate {
	if absoluteDate < 654415 {
		return FrenchDate{0, 0, 0}
	}
	approx := math.Floor((absoluteDate - 654414) / 366)
	f := func(float64) float64 { return 1 }
	year := approx + sum(f, approx,
		func(y float64) bool {
			return absoluteDate >= AbsoluteFromFrench(FrenchDate{y + 1, vendémiaire, 1})
		})
	month := 1 + sum(f, 1,
		func(m float64) bool {
			return absoluteDate > AbsoluteFromFrench(FrenchDate{year, m, FrenchLastDayOfMonth(m, year)})
		})
	day := absoluteDate - (AbsoluteFromFrench(FrenchDate{year, month, 1}) - 1)
	return FrenchDate{year, month, day}
}

// The Old Hindu Calendars

// Old Hindu solar date type
type OldHinduSolarDate struct {
	Year  float64
	Month float64
	Day   float64
}

// Old Hindu lunar date type
type OldHinduLunarDate struct {
	Year      float64
	Month     float64
	LeapMonth bool
	Day       float64
}

// Old Hindu solar months
const (
	mesha     = 1
	vrshabha  = 2
	mithuna   = 3
	karka     = 4
	simha     = 5
	kanya     = 6
	tula      = 7
	vrischika = 8
	dhanus    = 9
	makara    = 10
	kumbha    = 11
	mina      = 12
)

// Old Hindu lunar months
const (
	chaitra    = 1
	vaisakha   = 2
	jyaishtha  = 3
	ashadha    = 4
	sravana    = 5
	bhadrapada = 6
	asvina     = 7
	karttika   = 8
	margasira  = 9
	pausha     = 10
	magha      = 11
	phalguna   = 12
)

var SolarSiderealYear = add(big.NewRat(365, 1), big.NewRat(279457, 1080000))
var SolarMonth = div(SolarSiderealYear, big.NewRat(12, 1))
var LunarSiderealMonth = add(big.NewRat(27, 1), big.NewRat(4644439, 14438334))
var LunarSynodicMonth = add(big.NewRat(29, 1), big.NewRat(7087771, 13358334))

// SolarLongitude returns the position of the sun (in degrees) for a
// given moment (day and fraction of a day).
func SolarLongitude(t *big.Rat) (degrees *big.Rat) {
	return mult(modr(div(t, SolarSiderealYear), big.NewRat(1, 1)), big.NewRat(360, 1))
}

// Zodiac returns the zodiacal sign for a given moment (day and fraction
// of day).
func Zodiac(t *big.Rat) (zodiac float64) {
	zodiac = quotient(SolarLongitude(t), big.NewRat(30, 1))
	return zodiac + 1
}

// OldHinduSolarFromAbsolute computes the Old Hindu solar date corresponding
// to a given absolute (fixed) date.
func OldHinduSolarFromAbsolute(absoluteDate float64) OldHinduSolarDate {
	hdate := add(big.NewRat(int64(absoluteDate), 1), big.NewRat(1132959, 1), big.NewRat(1, 4))
	year := quotient(hdate, SolarSiderealYear)
	month := Zodiac(hdate)
	day := floorf(modr(hdate, SolarMonth)) + 1
	return OldHinduSolarDate{year, month, day}
}

// AbsoluteFromOldHinduSolar returns the absolute (fixed) date from a given
// Old Hindu solar date.
func AbsoluteFromOldHinduSolar(d OldHinduSolarDate) (absoluteDate float64) {
	year := big.NewRat(int64(d.Year), 1)
	month := big.NewRat(int64(d.Month), 1)
	day := big.NewRat(int64(d.Day), 1)
	one := big.NewRat(1, 1)
	result, _ := add(
		mult(year, SolarSiderealYear),
		mult(sub(month, one), SolarMonth),
		day,
		big.NewRat(-1, 4),
		big.NewRat(-1132959, 1)).Float64()
	return math.Floor(result)
}

// LunarLongitude returns the sidereal longitude of the moon (in degrees) at
// a given moment (date and fraction of a day).
func LunarLongitude(t *big.Rat) (degrees *big.Rat) {
	return mult(modr(div(t, LunarSiderealMonth), big.NewRat(1, 1)), big.NewRat(360, 1))
}

// LunarPhase computes the lunar phase of the moon for a given moment (date
// and fraction of a day).
func LunarPhase(t *big.Rat) (phase float64) {
	return 1 + quotient(
		modr(
			sub(LunarLongitude(t), SolarLongitude(t)),
			big.NewRat(360, 1)),
		big.NewRat(12, 1))
}

// NewMoon determines the time of the most recent new moon for a given moment
// (date and fraction of day).
func NewMoon(t *big.Rat) *big.Rat {
	return sub(t, modr(t, LunarSynodicMonth))
}

// OldHinduLunarFromAbsolute returns the Old Hindu lunar date corresponding to
// a given absolute (fixed) date.
func OldHinduLunarFromAbsolute(absoluteDate float64) OldHinduLunarDate {
	hdate := big.NewRat(int64(absoluteDate+1132959), 1)
	sunrise := add(hdate, big.NewRat(1, 4))
	lastNewMoon := NewMoon(sunrise)
	nextNewMoon := add(lastNewMoon, LunarSynodicMonth)
	month := amod(Zodiac(lastNewMoon)+1, 12)
	day := LunarPhase(sunrise)
	leapMonth := Zodiac(lastNewMoon) == Zodiac(nextNewMoon)
	nextMonth := nextNewMoon
	if leapMonth {
		nextMonth = add(nextMonth, LunarSynodicMonth)
	} else {
		nextMonth = add(nextMonth, big.NewRat(0, 1))
	}
	year := quotient(nextMonth, SolarSiderealYear)
	return OldHinduLunarDate{year, month, leapMonth, day}
}

// OldHinduLunarPrecedes returns true if a given Hindu lunar date d1 precedes
// (i.e. is smaller than) a given Hindu lunar date d2, and false otherwise.
func OldHinduLunarPrecedes(d1, d2 OldHinduLunarDate) bool {
	year_1 := d1.Year
	year_2 := d2.Year
	month_1 := d1.Month
	month_2 := d2.Month
	leap_1 := d1.LeapMonth
	leap_2 := d2.LeapMonth
	day_1 := d1.Day
	day_2 := d2.Day
	return (year_1 < year_2 ||
		(year_1 == year_2 &&
			(month_1 < month_2 ||
				(month_1 == month_2 &&
					(leap_1 && !leap_2 ||
						(leap_1 == leap_2 &&
							day_1 < day_2))))))
}

// AbsoluteFromOldHinduLunar returns the absolute (fixed) date corresponding
// to a given Old Hindu lunar date.
func AbsoluteFromOldHinduLunar(d OldHinduLunarDate) (absoluteDate float64) {
	years := d.Year
	months := d.Month - 2
	approx := floorf(mult(big.NewRat(int64(years), 1), SolarSiderealYear)) +
		floorf(mult(big.NewRat(int64(months), 1), LunarSynodicMonth)) +
		(-1132959)
	try := approx + sum(
		func(float64) float64 { return 1 },
		approx,
		func(i float64) bool { return OldHinduLunarPrecedes(OldHinduLunarFromAbsolute(i), d) })
	if OldHinduLunarFromAbsolute(try) == d {
		return try
	} else {
		return math.NaN()
	}
}
