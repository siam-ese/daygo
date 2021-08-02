# daygo

A handle date-time Go library inspire for [Dayjs](https://github.com/iamkun/dayjs/)

```go
daygo.Now().Add(1, daygo.Month).EndOf(daygo.Day).Format("YYYY-MM-DD HH:mm:ss")
```

- 😀Immutable
- 🔥Chainable

## Installation

```console
go get -u github.com/siam-ese/daygo@v0.1.3
```

## Api

Now()

A day.D by now time

```go
daygo.Now()
```

New(time.Time)

A day.D by time.Time

```go
daygo.New(time.Time)
```

Format(string)

A day.D by [ISO](https://zh.wikipedia.org/zh-tw/ISO_8601) format string

```go
day.Format(`2021-08-02T05:53:12`)
day.Format(`20210802T055312`)
```

List([]int)

A day.D by int list, default values are used for unset items

```go
// [year is required, month = 1, day = 1, hour = 0, minute = 0, second = 0]
day.List([2021, 8]) // 2021-8-1 00:00:00
day.List([2021, 8, 2, 19, 13, 54])
```

Set translator to daygo
Locale(locale.translator)

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

D.Time()

```go
D.Time() // return time.Time from D
```
