package day

import (
	"regexp"
)

func parseT(t string) []string {
	c := regexp.MustCompile(`^(\d{4})[-\/]?(\d{1,2})?[-\/]?(\d{0,2})[Tt\s]*(\d{1,2})?:?(\d{1,2})?:?(\d{1,2})?[.:]?(\d+)?$`)

	ret := c.FindStringSubmatch(t)

	if len(ret) <= 0 {
		return nil
	}

	return ret
}

func parseList(list []int) (year, month, day, hour, minute, second int) {
	l := len(list)
	if l > 0 {
		year = list[0]
	}
	if l > 1 {
		month = list[1]
	}
	if l > 2 {
		day = list[2]
	}
	if l > 3 {
		hour = list[3]
	}
	if l > 4 {
		minute = list[4]
	}
	if l > 5 {
		second = list[5]
	}

	return year, month, day, hour, minute, second
}
