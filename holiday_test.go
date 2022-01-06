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

// create reference holiday dates
var holidays = createTestHolidays()

// run tests
func TestIndependenceDay(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.IndependenceDay[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := IndependenceDay(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLaborDay(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.LaborDay[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := LaborDay(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemorialDay(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.MemorialDay[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := MemorialDay(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaylightSavingsStart(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.DaylightSavingsStart[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := DaylightSavingsStart(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaylightSavingsEnd(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.DaylightSavingsEnd[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := DaylightSavingsEnd(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChristmas(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.Christmas[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := Christmas(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdvent(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.Advent[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := Advent(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEpiphany(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.Epiphany[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := Epiphany(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEasternOrthodoxChristmas(t *testing.T) {
	tests := make([]struct {
		year float64
		want []float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.EasternOrthodoxChristmas[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := EasternOrthodoxChristmas(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNicaeanRuleEaster(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.NicaeanRuleEaster[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := NicaeanRuleEaster(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEaster(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.Easter[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := Easter(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPentecost(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.Pentecost[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := Pentecost(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMuladAlNabi(t *testing.T) {
	tests := make([]struct {
		year float64
		want []float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.MuladAlNabi[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := MuladAlNabi(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYomKippur(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.YomKippur[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := YomKippur(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassover(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.Passover[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := Passover(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPurim(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.Purim[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := Purim(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaAnitEsther(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.TaAnitEsther[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := TaAnitEsther(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTishaBAv(t *testing.T) {
	tests := make([]struct {
		year float64
		want float64
	}, len(holidays.Year))
	for i, year := range holidays.Year {
		tests[i].year = year
		tests[i].want = holidays.TishaBAv[i]
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%.0f", tt.year)
		t.Run(testname, func(t *testing.T) {
			got := TishaBAv(tt.year)
			t.Logf("got %v, want %v", got, tt.want)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// hebrew birthday
// hebrew yahrzeit
