# libcalendar - Calendrical calculations in Go

## About
_libcalendar_ translates into Go the Lisp functions described and presented in: 

- [Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical Calculations", Software - Practice and Experience, 20 (9), 899--928.](https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.17.4274)
- [Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993. "Calendrical Calculations, II: Three Historical Calendars", Software - Practice & Experience, 23 (4), 383--404.](https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.9215) The Lisp source code can be found at https://www.cs.tau.ac.il/~nachum/calendar-book/papers/.

_libcalendar_ allows the computation of and conversion between dates from 11 calendars: Gregorian, ISO, Julian, Islamic, Hebrew, Mayan (long count, haab, tzolkin), French Revolutionary, and Old Hindu (solar, lunar).


## Installation

If you have a toy project where you want to use _libcalendar_, clone this repository and add a [replace directive](https://go.dev/ref/mod#go-mod-file-replace) to your project's `go.mod` file with the appropriate path to the location where you cloned _libcalendar_, e.g.

```go
module toyProjectName

go 1.17

replace staudtlex.de/libcalendar => /custom/path/to/libcalendar

// The `require` statement will be automatically inserted when running `go mod tidy`. No need to add it manually
require staudtlex.de/libcalendar v0.0.0.-...
```

## Examples
A basic example application is located in the `example` subdirectory, further examples may be added in the future.

## Limitations
The primary motivation for writing _libcalendar_ was to take first steps in understanding calendar-related algorithms and Go programming. 

- Note that the functions implemented in _libcalendar_ do not generally work for absolute dates smaller than 1 (except the Mayan calendars). 

- Furthermore, note that the Islamic and French Revolutionary calendar functions do not work with dates prior to their respective epochs. If provided with such dates, the functions will return invalid results.

- `DaylightSavingsStart` and `DaylightSavingsEnd` use the US rules for determining start and end of DST which are in place since 2007, whereas the corresponding Lisp-functions use the pre-2007 rules.
