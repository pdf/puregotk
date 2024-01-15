// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// An opaque structure that represents a date and time, including a time zone.
type DateTime struct {
}

func (x *DateTime) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewDateTime func(*TimeZone, int, int, int, int, int, float64) *DateTime

// Creates a new #GDateTime corresponding to the given date and time in
// the time zone @tz.
//
// The @year must be between 1 and 9999, @month between 1 and 12 and @day
// between 1 and 28, 29, 30 or 31 depending on the month and the year.
//
// @hour must be between 0 and 23 and @minute must be between 0 and 59.
//
// @seconds must be at least 0.0 and must be strictly less than 60.0.
// It will be rounded down to the nearest microsecond.
//
// If the given time is not representable in the given time zone (for
// example, 02:30 on March 14th 2010 in Toronto, due to daylight savings
// time) then the time will be rounded up to the nearest existing time
// (in this case, 03:00).  If this matters to you then you should verify
// the return value for containing the same as the numbers you gave.
//
// In the case that the given time is ambiguous in the given time zone
// (for example, 01:30 on November 7th 2010 in Toronto, due to daylight
// savings time) then the time falling within standard (ie:
// non-daylight) time is taken.
//
// It not considered a programmer error for the values to this function
// to be out of range, but in the case that they are, the function will
// return %NULL.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
func NewDateTime(TzVar *TimeZone, YearVar int, MonthVar int, DayVar int, HourVar int, MinuteVar int, SecondsVar float64) *DateTime {

	cret := xNewDateTime(TzVar, YearVar, MonthVar, DayVar, HourVar, MinuteVar, SecondsVar)
	return cret
}

var xNewDateTimeFromIso8601 func(string, *TimeZone) *DateTime

// Creates a #GDateTime corresponding to the given
// [ISO 8601 formatted string](https://en.wikipedia.org/wiki/ISO_8601)
// @text. ISO 8601 strings of the form &lt;date&gt;&lt;sep&gt;&lt;time&gt;&lt;tz&gt; are supported, with
// some extensions from [RFC 3339](https://tools.ietf.org/html/rfc3339) as
// mentioned below.
//
// Note that as #GDateTime "is oblivious to leap seconds", leap seconds information
// in an ISO-8601 string will be ignored, so a `23:59:60` time would be parsed as
// `23:59:59`.
//
// &lt;sep&gt; is the separator and can be either 'T', 't' or ' '. The latter two
// separators are an extension from
// [RFC 3339](https://tools.ietf.org/html/rfc3339#section-5.6).
//
// &lt;date&gt; is in the form:
//
//   - `YYYY-MM-DD` - Year/month/day, e.g. 2016-08-24.
//   - `YYYYMMDD` - Same as above without dividers.
//   - `YYYY-DDD` - Ordinal day where DDD is from 001 to 366, e.g. 2016-237.
//   - `YYYYDDD` - Same as above without dividers.
//   - `YYYY-Www-D` - Week day where ww is from 01 to 52 and D from 1-7,
//     e.g. 2016-W34-3.
//   - `YYYYWwwD` - Same as above without dividers.
//
// &lt;time&gt; is in the form:
//
// - `hh:mm:ss(.sss)` - Hours, minutes, seconds (subseconds), e.g. 22:10:42.123.
// - `hhmmss(.sss)` - Same as above without dividers.
//
// &lt;tz&gt; is an optional timezone suffix of the form:
//
// - `Z` - UTC.
// - `+hh:mm` or `-hh:mm` - Offset from UTC in hours and minutes, e.g. +12:00.
// - `+hh` or `-hh` - Offset from UTC in hours, e.g. +12.
//
// If the timezone is not provided in @text it must be provided in @default_tz
// (this field is otherwise ignored).
//
// This call can fail (returning %NULL) if @text is not a valid ISO 8601
// formatted string.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
func NewDateTimeFromIso8601(TextVar string, DefaultTzVar *TimeZone) *DateTime {

	cret := xNewDateTimeFromIso8601(TextVar, DefaultTzVar)
	return cret
}

var xNewDateTimeFromTimevalLocal func(*TimeVal) *DateTime

// Creates a #GDateTime corresponding to the given #GTimeVal @tv in the
// local time zone.
//
// The time contained in a #GTimeVal is always stored in the form of
// seconds elapsed since 1970-01-01 00:00:00 UTC, regardless of the
// local time offset.
//
// This call can fail (returning %NULL) if @tv represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
func NewDateTimeFromTimevalLocal(TvVar *TimeVal) *DateTime {

	cret := xNewDateTimeFromTimevalLocal(TvVar)
	return cret
}

