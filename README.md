# daygo

A handle date-time Go library inspire for [Dayjs](https://github.com/iamkun/dayjs/)

```go
daygo.Now().Add(1, daygo.Month).EndOf(daygo.Day).Format("YYYY-MM-DD HH:mm:ss")
```

- 😀Immutable
- 🔥Chainable

## Installation

```console
go get -u github.com/siam-ese/daygo
```

## Api

### Now()

Get a day.D by now time

```go
daygo.Now()
```

### New(time.Time)

Get a day.D by time.Time

```go
daygo.New(time.Time)
```

### Parse(string)

Get a day.D by [ISO](https://zh.wikipedia.org/zh-tw/ISO_8601) format string

```go
daygo.Parse(`2021-08-02T05:53:12`)
daygo.Parse(`20210802T055312`)
```

### List([]int)

Get a day.D by int list, default values are used for unset items

```go
// [year is required, month = 1, day = 1, hour = 0, minute = 0, second = 0]
daygo.List([2021, 8]) // 2021-8-1 00:00:00
daygo.List([2021, 8, 2, 19, 13, 54])
```

### Unix(millisecond int)

Get a day.D by millisecond

```go
daygo.Unix(1627542959340)
```

### Locale(locale.translator)

Set translator to daygo

```go
import daygo
import daygo/locale

daygo.Locale(locale.EN)
// create your translator
customTranslator = locale.Translator{
    WeekMap: int[7]
    MonthMap: int[12]
}
daygo.Locale(customTranslator)
```

## Api daygo.D

```go
package daygo

type D struct {
    Year     int
	Month    time.Month
	Day      int
	Hour     int
	Minute   int
	Second   int
	Unix     int64 // millisecond
	UnixNano int64
	WeekDay  time.Weekday
}
```

### D.Time()

```go
D.Time() // return time.Time from D
```

### D.Add(value int, unit daygo.Unit)

### D.Subtract(value int, unit daygo.Unit)

### D.Set(value int, unit daygo.Unit)

return a change time new daygo.D

```go
D.Add(1, daygo.Year) // add 1 year
D.Add(1, daygo.Month)
D.Add(1, daygo.Day)
D.Add(1, daygo.Hour)
D.Add(1, daygo.Minute)
D.Add(1, daygo.Second)
D.Add(1, daygo.WeekDay) // Monday -> Tuesday, if Sunday Add 1 to Monday

D.Subtract(1, daygo.Year)
D.Subtract(1, daygo.Month)
D.Subtract(1, daygo.Day)
D.Subtract(1, daygo.Hour)
D.Subtract(1, daygo.Minute)
D.Subtract(1, daygo.Second)
D.Subtract(1, daygo.WeekDay)

D.Set(2020, daygo.Year) // set year to 2020
D.Set(1, daygo.Month)
D.Set(1, daygo.Day)
D.Set(1, daygo.Hour)
D.Set(1, daygo.Minute)
D.Set(1, daygo.Second)
D.Set(1, daygo.WeekDay)
// set method alias
D.SetYear(2020) // set year to 2020
D.SetMonth(1)
D.SetDay(1)
D.SetMinute(1)
D.SetHour(1)
D.SetSecond(1)
D.SetWeekDay(1)
```

### D.Format(template string)

```go
// time 2021-8-3 10:02:55
D.Format("YYYY年MM月DD日，HH时mm分ss秒") // return 2021年08月03日，10时02分55秒

template
YYYY // 2021 Four-digit year
YY // 21 Two-digit year
M // The month, beginning at 1
MM // The month, 2-digits
MMMM // The month name by translator, here is August
D // The day of the month
DD // The day of the month, 2-digits
h // The hour, 12-hour clock
hh // The hour, 12-hour clock, 2-digits
H // The hour 24-hour
HH // The hour, 24-hour 2-digits
m // The minute
mm // The minute, 2-digits
s // The second
ss // The second, 2-digits
SSS // The millisecond, 3-digits 000-999
d // The number of Weekday 1-7
dd // The Weekday name by translator, here is Tuesday

```

### D.StartOf(unit daygo.Unit)

### D.EndOf(unit daygo.Unit)

```go
Support: Year, Month, Day, Hour, Minute, Second
// time 2021-8-3 10:02:55
D.StartOf(daygo.Year) // 2021-1-1 00:00:00 Nanosecond is 0e9
D.EndOf(daygo.Year) // 2021-12-31 23:59:59 Nanosecond is nine bits 9

// Of second, is only set millisecond or nanosecond
D.StartOf(daygo.Second)
D.EndOf(daygo.Second)
```

### D.DaysInMonth()

Get days by day.D current Month

```go
// time 2021-08
D.DaysInMonth() // 31
// time 2021-02
D.DaysInMonth() // 28

```

### From(b day.D) time.Duration

return a time.Duration by current day.D sub b-day.D

```go
a := daygo.Now()
b := daygo.Now().Subtract(1, daygo.Minute)
a.From(b) // a - b, return time.Duration
```

### D.UTC()

### D.Local()

Get a new day.D by UTC or Local

```
D.UTC()
D.Local()
```
