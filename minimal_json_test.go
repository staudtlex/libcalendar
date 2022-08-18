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

func ExampleDate_Json() {
	gregorianDate := lc.GregorianDate{Year: 2022, Month: 06, Day: 15}

	gregorian_json := gregorianDate.Date().Json()
	fmt.Println(gregorian_json)
}
func ExampleJsonToDate() {
	jsonString := `{
		"components": [2022,6,15],
		"calendar": "gregorian",
		"componentNames" :[],
		"monthNames": []
	}`
	fmt.Println(jsonString)
	unmarshal_json := lc.JsonToDate(jsonString)
	fmt.Println(unmarshal_json)

	// Convert a Date to an absolute date, and compute a Mayan Long count
	// date from that absolute date
	mayanLongCount :=
		lc.FromAbsolute(
			lc.AbsoluteFromDate(unmarshal_json), "mayanLongCount")
	fmt.Println("Mayan long count: ", mayanLongCount)
}