var xNewDateTimeFromTimevalUtc func(*TimeVal) *DateTime

// Creates a #GDateTime corresponding to the given #GTimeVal @tv in UTC.
//
// The time contained in a #GTimeVal is always stored in the form of
// seconds elapsed since 1970-01-01 00:00:00 UTC.
//
// This call can fail (returning %NULL) if @tv represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
func NewDateTimeFromTimevalUtc(TvVar *TimeVal) *DateTime {

	cret := xNewDateTimeFromTimevalUtc(TvVar)
	return cret
}

var xNewDateTimeFromUnixLocal func(int64) *DateTime

// Creates a #GDateTime corresponding to the given Unix time @t in the
// local time zone.
//
// Unix time is the number of seconds that have elapsed since 1970-01-01
// 00:00:00 UTC, regardless of the local time offset.
//
// This call can fail (returning %NULL) if @t represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
func NewDateTimeFromUnixLocal(TVar int64) *DateTime {

	cret := xNewDateTimeFromUnixLocal(TVar)
	return cret
}

var xNewDateTimeFromUnixUtc func(int64) *DateTime

// Creates a #GDateTime corresponding to the given Unix time @t in UTC.
//
// Unix time is the number of seconds that have elapsed since 1970-01-01
// 00:00:00 UTC.
//
// This call can fail (returning %NULL) if @t represents a time outside
// of the supported range of #GDateTime.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
func NewDateTimeFromUnixUtc(TVar int64) *DateTime {

	cret := xNewDateTimeFromUnixUtc(TVar)
	return cret
}

var xNewDateTimeLocal func(int, int, int, int, int, float64) *DateTime

// Creates a new #GDateTime corresponding to the given date and time in
// the local time zone.
//
// This call is equivalent to calling g_date_time_new() with the time
// zone returned by g_time_zone_new_local().
func NewDateTimeLocal(YearVar int, MonthVar int, DayVar int, HourVar int, MinuteVar int, SecondsVar float64) *DateTime {

	cret := xNewDateTimeLocal(YearVar, MonthVar, DayVar, HourVar, MinuteVar, SecondsVar)
	return cret
}

var xNewDateTimeNow func(*TimeZone) *DateTime

// Creates a #GDateTime corresponding to this exact instant in the given
// time zone @tz.  The time is as accurate as the system allows, to a
// maximum accuracy of 1 microsecond.
//
// This function will always succeed unless GLib is still being used after the
// year 9999.
//
// You should release the return value by calling g_date_time_unref()
// when you are done with it.
func NewDateTimeNow(TzVar *TimeZone) *DateTime {

	cret := xNewDateTimeNow(TzVar)
	return cret
}

var xNewDateTimeNowLocal func() *DateTime

// Creates a #GDateTime corresponding to this exact instant in the local
// time zone.
//
// This is equivalent to calling g_date_time_new_now() with the time
// zone returned by g_time_zone_new_local().
func NewDateTimeNowLocal() *DateTime {

	cret := xNewDateTimeNowLocal()
	return cret
}

var xNewDateTimeNowUtc func() *DateTime

// Creates a #GDateTime corresponding to this exact instant in UTC.
//
// This is equivalent to calling g_date_time_new_now() with the time
// zone returned by g_time_zone_new_utc().
func NewDateTimeNowUtc() *DateTime {

	cret := xNewDateTimeNowUtc()
	return cret
}

var xNewDateTimeUtc func(int, int, int, int, int, float64) *DateTime

// Creates a new #GDateTime corresponding to the given date and time in
// UTC.
//
// This call is equivalent to calling g_date_time_new() with the time
// zone returned by g_time_zone_new_utc().
func NewDateTimeUtc(YearVar int, MonthVar int, DayVar int, HourVar int, MinuteVar int, SecondsVar float64) *DateTime {

	cret := xNewDateTimeUtc(YearVar, MonthVar, DayVar, HourVar, MinuteVar, SecondsVar)
	return cret
}

var xDateTimeAdd func(uintptr, TimeSpan) *DateTime

// Creates a copy of @datetime and adds the specified timespan to the copy.
func (x *DateTime) Add(TimespanVar TimeSpan) *DateTime {

	cret := xDateTimeAdd(x.GoPointer(), TimespanVar)
	return cret
}

var xDateTimeAddDays func(uintptr, int) *DateTime

