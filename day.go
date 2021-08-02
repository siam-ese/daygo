package day

import (
	"day/locale"
	"errors"
	"fmt"
	"strconv"
	"time"
)

/**
day.D package mainly type
contains mostly method
*/
type D struct {
	time     time.Time
	Year     int
	Month    time.Month
	Day      int
	Hour     int
	Minute   int
	Second   int
	Unix     int64
	UnixNano int64
	WeekDay  time.Weekday
}

type Unit = int

// day bulit-in unit-type, use in Add, Subtract, Set, EndOf, StartOf methods
const (
	Year Unit = iota
	Month
	Day
	Hour
	Minute
	Second
	WeekDay
)

var translator locale.Translator

func init() {
	translator = locale.EN
}

func (d *D) fields() {
	time := d.time

	d.Year = time.Year()
	d.Month = time.Month()
	d.Day = time.Day()
	d.Hour = time.Hour()
	d.Minute = time.Minute()
	d.Second = time.Second()
	d.WeekDay = time.Weekday()

	unixNano := time.UnixNano()
	d.Unix = unixNano / 1e6
	d.UnixNano = unixNano
}

func createDay(time time.Time) *D {
	d := &D{
		time: time,
	}

	d.fields()
	return d
}

// get maximum days by year and month
func MonthDay(year int, month int) int {
	leapMonth := []int{1, 3, 5, 7, 8, 10, 12}
	if month == 2 {
		if IsLeapYear(year) {
			return 29
		} else {
			return 28
		}
	} else if intInSlice(leapMonth, month) {
		return 31
	} else {
		return 30
	}
}

// set translator of day, day/locale has zh-cn or en, default: en
func Locale(t locale.Translator) {
	translator = t
}

func (d *D) Time() time.Time {
	return d.time
}

// change Day date-time with value or unit
// value might int or -int
func (d *D) change(value int, unit Unit) *D {
	sec := int(time.Second)

	switch unit {
	case Year:
		return createDay(d.time.AddDate(value, 0, 0))
	case Month:
		month := int(d.Month) + value
		return createDay(time.Date(d.Year, time.Month(month), MonthDay(d.Year, month), d.Hour, d.Minute, d.Minute, d.SecondAfterUnixNano(), d.time.Location()))
	case Day:
		return createDay(d.time.AddDate(0, 0, value))
	case Hour:
		return createDay(d.time.Add(time.Duration(60 * 60 * value * sec)))
	case Minute:
		return createDay(d.time.Add(time.Duration(60 * value * sec)))
	case Second:
		return createDay(d.time.Add(time.Duration(value * sec)))
	}

	return d
}

// use time.Time create a new day.D
func New(time time.Time) *D {
	return createDay(time)
}

// use time.Now create a new day.D
func Now() *D {
	return createDay(time.Now())
}

// parse ISO date-time-string, use parse result create a new day.D
func Format(t string) (*D, error) {
	ret := parseT(t)
	if ret == nil {
		return nil, fmt.Errorf("formt failed: Could not parse string %s", t)
	}

	var list []int
	for _, r := range ret[1:] { // parse the strings to ints
		val, _ := strconv.Atoi(r)
		list = append(list, val)
	}

	year, month, Day, hour, minute, second := parseList(list)

	return createDay(time.Date(year, time.Month(month), Day, hour, minute, second, 0, time.Local)), nil

}

// use millisecond create a new day.D
func Unix(unix int) (*D, error) {
	unixStr := fmt.Sprintf("%v", unix)
	if len(unixStr) != 13 {
		return nil, errors.New("unix is 13 bit milliseconds")
	}

	sec := int64(unix / 1e3)

	nsecStr := fmt.Sprintf("%v", unix)
	nsec, _ := strconv.Atoi(nsecStr[len(nsecStr)-3:])

	return createDay(time.Unix(sec, int64(nsec))), nil
}

// use int slice creat a new day.D, unset items use default value
// parameter: list int[year, month, day, hour, minute, second]
// example: List([]int{2021, 8, 17}) => day.D year: 2021, month: 8, day: 17, hour: 0, minute: 0, second: 0
func List(list []int) *D {
	year, month, Day, hour, minute, second := parseList(list)
	return createDay(time.Date(year, time.Month(month), Day, hour, minute, second, 0, time.Local))
}

// judge year is leap year
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

// set value by specific unit
// example: D.set(2020, day.Year) // set year to 2020
func (d *D) Set(value int, unit Unit) *D {
	switch unit {
	case Year:
		return d.change(value-d.Year, Year)
	case Month:
		return d.change(value-int(d.Month), Month)
	case Day:
		return d.change(value-d.Day, Day)
	case Hour:
		return d.change(value-d.Hour, Hour)
	case Minute:
		return d.change(value-d.Minute, Minute)
	case Second:
		return d.change(value-d.Second, Second)
	case WeekDay:
		return d.change(value-int(d.WeekDay), Day)
	}

	return d
}

