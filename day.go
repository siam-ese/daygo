package day

import (
	"day/locale"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type day struct {
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

func (d *day) fields() {
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

func createDay(time time.Time) *day {
	d := &day{
		time: time,
	}

	d.fields()
	return d
}

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

func (d *day) Time() time.Time {
	return d.time
}

// change day date-time with value or unit
// value might int or -int
func (d *day) change(value int, unit Unit) *day {
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

func New(time time.Time) *day {
	return createDay(time)
}

func Now() *day {
	return createDay(time.Now())
}

func Format(t string) (*day, error) {
	ret := parseT(t)
	if ret == nil {
		return nil, errors.New("format parse failed")
	}

	var list []int
	for _, r := range ret[1:] { // parse the strings to ints
		val, _ := strconv.Atoi(r)
		list = append(list, val)
	}

	year, month, day, hour, minute, second := parseList(list)

	return createDay(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)), nil

}

func Unix(unix int) (*day, error) {

	unixStr := fmt.Sprintf("%v", unix)
	if len(unixStr) != 13 {
		return nil, errors.New("unix is 13 bit milliseconds")
	}

	sec := int64(unix / 1e3)

	nsecStr := fmt.Sprintf("%v", unix)
	nsec, _ := strconv.Atoi(nsecStr[len(nsecStr)-3:])

	return createDay(time.Unix(sec, int64(nsec))), nil
}

func List(list []int) *day {
	year, month, day, hour, minute, second := parseList(list)
	return createDay(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local))
}

func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func (d *day) Set(value int, unit Unit) *day {
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

func (d *day) Add(value int, unit Unit) *day {
	return d.change(value, unit)
}

func (d *day) Subtract(value int, unit Unit) *day {
	return d.change(-value, unit)
}

func (d *day) SetYear(value int) *day {
	return d.Set(value, Year)
}
func (d *day) SetMonth(value int) *day {
	return d.Set(value, Month)
}
func (d *day) SetDay(value int) *day {
	return d.Set(value, Day)
}
func (d *day) SetMinute(value int) *day {
	return d.Set(value, Minute)
}
func (d *day) SetHour(value int) *day {
	return d.Set(value, Hour)
}
func (d *day) SetSecond(value int) *day {
	return d.Set(value, Second)
}
func (d *day) SetWeekDay(value int) *day {
	return d.Set(value, WeekDay)
}

func (d *day) SecondAfterUnixNano() int {
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

func (d *day) Format(t string) string {
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

func (d *day) UTC() *day {
	return createDay(d.time.UTC())
}

func (d *day) Local() *day {
	return createDay(d.time.Local())
}

func (d *day) StartOf(unit Unit) *day {
	var year = d.Year
	var (
		month,
		day,
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
		day = d.Hour
	}
	if unit >= Minute {
		day = d.Minute
	}
	if unit >= Second {
		day = d.Second
	}

	return createDay(time.Date(year, time.Month(month), day, hour, minute, second, 0, d.time.Location()))
}

func intInSlice(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func (d *day) EndOf(unit Unit) *day {
	month := 12
	day := 31
	hour := 23
	minute := 59
	second := 59

	if unit >= Month {
		month = int(d.Month)
	}
	if unit >= Day {
		day = MonthDay(d.Year, month)
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

	return createDay(time.Date(d.Year, time.Month(month), day, hour, minute, second, int(9e8), d.time.Location()))
}