// Creates a copy of @datetime and adds the specified number of days to the
// copy. Add negative values to subtract days.
func (x *DateTime) AddDays(DaysVar int) *DateTime {

	cret := xDateTimeAddDays(x.GoPointer(), DaysVar)
	return cret
}

var xDateTimeAddFull func(uintptr, int, int, int, int, int, float64) *DateTime

// Creates a new #GDateTime adding the specified values to the current date and
// time in @datetime. Add negative values to subtract.
func (x *DateTime) AddFull(YearsVar int, MonthsVar int, DaysVar int, HoursVar int, MinutesVar int, SecondsVar float64) *DateTime {

	cret := xDateTimeAddFull(x.GoPointer(), YearsVar, MonthsVar, DaysVar, HoursVar, MinutesVar, SecondsVar)
	return cret
}

var xDateTimeAddHours func(uintptr, int) *DateTime

// Creates a copy of @datetime and adds the specified number of hours.
// Add negative values to subtract hours.
func (x *DateTime) AddHours(HoursVar int) *DateTime {

	cret := xDateTimeAddHours(x.GoPointer(), HoursVar)
	return cret
}

var xDateTimeAddMinutes func(uintptr, int) *DateTime

// Creates a copy of @datetime adding the specified number of minutes.
// Add negative values to subtract minutes.
func (x *DateTime) AddMinutes(MinutesVar int) *DateTime {

	cret := xDateTimeAddMinutes(x.GoPointer(), MinutesVar)
	return cret
}

var xDateTimeAddMonths func(uintptr, int) *DateTime

// Creates a copy of @datetime and adds the specified number of months to the
// copy. Add negative values to subtract months.
//
// The day of the month of the resulting #GDateTime is clamped to the number
// of days in the updated calendar month. For example, if adding 1 month to
// 31st January 2018, the result would be 28th February 2018. In 2020 (a leap
// year), the result would be 29th February.
func (x *DateTime) AddMonths(MonthsVar int) *DateTime {

	cret := xDateTimeAddMonths(x.GoPointer(), MonthsVar)
	return cret
}

var xDateTimeAddSeconds func(uintptr, float64) *DateTime

// Creates a copy of @datetime and adds the specified number of seconds.
// Add negative values to subtract seconds.
func (x *DateTime) AddSeconds(SecondsVar float64) *DateTime {

	cret := xDateTimeAddSeconds(x.GoPointer(), SecondsVar)
	return cret
}

var xDateTimeAddWeeks func(uintptr, int) *DateTime

// Creates a copy of @datetime and adds the specified number of weeks to the
// copy. Add negative values to subtract weeks.
func (x *DateTime) AddWeeks(WeeksVar int) *DateTime {

	cret := xDateTimeAddWeeks(x.GoPointer(), WeeksVar)
	return cret
}

var xDateTimeAddYears func(uintptr, int) *DateTime

// Creates a copy of @datetime and adds the specified number of years to the
// copy. Add negative values to subtract years.
//
// As with g_date_time_add_months(), if the resulting date would be 29th
// February on a non-leap year, the day will be clamped to 28th February.
func (x *DateTime) AddYears(YearsVar int) *DateTime {

	cret := xDateTimeAddYears(x.GoPointer(), YearsVar)
	return cret
}

var xDateTimeCompare func(uintptr, uintptr) int

// A comparison function for #GDateTimes that is suitable
// as a #GCompareFunc. Both #GDateTimes must be non-%NULL.
func (x *DateTime) Compare(Dt2Var uintptr) int {

	cret := xDateTimeCompare(x.GoPointer(), Dt2Var)
	return cret
}

var xDateTimeDifference func(uintptr, *DateTime) TimeSpan

// Calculates the difference in time between @end and @begin.  The
// #GTimeSpan that is returned is effectively @end - @begin (ie:
// positive if the first parameter is larger).
func (x *DateTime) Difference(BeginVar *DateTime) TimeSpan {

	cret := xDateTimeDifference(x.GoPointer(), BeginVar)
	return cret
}

var xDateTimeEqual func(uintptr, uintptr) bool

// Checks to see if @dt1 and @dt2 are equal.
//
// Equal here means that they represent the same moment after converting
// them to the same time zone.
func (x *DateTime) Equal(Dt2Var uintptr) bool {

	cret := xDateTimeEqual(x.GoPointer(), Dt2Var)
	return cret
}

var xDateTimeFormat func(uintptr, string) string

