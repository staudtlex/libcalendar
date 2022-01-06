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
//   Calculations", Software---Practice and Experience, 20 (9), 899--928.
// - Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993.
//   "Calendrical Calculations, II: Three Historical Calendars",
//   Software---Practice & Experience, 23 (4), 383--404.

package libcalendar

import (
	"math"
	"math/big"
)

// The Gregorian Calendar

// Gregorian and Julian months
const (
	january   float64 = 1
	february  float64 = 2
	march     float64 = 3
	april     float64 = 4
	may       float64 = 5
	june      float64 = 6
	july      float64 = 7
	august    float64 = 8
	september float64 = 9
	october   float64 = 10
	november  float64 = 11
	december  float64 = 12
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
func AbsoluteFromGregorian(d GregorianDate) (date float64) {
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
func GregorianFromAbsolute(date float64) GregorianDate {
	// compute year
	d_0 := date - 1
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
		return date > AbsoluteFromGregorian(d)
	}
	month := sum(f, 1, p) + 1
	// compute day
	day := date - (AbsoluteFromGregorian(GregorianDate{year, month, 1}) - 1)
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
func KDayOnOrBefore(date float64, k float64) float64 {
	return date - mod(date-k, 7)
}

// AbsoluteFromIso computes the absolute (fixed) date from an ISO date.
func AbsoluteFromIso(d IsoDate) (date float64) {
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
func IsoFromAbsolute(date float64) IsoDate {
	approx := GregorianFromAbsolute(date - 3).Year
	year := 0.0
	if date >= AbsoluteFromIso(IsoDate{approx + 1, 1, 1}) {
		year = approx + 1
	} else {
		year = approx
	}
	week := 1 + math.Floor(
		(date-AbsoluteFromIso(IsoDate{year, 1, 1}))/
			7)
	day := 0.0
	if mod(date, 7) == 0 {
		day = 7
	} else {
		day = mod(date, 7)
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
type JulianDate = GregorianDate

// AbsoluteFromJulian computes the absolute (fixed) date corresponding to a
// given Julian date.
func AbsoluteFromJulian(d JulianDate) (date float64) {
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

func JulianFromAbsolute(date float64) JulianDate {
	approx := math.Floor((date + 2) / 366)
	f := func(float64) float64 { return 1 }
	p := func(y float64) bool {
		return date >= AbsoluteFromJulian(JulianDate{y + 1, january, 1})
	}
	year := approx + sum(f, approx, p)
	month := 1 + sum(f, 1, func(m float64) bool {
		return date > AbsoluteFromJulian(
			JulianDate{year, m, LastDayOfJulianMonth(m, year)})
	})
	day := date - (AbsoluteFromJulian(JulianDate{year, month, 1}) - 1)
	return JulianDate{year, month, day}
}

// The Islamic Calendar

// Islamic months
const (
	muharram    float64 = 1
	safar       float64 = 2
	rabi_i      float64 = 3
	rabi_ii     float64 = 4
	jumada_i    float64 = 5
	jumada_ii   float64 = 6
	rajab       float64 = 7
	shaBan      float64 = 8
	ramadan     float64 = 9
	shawwal     float64 = 10
	dhuAlQada   float64 = 11
	dhuAlHijjah float64 = 12
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
func AbsoluteFromIslamic(d IslamicDate) (date float64) {
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
func IslamicFromAbsolute(date float64) IslamicDate {
	if date <= 227014 {
		return IslamicDate{0, 0, 0}
	}
	approx := math.Floor((date - 227014) / 355)
	f := func(float64) float64 { return 1 }
	p := func(y float64) bool {
		return date >= AbsoluteFromIslamic(IslamicDate{y + 1, muharram, 1})
	}
	year := approx + sum(f, approx, p)
	month := 1 + sum(f, 1, func(m float64) bool {
		return date > AbsoluteFromIslamic(
			IslamicDate{year, m, LastDayOfIslamicMonth(m, year)})
	})
	day := date - (AbsoluteFromIslamic(IslamicDate{year, month, 1}) - 1)
	return IslamicDate{year, month, day}
}

// The Hebrew Calendar

// Hebrew months
const (
	nisan        float64 = 1
	iyyar        float64 = 2
	sivan        float64 = 3
	tammuz       float64 = 4
	av           float64 = 5
	elul         float64 = 6
	tishri       float64 = 7
	heshvan      float64 = 8
	kislev       float64 = 9
	teveth       float64 = 10
	shevat       float64 = 11
	adar_i_leap  float64 = 12 // only in leap years, otherwise adar_i_leap is skipped
	adar_ii      float64 = 12
	adar_ii_leap float64 = 13 // only in leap years, otherwise adar_i_leap is skipped
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
func AbsoluteFromHebrew(d HebrewDate) (date float64) {
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
func HebrewFromAbsolute(date float64) HebrewDate {
	approx := math.Floor((date + 1373429) / 366)
	f := func(float64) float64 { return 1 }
	year := approx + sum(f, approx, func(y float64) bool {
		return date >= AbsoluteFromHebrew(HebrewDate{y + 1, 7, 1})
	})
	start := 0.0
	if (date < AbsoluteFromHebrew(HebrewDate{year, 1, 1})) {
		start = 7
	} else {
		start = 1
	}
	month := start + sum(f, start, func(m float64) bool {
		return date > AbsoluteFromHebrew(HebrewDate{year, m, LastDayOfHebrewMonth(m, year)})
	})
	day := date - (AbsoluteFromHebrew(HebrewDate{year, month, 1}) - 1)
	return HebrewDate{year, month, day}
}

// The Mayan Calendars

// Mayan haab months
const (
	pop    float64 = 1
	uo     float64 = 2
	zip    float64 = 3
	zotz   float64 = 4
	tzec   float64 = 5
	xul    float64 = 6
	yaxkin float64 = 7
	mol    float64 = 8
	chen   float64 = 9
	yax    float64 = 10
	zac    float64 = 11
	ceh    float64 = 12
	mac    float64 = 13
	kankin float64 = 14
	muan   float64 = 15
	pax    float64 = 16
	kayab  float64 = 17
	cumku  float64 = 18
)

// Mayan tzolkin names
const (
	imix    float64 = 1
	ik      float64 = 2
	akbal   float64 = 3
	kan     float64 = 4
	chiccan float64 = 5
	cimi    float64 = 6
	manik   float64 = 7
	lamat   float64 = 8
	muluc   float64 = 9
	oc      float64 = 10
	chuen   float64 = 11
	eb      float64 = 12
	ben     float64 = 13
	ix      float64 = 14
	men     float64 = 15
	cib     float64 = 16
	caban   float64 = 17
	etznab  float64 = 18
	cauac   float64 = 19
	ahau    float64 = 20
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
func MayanLongCountFromAbsolute(date float64) MayanLongCount {
	longCount := date + MayanDaysBeforeAbsoluteZero
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
func MayanHaabFromAbsolute(date float64) MayanHaabDate {
	longCount := date + MayanDaysBeforeAbsoluteZero
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
func MayanTzolkinFromAbsolute(date float64) MayanTzolkinDate {
	longCount := date + MayanDaysBeforeAbsoluteZero
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
func MayanHaabTzolkinOnOrBefore(haab MayanHaabDate, tzolkin MayanTzolkinDate, d float64) (date float64) {
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
	vendémiaire float64 = 1
	brumaire    float64 = 2
	frimaire    float64 = 3
	nivôse      float64 = 4
	pluviôse    float64 = 5
	ventôse     float64 = 6
	germinal    float64 = 7
	floréal     float64 = 8
	prairial    float64 = 9
	messidor    float64 = 10
	thermidor   float64 = 11
	fructidor   float64 = 12
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
func AbsoluteFromFrench(d FrenchDate) (date float64) {
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
func FrenchFromAbsolute(date float64) FrenchDate {
	if date < 654415 {
		return FrenchDate{0, 0, 0}
	}
	approx := math.Floor((date - 654414) / 366)
	f := func(float64) float64 { return 1 }
	year := approx + sum(f, approx,
		func(y float64) bool {
			return date >= AbsoluteFromFrench(FrenchDate{y + 1, vendémiaire, 1})
		})
	month := 1 + sum(f, 1,
		func(m float64) bool {
			return date > AbsoluteFromFrench(FrenchDate{year, m, FrenchLastDayOfMonth(m, year)})
		})
	day := date - (AbsoluteFromFrench(FrenchDate{year, month, 1}) - 1)
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

/*
// Old Hindu solar months
const (
	mesha     float64 = 1
	vrshabha  float64 = 2
	mithuna   float64 = 3
	karka     float64 = 4
	simha     float64 = 5
	kanya     float64 = 6
	tula      float64 = 7
	vrischika float64 = 8
	dhanus    float64 = 9
	makara    float64 = 10
	kumbha    float64 = 11
	mina      float64 = 12
)

// Old Hindu lunar months
const (
	chaitra    float64 = 1
	vaisakha   float64 = 2
	jyaishtha  float64 = 3
	ashadha    float64 = 4
	sravana    float64 = 5
	bhadrapada float64 = 6
	asvina     float64 = 7
	karttika   float64 = 8
	margasira  float64 = 9
	pausha     float64 = 10
	magha      float64 = 11
	phalguna   float64 = 12
)
*/

var SolarSiderealYear = Add(Ratio(365), Ratio(279457, 1080000))
var SolarMonth = Div(SolarSiderealYear, Ratio(12))
var LunarSiderealMonth = Add(Ratio(27), Ratio(4644439, 14438334))
var LunarSynodicMonth = Add(Ratio(29), Ratio(7087771, 13358334))

// SolarLongitude returns the position of the sun (in degrees) for a
// given moment (day and fraction of a day).
func SolarLongitude(t *big.Rat) (degrees *big.Rat) {
	return Mult(Modr(Div(t, SolarSiderealYear), Ratio(1)), Ratio(360))
}

// Zodiac returns the zodiacal sign for a given moment (day and fraction
// of day).
func Zodiac(t *big.Rat) (zodiac float64) {
	zodiac = Quotient(SolarLongitude(t), Ratio(30))
	return zodiac + 1
}

// OldHinduSolarFromAbsolute computes the Old Hindu solar date corresponding
// to a given absolute (fixed) date.
func OldHinduSolarFromAbsolute(date float64) OldHinduSolarDate {
	hdate := Add(Ratio(date), Ratio(1132959), Ratio(1, 4))
	year := Quotient(hdate, SolarSiderealYear)
	month := Zodiac(hdate)
	day := Floorf(Modr(hdate, SolarMonth)) + 1
	return OldHinduSolarDate{year, month, day}
}

// AbsoluteFromOldHinduSolar returns the absolute (fixed) date from a given
// Old Hindu solar date.
func AbsoluteFromOldHinduSolar(d OldHinduSolarDate) (date float64) {
	year := Ratio(d.Year)
	month := Ratio(d.Month)
	day := Ratio(d.Day)
	one := Ratio(1)
	result, _ := Add(
		Mult(year, SolarSiderealYear),
		Mult(Sub(month, one), SolarMonth),
		day,
		Ratio(-1, 4),
		Ratio(-1132959)).Float64()
	return math.Floor(result)
}

// LunarLongitude returns the sidereal longitude of the moon (in degrees) at
// a given moment (date and fraction of a day).
func LunarLongitude(t *big.Rat) (degrees *big.Rat) {
	return Mult(Modr(Div(t, LunarSiderealMonth), Ratio(1)), Ratio(360))
}

// LunarPhase computes the lunar phase of the moon for a given moment (date
// and fraction of a day).
func LunarPhase(t *big.Rat) (phase float64) {
	return 1 + Quotient(
		Modr(
			Sub(LunarLongitude(t), SolarLongitude(t)),
			Ratio(360)),
		Ratio(12))
}

// NewMoon determines the time of the most recent new moon for a given moment
// (date and fraction of day).
func NewMoon(t *big.Rat) *big.Rat {
	return Sub(t, Modr(t, LunarSynodicMonth))
}

// OldHinduLunarFromAbsolute returnsthe Old Hindu lunar date corresponding to
// a given absolute (fixed) date.
func OldHinduLunarFromAbsolute(date float64) OldHinduLunarDate {
	hdate := Ratio(date + 1132959)
	sunrise := Add(hdate, Ratio(1, 4))
	lastNewMoon := NewMoon(sunrise)
	nextNewMoon := Add(lastNewMoon, LunarSynodicMonth)
	month := amod(Zodiac(lastNewMoon)+1, 12)
	day := LunarPhase(sunrise)
	leapMonth := Zodiac(lastNewMoon) == Zodiac(nextNewMoon)
	nextMonth := nextNewMoon
	if leapMonth {
		nextMonth = Add(nextMonth, LunarSynodicMonth)
	} else {
		nextMonth = Add(nextMonth, Ratio(0))
	}
	year := Quotient(nextMonth, SolarSiderealYear)
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
func AbsoluteFromOldHinduLunar(d OldHinduLunarDate) (date float64) {
	years := d.Year
	months := d.Month - 2
	approx := Floorf(Mult(Ratio(years), SolarSiderealYear)) +
		Floorf(Mult(Ratio(months), LunarSynodicMonth)) +
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
