package day

import (
	"day/template"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

type day struct {
	time    time.Time
	Year    int
	Month   time.Month
	Day     int
	Hour    int
	Minute  int
	Second  int
	Unix    int64
	WeekDay time.Weekday
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

func (d *day) fields() {
	time := d.time

	d.Year = time.Year()
	d.Month = time.Month()
	d.Day = time.Day()
	d.Hour = time.Hour()
	d.Minute = time.Minute()
	d.Second = time.Second()
	d.Unix = time.UnixNano() / 1e6
	d.WeekDay = time.Weekday()
}

func createDay(time time.Time) *day {
	d := &day{
		time: time,
	}

	d.fields()
	return d
}

func (d *day) change(value int, unit Unit) *day {
	sec := int(1e9)
	switch unit {
	case Year:
		return createDay(d.time.AddDate(value, 0, 0))
	case Month:
		return createDay(d.time.AddDate(0, value, 0))
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

	return createDay(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)), nil

}

func Unix(unix int) (*day, error) {
	valid, _ := regexp.MatchString(`\d{13}`, fmt.Sprintf("%v", unix))

	if !valid {
		return nil, errors.New("unix is 13 bit milliseconds")
	}

	sec := int64(unix / 1e3)

	nsecStr := fmt.Sprintf("%v", unix)
	nsec, _ := strconv.Atoi(nsecStr[len(nsecStr)-3:])

	return createDay(time.Unix(sec, int64(nsec))), nil
}

func List(list []int) *day {
	year, month, day, hour, minute, second := parseList(list)
	return createDay(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC))
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

func (d *day) Format(t string) string {
	c := regexp.MustCompile(`/\[([^\]]+)]|Y{1,4}|M{1,4}|D{1,2}|d{1,4}|H{1,2}|h{1,2}|a|A|m{1,2}|s{1,2}|Z{1,2}|SSS/g`)

	ret := c.ReplaceAllStringFunc(t, func(substr string) string {
		switch substr {
		case template.Year:
			return fmt.Sprintf("%v", d.Year)
		case template.Month:
			return fmt.Sprintf("%v", int(d.Month))
		case template.Day:
			return fmt.Sprintf("%v", d.Day)
		case template.Hour12:
			return fmt.Sprintf("%v", d.Hour%12)
		case template.Hour24:
			return fmt.Sprintf("%v", d.Hour)
		case template.Minute:
			return fmt.Sprintf("%v", d.Minute)
		case template.Second:
			return fmt.Sprintf("%v", d.Second)
		}

		return ""
	})
	log.Println("format result:", ret)
	return ret
}