// Creates a newly allocated string representing the requested @format.
//
// The format strings understood by this function are a subset of the
// strftime() format language as specified by C99.  The \%D, \%U and \%W
// conversions are not supported, nor is the 'E' modifier.  The GNU
// extensions \%k, \%l, \%s and \%P are supported, however, as are the
// '0', '_' and '-' modifiers. The Python extension \%f is also supported.
//
// In contrast to strftime(), this function always produces a UTF-8
// string, regardless of the current locale.  Note that the rendering of
// many formats is locale-dependent and may not match the strftime()
// output exactly.
//
// The following format specifiers are supported:
//
//   - \%a: the abbreviated weekday name according to the current locale
//   - \%A: the full weekday name according to the current locale
//   - \%b: the abbreviated month name according to the current locale
//   - \%B: the full month name according to the current locale
//   - \%c: the preferred date and time representation for the current locale
//   - \%C: the century number (year/100) as a 2-digit integer (00-99)
//   - \%d: the day of the month as a decimal number (range 01 to 31)
//   - \%e: the day of the month as a decimal number (range  1 to 31)
//   - \%F: equivalent to `%Y-%m-%d` (the ISO 8601 date format)
//   - \%g: the last two digits of the ISO 8601 week-based year as a
//     decimal number (00-99). This works well with \%V and \%u.
//   - \%G: the ISO 8601 week-based year as a decimal number. This works
//     well with \%V and \%u.
//   - \%h: equivalent to \%b
//   - \%H: the hour as a decimal number using a 24-hour clock (range 00 to 23)
//   - \%I: the hour as a decimal number using a 12-hour clock (range 01 to 12)
//   - \%j: the day of the year as a decimal number (range 001 to 366)
//   - \%k: the hour (24-hour clock) as a decimal number (range 0 to 23);
//     single digits are preceded by a blank
//   - \%l: the hour (12-hour clock) as a decimal number (range 1 to 12);
//     single digits are preceded by a blank
//   - \%m: the month as a decimal number (range 01 to 12)
//   - \%M: the minute as a decimal number (range 00 to 59)
//   - \%f: the microsecond as a decimal number (range 000000 to 999999)
//   - \%p: either "AM" or "PM" according to the given time value, or the
//     corresponding  strings for the current locale.  Noon is treated as
//     "PM" and midnight as "AM". Use of this format specifier is discouraged, as
//     many locales have no concept of AM/PM formatting. Use \%c or \%X instead.
//   - \%P: like \%p but lowercase: "am" or "pm" or a corresponding string for
//     the current locale. Use of this format specifier is discouraged, as
//     many locales have no concept of AM/PM formatting. Use \%c or \%X instead.
//   - \%r: the time in a.m. or p.m. notation. Use of this format specifier is
//     discouraged, as many locales have no concept of AM/PM formatting. Use \%c
//     or \%X instead.
//   - \%R: the time in 24-hour notation (\%H:\%M)
//   - \%s: the number of seconds since the Epoch, that is, since 1970-01-01
//     00:00:00 UTC
//   - \%S: the second as a decimal number (range 00 to 60)
//   - \%t: a tab character
//   - \%T: the time in 24-hour notation with seconds (\%H:\%M:\%S)
//   - \%u: the ISO 8601 standard day of the week as a decimal, range 1 to 7,
//     Monday being 1. This works well with \%G and \%V.
//   - \%V: the ISO 8601 standard week number of the current year as a decimal
//     number, range 01 to 53, where week 1 is the first week that has at
//     least 4 days in the new year. See g_date_time_get_week_of_year().
//     This works well with \%G and \%u.
//   - \%w: the day of the week as a decimal, range 0 to 6, Sunday being 0.
//     This is not the ISO 8601 standard format -- use \%u instead.
//   - \%x: the preferred date representation for the current locale without
//     the time
//   - \%X: the preferred time representation for the current locale without
//     the date
//   - \%y: the year as a decimal number without the century
//   - \%Y: the year as a decimal number including the century
//   - \%z: the time zone as an offset from UTC (+hhmm)
//   - \%:z: the time zone as an offset from UTC (+hh:mm).
//     This is a gnulib strftime() extension. Since: 2.38
//   - \%::z: the time zone as an offset from UTC (+hh:mm:ss). This is a
//     gnulib strftime() extension. Since: 2.38
//   - \%:::z: the time zone as an offset from UTC, with : to necessary
//     precision (e.g., -04, +05:30). This is a gnulib strftime() extension. Since: 2.38
//   - \%Z: the time zone or name or abbreviation
//   - \%\%: a literal \% character
//
// Some conversion specifications can be modified by preceding the
// conversion specifier by one or more modifier characters. The
// following modifiers are supported for many of the numeric
// conversions:
//
//   - O: Use alternative numeric symbols, if the current locale supports those.
//   - _: Pad a numeric result with spaces. This overrides the default padding
//     for the specifier.
//   - -: Do not pad a numeric result. This overrides the default padding
//     for the specifier.
//   - 0: Pad a numeric result with zeros. This overrides the default padding
//     for the specifier.
//
// Additionally, when O is used with B, b, or h, it produces the alternative
// form of a month name. The alternative form should be used when the month
// name is used without a day number (e.g., standalone). It is required in
// some languages (Baltic, Slavic, Greek, and more) due to their grammatical
// rules. For other languages there is no difference. \%OB is a GNU and BSD
// strftime() extension expected to be added to the future POSIX specification,
// \%Ob and \%Oh are GNU strftime() extensions. Since: 2.56
func (x *DateTime) Format(FormatVar string) string {

	cret := xDateTimeFormat(x.GoPointer(), FormatVar)
	return cret
}

