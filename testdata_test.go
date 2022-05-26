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
)

// ========================================================
// Structs holding sample data
// ========================================================

type rawSampleDates struct {
	Rd             []float64   `json:"rd"`
	Gregorian      [][]float64 `json:"gregorian"`
	Julian         [][]float64 `json:"julian"`
	Iso            [][]float64 `json:"iso"`
	Islamic        [][]float64 `json:"islamic"`
	Hebrew         [][]float64 `json:"hebrew"`
	MayanLongCount [][]float64 `json:"mayanlongcount"`
	MayanHaab      [][]float64 `json:"mayanhaab"`
	MayanTzolkin   [][]float64 `json:"mayantzolkin"`
	French         [][]float64 `json:"french"`
	OldHinduSolar  [][]float64 `json:"oldhindusolar"`
	OldHinduLunar  [][]float64 `json:"oldhindulunar"`
}

type sampleDates struct {
	Rd             []float64
	Gregorian      []GregorianDate
	Julian         []JulianDate
	Iso            []IsoDate
	Islamic        []IslamicDate
	Hebrew         []HebrewDate
	MayanLongCount []MayanLongCount
	MayanHaab      []MayanHaabDate
	MayanTzolkin   []MayanTzolkinDate
	French         []FrenchDate
	OldHinduSolar  []OldHinduSolarDate
	OldHinduLunar  []OldHinduLunarDate
}

type sampleHolidays struct {
	Year                     []float64   `json:"year"`                     //
	IndependenceDay          []float64   `json:"independenceDay"`          //
	LaborDay                 []float64   `json:"laborDay"`                 //
	MemorialDay              []float64   `json:"memorialDay"`              //
	DaylightSavingsStart     []float64   `json:"daylightSavingsStart"`     //
	DaylightSavingsEnd       []float64   `json:"daylightSavingsEnd"`       //
	Christmas                []float64   `json:"christmas"`                //
	Advent                   []float64   `json:"advent"`                   //
	Epiphany                 []float64   `json:"epiphany"`                 //
	EasternOrthodoxChristmas [][]float64 `json:"easternOrthodoxChristmas"` //
	NicaeanRuleEaster        []float64   `json:"nicaeanRuleEaster"`        //
	Easter                   []float64   `json:"easter"`                   //
	Pentecost                []float64   `json:"pentecost"`                //
	MuladAlNabi              [][]float64 `json:"muladAlNabi"`              //
	YomKippur                []float64   `json:"yomKippur"`                //
	Passover                 []float64   `json:"passover"`                 //
	Purim                    []float64   `json:"purim"`                    //
	TaAnitEsther             []float64   `json:"taAnitEsther"`             //
	TishaBAv                 []float64   `json:"tishaBAv"`                 //
}

// ========================================================
// create date objects from float64-slices
// ========================================================
// createGregorian takes a []float64{month, day, year} and returns
// a GregorianDate.
func createGregorian(s []float64) GregorianDate {
	return GregorianDate{
		Year:  s[0],
		Month: s[1],
		Day:   s[2],
	}
}

// createJulian takes a []float64{month, day, year} and returns
// a JulianDate.
func createJulian(s []float64) JulianDate {
	return JulianDate{
		Year:  s[0],
		Month: s[1],
		Day:   s[2],
	}
}

// createIslamic takes a []float64{month, day, year} and returns
// an IslamicDate.
func createIslamic(s []float64) IslamicDate {
	return IslamicDate{
		Year:  s[0],
		Month: s[1],
		Day:   s[2],
	}
}

// createHebrew takes a []float64{month, day, year} and returns a HebrewDate.
func createHebrew(s []float64) HebrewDate {
	return HebrewDate{
		Year:  s[0],
		Month: s[1],
		Day:   s[2],
	}
}

// createIso takes a []float64{week, day, year} and returns an IsoDate.
func createIso(s []float64) IsoDate {
	return IsoDate{
		Year: s[0],
		Week: s[1],
		Day:  s[2],
	}
}

func createFrench(s []float64) FrenchDate {
	return FrenchDate{
		Year:  s[0],
		Month: s[1],
		Day:   s[2],
	}
}

// createLongCount takes a []float64{baktun, katun, tun, uinal, kin} and
// returns a MayanLongCount.
func createMayanLongCount(s []float64) MayanLongCount {
	return MayanLongCount{
		Baktun: s[0],
		Katun:  s[1],
		Tun:    s[2],
		Uinal:  s[3],
		Kin:    s[4],
	}
}

// createHaab takes a []float64{day, month} and returns a HaabDate.
func createMayanHaab(s []float64) MayanHaabDate {
	return MayanHaabDate{
		Month: s[0],
		Day:   s[1],
	}
}

// createTzolkin takes a []float64{number, name} and returns a TzolkinDate.
// Use this function to create Mayan TzolkinDate or Aztec TonalpohualliDate.
func createMayanTzolkin(s []float64) MayanTzolkinDate {
	return MayanTzolkinDate{
		Number: s[0],
		Name:   s[1],
	}
}

// createOldHinduLunar takes a []float64{month, leap, day, year} and returns an
// OldHinduLunarDate.
func createOldHinduLunar(s []float64) OldHinduLunarDate {
	return OldHinduLunarDate{
		Year:      s[0],
		Month:     s[1],
		LeapMonth: s[2] == 1,
		Day:       s[3],
	}
}

// createOldHinduSolar takes a []float64{year, month, day} and returns an
// OldHinduSolarDate.
func createOldHinduSolar(s []float64) OldHinduSolarDate {
	return OldHinduSolarDate{
		Year:  s[0],
		Month: s[1],
		Day:   s[2],
	}
}

