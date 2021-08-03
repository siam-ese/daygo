package daygo

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/siam-ese/daygo/locale"
)

func ErrorF(method string) func(...interface{}) string {
	t := `[%s] test failed`
	return func(vals ...interface{}) string {
		vals = append(vals, method)
		if len(vals) > 0 {
			t += ",params: ("
			for i := 0; i < len(vals)-1; i++ {
				t += "%s, "
			}
			t += ")"
		}

		return fmt.Sprintf(t, vals...)
	}
}

func TestNow(t *testing.T) {
	f := ErrorF("Now")

	d := Now()
	if d.Day != time.Now().Day() {
		t.Error(f())
	}
}

func TestFormat(t *testing.T) {
	f := ErrorF("Format")
	p1 := "2021-07-28 12:26:36"
	d, _ := Format(p1)

	if d.Year != 2021 {
		t.Error(f(p1))
	}
}

func TestUnix(t *testing.T) {
	param1 := 1627542959340
	d, _ := Unix(param1)
	f := ErrorF("Unix")
	if d.Year != 2021 && d.Minute != 15 {
		t.Error(f(param1))
	}

	param2 := 162754295934 // 14 bit
	_, err := Unix(param2)
	if err == nil {
		t.Error("Unix uncaught error", 162754295934)
	}
}

func TestDayAdd(t *testing.T) {
	p1 := "2021-01-31 19:02:59"
	d, _ := Format(p1)

	f := ErrorF("Day.Add")
	if d.Add(1, Year).Year != 2022 {
		t.Error(f(p1), "Add 1 Year")
	}
	addMonth := d.Add(1, Month)
	if addMonth.Month != 2 && addMonth.Day != 28 { // because maximum days of February is 28
		t.Error(f(p1), "Add 1 Month")
	}
	if d.Add(1, Day).Day != 1 { // 1.31 + 1 = 2.1
		t.Error(f(p1), "Add 1 Day")
	}
	if d.Add(1, Hour).Hour != 20 {
		t.Error(f(p1), "Add 1 Day")
	}
	if d.Add(1, Minute).Minute != 3 {
		t.Error(f(p1), "Add 1 Minute")
	}
	if d.Add(1, Second).Second != 0 { // 59 + 1 = 1m0s
		t.Error(f(p1), "Add 1 Second")
	}
}

func TestDaySubtract(t *testing.T) {
	p1 := "2021-01-31 19:02:00"
	d, _ := Format(p1)

	f := ErrorF("Day.Subtract")
	if d.Subtract(1, Year).Year != 2020 {
		t.Error(f(p1), "Add 1 Year")
	}
	if d.Subtract(1, Month).Month != 12 { // because maximum days of February is 28
		t.Error(f(p1), "Add 1 Month")
	}
	if d.Subtract(1, Day).Day != 30 { // 1.31 + 1 = 2.1
		t.Error(f(p1), "Add 1 Day")
	}
	if d.Subtract(1, Hour).Hour != 18 {
		t.Error(f(p1), "Add 1 Day")
	}
	if d.Subtract(1, Minute).Minute != 1 {
		t.Error(f(p1), "Add 1 Minute")
	}
	if d.Subtract(1, Second).Second != 59 { // 59 + 1 = 1m0s
		t.Error(f(p1), "Add 1 Second")
	}
}

func TestDayFormat(t *testing.T) {
	d, _ := Format("2021-07-30 10:23:59")
	f := ErrorF("Day.Format")
	p1 := "YYYY年MM月DD日，HH时mm分ss秒"
	if d.Format(p1) != "2021年07月30日，10时23分59秒" {
		t.Error(f(p1))
	}

	p2 := "SSS,MMMM,dd"
	if d.Format(p2) != "000,July,Friday" {
		t.Error(f(p2))
	}
}

func TestLocale(t *testing.T) {
	Locale(locale.ZH_CN)
	f := ErrorF("Locale")
	d, _ := Format("2021-07-30 10:23:59")
	p1 := "MMMM,dd"

	if d.Format(p1) != "七月,星期五" {
		t.Error(f(p1))
	}

}

func TestDayStartOf(t *testing.T) {
	d, _ := Format("2021-07-30 10:23:59")

	sfYear := d.StartOf(Year)
	if !(sfYear.Year == 2021 &&
		sfYear.Month == 1 &&
		sfYear.Day == 1 &&
		sfYear.Hour == 0 &&
		sfYear.Minute == 0 &&
		sfYear.Second == 0) {
		t.Error("Day.StartOf test failed: StartOf Year")
	}
	sfMonth := d.StartOf(Month)
	if !(sfMonth.Year == 2021 &&
		sfMonth.Month == 7 &&
		sfMonth.Day == 1 &&
		sfMonth.Hour == 0 &&
		sfMonth.Minute == 0 &&
		sfMonth.Second == 0) {
		t.Error("Day.StartOf test failed: StartOf Month")
	}

	sfDay := d.StartOf(Day)
	if !(sfDay.Year == 2021 &&
		sfDay.Month == 7 &&
		sfDay.Day == 30 &&
		sfDay.Hour == 0 &&
		sfDay.Minute == 0 &&
		sfDay.Second == 0) {
		t.Error("Day.StartOf test failed: StartOf Day")
	}

	sfHour := d.StartOf(Hour)
	if !(sfHour.Year == 2021 &&
		sfHour.Month == 7 &&
		sfHour.Day == 30 &&
		sfHour.Hour == 10 &&
		sfHour.Minute == 0 &&
		sfHour.Second == 0) {
		t.Error("Day.StartOf test failed: StartOf Hour")
	}

	sfMinute := d.StartOf(Minute)
	if !(sfMinute.Year == 2021 &&
		sfMinute.Month == 7 &&
		sfMinute.Day == 30 &&
		sfMinute.Hour == 10 &&
		sfMinute.Minute == 23 &&
		sfMinute.Second == 0) {
		t.Error("Day.StartOf test failed: StartOf Minute")
	}

	sfSecond := d.StartOf(Second)

	unixNanoStr := fmt.Sprintf("%v", sfSecond.UnixNano)
	unixNanoStr = unixNanoStr[len(unixNanoStr)-9:]
	unixNano, _ := strconv.Atoi(unixNanoStr)

	if !(sfSecond.Year == 2021 &&
		sfSecond.Month == 7 &&
		sfSecond.Day == 30 &&
		sfSecond.Hour == 10 &&
		sfSecond.Minute == 23 &&
		sfSecond.Second == 59 &&
		unixNano == 0) {
		t.Error("Day.StartOf test failed: StartOf Second")
	}
}