var xDateTimeFormatIso8601 func(uintptr) string

// Format @datetime in [ISO 8601 format](https://en.wikipedia.org/wiki/ISO_8601),
// including the date, time and time zone, and return that as a UTF-8 encoded
// string.
//
// Since GLib 2.66, this will output to sub-second precision if needed.
func (x *DateTime) FormatIso8601() string {

	cret := xDateTimeFormatIso8601(x.GoPointer())
	return cret
}

var xDateTimeGetDayOfMonth func(uintptr) int

// Retrieves the day of the month represented by @datetime in the gregorian
// calendar.
func (x *DateTime) GetDayOfMonth() int {

	cret := xDateTimeGetDayOfMonth(x.GoPointer())
	return cret
}

var xDateTimeGetDayOfWeek func(uintptr) int

// Retrieves the ISO 8601 day of the week on which @datetime falls (1 is
// Monday, 2 is Tuesday... 7 is Sunday).
func (x *DateTime) GetDayOfWeek() int {

	cret := xDateTimeGetDayOfWeek(x.GoPointer())
	return cret
}

var xDateTimeGetDayOfYear func(uintptr) int

// Retrieves the day of the year represented by @datetime in the Gregorian
// calendar.
func (x *DateTime) GetDayOfYear() int {

	cret := xDateTimeGetDayOfYear(x.GoPointer())
	return cret
}

var xDateTimeGetHour func(uintptr) int

// Retrieves the hour of the day represented by @datetime
func (x *DateTime) GetHour() int {

	cret := xDateTimeGetHour(x.GoPointer())
	return cret
}

var xDateTimeGetMicrosecond func(uintptr) int

// Retrieves the microsecond of the date represented by @datetime
func (x *DateTime) GetMicrosecond() int {

	cret := xDateTimeGetMicrosecond(x.GoPointer())
	return cret
}

var xDateTimeGetMinute func(uintptr) int

// Retrieves the minute of the hour represented by @datetime
func (x *DateTime) GetMinute() int {

	cret := xDateTimeGetMinute(x.GoPointer())
	return cret
}

var xDateTimeGetMonth func(uintptr) int

// Retrieves the month of the year represented by @datetime in the Gregorian
// calendar.
func (x *DateTime) GetMonth() int {

	cret := xDateTimeGetMonth(x.GoPointer())
	return cret
}

var xDateTimeGetSecond func(uintptr) int

// Retrieves the second of the minute represented by @datetime
func (x *DateTime) GetSecond() int {

	cret := xDateTimeGetSecond(x.GoPointer())
	return cret
}

var xDateTimeGetSeconds func(uintptr) float64

// Retrieves the number of seconds since the start of the last minute,
// including the fractional part.
func (x *DateTime) GetSeconds() float64 {

	cret := xDateTimeGetSeconds(x.GoPointer())
	return cret
}

var xDateTimeGetTimezone func(uintptr) *TimeZone

// Get the time zone for this @datetime.
func (x *DateTime) GetTimezone() *TimeZone {

	cret := xDateTimeGetTimezone(x.GoPointer())
	return cret
}

var xDateTimeGetTimezoneAbbreviation func(uintptr) string

