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

// Gregorian and Julian calendar
var gregorianMonths = map[float64]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func (d GregorianDate) String() string {
	return fmt.Sprintf("%v %v %v", d.Day, gregorianMonths[d.Month], d.Year)
}

// ISO calendar
func (d IsoDate) String() string {
	return fmt.Sprintf("%v-W%v-%v", d.Year, fmt.Sprintf("%02d", int(d.Week)), d.Day)
}

// Islamic calendar
var islamicMonths = map[float64]string{
	1:  "Muharram",
	2:  "Safar",
	3:  "Rabi I",
	4:  "Rabi II",
	5:  "Jumada I",
	6:  "Jumada II",
	7:  "Rajab",
	8:  "Sha' Ban",
	9:  "Ramadan",
	10: "Shawwal",
	11: "Dhu al-Qada",
	12: "Dhu al-Hijjah",
}

func (d IslamicDate) String() string {
	return fmt.Sprintf("%v %v %v", d.Day, islamicMonths[d.Month], d.Year)
}

// Hebrew calendar
var hebrewMonths = map[float64]string{
	1:  "Nisan",
	2:  "Iyyar",
	3:  "Sivan",
	4:  "Tammuz",
	5:  "Av",
	6:  "Elul",
	7:  "Tishri",
	8:  "Heshvan",
	9:  "Kislev",
	10: "Teveth",
	11: "Shevat",
	12: "Adar ",
}

func (d HebrewDate) String() string {
	month := ""
	switch { // deal with leap years and intercalary month Adar I
	case d.Month == 12 && HebrewLeapYear(d.Year):
		month = "Adar I"
	case d.Month == 13 && HebrewLeapYear(d.Year):
		month = "Adar II"
	default:
		month = hebrewMonths[d.Month]
	}
	return fmt.Sprintf("%v %v %v", d.Day, month, d.Year)
}

// Mayan calendars
var mayanHaabMonths = map[float64]string{
	1:  "Pop",
	2:  "Uo",
	3:  "Zip",
	4:  "Zotz",
	5:  "Tzec",
	6:  "Xul",
	7:  "Yaxkin",
	8:  "Mol",
	9:  "Chen",
	10: "Yax",
	11: "Zac",
	12: "Ceh",
	13: "Mac",
	14: "Kankin",
	15: "Muan",
	16: "Pax",
	17: "Kayab",
	18: "Cumku",
}

var mayanTzolkinNames = map[float64]string{
	1:  "Imix",
	2:  "Ik",
	3:  "Akbal",
	4:  "Kan",
	5:  "Chiccan",
	6:  "Cimi",
	7:  "Manik",
	8:  "Lamat",
	9:  "Muluc",
	10: "Oc",
	11: "Chuen",
	12: "Eb",
	13: "Ben",
	14: "Ix",
	15: "Men",
	16: "Cib",
	17: "Caban",
	18: "Etznab",
	19: "Cauac",
	20: "Ahau",
}

func (d MayanLongCount) String() string {
	return fmt.Sprintf("%v.%v.%v.%v.%v", d.Baktun, d.Katun, d.Tun, d.Uinal, d.Kin)
}

func (d MayanHaabDate) String() string {
	return fmt.Sprintf("%v %v", d.Day, mayanHaabMonths[d.Month])
}

func (d MayanTzolkinDate) String() string {
	return fmt.Sprintf("%v %v", d.Number, mayanTzolkinNames[d.Name])
}

// French Revolutionary calendar
var frenchMonths = map[float64]string{
	1:  "Vendémiaire",
	2:  "Brumaire",
	3:  "Frimaire",
	4:  "Nivôse",
	5:  "Pluviôse",
	6:  "Ventôse",
	7:  "Germinal",
	8:  "Floréal",
	9:  "Prairial",
	10: "Messidor",
	11: "Thermidor",
	12: "Fructidor",
	13: "jour complémentaire",
}

func (d FrenchDate) String() string {
	return fmt.Sprintf("%v %v an %v", d.Day, frenchMonths[d.Month], d.Year)
}

// Old Hindu calendars
var hinduSolarMonths = map[float64]string{
	1:  "Mesha",
	2:  "Vrshabha",
	3:  "Mithuna",
	4:  "Karka",
	5:  "Simha",
	6:  "Kanya",
	7:  "Tula",
	8:  "Vrischika",
	9:  "Dhanus",
	10: "Makara",
	11: "Kumbha",
	12: "Mina",
}

func (d OldHinduSolarDate) String() string {
	return fmt.Sprintf("%v %v %v", d.Day, hinduSolarMonths[d.Month], d.Year)
}

var hinduLunarMonths = map[float64]string{
	1:  "Chaitra",
	2:  "Vaisakha",
	3:  "Jyaishtha",
	4:  "Ashadha",
	5:  "Sravana",
	6:  "Bhadrapada",
	7:  "Asvina",
	8:  "Kartika",
	9:  "Margasira",
	10: "Pausha",
	11: "Magha",
	12: "Phalguna",
}

func (d OldHinduLunarDate) String() string {
	return fmt.Sprintf("%v %v %v", d.Day, hinduLunarMonths[d.Month], d.Year)
}