// ========================================================
// create date objects of slices of float64-slices
// ========================================================
func mapGregorian(s [][]float64) []GregorianDate {
	dates := make([]GregorianDate, len(s))
	for i, d := range s {
		dates[i] = createGregorian(d)
	}
	return dates
}

func mapIso(s [][]float64) []IsoDate {
	dates := make([]IsoDate, len(s))
	for i, d := range s {
		dates[i] = createIso(d)
	}
	return dates
}

func mapJulian(s [][]float64) []JulianDate {
	dates := make([]JulianDate, len(s))
	for i, d := range s {
		dates[i] = createJulian(d)
	}
	return dates
}

func mapIslamic(s [][]float64) []IslamicDate {
	dates := make([]IslamicDate, len(s))
	for i, d := range s {
		dates[i] = createIslamic(d)
	}
	return dates
}

func mapHebrew(s [][]float64) []HebrewDate {
	dates := make([]HebrewDate, len(s))
	for i, d := range s {
		dates[i] = createHebrew(d)
	}
	return dates
}

func mapFrench(s [][]float64) []FrenchDate {
	dates := make([]FrenchDate, len(s))
	for i, d := range s {
		dates[i] = createFrench(d)
	}
	return dates
}

func mapMayanLongCount(s [][]float64) []MayanLongCount {
	dates := make([]MayanLongCount, len(s))
	for i, d := range s {
		dates[i] = createMayanLongCount(d)
	}
	return dates
}

func mapMayanHaab(s [][]float64) []MayanHaabDate {
	dates := make([]MayanHaabDate, len(s))
	for i, d := range s {
		dates[i] = createMayanHaab(d)
	}
	return dates
}

func mapMayanTzolkin(s [][]float64) []MayanTzolkinDate {
	dates := make([]MayanTzolkinDate, len(s))
	for i, d := range s {
		dates[i] = createMayanTzolkin(d)
	}
	return dates
}

func mapOldHinduSolar(s [][]float64) []OldHinduSolarDate {
	dates := make([]OldHinduSolarDate, len(s))
	for i, d := range s {
		dates[i] = createOldHinduSolar(d)
	}
	return dates
}

func mapOldHinduLunar(s [][]float64) []OldHinduLunarDate {
	dates := make([]OldHinduLunarDate, len(s))
	for i, d := range s {
		dates[i] = createOldHinduLunar(d)
	}
	return dates
}

// ========================================================
// Functions to load JSON-data and convert it to
// libcalendar sample dates
// ========================================================
func equal(s1, s2 []float64) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, elem := range s1 {
		if elem != s2[i] {
			return false
		}
	}
	return true
}

// unmarshalDates loads a json-File which contains the sample data against
// which the libcalendar-functions are being tested.
// unmarshalDates returns a struct of type rawSampleDates.
// The returned struct contains the sample data stored as slices of float64
// or []float64{}.
func unmarshalDates(jsonString string) rawSampleDates {
	result := rawSampleDates{}
	json.Unmarshal([]byte(jsonString), &result)
	return result
}

func unmarshalHolidays(jsonString string) sampleHolidays {
	result := sampleHolidays{}
	json.Unmarshal([]byte(jsonString), &result)
	//fmt.Println(result)
	return result
}

// parseReferenceDates creates the appropriate date structs from slices of
// float64 and []float64 stored in a rawSampleDates-struct.
// It returns a struct of type sampleDates which contain slices of date
// structs.
func parseReferenceDates(d rawSampleDates) sampleDates {
	return sampleDates{
		Rd:             d.Rd,
		Gregorian:      mapGregorian(d.Gregorian),
		Julian:         mapJulian(d.Julian),
		Iso:            mapIso(d.Iso),
		Islamic:        mapIslamic(d.Islamic),
		Hebrew:         mapHebrew(d.Hebrew),
		MayanLongCount: mapMayanLongCount(d.MayanLongCount),
		MayanHaab:      mapMayanHaab(d.MayanHaab),
		MayanTzolkin:   mapMayanTzolkin(d.MayanTzolkin),
		French:         mapFrench(d.French),
		OldHinduSolar:  mapOldHinduSolar(d.OldHinduSolar),
		OldHinduLunar:  mapOldHinduLunar(d.OldHinduLunar),
	}
}

func parseHolidays(d sampleHolidays) sampleHolidays {
	holidays := d
	holidays.EasternOrthodoxChristmas =
		applyToSlice(zeroToEmptyList, d.EasternOrthodoxChristmas)
	return holidays
}

func applyToSlice(f func([]float64) []float64, s [][]float64) [][]float64 {
	result := make([][]float64, len(s))
	for i, elem := range s {
		result[i] = f(elem)
	}
	return result
}

func zeroToEmptyList(l []float64) (list []float64) {
	if equal(l, []float64{0.0}) {
		return []float64{}
	} else {
		return l
	}
}

// ========================================================
// Load sample data
// ========================================================
// dates holds sample dates created with calendar.l for the range of years
// provided in table 1 of Reingold/Dershowitz (2018), appendix C, pp. 447-453.
// createTestDates returns the reference dates against which the calendar
// functions are being tested.
func createTestDates() (dates sampleDates) {
	dates = parseReferenceDates(unmarshalDates(referenceDates))
	return dates
}

// createTestHolidays returns the reference dates of holidays against which
// the holiday functions are being tested. The reference dates hold holiday
// dates for the years 1900-2199.
func createTestHolidays() (holidays sampleHolidays) {
	holidays = parseHolidays(unmarshalHolidays(referenceHolidays))
	return holidays
}