// Determines the time zone abbreviation to be used at the time and in
// the time zone of @datetime.
//
// For example, in Toronto this is currently "EST" during the winter
// months and "EDT" during the summer months when daylight savings
// time is in effect.
func (x *DateTime) GetTimezoneAbbreviation() string {

	cret := xDateTimeGetTimezoneAbbreviation(x.GoPointer())
	return cret
}

var xDateTimeGetUtcOffset func(uintptr) TimeSpan

// Determines the offset to UTC in effect at the time and in the time
// zone of @datetime.
//
// The offset is the number of microseconds that you add to UTC time to
// arrive at local time for the time zone (ie: negative numbers for time
// zones west of GMT, positive numbers for east).
//
// If @datetime represents UTC time, then the offset is always zero.
func (x *DateTime) GetUtcOffset() TimeSpan {

	cret := xDateTimeGetUtcOffset(x.GoPointer())
	return cret
}

var xDateTimeGetWeekNumberingYear func(uintptr) int

// Returns the ISO 8601 week-numbering year in which the week containing
// @datetime falls.
//
// This function, taken together with g_date_time_get_week_of_year() and
// g_date_time_get_day_of_week() can be used to determine the full ISO
// week date on which @datetime falls.
//
// This is usually equal to the normal Gregorian year (as returned by
// g_date_time_get_year()), except as detailed below:
//
// For Thursday, the week-numbering year is always equal to the usual
// calendar year.  For other days, the number is such that every day
// within a complete week (Monday to Sunday) is contained within the
// same week-numbering year.
//
// For Monday, Tuesday and Wednesday occurring near the end of the year,
// this may mean that the week-numbering year is one greater than the
// calendar year (so that these days have the same week-numbering year
// as the Thursday occurring early in the next year).
//
// For Friday, Saturday and Sunday occurring near the start of the year,
// this may mean that the week-numbering year is one less than the
// calendar year (so that these days have the same week-numbering year
// as the Thursday occurring late in the previous year).
//
// An equivalent description is that the week-numbering year is equal to
// the calendar year containing the majority of the days in the current
// week (Monday to Sunday).
//
// Note that January 1 0001 in the proleptic Gregorian calendar is a
// Monday, so this function never returns 0.
func (x *DateTime) GetWeekNumberingYear() int {

	cret := xDateTimeGetWeekNumberingYear(x.GoPointer())
	return cret
}

var xDateTimeGetWeekOfYear func(uintptr) int

// Returns the ISO 8601 week number for the week containing @datetime.
// The ISO 8601 week number is the same for every day of the week (from
// Moday through Sunday).  That can produce some unusual results
// (described below).
//
// The first week of the year is week 1.  This is the week that contains
// the first Thursday of the year.  Equivalently, this is the first week
// that has more than 4 of its days falling within the calendar year.
//
// The value 0 is never returned by this function.  Days contained
// within a year but occurring before the first ISO 8601 week of that
// year are considered as being contained in the last week of the
// previous year.  Similarly, the final days of a calendar year may be
// considered as being part of the first ISO 8601 week of the next year
// if 4 or more days of that week are contained within the new year.
func (x *DateTime) GetWeekOfYear() int {

	cret := xDateTimeGetWeekOfYear(x.GoPointer())
	return cret
}

var xDateTimeGetYear func(uintptr) int

// Retrieves the year represented by @datetime in the Gregorian calendar.
func (x *DateTime) GetYear() int {

	cret := xDateTimeGetYear(x.GoPointer())
	return cret
}

var xDateTimeGetYmd func(uintptr, int, int, int)

// Retrieves the Gregorian day, month, and year of a given #GDateTime.
func (x *DateTime) GetYmd(YearVar int, MonthVar int, DayVar int) {

	xDateTimeGetYmd(x.GoPointer(), YearVar, MonthVar, DayVar)

}

var xDateTimeHash func(uintptr) uint

// Hashes @datetime into a #guint, suitable for use within #GHashTable.
func (x *DateTime) Hash() uint {

	cret := xDateTimeHash(x.GoPointer())
	return cret
}

var xDateTimeIsDaylightSavings func(uintptr) bool

// Determines if daylight savings time is in effect at the time and in
// the time zone of @datetime.
func (x *DateTime) IsDaylightSavings() bool {

	cret := xDateTimeIsDaylightSavings(x.GoPointer())
	return cret
}

var xDateTimeRef func(uintptr) *DateTime

// Atomically increments the reference count of @datetime by one.
func (x *DateTime) Ref() *DateTime {

	cret := xDateTimeRef(x.GoPointer())
	return cret
}

var xDateTimeToLocal func(uintptr) *DateTime