// add value by specific unit
// example: List([]int{2020}).Add(2, day.Year) // 2020 year add to 2022 year
func (d *D) Add(value int, unit Unit) *D {
	return d.change(value, unit)
}

// add value by specific unit
// example: List([]int{2020}).Subtract(2, day.Year) // 2020 year sub to 2018 year
func (d *D) Subtract(value int, unit Unit) *D {
	return d.change(-value, unit)
}

func (d *D) SetYear(value int) *D {
	return d.Set(value, Year)
}
func (d *D) SetMonth(value int) *D {
	return d.Set(value, Month)
}
func (d *D) SetDay(value int) *D {
	return d.Set(value, Day)
}
func (d *D) SetMinute(value int) *D {
	return d.Set(value, Minute)
}
func (d *D) SetHour(value int) *D {
	return d.Set(value, Hour)
}
func (d *D) SetSecond(value int) *D {
	return d.Set(value, Second)
}
func (d *D) SetWeekDay(value int) *D {
	return d.Set(value, WeekDay)
}

// slice second after 9 bit unixnano
// example: 1627637214-376669500 => 376669500 return unixnano
func (d *D) SecondAfterUnixNano() int {
	str := fmt.Sprintf("%v", d.UnixNano)
	ret, _ := strconv.Atoi(str[len(str)-9:])

	return ret
}

func fillZero(value int) string {
	if value < 10 {
		return fmt.Sprintf("0%v", value)
	} else {
		return fmt.Sprintf("%v", value)
	}
}

func (d *D) Format(t string) string {
	ret := formatRe.ReplaceAllStringFunc(t, func(substr string) string {
		switch substr {
		case "YYYY":
			return fmt.Sprintf("%v", d.Year)
		case "YY":
			year := fmt.Sprintf("%v", d.Year)
			return year[len(year)-2:]
		case "M":
			return fmt.Sprintf("%v", int(d.Month))
		case "MM":
			return fillZero(int(d.Month))
		case "D":
			return fmt.Sprintf("%v", d.Day)
		case "DD":
			return fillZero(d.Day)
		case "h":
			return fmt.Sprintf("%v", d.Hour%12)
		case "hh":
			return fillZero(d.Hour % 12)
		case "H":
			return fmt.Sprintf("%v", d.Hour)
		case "HH":
			return fillZero(d.Hour)
		case "m":
			return fmt.Sprintf("%v", d.Minute)
		case "mm":
			return fillZero(d.Minute)
		case "s":
			return fmt.Sprintf("%v", d.Second)
		case "ss":
			return fillZero(d.Second)
		case "SSS":
			return fmt.Sprintf("%v", d.Unix)
		case "d":
			return fmt.Sprintf("%v", int(d.WeekDay))
		case "dd":
			return translator.WT(int(d.WeekDay))
		}

		return substr
	})

	return ret
}

// return current day.D from UTC time
func (d *D) UTC() *D {
	return createDay(d.time.UTC())
}

// return  current day.D from local time
func (d *D) Local() *D {
	return createDay(d.time.Local())
}

func intInSlice(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

// set date-time startOf special unit
func (d *D) StartOf(unit Unit) *D {
	year := d.Year
	month := 1
	day := 1

	var (
		hour,
		minute,
		second int
	)

	if unit >= Month {
		month = int(d.Month)
	}
	if unit >= Day {
		day = d.Day
	}
	if unit >= Hour {
		hour = d.Hour
	}
	if unit >= Minute {
		minute = d.Minute
	}
	if unit >= Second {
		second = d.Second
	}

	return createDay(time.Date(year, time.Month(month), day, hour, minute, second, 0, d.time.Location()))
}

// set date-time endOf special unit
func (d *D) EndOf(unit Unit) *D {
	month := 12
	day := MonthDay(d.Year, month)
	hour := 23
	minute := 59
	second := 59

	if unit >= Month {
		month = int(d.Month)
	}
	if unit >= Day {
		day = d.Day
	}
	if unit >= Hour {
		hour = d.Hour
	}
	if unit >= Minute {
		minute = d.Minute
	}
	if unit >= Second {
		second = d.Second
	}

	return createDay(time.Date(d.Year, time.Month(month), day, hour, minute, second, 999999999, d.time.Location()))
}

// returun days for month
func (d *D) DaysInMonth() int {
	return MonthDay(d.Year, int(d.Month))
}

// a.From(b) return a time.Duration, a.time sub b.time
func (d *D) From(d2 *D) time.Duration {
	return d.time.Sub(d2.time)
}