func TestDayEndOf(t *testing.T) {
	d, _ := Format("2021-07-30 10:23:20")

	ef := d.EndOf(Year)
	if !(ef.Year == 2021 &&
		ef.Month == 12 &&
		ef.Day == 31 &&
		ef.Hour == 23 &&
		ef.Minute == 59 &&
		ef.Second == 59) {
		t.Error("Day.EndOf test failed: EndOf Year")
	}

	ef = d.EndOf(Month)
	if !(ef.Year == 2021 &&
		ef.Month == 7 &&
		ef.Day == 31 &&
		ef.Hour == 23 &&
		ef.Minute == 59 &&
		ef.Second == 59) {
		t.Error("Day.EndOf test failed: EndOf Month")
	}

	ef = d.EndOf(Day)
	if !(ef.Year == 2021 &&
		ef.Month == 7 &&
		ef.Day == 30 &&
		ef.Hour == 23 &&
		ef.Minute == 59 &&
		ef.Second == 59) {
		t.Error("Day.EndOf test failed: EndOf Day")
	}

	ef = d.EndOf(Hour)
	if !(ef.Year == 2021 &&
		ef.Month == 7 &&
		ef.Day == 30 &&
		ef.Hour == 10 &&
		ef.Minute == 59 &&
		ef.Second == 59) {
		t.Error("Day.EndOf test failed: EndOf Hour")
	}

	ef = d.EndOf(Minute)
	if !(ef.Year == 2021 &&
		ef.Month == 7 &&
		ef.Day == 30 &&
		ef.Hour == 10 &&
		ef.Minute == 23 &&
		ef.Second == 59) {
		t.Error("Day.EndOf test failed: EndOf Minute")
	}

	ef = d.EndOf(Second)

	unixNanoStr := fmt.Sprintf("%v", ef.UnixNano)
	unixNanoStr = unixNanoStr[len(unixNanoStr)-9:]
	unixNano, _ := strconv.Atoi(unixNanoStr)

	if !(ef.Year == 2021 &&
		ef.Month == 7 &&
		ef.Day == 30 &&
		ef.Hour == 10 &&
		ef.Minute == 23 &&
		ef.Second == 20 &&
		unixNano == 999999999) {
		t.Error("Day.EndOf test failed: EndOf Second")
	}
}

func TestDaySet(t *testing.T) {
	d, _ := Format("2021-02-17 10:23:20")
	f := ErrorF("Day.Set")
	if d.Set(2024, Year).Year != 2024 {
		t.Error(f(), "Set Year")
	}
	if d.Set(3, Month).Month != 3 {
		t.Error(f(), "Set Month")
	}
	if d.Set(29, Day).Day != 1 {
		t.Error(f(), "Set Day")
	}
	if d.Set(3, Hour).Hour != 3 {
		t.Error(f(), "Set Month")
	}
	if d.Set(1, Minute).Minute != 1 {
		t.Error(f(), "Set Minute")
	}
	if d.Set(59, Second).Second != 59 {
		t.Error(f(), "Set Second")
	}
}

func TestIsLeapYear(t *testing.T) {
	f := ErrorF("IsLeapYear")
	if IsLeapYear(2021) != false && IsLeapYear(2100) != false {
		t.Error(f(2021))
		t.Error(f(2100))
	}
	if !IsLeapYear(2020) && !IsLeapYear(2400) {
		t.Error(f(2020))
		t.Error(f(2400))
	}
}

func TestFillZero(t *testing.T) {
	str1 := fillZero(8)
	str2 := fillZero(12)
	f := ErrorF("fill zero")
	if str1 != "08" {
		t.Error(f(8))
	}
	if str2 != "12" {
		t.Error(f(12))
	}
}
func TestDaySecondAfterUnixNano(t *testing.T) {
	tim := time.Unix(0, 1627637214376669500)
	d := New(tim)
	un := d.SecondAfterUnixNano()
	if un != 376669500 {
		t.Error("SecondAfterUnixNano test failed")
	}
}

func TestDayDaysInMonth(t *testing.T) {
	d1 := List([]int{2021, 8})
	d2 := List([]int{2021, 2})
	d3 := List([]int{2020, 2})
	f := ErrorF("day.DaysInMonth")

	if d1.DaysInMonth() != 31 {
		t.Error(f(), "2021, 8")
	}
	if d2.DaysInMonth() != 28 {
		t.Error(f(), "2021, 2")
	}
	if d3.DaysInMonth() != 29 {
		t.Error(f(), "2020, 2")
	}

}