// Creates a new #GDateTime corresponding to the same instant in time as
// @datetime, but in the local time zone.
//
// This call is equivalent to calling g_date_time_to_timezone() with the
// time zone returned by g_time_zone_new_local().
func (x *DateTime) ToLocal() *DateTime {

	cret := xDateTimeToLocal(x.GoPointer())
	return cret
}

var xDateTimeToTimeval func(uintptr, *TimeVal) bool

// Stores the instant in time that @datetime represents into @tv.
//
// The time contained in a #GTimeVal is always stored in the form of
// seconds elapsed since 1970-01-01 00:00:00 UTC, regardless of the time
// zone associated with @datetime.
//
// On systems where 'long' is 32bit (ie: all 32bit systems and all
// Windows systems), a #GTimeVal is incapable of storing the entire
// range of values that #GDateTime is capable of expressing.  On those
// systems, this function returns %FALSE to indicate that the time is
// out of range.
//
// On systems where 'long' is 64bit, this function never fails.
func (x *DateTime) ToTimeval(TvVar *TimeVal) bool {

	cret := xDateTimeToTimeval(x.GoPointer(), TvVar)
	return cret
}

var xDateTimeToTimezone func(uintptr, *TimeZone) *DateTime

// Create a new #GDateTime corresponding to the same instant in time as
// @datetime, but in the time zone @tz.
//
// This call can fail in the case that the time goes out of bounds.  For
// example, converting 0001-01-01 00:00:00 UTC to a time zone west of
// Greenwich will fail (due to the year 0 being out of range).
func (x *DateTime) ToTimezone(TzVar *TimeZone) *DateTime {

	cret := xDateTimeToTimezone(x.GoPointer(), TzVar)
	return cret
}

var xDateTimeToUnix func(uintptr) int64

// Gives the Unix time corresponding to @datetime, rounding down to the
// nearest second.
//
// Unix time is the number of seconds that have elapsed since 1970-01-01
// 00:00:00 UTC, regardless of the time zone associated with @datetime.
func (x *DateTime) ToUnix() int64 {

	cret := xDateTimeToUnix(x.GoPointer())
	return cret
}

var xDateTimeToUtc func(uintptr) *DateTime

// Creates a new #GDateTime corresponding to the same instant in time as
// @datetime, but in UTC.
//
// This call is equivalent to calling g_date_time_to_timezone() with the
// time zone returned by g_time_zone_new_utc().
func (x *DateTime) ToUtc() *DateTime {

	cret := xDateTimeToUtc(x.GoPointer())
	return cret
}

var xDateTimeUnref func(uintptr)

// Atomically decrements the reference count of @datetime by one.
//
// When the reference count reaches zero, the resources allocated by
// @datetime are freed
func (x *DateTime) Unref() {

	xDateTimeUnref(x.GoPointer())

}

// A value representing an interval of time, in microseconds.
type TimeSpan = int64

