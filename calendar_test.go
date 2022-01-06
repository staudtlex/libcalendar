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
	"fmt"
	"testing"
)

// Create reference dates
var dates = createTestDates()

// Run tests

// Gregorian calendar
func TestAbsoluteFromGregorian(t *testing.T) {
	tests := make([]struct {
		date GregorianDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.Gregorian {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromGregorian(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGregorianFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want GregorianDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.Gregorian[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := GregorianFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// ISO calendar
func TestAbsoluteFromIso(t *testing.T) {
	tests := make([]struct {
		date IsoDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.Iso {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromIso(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsoFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want IsoDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.Iso[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := IsoFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Julian calendar
func TestAbsoluteFromJulian(t *testing.T) {
	tests := make([]struct {
		date JulianDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.Julian {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromJulian(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJulianFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want JulianDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.Julian[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := JulianFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Islamic calendar
func TestAbsoluteFromIslamic(t *testing.T) {
	tests := make([]struct {
		date IslamicDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.Islamic {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromIslamic(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIslamicFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want IslamicDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.Islamic[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := IslamicFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Hebrew calendar
func TestAbsoluteFromHebrew(t *testing.T) {
	tests := make([]struct {
		date HebrewDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.Hebrew {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromHebrew(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHebrewFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want HebrewDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.Hebrew[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := HebrewFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Mayan calendars
func TestAbsoluteFromMayanLongCount(t *testing.T) {
	tests := make([]struct {
		date MayanLongCount
		want float64
	}, len(dates.Rd))
	for i, date := range dates.MayanLongCount {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromMayanLongCount(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMayanLongCountFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want MayanLongCount
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.MayanLongCount[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := MayanLongCountFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMayanHaabFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want MayanHaabDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.MayanHaab[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := MayanHaabFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMayanTzolkinFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want MayanTzolkinDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.MayanTzolkin[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := MayanTzolkinFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// French Revolutionary calendar
func TestAbsoluteFromFrench(t *testing.T) {
	tests := make([]struct {
		date FrenchDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.French {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromFrench(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrenchFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want FrenchDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.French[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := FrenchFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// Old Hindu calendars
func TestAbsoluteFromOldHinduSolar(t *testing.T) {
	tests := make([]struct {
		date OldHinduSolarDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.OldHinduSolar {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromOldHinduSolar(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOldHinduSolarFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want OldHinduSolarDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.OldHinduSolar[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := OldHinduSolarFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsoluteFromOldHinduLunar(t *testing.T) {
	tests := make([]struct {
		date OldHinduLunarDate
		want float64
	}, len(dates.Rd))
	for i, date := range dates.OldHinduLunar {
		tests[i].date = date
		tests[i].want = dates.Rd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.date)
		t.Run(testname, func(t *testing.T) {
			got := AbsoluteFromOldHinduLunar(tt.date)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOldHinduLunarFromAbsolute(t *testing.T) {
	tests := make([]struct {
		rd   float64
		want OldHinduLunarDate
	}, len(dates.Rd))
	for i, rd := range dates.Rd {
		tests[i].rd = rd
		tests[i].want = dates.OldHinduLunar[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.rd)
		t.Run(testname, func(t *testing.T) {
			got := OldHinduLunarFromAbsolute(tt.rd)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
