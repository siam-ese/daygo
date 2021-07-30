package day

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestDayNow(t *testing.T) {
	d := Now()
	if d.Day != time.Now().Day() {
		t.Error("day.Day exception")
	}
}

// func TestParseList(t *testing.T) {
// 	_, _, _, _, _, _, err := parseList([]int{2021, 13, 31})

// 	if err == nil {
// 		t.Error("not caught [parseList]: month error")
// 	}
// }

func TestFormat(t *testing.T) {
	d, _ := Format("2021-07-28 12:26:36")
	if d.Year != 2021 {
		t.Error("Format error", d.Year)
	}
}

func TestUnix(t *testing.T) {
	d, _ := Unix(1627542959340)
	if d.Year != 2021 && d.Minute != 15 {
		t.Error(fmt.Printf("unix parse error year: %v, minute: %v", d.Year, d.Minute))
	}

	param2 := 162754295934
	_, err := Unix(param2)
	if err != nil {
		log.Println("unix exception error-1:", err)
	} else {
		t.Error("unix incorrect parse Parameter is:", param2)
	}
}

func TestAdd(t *testing.T) {
	d, _ := Format("2021-07-29 19:02:60")
	oldY := d.Year
	oldM := d.Month
	oldD := d.Day
	oldH := d.Hour
	oldm := d.Minute
	olds := d.Second

	newY := d.Add(1, Year).Year
	newM := d.Add(1, Month).Month
	newD := d.Add(1, Day).Day
	newH := d.Add(1, Hour).Hour
	newm := d.Add(1, Minute).Minute
	news := d.Add(1, Second).Second
	log.Println("weekday \n", int(d.WeekDay))
	log.Println(fmt.Sprintf("Year old: %v, new: %v", oldY, newY))
	log.Println(fmt.Sprintf("Month old: %v, new: %v", oldM, newM))
	log.Println(fmt.Sprintf("Day old: %v, new: %v", oldD, newD))
	log.Println(fmt.Sprintf("Hour old: %v, new: %v", oldH, newH))
	log.Println(fmt.Sprintf("Minute old: %v, new: %v", oldm, newm))
	log.Println(fmt.Sprintf("Second old: %v, new: %v", olds, news))
}

func TestDayFormat(t *testing.T) {
	d := Now()
	d.Format("YYYY年MM月DD日，HH时mm分ss秒")
}