const (
	// Evaluates to a time span of one day.
	TIME_SPAN_DAY int64 = 86400000000
	// Evaluates to a time span of one hour.
	TIME_SPAN_HOUR int64 = 3600000000
	// Evaluates to a time span of one millisecond.
	TIME_SPAN_MILLISECOND int64 = 1000
	// Evaluates to a time span of one minute.
	TIME_SPAN_MINUTE int64 = 60000000
	// Evaluates to a time span of one second.
	TIME_SPAN_SECOND int64 = 1000000
)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDateTime, lib, "g_date_time_new")
	core.PuregoSafeRegister(&xNewDateTimeFromIso8601, lib, "g_date_time_new_from_iso8601")
	core.PuregoSafeRegister(&xNewDateTimeFromTimevalLocal, lib, "g_date_time_new_from_timeval_local")
	core.PuregoSafeRegister(&xNewDateTimeFromTimevalUtc, lib, "g_date_time_new_from_timeval_utc")
	core.PuregoSafeRegister(&xNewDateTimeFromUnixLocal, lib, "g_date_time_new_from_unix_local")
	core.PuregoSafeRegister(&xNewDateTimeFromUnixUtc, lib, "g_date_time_new_from_unix_utc")
	core.PuregoSafeRegister(&xNewDateTimeLocal, lib, "g_date_time_new_local")
	core.PuregoSafeRegister(&xNewDateTimeNow, lib, "g_date_time_new_now")
	core.PuregoSafeRegister(&xNewDateTimeNowLocal, lib, "g_date_time_new_now_local")
	core.PuregoSafeRegister(&xNewDateTimeNowUtc, lib, "g_date_time_new_now_utc")
	core.PuregoSafeRegister(&xNewDateTimeUtc, lib, "g_date_time_new_utc")

	core.PuregoSafeRegister(&xDateTimeAdd, lib, "g_date_time_add")
	core.PuregoSafeRegister(&xDateTimeAddDays, lib, "g_date_time_add_days")
	core.PuregoSafeRegister(&xDateTimeAddFull, lib, "g_date_time_add_full")
	core.PuregoSafeRegister(&xDateTimeAddHours, lib, "g_date_time_add_hours")
	core.PuregoSafeRegister(&xDateTimeAddMinutes, lib, "g_date_time_add_minutes")
	core.PuregoSafeRegister(&xDateTimeAddMonths, lib, "g_date_time_add_months")
	core.PuregoSafeRegister(&xDateTimeAddSeconds, lib, "g_date_time_add_seconds")
	core.PuregoSafeRegister(&xDateTimeAddWeeks, lib, "g_date_time_add_weeks")
	core.PuregoSafeRegister(&xDateTimeAddYears, lib, "g_date_time_add_years")
	core.PuregoSafeRegister(&xDateTimeCompare, lib, "g_date_time_compare")
	core.PuregoSafeRegister(&xDateTimeDifference, lib, "g_date_time_difference")
	core.PuregoSafeRegister(&xDateTimeEqual, lib, "g_date_time_equal")
	core.PuregoSafeRegister(&xDateTimeFormat, lib, "g_date_time_format")
	core.PuregoSafeRegister(&xDateTimeFormatIso8601, lib, "g_date_time_format_iso8601")
	core.PuregoSafeRegister(&xDateTimeGetDayOfMonth, lib, "g_date_time_get_day_of_month")
	core.PuregoSafeRegister(&xDateTimeGetDayOfWeek, lib, "g_date_time_get_day_of_week")
	core.PuregoSafeRegister(&xDateTimeGetDayOfYear, lib, "g_date_time_get_day_of_year")
	core.PuregoSafeRegister(&xDateTimeGetHour, lib, "g_date_time_get_hour")
	core.PuregoSafeRegister(&xDateTimeGetMicrosecond, lib, "g_date_time_get_microsecond")
	core.PuregoSafeRegister(&xDateTimeGetMinute, lib, "g_date_time_get_minute")
	core.PuregoSafeRegister(&xDateTimeGetMonth, lib, "g_date_time_get_month")
	core.PuregoSafeRegister(&xDateTimeGetSecond, lib, "g_date_time_get_second")
	core.PuregoSafeRegister(&xDateTimeGetSeconds, lib, "g_date_time_get_seconds")
	core.PuregoSafeRegister(&xDateTimeGetTimezone, lib, "g_date_time_get_timezone")
	core.PuregoSafeRegister(&xDateTimeGetTimezoneAbbreviation, lib, "g_date_time_get_timezone_abbreviation")
	core.PuregoSafeRegister(&xDateTimeGetUtcOffset, lib, "g_date_time_get_utc_offset")
	core.PuregoSafeRegister(&xDateTimeGetWeekNumberingYear, lib, "g_date_time_get_week_numbering_year")
	core.PuregoSafeRegister(&xDateTimeGetWeekOfYear, lib, "g_date_time_get_week_of_year")
	core.PuregoSafeRegister(&xDateTimeGetYear, lib, "g_date_time_get_year")
	core.PuregoSafeRegister(&xDateTimeGetYmd, lib, "g_date_time_get_ymd")
	core.PuregoSafeRegister(&xDateTimeHash, lib, "g_date_time_hash")
	core.PuregoSafeRegister(&xDateTimeIsDaylightSavings, lib, "g_date_time_is_daylight_savings")
	core.PuregoSafeRegister(&xDateTimeRef, lib, "g_date_time_ref")
	core.PuregoSafeRegister(&xDateTimeToLocal, lib, "g_date_time_to_local")
	core.PuregoSafeRegister(&xDateTimeToTimeval, lib, "g_date_time_to_timeval")
	core.PuregoSafeRegister(&xDateTimeToTimezone, lib, "g_date_time_to_timezone")
	core.PuregoSafeRegister(&xDateTimeToUnix, lib, "g_date_time_to_unix")
	core.PuregoSafeRegister(&xDateTimeToUtc, lib, "g_date_time_to_utc")
	core.PuregoSafeRegister(&xDateTimeUnref, lib, "g_date_time_unref")

}
